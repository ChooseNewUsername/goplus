package Object

type UserAttrFunc func(user *User)
type UserAttrFuncs []UserAttrFunc

func (this UserAttrFuncs)apply(u *User)   {
	for _,f:=range this {
		f(u)
	}
}


func WithId(id int) UserAttrFunc {
	return func(u *User) {
		u.Id = id
	}
}
func WithName(name string) UserAttrFunc  {
	return func(u *User) {
		u.Name = name
	}
}