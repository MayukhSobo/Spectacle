package db

// Database defines the interface for database operations.
// Implementations provide specific database type functionality.
type Database interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
	Close() error
}
