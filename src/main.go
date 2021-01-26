package main

import (
	"fmt"
	"goplus/src/Map"
)

func chang(user Map.User)  {
	user["cc"] = "ad"
}
func main()  {
	u:= Map.NewUser()
	u.Set("id",1).Set("cc","dd")
	fmt.Println(u)
}
