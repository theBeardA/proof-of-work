package blockchain

import (
	"sync"
)

var wg sync.WaitGroup

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}

	chain.Chain = append(chain.Chain, blk)

	// TODO
}

func (chain Blockchain) IsValid() bool {
	// TODO
	check := true
	validator_ch := make(chan bool, 6)
	defer close(validator_ch)
	wg.Add(6)
	go initBlock(chain.Chain[0], validator_ch)
	go DifValue(chain.Chain, validator_ch)
	go GenValue(chain.Chain, validator_ch)
	go PHashValue(chain.Chain, validator_ch)
	go HashValue(chain.Chain, validator_ch)
	go hashDiff(chain.Chain, validator_ch)
	wg.Wait()
	for i := 0; i < 6; i++ {
		res := <-validator_ch
		if res == true {
			check = true && check
		} else {
			check = false && check
		}

	}

	return check

}

//supporting goroutines
//The initial block has previous hash all null bytes and is generation zero.
func initBlock(blk Block, ch chan bool) {
	defer wg.Done()
	counter := 0
	for _, i := range blk.PrevHash {
		if i == 0 {
			counter++
		}
	}
	if counter == len(blk.PrevHash) && blk.Generation == 0 {
		ch <- true
	}
}

//Each block has the same difficulty value.
func DifValue(chain []Block, ch chan bool) {
	defer wg.Done()
	check := true
	diff := chain[0].Difficulty
	for _, i := range chain {
		if i.Difficulty == diff {
			check = true && check
		} else {
			check = false && check
		}
	}
	ch <- check
}

//Each block has a generation value that is one more than the previous block.
func GenValue(chain []Block, ch chan bool) {
	defer wg.Done()
	check := true
	Gen1 := chain[0].Generation
	for pos, i := range chain {
		if i.Generation == Gen1+uint64(pos) {
			check = true && check
		} else {
			check = false && check
		}
	}
	ch <- check
}

//Each block's previous hash matches the previous block's hash.
func PHashValue(chain []Block, ch chan bool) {
	defer wg.Done()
	check := true

	for i := 1; i < len(chain); i++ {
		pHash := string(chain[i].PrevHash)
		Hash := string(chain[i-1].Hash)
		if pHash == Hash {
			check = true && check
		} else {
			check = false && check
		}
	}
	ch <- check
}

//Each block's hash value actually matches its contents.
func HashValue(chain []Block, ch chan bool) {
	defer wg.Done()
	check := true

	for _, i := range chain {
		Hash_Saved := string(i.Hash)
		Hash_comp := string(i.CalcHash())
		if Hash_Saved == Hash_comp {
			check = true && check
		} else {
			check = false && check
		}
	}
	ch <- check
}

//Each block's hash value ends in difficulty null bytes.
func hashDiff(chain []Block, ch chan bool) {
	defer wg.Done()
	check := true

	for _, i := range chain {
		if i.ValidHash() == true {
			check = true && check
		} else {
			check = false && check
		}
	}
	ch <- check
}
