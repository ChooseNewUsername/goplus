package String

import (
	"fmt"
	"strconv"
)

type String string

func (this String)Len() int  {
	return len(this)
}

func (this String)Each(f func(item string))  {
	for _,c:=range this{

		f(fmt.Sprintf("%c",c));
	}
}

func From(str string) String  {
	return String(str)
}
func FromInt(n int) String  {
	return String(strconv.Itoa(n))
}