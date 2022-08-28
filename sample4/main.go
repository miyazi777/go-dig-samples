package main

import (
	"fmt"

	"go.uber.org/dig"
)

// digInを使った場合の実装
type Usecase interface {
	use()
}

type FooRepository interface {
	foo()
}

type BarRepository interface {
	bar()
}

// dig.Inを定義してパラメータオブジェクト
type inUsecase struct {
	dig.In
	Foo FooRepository
	Bar BarRepository
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
	c.Provide(func(u inUsecase) Usecase {
		return &usecase{
			Foo: u.Foo,
			Bar: u.Bar,
		}
	})
	c.Invoke((func(u Usecase) {
		u.use()
	}))
}
