package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"

	"github.com/syndtr/goleveldb/leveldb"
)

func randString() string {
	num := rand.Uint32()
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, num)
	hash := md5.Sum(bs)
	return hex.EncodeToString(hash[:])
}

func main() {
	rand.Seed(0)
	keys := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		keys = append(keys, randString())
	}

	db, _ := leveldb.OpenFile("./test_db", nil)
	defer db.Close()
	bs := make([]byte, 8)
	n := 10000
	for i := 0; i < n; i++ {
		for _, key := range keys {
			num := rand.Int63()
			binary.LittleEndian.PutUint64(bs, uint64(num))
			db.Put([]byte(key), bs, nil)
			//fmt.Printf("%s %d\n", key, uint64(num))
		}
		fmt.Printf("%d / %d\n", i, n)
	}
}
