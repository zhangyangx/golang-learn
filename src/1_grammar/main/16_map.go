package main

import (
	"fmt"
	"sort"
)

// map
// 1.map定义
// 2.map增删改查
// 3.map遍历
// 4.map排序
func main() {

	// 定义map
	var firstMap = make(map[string]string, 10)
	// 添加数据
	firstMap["1"] = "hello"
	firstMap["2"] = "world"
	firstMap["3"] = "!"
	fmt.Println(firstMap)
	fmt.Println(firstMap["1"])

	// 定义map
	cityMap := make(map[int]string)
	cityMap[0] = "北京"
	cityMap[1] = "成都"
	fmt.Println(cityMap[1])

	// 定义map
	gameMap := map[int]string{
		1: "dnf",
		2: "csgo",
		3: "lol",
	}
	fmt.Println(gameMap[1])

	// 删除key（如果要删除所有key，需要遍历map，要么make一个新的空map给当前map变量，让之前的数据成为垃圾，被gc回收）
	delete(gameMap, 1)
	fmt.Println("获取被删除的数据：", gameMap[1])

	// 查询key
	val, ok := gameMap[1]
	if ok {
		fmt.Println("存在key为2的值：", val)
	} else {
		fmt.Println("不存在key为2的值")
	}

	// map遍历
	for key, val := range gameMap {
		fmt.Printf("key:%v,value:%v \n", key, val)
	}

	// map排序
	// 需要先将map的所有key放到一个切片，对切片进行排序，然后按照排序的key取值
	var keys []int
	for k, _ := range gameMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("k:", k, "v:", gameMap[k])
	}
}
