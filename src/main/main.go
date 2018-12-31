package main

import (
	"blockchain"
	"encoding/hex"
	"fmt"
)

func main() {

	//b0 := blockchain.Initial(2)
	//b0.Proof = 242278
	//b0.Mine()
	//b1 := b0.Next("message")
	//b1.Mine()

	/*fmt.Println(b1.Generation)
	fmt.Println(b1.Difficulty)
	fmt.Println(b1.Data)*/
	//fmt.Println(b1.PrevHash)

	//fmt.Println(len(h))

	//fmt.Println(hex.EncodeToString(h))
	//fmt.Println(hex.EncodeToString(b1.CalcHash()))

	//var coin blockchain.Blockchain
	//var coin blockchain.Blockchain
	b0 := blockchain.Initial(2)
	b0.Mine(1)
	//fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	//fmt.Println(b0.ValidHash())
	//coin.Add(b0)
	//var b1 blockchain.Block
	//coin.Add(b0)
	for i := 0; i < 10; i++ {

		b0 = b0.Next("Hello")
		b0.Mine(1)
		//fmt.Println("3")
		fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
		//coin.Add(b0)

	}
	//fmt.Print(coin)
	//fmt.Println(coin.IsValid())

	//b51 := b0.Next("Hello")
	//b0.Mine(1)
	//fmt.Println("3")
	//fmt.Println(b51.Proof, b51.Generation)

	//coin.Add(b1)
	//fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	//b2 := b1.Next("this is not interesting")
	//b2.Mine(1)
	//coin.Add(b2)
	//fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))*/
	//fmt.Println(coin.Chain[0])
	//fmt.Println(coin)
	/*b0 := blockchain.Initial(2)
	b0.Mine(1)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
	b3 := b2.Next("this is an interesting message")
	b3.Mine(1)
	fmt.Println(b3.Proof, hex.EncodeToString(b3.Hash))
	b4 := b3.Next("this is not interesting")
	b4.Mine(1)
	fmt.Println(b4.Proof, hex.EncodeToString(b4.Hash))
	b5 := b4.Next("this is an interesting message")
	b5.Mine(1)
	fmt.Println(b5.Proof, hex.EncodeToString(b5.Hash))
	b6 := b5.Next("this is not interesting")
	b6.Mine(1)
	fmt.Println(b6.Proof, hex.EncodeToString(b6.Hash))
	b7 := b6.Next("this is an interesting message")
	b7.Mine(1)
	fmt.Println(b7.Proof, hex.EncodeToString(b7.Hash))
	b8 := b7.Next("this is not interesting")
	b8.Mine(1)
	fmt.Println(b8.Proof, hex.EncodeToString(b8.Hash))
	b9 := b8.Next("this is an interesting message")
	b9.Mine(1)
	fmt.Println(b9.Proof, hex.EncodeToString(b9.Hash))
	b10 := b9.Next("this is not interesting")
	b10.Mine(1)
	fmt.Println(b10.Proof, hex.EncodeToString(b10.Hash))
	b11 := b10.Next("this is not interesting")
	b11.Mine(1)
	fmt.Println(b11.Proof, hex.EncodeToString(b11.Hash))
	b12 := b11.Next("this is not interesting")
	b12.Mine(1)
	fmt.Println(b12.Proof, hex.EncodeToString(b12.Hash))*/

}
