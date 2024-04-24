package database

type Database interface {
	DbInsert(b []byte) string
}
