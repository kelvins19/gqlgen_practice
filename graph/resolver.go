package graph

import (
	"github.com/kelvins19/BCX_BE/storage"
	"github.com/uptrace/bun"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB             *bun.DB
	StorageService storage.StorageInterface
}
