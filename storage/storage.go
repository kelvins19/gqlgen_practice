package storage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/kelvins19/BCX_BE/entity"
	"github.com/kelvins19/BCX_BE/graph/model"
	"github.com/kelvins19/BCX_BE/helper"
	"github.com/uptrace/bun"
)

type StorageInterface interface {
	GetCategory(c context.Context, ids []int) ([]*model.Categories, error)
	GetProduct(c context.Context, id int) ([]*model.Products, error)
}

// import graph gophers with your other imports

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// DataReader reads data from a database
type CategoryReader struct {
	conn *bun.DB
}

type ProductReader struct {
	conn *bun.DB
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	CategoryLoader *dataloader.Loader[int, *model.Categories]
	ProductLoader  *dataloader.Loader[int, []*model.Products]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(conn *bun.DB) *Loaders {
	// define the data loader
	categoryReader := &CategoryReader{conn: conn}
	productReader := &ProductReader{conn: conn}
	loaders := &Loaders{
		CategoryLoader: dataloader.NewBatchedLoader(categoryReader.GetCategories),
		ProductLoader:  dataloader.NewBatchedLoader(productReader.GetProducts),
	}
	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUsers implements a batch function that can retrieve many users by ID,
// for use in a dataloader
func (r *CategoryReader) GetCategories(ctx context.Context, keys []int) []*dataloader.Result[*model.Categories] {
	// read all requested users in a single query
	keySlice := helper.SliceToSql(keys, "(")

	list := []entity.Categories{}

	err := r.conn.NewSelect().Model(&list).Where(fmt.Sprintf("id in %v", keySlice)).Order("id asc").Scan(ctx)
	if err != nil {
		panic(err)
	}

	categories := map[int]*model.Categories{}
	for _, v := range list {
		if categories[v.ID] == nil {
			categories[v.ID] = &model.Categories{
				ID:          v.ID,
				Name:        v.Name,
				Description: v.Description,
			}
		}
	}

	output := make([]*dataloader.Result[*model.Categories], len(keys))
	for index, categoryKeys := range keys {
		productList, ok := categories[categoryKeys]
		if ok {
			output[index] = &dataloader.Result[*model.Categories]{Data: productList, Error: nil}
		} else {
			placeholder := &model.Categories{}
			output[index] = &dataloader.Result[*model.Categories]{Data: placeholder, Error: nil}
		}
	}
	return output
}

func (r *ProductReader) GetProducts(c context.Context, keys []int) []*dataloader.Result[[]*model.Products] {
	keySlice := helper.SliceToString(keys)

	list := []entity.Products{}

	err := r.conn.NewSelect().Model(&list).Where(fmt.Sprintf("categories && array[%v]", keySlice)).Order("id asc").Scan(c)
	if err != nil {
		panic(err)
	}

	products := map[int][]*model.Products{}
	for _, v := range list {
		for _, cates := range v.Categories {
			products[cates] = append(products[cates], &model.Products{
				ID:           v.ID,
				Name:         v.Name,
				Description:  v.Description,
				CategoriesId: v.Categories,
				Price:        v.Price,
			})
		}
	}

	output := make([]*dataloader.Result[[]*model.Products], len(keys))
	for index, productKeys := range keys {
		productList, ok := products[productKeys]
		if ok {
			output[index] = &dataloader.Result[[]*model.Products]{Data: productList, Error: nil}
		} else {
			placeholder := []*model.Products{}
			output[index] = &dataloader.Result[[]*model.Products]{Data: placeholder, Error: nil}
		}
	}
	return output
}

func (r *CategoryReader) GetCategory(c context.Context, ids []int) ([]*model.Categories, error) {
	loaders := For(c)

	thunk := loaders.CategoryLoader.LoadMany(c, ids)
	raw, _ := thunk()
	return raw, nil
}

func (r *ProductReader) GetProduct(c context.Context, id int) ([]*model.Products, error) {
	loaders := For(c)
	thunk := loaders.ProductLoader.Load(c, id)
	raw, err := thunk()
	if err != nil {
		return nil, err
	}
	return raw, nil
}
