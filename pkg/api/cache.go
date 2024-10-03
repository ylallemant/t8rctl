package api

import "time"

type Cache interface {
	Path() string
	Exists() bool
	Valid() bool
	Size() int64
	Age() time.Duration
	TTL() time.Duration
	Expires() time.Time
	Read() ([]byte, error)
	Write(content []byte) error
	Purge() error
}
