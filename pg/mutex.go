package pg

import (
	"fmt"
	"sync"
)

type ShardMemory struct {
	lock   sync.Mutex
	rwLock sync.RWMutex
	value  string
}

func initShardMemory() *ShardMemory {
	return &ShardMemory{sync.Mutex{}, sync.RWMutex{}, "init value"}
}

func modifyValue() {
	shardMemory := initShardMemory()
	shardMemory.lock.Lock()
	shardMemory.value = "change value"
	shardMemory.lock.Unlock()
	fmt.Println(shardMemory.value)
}

func main() {
	modifyValue()
}
