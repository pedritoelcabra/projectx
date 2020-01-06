package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/container"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
)

type chunk struct {
	tiles               []*Tile
	Vegetation          []int
	ChunkData           *container.Container
	Location            tiling.Coord
	Generated           bool
	queuedForGeneration bool
	isPreloaded         bool
	terrainImage        *ebiten.Image
	sector              *Sector
	SectorId            SectorKey
}

func NewChunk(location tiling.Coord) *chunk {
	aChunk := &chunk{}
	aChunk.isPreloaded = false
	aChunk.terrainImage = nil
	aChunk.ChunkData = container.NewContainer()
	aChunk.Preload(location)
	return aChunk
}

func (ch *chunk) GetSector() *Sector {
	return ch.sector
}

func (ch *chunk) SetSector(sector *Sector) {
	ch.sector = sector
}

func (ch *chunk) Preload(location tiling.Coord) {
	if ch.isPreloaded {
		return
	}
	ch.tiles = make([]*Tile, ChunkSize*ChunkSize)
	ch.Location = location
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			tileX := (ch.Location.X() * ChunkSize) + x
			tileY := (ch.Location.Y() * ChunkSize) + y
			tileLocation := tiling.NewCoord(tileX, tileY)
			tileIndex := ch.tileIndex(tileX, tileY)
			aTile := NewTile()
			aTile.coordinates = tileLocation

			centerX := float64(tileX) * tiling.TileHorizontalSeparation
			centerY := float64(tileY) * tiling.TileHeight
			if x%2 > 0 {
				centerY += tiling.TileHeight / 2
			}
			renderX := centerX - tiling.TileWidth/2
			renderY := centerY - tiling.TileHeight/2
			aTile.SetF(RenderX, renderX)
			aTile.SetF(RenderY, renderY)
			renderDoubleX := centerX - tiling.TileWidth
			renderDoubleY := centerY - tiling.TileHeight
			aTile.SetF(RenderDoubleX, renderDoubleX)
			aTile.SetF(RenderDoubleY, renderDoubleY)
			aTile.SetF(CenterX, centerX)
			aTile.SetF(CenterY, centerY)
			ch.tiles[tileIndex] = aTile
		}
	}
	if ch.IsGenerated() {
		for index, tile := range ch.tiles {
			if ch.Vegetation[index] != 0 {
				tile.Set(Flora, ch.Vegetation[index])
			}
		}
	}
	ch.RunOnAllTiles(func(t *Tile) {
		t.InitializeTile()
	})
	ch.PreloadChunkData()
	ch.isPreloaded = true
}

func (ch *chunk) PreSave() {
	ch.Vegetation = make([]int, ChunkSize*ChunkSize)
	for index, tile := range ch.tiles {
		ch.Vegetation[index] = tile.Get(Flora)
	}
}

func (ch *chunk) PreloadChunkData() {
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

func (ch *chunk) IsGenerated() bool {
	return ch.Generated
}

func (ch *chunk) IsQueueForGeneration() bool {
	return ch.queuedForGeneration
}

func (ch *chunk) RunOnAllTiles(f func(t *Tile)) {
	for _, c := range ch.tiles {
		f(c)
	}
}

func (ch *chunk) Tile(tileCoord tiling.Coord) *Tile {
	return ch.tiles[ch.tileIndex(tileCoord.X(), tileCoord.Y())]
}

func (ch *chunk) tileIndex(x, y int) int {
	x -= ch.Location.X() * ChunkSize
	y -= ch.Location.Y() * ChunkSize
	return (x * ChunkSize) + y
}

func (ch *chunk) SetImage(image *ebiten.Image) {
	ch.terrainImage = image
}

func (ch *chunk) GetImage() *ebiten.Image {
	return ch.terrainImage
}

func (ch *chunk) GenerateImage() {
	if ch.GetImage() != nil {
		return
	}
	imageWidth := (ChunkSize + 1) * tiling.TileHorizontalSeparation
	imageHeight := (ChunkSize + 1) * tiling.TileHeight
	ch.terrainImage, _ = ebiten.NewImage(int(imageWidth), int(imageHeight), ebiten.FilterDefault)

	for _, t := range ch.tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(tiling.TileWidthScale, tiling.TileHeightScale)
		xOff := tiling.TileHorizontalSeparation * float64(ch.Location.X()*ChunkSize)
		yOff := tiling.TileHeight * float64(ch.Location.Y()*ChunkSize)
		localX := t.GetF(RenderX) - xOff + (tiling.TileWidth / 2)
		localY := t.GetF(RenderY) - yOff + (tiling.TileHeight / 2)
		gfx.DrawHexTerrainToImage(localX, localY, t.Get(TerrainBase), ch.terrainImage, op)
	}
}

func (ch *chunk) FirstTile() *Tile {
	return ch.tiles[0]
}
