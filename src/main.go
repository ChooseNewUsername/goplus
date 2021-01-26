package main

import (
	"fmt"
	"goplus/src/Object"
)


func main()  {
	u:=Object.NewUser(Object.WithId(111),Object.WithName("dd"))

	fmt.Printf("%+v",u)
}
