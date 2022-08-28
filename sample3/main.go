package main

import (
	"fmt"

	"go.uber.org/dig"
)

type Usecase interface {
	use()
}

type FooRepository interface {
	foo()
}

type BarRepository interface {
	bar()
}

type usecase struct {
	Foo FooRepository
	Bar BarRepository
}

func (u *usecase) use() {
	u.Foo.foo()
	u.Bar.bar()
}

type fooRepo struct {
	f string
}

func (u *fooRepo) foo() { fmt.Println(u.f) }

type barRepo struct {
	b string
}

func (u *barRepo) bar() { fmt.Println(u.b) }

func main() {
	c := dig.New()
	c.Provide(func() FooRepository { return &fooRepo{f: "Foo"} })
	c.Provide(func() BarRepository { return &barRepo{b: "Bar"} })
	c.Provide(func(f FooRepository, b BarRepository) Usecase {
		return &usecase{
			Foo: f,
			Bar: b,
		}
	})
	c.Invoke(func(u Usecase) {
		u.use()
	})
}
