package world

type SectorConnection struct {
	sectorA SectorKey
	sectorB SectorKey
}

func ConnectSectors(sectorA *Sector, sectorB *Sector) {
	aConnection := SectorConnection{}
	aConnection.sectorA = sectorA.Id
	aConnection.sectorB = sectorB.Id
	sectorA.AddNearbySector(sectorB.Id, aConnection)

	bConnection := SectorConnection{}
	bConnection.sectorA = sectorB.Id
	bConnection.sectorB = sectorA.Id
	sectorB.AddNearbySector(sectorA.Id, bConnection)
}
