package contracts

type ITerrainGenerator interface {
	Generate(terrainType TerrainType, width, height int) ITerrain
}
