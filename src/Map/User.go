package Map

import "fmt"

type User map[string]interface{}

func NewUser()User  {
	return make(map[string]interface{})
}
func (this User)Set(k string, v interface{})User  {
	this[k] = v
	return this
}
func (this User) String() string  {
	str:=""
	for k,v:=range this{
		str+=fmt.Sprintf("%v->%v\n",k,v)
	}
	return str
}