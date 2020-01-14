package world

import "github.com/pedritoelcabra/projectx/src/world/container"

type SaveGameData struct {
	Seed          int
	Tick          int
	Player        Player
	Grid          Grid
	WorldEntities *Entities
	Data          *container.Container
}
