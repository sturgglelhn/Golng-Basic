package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
/*type Student struct {
	ID     int
	Gender string
	Name   string
}*/

type Student struct {
	ID     int `json:"id"`
	Gender string
	name   string
}

// Class 班级

type Class struct {
	Title    string
	Students []*Student
}

func main() {

	/*
		结构体与JSON序列化
	*/
	/*c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}

	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	// JSON序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("%#v\n", c1)*/

	/*
		结构体标签（Tag）
	*/
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)
}

/*
json:{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Geer":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":""Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Namestu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}
&main.Class{Title:"101", Students:[]*main.Student{(*main.Student)(0xc000074a50), (*main.Student)(0xc000074a80), (*main.Student)(0xc000074ab0), (*main.Student)(0xc000074ae0), (*main.Student)(0xc000074b40), (*main.Student)(0xc000074b70), (*main.Student)(0xc000074ba0), (*main.Student)(0xc000074bd0), (*main.Student)(0xc000074c00), (*main.Student)(0xc000074c30)}}
*/

/*
json str:{"id":1,"Gender":"女"}
*/
