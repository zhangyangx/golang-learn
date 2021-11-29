package model

// 结构体名称小写开头，其他包无法直接访问
type student struct {
	Id   int
	Name string
	// 字段小写，其他包无法直接获取
	age int
}

// NewInstance 结构体名称小写开头，其他包需要此函数才能获取student结构体实例
func NewInstance(id int, name string, age int) *student {
	return &student{Id: id, Name: name, age: age}
}

// GetAge age字段小写开头，其他包需要此方法获取age的值
func (stu student) GetAge() int {
	return stu.age
}
