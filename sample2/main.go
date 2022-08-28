package main

import (
	"fmt"

	"go.uber.org/dig"
)

// sample1をdigを使って書き換えた例
// UseCase -> Repositoryという関係でUseCaseはRepositoryに依存している状態

type Usecase interface {
	Use()
}

type Repository interface {
	RepoPrint()
}

type usecase struct {
	repo Repository
}

func NewUsecase(r Repository) Usecase {
	return &usecase{repo: r}
}

func (u usecase) Use() {
	u.repo.RepoPrint()
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) RepoPrint() {
	fmt.Println("2")
}

type repository2 struct{}

func NewRepository2() Repository {
	return &repository2{}
}

func (r repository2) RepoPrint() {
	fmt.Println("3")
}

func main() {
	c := dig.New() // DIコンテナを作成
	// オブジェクトの登録
	c.Provide(NewUsecase)     // Usecaseを返却している。NewUsecaseは引数にRepository型を要求しているのでNewRepositoryで返却されるRepositoryがinjectされる
	c.Provide(NewRepository2) // Repositoryを返却している。ここをrepositoryにすれば、2が表示、ここをrepositoryにすれば、3が表示される
	// オブジェクトの構築・実行
	c.Invoke(func(u Usecase) {
		u.Use()
	})
}
