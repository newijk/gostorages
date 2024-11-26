package gostorages

import (
	"context"
	"errors"
	"io"
	"time"
)

// Storage is the storage interface.
type Storage interface {
	Save(ctx context.Context, content io.Reader, path string) error
	Stat(ctx context.Context, path string) (*Stat, error)
	Open(ctx context.Context, path string) (io.ReadCloser, error)
	OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *Stat, error)
	Delete(ctx context.Context, path string) error
	URL(ctx context.Context, path string) (string, error) // return a signed URL
}

// Stat contains metadata about content stored in storage.
type Stat struct {
	ModifiedTime time.Time
	Size         int64
}

// ErrNotExist is a sentinel error returned by the Open and the Stat methods.
var ErrNotExist = errors.New("does not exist")
