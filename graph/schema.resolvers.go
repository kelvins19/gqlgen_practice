package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/kelvins19/BCX_BE/entity"
	"github.com/kelvins19/BCX_BE/graph/model"
	"github.com/kelvins19/BCX_BE/helper"
)

// Products is the resolver for the products field.
func (r *categoriesResolver) Products(ctx context.Context, obj *model.Categories) ([]*model.Products, error) {
	data := []*entity.Products{}

	whereQuery := fmt.Sprintf("%d = any (categories)", obj.ID)
	err := r.DB.NewSelect().Model(&data).Where(whereQuery).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, err
	}

	lists := []*model.Products{}
	for _, v := range data {
		lists = append(lists, &model.Products{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		})
	}
	return lists, nil
	// return r.StorageService.GetProduct(ctx, obj.ID)
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Products, error) {
	product := entity.Products{
		Name:        input.Name,
		Description: &input.Description,
		Categories:  input.Categories,
		Price:       input.Price,
	}

	_, err := r.DB.NewInsert().Model(&product).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("error inserting new product: %v", err)
	}

	newModel := model.Products{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		CategoriesId: product.Categories,
		Price:        product.Price,
	}

	return &newModel, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id int, input model.NewProduct) (*model.Products, error) {
	// panic(fmt.Errorf("not implemented: UpdateProduct - updateProduct"))
	data := []*entity.Products{}
	whereValue := helper.SliceToSql([]int{id}, "(")
	err := r.DB.NewSelect().Model(&data).Where(fmt.Sprintf("id in %s", whereValue)).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error finding products: %v", err)
	}

	data[0].Name = input.Name
	data[0].Description = &input.Description
	data[0].Price = input.Price
	data[0].Categories = input.Categories

	_, err = r.DB.NewUpdate().
		Model(&data).
		SetColumn("name", "?", input.Name).
		SetColumn("description", "?", &input.Description).
		SetColumn("categories", "?", helper.SliceToSql(input.Categories, "{")).
		SetColumn("price", "?", fmt.Sprintf("%d", input.Price)).
		SetColumn("updated_at", "?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	newModel := model.Products{
		ID:           id,
		Name:         data[0].Name,
		Description:  data[0].Description,
		Price:        data[0].Price,
		CategoriesId: data[0].Categories,
	}

	return &newModel, nil
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id int) (*model.Products, error) {
	data := []*entity.Products{}
	whereValue := helper.SliceToSql([]int{id}, "(")
	err := r.DB.NewSelect().Model(&data).Where(fmt.Sprintf("id in %s", whereValue)).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error finding products: %v", err)
	}

	ProductsModel := entity.Products{}
	_, err = r.DB.NewDelete().
		Model(&ProductsModel).
		Where("id = ?", id).
		Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("error deleting products: %v", err)
	}

	newModel := model.Products{
		ID:           data[0].ID,
		Name:         data[0].Name,
		Description:  data[0].Description,
		Price:        data[0].Price,
		CategoriesId: data[0].Categories,
	}
	return &newModel, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Categories, error) {
	category := entity.Categories{
		Name:        input.Name,
		Description: &input.Description,
	}

	_, err := r.DB.NewInsert().Model(&category).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("error inserting new category: %v", err)
	}

	newModel := model.Categories{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &newModel, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id int, input model.NewCategory) (*model.Categories, error) {
	data := []*entity.Categories{}
	whereValue := helper.SliceToSql([]int{id}, "(")
	err := r.DB.NewSelect().Model(&data).Where(fmt.Sprintf("id in %s", whereValue)).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error finding categories: %v", err)
	}

	data[0].Name = input.Name
	data[0].Description = &input.Description

	_, err = r.DB.NewUpdate().
		Model(&data).
		SetColumn("name", "?", input.Name).
		SetColumn("description", "?", &input.Description).
		SetColumn("updated_at", "?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	newModel := model.Categories{
		ID:          id,
		Name:        data[0].Name,
		Description: data[0].Description,
	}

	return &newModel, nil
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id int) (*model.Categories, error) {
	data := []*entity.Categories{}
	whereValue := helper.SliceToSql([]int{id}, "(")
	err := r.DB.NewSelect().Model(&data).Where(fmt.Sprintf("id in %s", whereValue)).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error finding categories: %v", err)
	}

	CategoriesModel := entity.Categories{}
	_, err = r.DB.NewDelete().
		Model(&CategoriesModel).
		Where("id = ?", id).
		Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("error deleting categories: %v", err)
	}

	newModel := model.Categories{
		ID:          data[0].ID,
		Name:        data[0].Name,
		Description: data[0].Description,
	}
	return &newModel, nil
}

