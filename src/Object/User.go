package Object

type User struct {
	Id,Sex int
	Name string
}

func NewUser(fs ...func(u *User)) *User {
	u := new(User)
	for _,f:=range fs {
		f(u)
	}
	return  u
}
func WithId(id int)func(u *User)  {
	return func(u *User) {
		u.Id = id
	}
}
func WithName(name string)func(u *User)  {
	return func(u *User) {
		u.Name = name
	}
}