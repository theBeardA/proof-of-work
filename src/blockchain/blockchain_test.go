package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: some useful tests of Blocks
// The block chain created here can be tested for few function testin, as multiple functions are used here
func CreateBlockChain() Blockchain {
	var coin Blockchain
	b0 := Initial(2)
	b0.Mine(1)
	coin.Add(b0)
	for i := 0; i < 100; i++ {
		b0 = b0.Next("Hello")
		b0.Mine(1)
		coin.Add(b0)
	}
	return coin
}

func TestNext(t *testing.T) {

	coin := CreateBlockChain()
	for i := 1; i < 100; i++ {
		assert.Equal(t, coin.Chain[i].Difficulty, coin.Chain[i-1].Difficulty)
		assert.Equal(t, (coin.Chain[i].Generation - 1), coin.Chain[i-1].Generation)
		assert.Equal(t, (coin.Chain[i].PrevHash), coin.Chain[i-1].Hash)
	}
}

func TestMinedHash(t *testing.T) {
	coin := CreateBlockChain()
	for i := 0; i < 100; i++ {
		if !assert.Equal(t, coin.Chain[i].Hash, coin.Chain[i].CalcHash()) {
			t.Errorf("The block number %d was mined incorrectly ", i)
		}
	}
}

func TestValidHash(t *testing.T) {
	coin := CreateBlockChain()
	for i := 0; i < 100; i++ {
		if !assert.Equal(t, true, coin.Chain[i].ValidHash()) {
			t.Error("invalid Hash")
		}
	}
}

func TestMine(t *testing.T) {
	coin := CreateBlockChain()
	for i := 0; i < 10; i++ {
		blk := coin.Chain[i]
		blk.MineRange(0, uint64(4*1<<(10*blk.Difficulty)), 10, 4321)
		if !assert.NotEmpty(t, coin.Chain[i].Proof) {
			t.Error("The value is out of range, use higher range")
		}

	}
}
