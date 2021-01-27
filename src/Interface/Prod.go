package Interface

import "fmt"

type Prod struct {
	Id int
	Name string
}
func NewProdService() *Prod  {
	return new(Prod)
}
func (this *Prod )Save() IService  {
	fmt.Println("商品入库")
	return  this
}