// Categories is the resolver for the categories field.
func (r *productsResolver) Categories(ctx context.Context, obj *model.Products) ([]*model.Categories, error) {
	data := []*entity.Categories{}

	whereValue := helper.SliceToSql(obj.CategoriesId, "(")
	err := r.DB.NewSelect().Model(&data).Where(fmt.Sprintf("id in %s", whereValue)).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, err
	}

	lists := []*model.Categories{}
	for _, v := range data {
		lists = append(lists, &model.Categories{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	return lists, nil
	// return r.StorageService.GetCategory(ctx, obj.CategoriesId)
}

// GetCategories is the resolver for the getCategories field.
func (r *queryResolver) GetCategories(ctx context.Context) ([]*model.Categories, error) {
	data := []*entity.Categories{}
	err := r.DB.NewSelect().Model(&data).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, err
	}

	lists := []*model.Categories{}
	for _, v := range data {
		lists = append(lists, &model.Categories{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	return lists, nil
}

// GetSingleCategories is the resolver for the getSingleCategories field.
func (r *queryResolver) GetSingleCategories(ctx context.Context, id int) (*model.Categories, error) {
	data := []*entity.Categories{}
	whereValue := helper.SliceToSql([]int{id}, "(")
	err := r.DB.NewSelect().Model(&data).Where(fmt.Sprintf("id in %s", whereValue)).Order("id asc").Scan(ctx)

	if err != nil {
		return nil, err
	}

	list := []*model.Categories{}
	for _, v := range data {
		list = append(list, &model.Categories{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		})
	}

	return list[0], nil
}

// GetProducts is the resolver for the getProducts field.
func (r *queryResolver) GetProducts(ctx context.Context, categoryID *int, name *string) ([]*model.Products, error) {
	if categoryID == nil {
		data := []*entity.Products{}
		err := r.DB.NewSelect().Model(&data).Order("id asc").Scan(ctx)
		if err != nil {
			return nil, err
		}

		list := []*model.Products{}
		for _, v := range data {
			list = append(list, &model.Products{
				ID:           v.ID,
				Name:         v.Name,
				Description:  v.Description,
				CategoriesId: v.Categories,
				Price:        v.Price,
			})
		}
		return list, nil
	}

	// data, err := serv.ProductsRepo.GetByCategory(c, *category, *name)
	regex, err := regexp.Compile("[a-zA-Z0-9]")
	if err != nil {
		return nil, err
	}

	selecteds := regex.FindAllString(*name, -1)
	cleanStr := strings.Join(selecteds, "")
	whereQuery := fmt.Sprintf("%d = any (categories)", categoryID)
	if len(*name) > 0 {
		whereQuery += fmt.Sprintf(" and name ilike '%%%s%%'", cleanStr)
	}

	data := []*entity.Products{}
	err = r.DB.NewSelect().Model(&data).Where(whereQuery).Order("id asc").Scan(ctx)
	if err != nil {
		return nil, err
	}

	list := []*model.Products{}
	for _, v := range data {
		list = append(list, &model.Products{
			ID:           v.ID,
			Name:         v.Name,
			Description:  v.Description,
			CategoriesId: v.Categories,
			Price:        v.Price,
		})
	}

	return list, nil
}

// GetSingleProducts is the resolver for the getSingleProducts field.
func (r *queryResolver) GetSingleProducts(ctx context.Context, id int) (*model.Products, error) {
	// panic(fmt.Errorf("not implemented: GetSingleProducts - getSingleProducts"))
	data := []*entity.Products{}
	whereValue := helper.SliceToSql([]int{id}, "(")
	err := r.DB.NewSelect().Model(&data).Where(fmt.Sprintf("id in %s", whereValue)).Order("id asc").Scan(ctx)

	if err != nil {
		return nil, err
	}

	list := []*model.Products{}
	for _, v := range data {
		list = append(list, &model.Products{
			ID:           v.ID,
			Name:         v.Name,
			Description:  v.Description,
			CategoriesId: v.Categories,
			Price:        v.Price,
		})
	}

	return list[0], nil
}

// Categories returns CategoriesResolver implementation.
func (r *Resolver) Categories() CategoriesResolver { return &categoriesResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Products returns ProductsResolver implementation.
func (r *Resolver) Products() ProductsResolver { return &productsResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoriesResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type productsResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
