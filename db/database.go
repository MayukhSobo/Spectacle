package db

type Database interface {
	Connect()
	Ping()
	Execute()
	Close()
}
