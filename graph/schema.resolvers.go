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
	panic(fmt.Errorf("ItemByName not implemented"))
}

func (r *queryResolver) Dishes(ctx context.Context) ([]*model.Dish, error) {
	panic(fmt.Errorf("Dishes not implemented"))
}

func (r *queryResolver) Dish(ctx context.Context, id string) (*model.Dish, error) {
	panic(fmt.Errorf("Dish not implemented"))
}

func (r *queryResolver) InventoryItems(ctx context.Context) ([]*model.InventoryItem, error) {
	panic(fmt.Errorf("InventoryItems not implemented"))
}

func (r *queryResolver) InventoryItem(ctx context.Context, id string) (*model.InventoryItem, error) {
	panic(fmt.Errorf("InventoryItem not implemented"))
}

func (r *queryResolver) ItemLocations(ctx context.Context) ([]*model.ItemLocation, error) {
	panic(fmt.Errorf("ItemLocations not implemented"))
}

func (r *queryResolver) ItemCategories(ctx context.Context) ([]*model.ItemCategory, error) {
	panic(fmt.Errorf("ItemCategories not implemented"))
}

func (r *queryResolver) DishTags(ctx context.Context) ([]*model.DishTag, error) {
	panic(fmt.Errorf("DishTags not implemented"))
}

func (r *queryResolver) Purchases(ctx context.Context) ([]*model.Purchase, error) {
	panic(fmt.Errorf("Purchases not implemented"))
}

func (r *queryResolver) Purchase(ctx context.Context, id string) (*model.Purchase, error) {
	panic(fmt.Errorf("Purchase not implemented"))
}

func (r *queryResolver) PurchaseLocations(ctx context.Context) ([]*model.PurchaseLocation, error) {
	panic(fmt.Errorf("PurchaseLocations not implemented"))
}

func (r *mutationResolver) DeleteItem(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("DeleteItem not implemented"))
}

func (r *mutationResolver) AddPurchase(ctx context.Context, date string, location string) (*model.Purchase, error) {
	panic(fmt.Errorf("AddPurchase not implemented"))
}

func (r *mutationResolver) DeletePurchase(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("DeletePurchase not implemented"))
}

func (r *mutationResolver) AddPurchaseItem(ctx context.Context, purchaseID string, name string, price *float64, weightAmount *float64, weightUnit *string, quantityAmount *float64, quantityUnit *string, number int, itemType string) (*model.PurchaseItem, error) {
	panic(fmt.Errorf("AddPurchaseItem not implemented"))
}

func (r *mutationResolver) UpdatePurchaseItem(ctx context.Context, id string, name string, price *float64, weightAmount *float64, weightUnit *string, quantityAmount *float64, quantityUnit *string) (*model.PurchaseItem, error) {
	panic(fmt.Errorf("UpdatePurchaseItem not implemented"))
}

func (r *mutationResolver) DeletePurchaseItem(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("DeletePurchaseItem not implemented"))
}

func (r *mutationResolver) AddInventoryItem(ctx context.Context, name string, addDate *string, expiration *string, amount *string, defaultShelflife *string, category *string, location *string, itemType string, number int) (*model.InventoryItem, error) {
	panic(fmt.Errorf("AddInventoryItem not implemented"))
}

func (r *mutationResolver) UpdateInventoryItem(ctx context.Context, id string, addDate *string, expiration *string, amount *string, location *string, category *string, itemType *string) (*model.InventoryItem, error) {
	panic(fmt.Errorf("UpdateInventoryItem not implemented"))
}

func (r *mutationResolver) DeleteInventoryItem(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("DeleteInventoryItem not implemented"))
}

func (r *mutationResolver) EditItem(ctx context.Context, id string, name string, categoryID *int, defaultLocationID *int, defaultShelflife *int, itemType string, countsAs []*string) (*model.Item, error) {
	panic(fmt.Errorf("EditItem not implemented"))
}

func (r *mutationResolver) AddDish(ctx context.Context, name string, tags []*string, isActive bool, ingredientSets []*model.IngredientSetInput) (*model.Dish, error) {
	panic(fmt.Errorf("AddDish not implemented"))
}

func (r *mutationResolver) UpdateDish(ctx context.Context, id string, name string, tags []*string, isActive bool, ingredientSets []*model.IngredientSetInput) (*model.Dish, error) {
	panic(fmt.Errorf("UpdateDish not implemented"))
}

func (r *mutationResolver) DeleteDish(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("DeleteDish not implemented"))
}

func (r *mutationResolver) AddDishDate(ctx context.Context, dishID string, date string) (*model.DishDate, error) {
	panic(fmt.Errorf("AddDishDate not implemented"))
}

