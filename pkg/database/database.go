package database

type Database interface {
	DbInsert(d interface{}) (string, error)
}
