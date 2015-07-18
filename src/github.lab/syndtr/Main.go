package main

import (
	"fmt"
	"github.lab/syndtr/goleveldb/leveldb"
	"strconv"
	"time"
)

func main() {
	db, err := leveldb.OpenFile("/usr/local/leveldbdata", nil)
	if nil != err {
		fmt.Println(err)
	}

	for i := 0; i < 10; i++ {
		nano := time.Now().Nanosecond()

		db.Put([]byte("ceshi"+strconv.Itoa(nano)), []byte("ceshi"+strconv.Itoa(nano)), nil)
		if nil != err {
			fmt.Println(err)
		}
	}

	iter := db.NewIterator(nil)

	// for iter.Next() {
	// 	key, value := iter.Key(), iter.Value()
	// 	fmt.Printf("(%s,%s)\n", key, value)
	// }

	iter.Release()

	err = iter.Error()

	if nil != err {
		fmt.Println(err)
	}

	defer db.Close()
}
