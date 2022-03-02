package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord - No matching record
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials - invalid creds
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail - duplicate email
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet - struct for snippets
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User - define a new User type. Notice how the field names and types align
// with the columns in the database `users` table?
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
