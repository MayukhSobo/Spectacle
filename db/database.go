package db

type Database interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
	Close() error
}
