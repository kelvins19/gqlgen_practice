package entity

import "time"

type BaseModel struct {
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Products struct {
	ID          int     `bun:"id,pk,autoincrement"`
	Name        string  `bun:"name"`
	Description *string `bun:"description"`
	Price       int     `bun:"price"`
	Categories  []int   `bun:"categories,array"`
	BaseModel
}

type Categories struct {
	ID          int     `bun:"id,pk,autoincrement"`
	Name        string  `bun:"name"`
	Description *string `bun:"description"`
	BaseModel

	Products []*Products `bun:"-"`
}
