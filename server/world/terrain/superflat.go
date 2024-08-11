package terrain

import (
	"math/rand"

	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

// A superflat chunk generator. A superflat world is just 4 layers. One bedrock, two dirt, and one grass block, so it is very easy to implement
type SuperflatTerrain struct {
}

func (SuperflatTerrain) NewChunk(cx, cz int32) chunk.Chunk {
	c := chunk.NewChunk(cx, cz)

	for x := int32(0); x < 16; x++ {
		for z := int32(0); z < 16; z++ {
			c.SetBlock(x, 4, z, grassBlock)
			c.SetBlock(x, 2, z, dirt)
			c.SetBlock(x, 1, z, dirt)
			c.SetBlock(x, 0, z, bedrock)
		}
	}

	return c
}

func (SuperflatTerrain) GenerateWorldSpawn() (x, y, z int32) {
	return rand.Int31n(160), 4, rand.Int31n(160)
}
