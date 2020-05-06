package world

import (
	container2 "github.com/pedritoelcabra/projectx/src/core/world/container"
)

type SaveGameData struct {
	Seed          int
	Tick          int
	Player        Player
	Grid          Grid
	WorldEntities *Entities
	Data          *container2.Container
}
