package database

type DatabaseService[T any] interface {
	GetInstance() T
	Close()
}
