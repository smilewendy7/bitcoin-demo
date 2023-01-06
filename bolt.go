package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	//1. 打开数据库, 如果没有数据库， Open 会创建一个数据库

	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Panic("打开数据失败")
	}

	// 将要操作数据库 （改写）
	//2. 找到抽屉 bucket (检查是否有抽屉，如果没有就创建)
	db.Update(func(tx *bolt.Tx) error {
		// 函数接受参数是- 一个*函数指针*
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			// 没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket失败")
			}

		}
		// 写
		bucket.Put([]byte("11111"), []byte("hello"))
		bucket.Put([]byte("22222"), []byte("world"))

		return nil
		// test.db
		// $ strings test.db - does not work

	})

	//3. 写数据
	//4. 读数据

}
