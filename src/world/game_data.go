package world

type SaveGameData struct {
	Seed     int
	Tick     int
	Player   Player
	Grid     Grid
	Entities EntityMap
	Sectors  SectorMap
}
