package main

import "github.com/syndtr/goleveldb/leveldb"

func main() {
	db, _ := leveldb.OpenFile("./test_db", nil)
	for i := 0; i < 10000; i++ {
		db.Put([]byte("key"), []byte("value"), nil)
	}
	defer db.Close()
}
