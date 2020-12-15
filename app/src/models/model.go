package models

type Model interface {
	Create()
	Save()
	Delete()
}