func (r *mutationResolver) DeleteDishDate(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("DeleteDishDate not implemented"))
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
	rows, err := db.Conn.Query(context.Background(), `
      WITH generic_items AS (
        SELECT gi.id AS itemID  
        FROM item gi
        INNER JOIN item_counts_as ica on ica.generic_item_id = gi.id
        WHERE ica.specific_item_id = $1
      )
      SELECT DISTINCT dish.id, dish.name, dish.is_active_dish 
      FROM item dish
      INNER JOIN ingredient_set ings ON ings.parent_item_id = dish.id
      INNER JOIN ingredient ing ON ing.ingredient_set_id = ings.id
      INNER JOIN item i ON i.id = ing.item_id
      WHERE i.id IN ((SELECT UNNEST(ARRAY_APPEND(ARRAY_AGG(itemID), $1)) 
                      FROM generic_items))
		`, obj.ID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []*model.Dish

	for rows.Next() {
		var dish model.Dish
		var dishID int

		err := rows.Scan(
			&dishID, 
			dish.Name, 
			dish.IsActiveDish, 
		)
		if err != nil {
			return nil, err
		}

		dish.ID = strconv.Itoa(dishID)
		dishes = append(dishes, &dish)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return dishes, nil
}

func (r *itemResolver) DefaultLocation(ctx context.Context, obj *model.Item) (*model.ItemLocation, error) {
	var location model.ItemLocation
	var id int
	err := db.Conn.QueryRow(context.Background(), `
      SELECT inventory_item_location.* 
      FROM inventory_item_location
      INNER JOIN item on item.default_location_id = inventory_item_location.id
      WHERE item.id = $1
		`, obj.ID).Scan(&id, &location.Name)

	location.ID = strconv.Itoa(id)
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &location, nil
}

func (r *itemResolver) Purchases(ctx context.Context, obj *model.Item) ([]*model.PurchaseItem, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT  
				id,
				price,
        weight_amount, 
        weight_unit,
        quantity_amount,
        quantity_unit,
        purchase_id
      FROM purchase_item
      WHERE item_id = $1
		`, obj.ID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []*model.PurchaseItem

	for rows.Next() {
		var purchase model.PurchaseItem
		var purchaseID int

		err := rows.Scan(
			&purchaseID, 
			purchase.Price, 
			purchase.WeightAmount, 
			purchase.WeightUnit,
			purchase.QuantityAmount,
			purchase.QuantityUnit,
			purchase.PurchaseID,
		)
		if err != nil {
			return nil, err
		}

		purchase.ID = strconv.Itoa(purchaseID)
		purchases = append(purchases, &purchase)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return purchases, nil
}

func (r *itemResolver) CountsAs(ctx context.Context, obj *model.Item) ([]*model.Item, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT generic.id, name, default_shelflife, item_type
      FROM item generic
      JOIN item_counts_as ica ON ica.generic_item_id = generic.id
      WHERE ica.specific_item_id = $1
		`, obj.ID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countsAsItems []*model.Item

	for rows.Next() {
		var item model.Item
		var itemID int

		err := rows.Scan(
			&itemID, 
			item.Name, 
			item.DefaultShelflife, 
			item.ItemType,
		)
		if err != nil {
			return nil, err
		}

		item.ID = strconv.Itoa(itemID)
		countsAsItems = append(countsAsItems, &item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return countsAsItems, nil
}

func (r *dishResolver) Tags(ctx context.Context, obj *model.Dish) ([]*model.DishTag, error) {
	panic(fmt.Errorf("Dish Tags not implemented"))
}

func (r *dishResolver) Dates(ctx context.Context, obj *model.Dish) ([]*model.DishDate, error) {
	panic(fmt.Errorf("Dish Dates not implemented"))
}

func (r *dishResolver) IngredientSets(ctx context.Context, obj *model.Dish) ([]*model.IngredientSet, error) {
	panic(fmt.Errorf("Dish IngredientSets not implemented"))
}

func (r *ingredientResolver) Item(ctx context.Context, obj *model.Ingredient) (*model.Item, error) {
	panic(fmt.Errorf("Ingredient Item not implemented"))
}

func (r *ingredientSetResolver) Ingredients(ctx context.Context, obj *model.IngredientSet) ([]*model.Ingredient, error) {
	panic(fmt.Errorf("IngredientSet Ingredients not implemented"))
}

func (r *inventoryItemResolver) Item(ctx context.Context, obj *model.InventoryItem) (*model.Item, error) {
	panic(fmt.Errorf("InventoryItem Item not implemented"))
}

func (r *inventoryItemResolver) Location(ctx context.Context, obj *model.InventoryItem) (*model.ItemLocation, error) {
	panic(fmt.Errorf("InventoryItem Location not implemented"))
}

func (r *purchaseResolver) Location(ctx context.Context, obj *model.Purchase) (*model.PurchaseLocation, error) {
	panic(fmt.Errorf("Purchase Location not implemented"))
}

func (r *purchaseResolver) Items(ctx context.Context, obj *model.Purchase) ([]*model.PurchaseItem, error) {
	panic(fmt.Errorf("Purchase Items not implemented"))
}

func (r *purchaseItemResolver) Item(ctx context.Context, obj *model.PurchaseItem) (*model.Item, error) {
	panic(fmt.Errorf("PurchaseItem Item not implemented"))
}

func (r *purchaseItemResolver) Purchase(ctx context.Context, obj *model.PurchaseItem) (*model.Purchase, error) {
	panic(fmt.Errorf("PurchaseItem Purchase not implemented"))
}

func (r *ingredientInputResolver) Item(ctx context.Context, obj *model.IngredientInput, data *model.IngredientItemInput) error {
	panic(fmt.Errorf("IngredientInput Item not implemented"))
}

func (r *ingredientSetInputResolver) Ingredients(ctx context.Context, obj *model.IngredientSetInput, data []*model.IngredientInput) error {
	panic(fmt.Errorf("IngredientSetInput Ingredients not implemented"))
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
