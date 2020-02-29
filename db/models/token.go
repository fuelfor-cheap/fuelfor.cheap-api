package models

// Token houses the authentication token returned from the API to act like a user on their behalf
type Token struct {
	Author int64  `pg:",pk"`
	Token  string `pg:",pk"`
}
