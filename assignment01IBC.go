package assignment01IBC

import ("fmt"
  "crypto/sha256"
  "encoding/hex"	
)
type Block struct {
	transaction string
	prevPointer *Block
	preHash string
	}

// Reference https://gobyexample.com/sha1-hashes
	func calculatehash(chainHead *Block) string{
		if chainHead==nil{
			return "nil"
		}else{
			rec:=chainHead.transaction+chainHead.preHash
			y:=sha256.New()
			y.Write([]byte(rec))
			hashcal:=y.Sum(nil)
			return hex.EncodeToString(hashcal)
	
		}
	}

func InsertBlock(transaction string, chainHead *Block) *Block {
	
	var block Block
	if (chainHead==nil){
		
		block.transaction=transaction
		block.preHash=calculatehash(chainHead)
		block.prevPointer=nil
		return &block

	}else{
		block.transaction=transaction
		block.preHash=calculatehash(chainHead)
		block.prevPointer=chainHead
		return &block
		}
}

func ListBlocks(chainHead *Block) {
	//var block Block
	block:=chainHead
	if(block==nil){
		fmt.Println("No transaction has ever done!! it is empty")
	}else{
	for ;block!=nil;{
		fmt.Println("The transaction done was-->"+block.transaction)
		block=block.prevPointer
	}
}
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	node:=chainHead
	var entry string
	var count int
	for ;node!=nil;{
		entry=node.transaction
		if(entry==oldTrans){
			fmt.Println("we found the entry ")
			node.transaction=newTrans
			count=count+1
		}
		chainHead=node
		chainHead=chainHead.prevPointer
		node=node.prevPointer
	}
	if(count==0){
		fmt.Println("sorry we couldn't find the requested transaction")
	}
}


func VerifyChain(chainHead *Block) {
	var hash1 string
	var count int
	for ;chainHead.prevPointer!=nil;{
		hash1=calculatehash(chainHead.prevPointer)
		if (chainHead.preHash!=hash1){ //the hash func has error
			fmt.Print("someone has changed the blockchain")
			fmt.Print(chainHead.prevPointer.transaction+" is effected")
			count=count+1
		}
		chainHead=chainHead.prevPointer
	}
	if(count==0){
		fmt.Print("Blockchain is safe")
	}
}