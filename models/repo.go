package models

type Repo interface {
	GetOne(int) interface{}
	GetAll() []interface{}
}
