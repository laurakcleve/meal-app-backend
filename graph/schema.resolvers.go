package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"laurakcleve/meal/db"
	"laurakcleve/meal/graph/generated"
	"laurakcleve/meal/graph/model"
	"strconv"
)

func (r *mutationResolver) DeleteItem(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddPurchase(ctx context.Context, date string, location string) (*model.Purchase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePurchase(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddPurchaseItem(ctx context.Context, purchaseID string, name string, price *float64, weightAmount *float64, weightUnit *string, quantityAmount *float64, quantityUnit *string, number int, itemType string) (*model.PurchaseItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePurchaseItem(ctx context.Context, id string, name string, price *float64, weightAmount *float64, weightUnit *string, quantityAmount *float64, quantityUnit *string) (*model.PurchaseItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePurchaseItem(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddInventoryItem(ctx context.Context, name string, addDate *string, expiration *string, amount *string, defaultShelflife *string, category *string, location *string, itemType string, number int) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateInventoryItem(ctx context.Context, id string, addDate *string, expiration *string, amount *string, location *string, category *string, itemType *string) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteInventoryItem(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditItem(ctx context.Context, id string, name string, categoryID *int, defaultLocationID *int, defaultShelflife *int, itemType string, countsAs []*string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddDish(ctx context.Context, name string, tags []*string, isActive bool, ingredientSets []*model.IngredientSetInput) (*model.Dish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDish(ctx context.Context, id string, name string, tags []*string, isActive bool, ingredientSets []*model.IngredientSetInput) (*model.Dish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDish(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddDishDate(ctx context.Context, dishID string, date string) (*model.DishDate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDishDate(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	rows, err := db.Conn.Query(context.Background(), `SELECT id, name FROM item`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*model.Item

	for rows.Next() {
		var item model.Item
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		item.ID = strconv.Itoa(id)
		item.Name = name
		items = append(items, &item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return items, nil
}

func (r *queryResolver) ItemByID(ctx context.Context, id string) (*model.Item, error) {
	var item model.Item
	idNum, _ := strconv.Atoi(id)
	err := db.Conn.QueryRow(context.Background(), `SELECT name FROM item WHERE id = $1`, idNum).Scan(&item.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &item, nil
}

func (r *queryResolver) ItemByName(ctx context.Context, name string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dishes(ctx context.Context) ([]*model.Dish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dish(ctx context.Context, id string) (*model.Dish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) InventoryItems(ctx context.Context) ([]*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) InventoryItem(ctx context.Context, id string) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemLocations(ctx context.Context) ([]*model.ItemLocation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemCategories(ctx context.Context) ([]*model.ItemCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DishTags(ctx context.Context) ([]*model.DishTag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Purchases(ctx context.Context) ([]*model.Purchase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Purchase(ctx context.Context, id string) (*model.Purchase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PurchaseLocations(ctx context.Context) ([]*model.PurchaseLocation, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
