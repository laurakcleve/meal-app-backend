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

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT id, name, default_shelflife AS "defaultShelflife", item_type AS "itemType"
      FROM item 
      ORDER BY name
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*model.Item

	for rows.Next() {
		var item model.Item
		var itemID int

		err := rows.Scan(&itemID, &item.Name, &item.DefaultShelflife, &item.ItemType)
		if err != nil {
			return nil, err
		}

		item.ID = strconv.Itoa(itemID)
		items = append(items, &item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return items, nil
}

func (r *queryResolver) ItemByID(ctx context.Context, id string) (*model.Item, error) {
	item := model.Item{
		ID: id,
	}
	idNum, _ := strconv.Atoi(id)

	err := db.Conn.QueryRow(context.Background(), `
			SELECT name, default_shelflife AS "defaultShelflife", item_type AS "itemType" 
			FROM item 
			WHERE id = $1
		`, idNum).Scan(&item.Name, &item.DefaultShelflife, &item.ItemType)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("ITEM:\n", item)

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

func (r *itemResolver) Category(ctx context.Context, obj *model.Item) (*model.ItemCategory, error) {
	var category model.ItemCategory
	var id int
	err := db.Conn.QueryRow(context.Background(), `
      SELECT item_category.*
      FROM item_category
      INNER JOIN item ON item.category_id = item_category.id
      WHERE item.id = $1
		`, obj.ID).Scan(&id, &category.Name)

	category.ID = strconv.Itoa(id)
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &category, nil
}

func (r *itemResolver) Dishes(ctx context.Context, obj *model.Item) ([]*model.Dish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *itemResolver) DefaultLocation(ctx context.Context, obj *model.Item) (*model.ItemLocation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *itemResolver) Purchases(ctx context.Context, obj *model.Item) ([]*model.PurchaseItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *itemResolver) CountsAs(ctx context.Context, obj *model.Item) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *dishResolver) Tags(ctx context.Context, obj *model.Dish) ([]*model.DishTag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *dishResolver) Dates(ctx context.Context, obj *model.Dish) ([]*model.DishDate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *dishResolver) IngredientSets(ctx context.Context, obj *model.Dish) ([]*model.IngredientSet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *ingredientResolver) Item(ctx context.Context, obj *model.Ingredient) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *ingredientSetResolver) Ingredients(ctx context.Context, obj *model.IngredientSet) ([]*model.Ingredient, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *inventoryItemResolver) Item(ctx context.Context, obj *model.InventoryItem) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *inventoryItemResolver) Location(ctx context.Context, obj *model.InventoryItem) (*model.ItemLocation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *purchaseResolver) Location(ctx context.Context, obj *model.Purchase) (*model.PurchaseLocation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *purchaseResolver) Items(ctx context.Context, obj *model.Purchase) ([]*model.PurchaseItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *purchaseItemResolver) Item(ctx context.Context, obj *model.PurchaseItem) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *purchaseItemResolver) Purchase(ctx context.Context, obj *model.PurchaseItem) (*model.Purchase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *ingredientInputResolver) Item(ctx context.Context, obj *model.IngredientInput, data *model.IngredientItemInput) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *ingredientSetInputResolver) Ingredients(ctx context.Context, obj *model.IngredientSetInput, data []*model.IngredientInput) error {
	panic(fmt.Errorf("not implemented"))
}

// Dish returns generated.DishResolver implementation.
func (r *Resolver) Dish() generated.DishResolver { return &dishResolver{r} }

// Ingredient returns generated.IngredientResolver implementation.
func (r *Resolver) Ingredient() generated.IngredientResolver { return &ingredientResolver{r} }

// IngredientSet returns generated.IngredientSetResolver implementation.
func (r *Resolver) IngredientSet() generated.IngredientSetResolver { return &ingredientSetResolver{r} }

// InventoryItem returns generated.InventoryItemResolver implementation.
func (r *Resolver) InventoryItem() generated.InventoryItemResolver { return &inventoryItemResolver{r} }

// Item returns generated.ItemResolver implementation.
func (r *Resolver) Item() generated.ItemResolver { return &itemResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Purchase returns generated.PurchaseResolver implementation.
func (r *Resolver) Purchase() generated.PurchaseResolver { return &purchaseResolver{r} }

// PurchaseItem returns generated.PurchaseItemResolver implementation.
func (r *Resolver) PurchaseItem() generated.PurchaseItemResolver { return &purchaseItemResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// IngredientInput returns generated.IngredientInputResolver implementation.
func (r *Resolver) IngredientInput() generated.IngredientInputResolver {
	return &ingredientInputResolver{r}
}

// IngredientSetInput returns generated.IngredientSetInputResolver implementation.
func (r *Resolver) IngredientSetInput() generated.IngredientSetInputResolver {
	return &ingredientSetInputResolver{r}
}

type dishResolver struct{ *Resolver }
type ingredientResolver struct{ *Resolver }
type ingredientSetResolver struct{ *Resolver }
type inventoryItemResolver struct{ *Resolver }
type itemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type purchaseResolver struct{ *Resolver }
type purchaseItemResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type ingredientInputResolver struct{ *Resolver }
type ingredientSetInputResolver struct{ *Resolver }
