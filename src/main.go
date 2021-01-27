package main

import (
	"goplus/src/Interface"
)

func SaveModel(service Interface.IService) Interface.IService {
	service.Save()
	return service
}
func main()  {
	SaveModel(Interface.NewProdService()).Save().Save()
}
