package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"flag"
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
	var verbose = flag.Bool("v", false, "verbose flag")
	var nkey = flag.Int("k", 100, "number of keys")
	var count = flag.Int("c", 10000, "put count")
	flag.Parse()

	rand.Seed(0)
	keys := make([]string, 0, *nkey)
	for i := 0; i < *nkey; i++ {
		keys = append(keys, randString())
	}

	db, _ := leveldb.OpenFile("./test_db", nil)
	defer db.Close()
	valueByte := make([]byte, 8)
	for i := 0; i < *count; i++ {
		for _, key := range keys {
			num := rand.Int63()
			value := uint64(num)
			binary.LittleEndian.PutUint64(valueByte, value)
			db.Put([]byte(key), valueByte, nil)
			if *verbose {
				fmt.Printf("PUT %s %d\n", key, value)
			}
		}
		if *verbose {
			fmt.Printf("%d / %d\n", i, *count)
		}
	}
}
