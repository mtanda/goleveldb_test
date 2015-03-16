package main

import "github.com/syndtr/goleveldb/leveldb"

func main() {
	db, _ := leveldb.OpenFile("./test_db", nil)
	db.Put([]byte("key"), []byte("value"), nil)
	defer db.Close()
}
