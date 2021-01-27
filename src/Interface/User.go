package Interface

import "fmt"

type User struct {
	Id int
	Name string
}

func NewUserService() *User  {
	return new(User)
}

func (this User )Save(data interface{})  {
	if user,ok:=data.(*User);ok {
		fmt.Printf(user.Name)
	}else {
		fmt.Println("参数错误")
	}

}