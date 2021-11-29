package model

type People struct {
	Name string
	age  int
}

func (people *People) SetAge(age int) {
	people.age = age
}

func (people People) GetAge() (age int) {
	return people.age
}

type Child struct {
	// 嵌套匿名结构体，则可以继承
	People
}

type Adult struct {
	// 嵌套匿名结构体，则可以继承
	People
	job string
}

func (adult *Adult) FindJob(job string) {
	adult.job = job
}
