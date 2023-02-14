package model

type Products struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Description  *string       `json:"description"`
	CategoriesId []int         `json:"categories_id"`
	Categories   []*Categories `json:"categories"`
	Price        int           `json:"price"`
}
