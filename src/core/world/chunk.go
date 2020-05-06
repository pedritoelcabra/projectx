package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	container2 "github.com/pedritoelcabra/projectx/src/core/world/container"
	tiling2 "github.com/pedritoelcabra/projectx/src/core/world/tiling"
	"github.com/pedritoelcabra/projectx/src/gfx"
)

type Chunk struct {
	tiles               []*Tile
	Features            []map[int]int
	ChunkData           *container2.Container
	Location            tiling2.Coord
	Generated           bool
	queuedForGeneration bool
	isPreloaded         bool
	terrainImage        *ebiten.Image
	sector              *Sector
	SectorId            SectorKey
	units               UnitArray
	unitsLastUpdated    int
}

var savableChunkData = []int{
	Flora,
	Road,
}

func NewChunk(location tiling2.Coord) *Chunk {
	aChunk := &Chunk{}
	aChunk.isPreloaded = false
	aChunk.terrainImage = nil
	aChunk.units = UnitArray{}
	aChunk.unitsLastUpdated = 0
	aChunk.ChunkData = container2.NewContainer()
	aChunk.Preload(location)
	return aChunk
}

func (ch *Chunk) GetSector() *Sector {
	return ch.sector
}

func (ch *Chunk) SetSector(sector *Sector) {
	ch.sector = sector
}

func (ch *Chunk) Preload(location tiling2.Coord) {
	if ch.isPreloaded {
		return
	}
	ch.tiles = make([]*Tile, ChunkSizeSquare)
	ch.Location = location
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			tileX := (ch.Location.X() * ChunkSize) + x
			tileY := (ch.Location.Y() * ChunkSize) + y
			tileLocation := tiling2.NewCoord(tileX, tileY)
			tileIndex := ch.tileIndex(tileX, tileY)
			aTile := NewTile()
			aTile.coordinates = tileLocation

			centerX := float64(tileX) * tiling2.TileHorizontalSeparation
			centerY := float64(tileY) * tiling2.TileHeight
			if x%2 > 0 {
				centerY += tiling2.TileHeight / 2
			}
			renderX := centerX - tiling2.TileWidth/2
			renderY := centerY - tiling2.TileHeight/2
			aTile.SetF(RenderX, renderX)
			aTile.SetF(RenderY, renderY)
			renderDoubleX := centerX - tiling2.TileWidth
			renderDoubleY := centerY - tiling2.TileHeight
			aTile.SetF(RenderDoubleX, renderDoubleX)
			aTile.SetF(RenderDoubleY, renderDoubleY)
			aTile.SetF(CenterX, centerX)
			aTile.SetF(CenterY, centerY)
			ch.tiles[tileIndex] = aTile
		}
	}
	if ch.IsGenerated() {
		for index, tile := range ch.tiles {
			for _, featureKey := range savableChunkData {
				tile.Set(featureKey, ch.Features[index][featureKey])
			}
		}
	}
	ch.RunOnAllTiles(func(t *Tile) {
		t.InitializeTile()
	})
	ch.PreloadChunkData()
	ch.isPreloaded = true
}

func (ch *Chunk) PreSave() {
	if !ch.IsGenerated() {
		return
	}
	ch.Features = make([]map[int]int, ChunkSize*ChunkSize)
	for index, tile := range ch.tiles {
		ch.Features[index] = make(map[int]int)
		for _, featureKey := range savableChunkData {
			ch.Features[index][featureKey] = tile.Get(featureKey)
		}
	}
}

func (ch *Chunk) PreloadChunkData() {
	totalHeight := 0
	maxHeight := 0
	minHeight := 0
	totalTiles := ChunkSize * ChunkSize
	for _, tile := range ch.tiles {
		tileHeight := tile.Get(Height)
		totalHeight += tileHeight
		if maxHeight < tileHeight {
			maxHeight = tileHeight
		}
		if minHeight > tileHeight {
			minHeight = tileHeight
		}
	}
	ch.ChunkData.Set(AvgHeight, totalHeight/totalTiles)
	ch.ChunkData.Set(MaxHeight, maxHeight)
	ch.ChunkData.Set(MinHeight, minHeight)
}

func (ch *Chunk) IsGenerated() bool {
	return ch.Generated
}

func (ch *Chunk) IsQueueForGeneration() bool {
	return ch.queuedForGeneration
}

func (ch *Chunk) RunOnAllTiles(f func(t *Tile)) {
	for _, c := range ch.tiles {
		f(c)
	}
}

func (ch *Chunk) Tile(tileCoord tiling2.Coord) *Tile {
	return ch.tiles[ch.tileIndex(tileCoord.X(), tileCoord.Y())]
}

