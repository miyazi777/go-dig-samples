package main

import "fmt"

// DIコンテナを使わない例

// UseCase.Use()を実行
// Repository.RepoPrint()を実行
//
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
	fmt.Println("1")
}

func main() {
	repo := NewRepository()
	usecase := NewUsecase(repo)
	usecase.Use()
}
