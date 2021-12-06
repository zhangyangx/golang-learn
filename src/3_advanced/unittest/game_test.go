package unittest

import (
	"testing"
)

// 单元测试
func TestGame_Store(t *testing.T) {
	game := Game{Name: "COD6", Price: 0, Evaluation: "good"}
	store := game.Store()
	if !store {
		t.Fatalf("测试失败")
	}
	t.Logf("测试成功")
}