func (ch *Chunk) tileIndex(x, y int) int {
	x -= ch.Location.X() * ChunkSize
	y -= ch.Location.Y() * ChunkSize
	return (x * ChunkSize) + y
}

func (ch *Chunk) SetImage(image *ebiten.Image) {
	ch.terrainImage = image
}

func (ch *Chunk) GetImage() *ebiten.Image {
	return ch.terrainImage
}

func (ch *Chunk) GenerateImage() {
	if ch.GetImage() != nil {
		return
	}
	imageWidth := (ChunkSize + 1) * tiling2.TileHorizontalSeparation
	imageHeight := (ChunkSize + 1) * tiling2.TileHeight
	ch.terrainImage, _ = ebiten.NewImage(int(imageWidth), int(imageHeight), ebiten.FilterDefault)

	for _, t := range ch.tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(tiling2.TileWidthScale, tiling2.TileHeightScale)
		xOff := tiling2.TileHorizontalSeparation * float64(ch.Location.X()*ChunkSize)
		yOff := tiling2.TileHeight * float64(ch.Location.Y()*ChunkSize)
		localX := t.GetF(RenderX) - xOff + (tiling2.TileWidth / 2)
		localY := t.GetF(RenderY) - yOff + (tiling2.TileHeight / 2)
		gfx.DrawHexTerrainToImage(localX, localY, t.Get(TerrainBase), ch.terrainImage, op)
	}
}

func (ch *Chunk) FirstTile() *Tile {
	return ch.tiles[0]
}

func (ch *Chunk) GenerateNPCs() {
	if !ch.IsGenerated() {
		return
	}
	if !ch.ShouldSpawnNPCs() {
		return
	}
	spawnTile := ch.GetRandomTile()
	if !ch.IsValidNPCSpawnTile(spawnTile) {
		return
	}
	def := ch.ChooseRandomNPCsTemplate()
	if def == nil {
		return
	}
	NewUnit(def.Name, tiling2.NewCoordF(spawnTile.GetRenderPos()))
	//logger.General("Generated a "+def.Name+" at "+spawnTile.GetCoord().ToString(), nil)
}

func (ch *Chunk) ChooseRandomNPCsTemplate() *defs.UnitDef {
	bestScore := -1
	bestDef := &defs.UnitDef{}
	for _, unit := range defs.UnitDefs() {
		if !unit.SpawnsWild {
			continue
		}
		thisScore := randomizer.RandomInt(0, 100)
		if thisScore > bestScore {
			bestScore = thisScore
			bestDef = unit
		}
	}
	if bestScore < 0 {
		return nil
	}
	return bestDef
}

func (ch *Chunk) IsValidNPCSpawnTile(tile *Tile) bool {
	if EntityShouldDraw(tile.GetRenderPos()) {
		return false
	}
	if tile.HasSector() {
		return false
	}
	if !tile.IsLand() || tile.IsImpassable() {
		return false
	}
	return true
}

func (ch *Chunk) GetRandomTile() *Tile {
	return ch.tiles[randomizer.RandomInt(0, ChunkSizeSquare-1)]
}

func (ch *Chunk) ShouldSpawnNPCs() bool {
	maxAmountOfNPCs := 1
	return len(ch.GetUnits()) <= maxAmountOfNPCs
}

func (ch *Chunk) RegisterUnit(unit *Unit) {
	ch.CheckUnitArray()
	ch.units = append(ch.units, unit)
}

func (ch *Chunk) GetUnits() UnitArray {
	ch.CheckUnitArray()
	return ch.units
}

func (ch *Chunk) CheckUnitArray() {
	if ch.unitsLastUpdated != theWorld.GetTick() {
		ch.units = UnitArray{}
		ch.unitsLastUpdated = theWorld.GetTick()
	}
}

func ChunksAroundTile(tile tiling2.Coord, radius int) []*Chunk {
	var chunks []*Chunk
	chunkCoord := theWorld.Grid.ChunkCoord(tile)
	for x := chunkCoord.X() - 3; x <= chunkCoord.X()+3; x++ {
		for y := chunkCoord.Y() - 3; y <= chunkCoord.Y()+3; y++ {
			chunks = append(chunks, theWorld.Grid.Chunk(tiling2.NewCoord(x, y)))
		}
	}
	return chunks
}

func ChunksAroundPlayer(radius int) []*Chunk {
	return ChunksAroundTile(tiling2.NewCoord(tiling2.PixelFToTileI(theWorld.PlayerUnit.GetPos())), radius)
}
