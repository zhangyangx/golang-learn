package main

import (
	"encoding/json"
	"fmt"
)

// json序列化
// 分别将结构体、map、数组等进行json序列化
// json反序列化
func main() {
	// 将结构体序列化
	monster := Monster{Name: "🐂", Age: 9000, Birthday: "2021.02.11", Sal: 0, Skill: "fk"}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误,", err)
	}
	fmt.Println("monster结构体序列化后：", string(data))

	// 将map进行序列化
	gameMap := make(map[int]string)
	gameMap[1] = "CSGO"
	gameMap[2] = "COD"
	gameMap[3] = "古墓丽影"
	data, err = json.Marshal(gameMap)
	if err != nil {
		fmt.Println("序列化错误,", err)
	}
	fmt.Println("gameMap结构体序列化后：", string(data))

	// 将切片进行序列化
	arr := []int{1, 2, 3}
	data, err = json.Marshal(arr)
	if err != nil {
		fmt.Println("序列化错误,", err)
	}
	fmt.Println("arr结构体序列化后：", string(data))

	// 将json字符串反序列化
	var bullDemon Monster
	err = json.Unmarshal([]byte("{\"Name\":\"🐂\",\"Age\":9000,\"Birthday\":\"2021.02.11\",\"Sal\":0,\"Skill\":\"fk\"}"), &bullDemon)
	if err != nil {
		fmt.Println("反序列化错误,", err)
	}
	fmt.Println(bullDemon)
}

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}
