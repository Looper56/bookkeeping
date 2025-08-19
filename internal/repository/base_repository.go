package repository

import "errors"

// basic error
var (
	ErrorSys        = errors.New("system error")
	ErrorNotfound   = errors.New("record not found")
	ErrorNotAllowed = errors.New("not allowed")
	ErrorTimeout    = errors.New("timed out")
)

// BaseRepository repository basic
type BaseRepository struct {
}
