package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/gfx"
	container2 "github.com/pedritoelcabra/projectx/src/world/container"
	tiling2 "github.com/pedritoelcabra/projectx/src/world/tiling"
)

type chunk struct {
	tiles               []*Tile
	ChunkData           *container2.Container
	Location            tiling2.Coord
	Generated           bool
	queuedForGeneration bool
	isPreloaded         bool
	terrainImage        *ebiten.Image
	sector              *Sector
	SectorId            SectorKey
}

func NewChunk(location tiling2.Coord) *chunk {
	aChunk := &chunk{}
	aChunk.isPreloaded = false
	aChunk.terrainImage = nil
	aChunk.ChunkData = container2.NewContainer()
	aChunk.Preload(location)
	return aChunk
}

func (ch *chunk) GetSector() *Sector {
	return ch.sector
}

func (ch *chunk) SetSector(sector *Sector) {
	ch.sector = sector
}

func (ch *chunk) Preload(location tiling2.Coord) {
	if ch.isPreloaded {
		return
	}
	ch.tiles = make([]*Tile, ChunkSize*ChunkSize)
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
			aTile.SetF(CenterX, centerX)
			aTile.SetF(CenterY, centerY)
			ch.tiles[tileIndex] = aTile
		}
	}
	ch.RunOnAllTiles(func(t *Tile) {
		t.InitializeTile()
	})
	ch.PreloadChunkData()
	ch.isPreloaded = true
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

func (ch *chunk) Tile(tileCoord tiling2.Coord) *Tile {
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

func (ch *chunk) FirstTile() *Tile {
	return ch.tiles[0]
}
