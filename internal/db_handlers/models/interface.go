package models

type Method interface {
	Create() error
	Update() error
	Delete() error
	Select() error
}
