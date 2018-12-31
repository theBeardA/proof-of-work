package blockchain

import (
	"work_queue"
)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	Chunk_start uint64
	Chunk_end   uint64
	New_Block   Block
}

func (m_worker *miningWorker) Run() interface{} {

	var new_Result MiningResult
	blk := m_worker.New_Block
	for i := m_worker.Chunk_start; i <= m_worker.Chunk_end; i++ {
		blk.Proof = i
		blk.Hash = blk.CalcHash()
		if blk.ValidHash() == true {
			new_Result.Proof = i
			new_Result.Found = true
			return new_Result
		} else {
			new_Result.Proof = i
			new_Result.Found = false
		}
	}
	return new_Result
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	new_queue := work_queue.Create(uint(workers), uint(chunks))
	//as not promised that number of chunks would be divisible by the total length
	// the loop runs upto equal chunks upto(chunks - 1), starting from zero so it says (chunks - 2)
	chunk_len := ((end - start) + 1) / chunks
	var i uint64
	for i = 0; i <= chunks-2; i++ {
		new_Worker := new(miningWorker)
		new_Worker.New_Block = blk
		new_Worker.Chunk_start = i * chunk_len
		new_Worker.Chunk_end = (chunk_len - 1) + i*chunk_len
		new_queue.Enqueue(new_Worker)
	}
	//the last chunk could be a little longer length than others
	new_Worker1 := new(miningWorker)
	new_Worker1.New_Block = blk
	new_Worker1.Chunk_start = (chunk_len * (chunks - 1))
	new_Worker1.Chunk_end = end
	new_queue.Enqueue(new_Worker1)
	for i = 0; i < chunks; i++ {
		v := <-new_queue.Results
		final_Result := v.(MiningResult)

		if final_Result.Found == true {
			new_queue.Shutdown()
			return final_Result
		}
	}
	return blk.MineRange(end+1, end^2, workers, chunks)

}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	//fmt.Println(mr.Found)
	//fmt.Println(mr.Proof)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found

}

// a simple method that loops over proof values starting at 0 until a valid one is found.
func (blk *Block) proof_checker() uint64 {
	check := false
	var proof uint64 = 0
	for i := proof; check == false; i++ {

		blk.Proof = proof
		blk.Hash = blk.CalcHash()
		if blk.ValidHash() == true {
			check = true
		}
	}
	return proof
}
