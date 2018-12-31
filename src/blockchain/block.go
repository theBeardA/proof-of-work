package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	Generation uint64
	Difficulty uint8
	Data       string
	PrevHash   []byte
	Hash       []byte
	Proof      uint64
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	b := new(Block)
	b.Generation = 0
	b.Difficulty = difficulty
	b.Data = ""
	for i := 0; i < 32; i++ {
		b.PrevHash = append(b.PrevHash, 0)
	}
	// TODO
	return *b

}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	b := new(Block)
	b.Generation = prev_block.Generation + 1
	b.Difficulty = prev_block.Difficulty
	b.Data = data
	b.PrevHash = prev_block.Hash
	return *b
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	prHash_str := hex.EncodeToString(blk.PrevHash)
	gen_str := strconv.FormatUint(uint64(blk.Generation), 10)
	dif_str := strconv.FormatUint(uint64(blk.Difficulty), 10)
	proof_str := strconv.FormatUint(uint64(blk.Proof), 10)
	Hash_str := prHash_str + ":" + gen_str + ":" + dif_str + ":" + blk.Data + ":" + proof_str

	hash_n := sha256.New()
	hash_n.Write([]byte(Hash_str))
	return hash_n.Sum(nil)
	// TODO
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO
	//check := false
	dif := blk.Difficulty
	var counter uint8 = 0
	hash_len := uint8(len(blk.Hash))
	hash_arr := blk.Hash
	for i := hash_len - 1; i >= hash_len-dif; i-- {
		if hash_arr[i] == 0 {
			counter = counter + 1
		} else {
			counter = 0
		}
	}
	if counter == dif {
		return true
	}
	return false
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
