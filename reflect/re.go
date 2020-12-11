package main

import (
	"fmt"
	"reflect"
)

/*
反射能获取到的方法和结构体中的字段只有
大写字母开头的
大写字母开头的
大写字母开头的
大写字母开头的
否则获取不到
 */

type Config struct {
	Address  string
	Port     int
	Username string
	Password string
}

func (c Config)GetAddress(){
	fmt.Println(c.Address)
}

func (c Config)SetAddress(s string){
	fmt.Println(s)
}

func refle(a interface{}) {
	t := reflect.TypeOf(a).Elem() //传指针的情况要用Elem才能获取到具体的类型
	//通过反射遍历结构体的字段信息
	for i := 0; i < t.NumField(); i++{
		filed := t.Field(i)
		fmt.Println(filed.Name,filed.Index,filed.Type,filed.Tag.Get("ini"))
	}

	//通过名称获取该字段在结构体中是否存在
	filed, ok := t.FieldByName("Address")
	if ok{
		fmt.Println(filed)
	}

	//通过反射遍历结构体的值
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++{
		value := v.Field(i)
		fmt.Println(value)
	}

	//通过名称获取字段对应的值
	value := v.FieldByName("Address")
	fmt.Println(value)
	//通过反射可以修改原有的值
	value.SetString("20.20.20.20")

	//通过反射遍历结构体中的方法
	for i:=0; i < t.NumMethod(); i++{
		m := t.Method(i)
		fmt.Println(m.Name,m.Type,m.Index)
	}

	//通过反射遍历结构体中的方法名
	for i := 0; i < v.NumMethod(); i++{
		value := v.Method(i)
		fmt.Println(value.Type(),value.Kind())
	}
	//通过反射调用
	v.MethodByName("GetAddress").Call(nil) //无参数传nil
	v.MethodByName("SetAddress").Call([]reflect.Value{reflect.ValueOf("100.111.140.222")})
}

func main() {
	conf := Config{
		Address:  "10.10.10.10",
		Port:     1010,
		Username: "root",
		Password: "iMC123",
	}
	refle(&conf)
}
