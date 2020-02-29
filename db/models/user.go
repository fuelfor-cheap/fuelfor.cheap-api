package models

// User has an id and associated email
type User struct {
	ID    int64 `pg:",pk"`
	Email string
}
