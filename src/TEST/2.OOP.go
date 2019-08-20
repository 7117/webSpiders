package main

import (
	"fmt"
)

// 子类
type Person struct {
	name string
	age int
}

func (person Person) getNameAndAge()(string,int){
	return person.name,person.age;
}

// 通过结构体进行继承
// 父类
type Student struct{
	Person 
	speciality string
}

func (student Student) getSpeciality()(string){
	return student.speciality
}

func main() {
	student := new(Student)
	// 对子类的属性进行赋值
	student.name="zhangsan";
	student.age=12;
	// 也可以进行赋值 
	// student.Person.name="zhangsan";
	// student.Person.age=12;
	// 通过struct进行继承方法 
	// 属性直接在struct里面进行继承
	name,age:=student.getNameAndAge()
	fmt.Println(name,age);

}
