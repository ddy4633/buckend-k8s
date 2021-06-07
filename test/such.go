package main

import "fmt"

type TableName interface {
	TableName() string
}
type TestRoleLogin struct {
	Id     string `model:"pk"`
}

func ( TestRoleLogin) TableName() string {
	return "test_role_login"
}
func main()  {
	originalModel:=TestRoleLogin{}
	var a interface{} =originalModel
	if _,ok:=a.(TableName);ok{
		fmt.Println("ok")
	}else {
		fmt.Println("!ok")
	}
}
