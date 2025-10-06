package option

// 选项模式

type User struct {
	name string
	age  uint8
}

type Option func(*User)

func NewUser(option ...Option) *User {
	user := &User{
		name: "",
		age:  100,
	}

	for _, op := range option {
		op(user)
	}
	return user
}

func WithName(name string) Option {
	return func(u *User) {
		u.name = name
	}
}
func WithAge(age uint8) Option {
	return func(u *User) {
		u.age = age
	}
}
