// redis
package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("test", []byte("hello "))
	if err != nil {
		panic(err)
	}
	res, err := client.Get("test")
	if err != nil {
		panic(err)
	}
	// 要求返回的是字符串  
	// 函数本身返回的是数组
	// 所以要string处理一下
	// 这样才会是字符串可查看
	fmt.Println(string(res))

	//test hmset
	f := make(map[string]interface{})
	f["name"] = "zhangsan"
	f["age"] = 12
	f["sex"] = "male"

	err = client.Hmset("test_hash", f)
	if err != nil {
		panic(err)
	}
	//test zset
	// 倒序集合
	// 参数：第一个是集合的键值  第二个是字符串数组的键值   第三参数是值
	_, err = client.Zadd("test_zset", []byte("beifeng"), 100)
	if err != nil {
		panic(err)
	}
}
