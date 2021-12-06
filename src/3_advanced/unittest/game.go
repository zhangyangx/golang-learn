package unittest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Game struct {
	Name       string
	Price      float64
	Evaluation string
}

// Store 将对象序列化后保存到本地
func (game Game) Store() bool {
	data, err := json.Marshal(game)
	if err != nil {
		return false
	}
	// 保存到文件
	filePath := "../../../tmp/game.ser"
	exists, err := PathExists(filePath)
	if err != nil {
		return false
	}
	if !exists {
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return false
	}
	return true
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// 文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		// 文件不存在
		return false, nil
	}
	// 未知错误
	return false, err
}
