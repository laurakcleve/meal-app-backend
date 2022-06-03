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

func (r *dishResolver) Tags(ctx context.Context, obj *model.Dish) ([]*model.DishTag, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT dish_tag.id, name
			FROM dish_tag
      INNER JOIN item_has_dish_tag ihdt ON ihdt.dish_tag_id = dish_tag.id
      WHERE ihdt.item_id = $1
		`, obj.ID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []*model.DishTag{}

	for rows.Next() {
		var tag model.DishTag
		var tagID int

		err := rows.Scan(&tagID, tag.Name)
		if err != nil {
			return nil, err
		}

		tag.ID = strconv.Itoa(tagID)
		tags = append(tags, &tag)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return tags, nil
}

func (r *dishResolver) Dates(ctx context.Context, obj *model.Dish) ([]*model.DishDate, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT id, date
			FROM dish_date 
      WHERE dish_id = $1
      ORDER BY date DESC
		`, obj.ID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dates := []*model.DishDate{}

	for rows.Next() {
		var date model.DishDate
		var dateID int

		err := rows.Scan(&dateID, date.Date)
		if err != nil {
			return nil, err
		}

		date.ID = strconv.Itoa(dateID)
		dates = append(dates, &date)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return dates, nil
}

func (r *dishResolver) IngredientSets(ctx context.Context, obj *model.Dish) ([]*model.IngredientSet, error) {
	rows, err := db.Conn.Query(context.Background(), `
		SELECT id, optional 
		FROM ingredient_set 
		WHERE parent_item_id = $1
	`, obj.ID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredientSets := []*model.IngredientSet{}

	for rows.Next() {
		var set model.IngredientSet
		var setID int

		err := rows.Scan(&setID, set.IsOptional)
		if err != nil {
			return nil, err
		}

		set.ID = strconv.Itoa(setID)
		ingredientSets = append(ingredientSets, &set)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return ingredientSets, nil
}

func (r *ingredientResolver) Item(ctx context.Context, obj *model.Ingredient) (*model.Item, error) {
	item := model.Item{}
	var tempID int
	
	err := db.Conn.QueryRow(context.Background(), `
		SELECT item.id, name, default_shelflife, item_type FROM item 
		INNER JOIN ingredient ON ingredient.item_id = item.id
		WHERE ingredient.id = $1
	`, obj.ID).Scan(
		tempID,
		item.Name,
		item.DefaultShelflife,
		item.ItemType,
	)

	if err != nil {
		return nil, err
	}

	item.ID = strconv.Itoa(tempID)

	return &item, nil
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

func (r *itemResolver) Category(ctx context.Context, obj *model.Item) (*model.ItemCategory, error) {
	category := model.ItemCategory{}
	var id int

	err := db.Conn.QueryRow(context.Background(), `
      SELECT item_category.*
      FROM item_category
      INNER JOIN item ON item.category_id = item_category.id
      WHERE item.id = $1
		`, obj.ID).Scan(&id, &category.Name)

	category.ID = strconv.Itoa(id)

	if err != nil { // True if no rows are returned
		fmt.Println(err)
		// Returning err prevents the data being returned
		// Category can be null, getting no rows back shouldn't cause error
		return nil, nil
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

	dishes := []*model.Dish{}

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

	if err != nil { // True if no rows are returned
		fmt.Println(err)
		// Returning err prevents the data being returned
		// Category can be null, getting no rows back shouldn't cause error
		return nil, nil
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

	purchases := []*model.PurchaseItem{}

	for rows.Next() {
		var purchase model.PurchaseItem
		var purchaseID int

		err := rows.Scan(
			&purchaseID,
			&purchase.Price,
			&purchase.WeightAmount,
			&purchase.WeightUnit,
			&purchase.QuantityAmount,
			&purchase.QuantityUnit,
			&purchase.PurchaseID,
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

	countsAsItems := []*model.Item{}

	for rows.Next() {
		var item model.Item
		var itemID int

		err := rows.Scan(
			&itemID,
			&item.Name,
			&item.DefaultShelflife,
			&item.ItemType,
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

func (r *mutationResolver) DeleteItem(ctx context.Context, id string) (*int, error) {
	result, err := db.Conn.Exec(context.Background(), `
      DELETE FROM item
      WHERE id = $1
	`, id)

	rows := new(int)
	*rows = int(result.RowsAffected())

	return rows, err
}

func (r *mutationResolver) AddPurchase(ctx context.Context, date string, location string) (*model.Purchase, error) {
	purchase := model.Purchase{}
	var purchaseID int

	err := db.Conn.QueryRow(context.Background(), `
      WITH retrieved_purchase_location_id AS (
        SELECT purchase_location_id_for_insert($2)
      )
      INSERT INTO purchase(date, location_id)
      SELECT 
        $1 AS date, 
        (SELECT * FROM retrieved_purchase_location_id) AS location_id
      RETURNING id, CAST(date AS TEXT)
	`, date, location).Scan(&purchaseID, &purchase.Date)

	purchase.ID = strconv.Itoa(purchaseID)

	return &purchase, err
}

func (r *mutationResolver) DeletePurchase(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("DeletePurchase not implemented"))
}

func (r *mutationResolver) AddPurchaseItem(ctx context.Context, purchaseID string, name string, price *float64, weightAmount *float64, weightUnit *string, quantityAmount *float64, quantityUnit *string, number int, itemType string) (*model.PurchaseItem, error) {
	purchaseItem := model.PurchaseItem{}
	var tempID int

	err := db.Conn.QueryRow(context.Background(), `
      WITH retrieved_item_id AS (
        SELECT item_id_for_insert($2, CAST($9 AS itemtype)) 
      )
      INSERT INTO purchase_item(
        purchase_id,
        item_id, 
        price, 
        weight_amount, 
        weight_unit, 
        quantity_amount, 
        quantity_unit)
      SELECT
        $1 AS purchase_id,
        (SELECT * FROM retrieved_item_id),
        $3 AS price,
        $4 AS weight_amount,
        $5 AS weight_unit,
        $6 AS quantity_amount,
        $7 AS quantity_unit
      FROM GENERATE_SERIES(1, $8)
      RETURNING id, price, weight_amount, weight_unit, quantity_amount, quantity_unit
	`, purchaseID,
		name,
		price,
		weightAmount,
		weightUnit,
		quantityAmount,
		quantityUnit,
		number,
		itemType,
	).Scan(
		&tempID,
		&purchaseItem.Price,
		&purchaseItem.WeightAmount,
		&purchaseItem.WeightUnit,
		&purchaseItem.QuantityAmount,
		&purchaseItem.QuantityUnit,
	)

	purchaseItem.ID = strconv.Itoa(tempID)

	return &purchaseItem, err
}

func (r *mutationResolver) UpdatePurchaseItem(ctx context.Context, id string, name string, price *float64, weightAmount *float64, weightUnit *string, quantityAmount *float64, quantityUnit *string) (*model.PurchaseItem, error) {
	panic(fmt.Errorf("UpdatePurchaseItem not implemented"))
}

func (r *mutationResolver) DeletePurchaseItem(ctx context.Context, id string) (*int, error) {
	result, err := db.Conn.Exec(context.Background(), `
      DELETE FROM purchase_item
      WHERE id = $1
	`, id)

	rows := new(int)
	*rows = int(result.RowsAffected())

	return rows, err
}

func (r *mutationResolver) AddInventoryItem(ctx context.Context, name string, addDate *string, expiration *string, amount *string, defaultShelflife *string, category *string, location *string, itemType string, number int) (*model.InventoryItem, error) {
	panic(fmt.Errorf("AddInventoryItem not implemented"))
}

func (r *mutationResolver) UpdateInventoryItem(ctx context.Context, id string, addDate *string, expiration *string, amount *string, location *string, category *string, itemType *string) (*model.InventoryItem, error) {
	panic(fmt.Errorf("UpdateInventoryItem not implemented"))
}

func (r *mutationResolver) DeleteInventoryItem(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("DeleteInventoryItem not implemented"))
}

func (r *mutationResolver) EditItem(ctx context.Context, id string, name string, categoryID *int, defaultLocationID *int, defaultShelflife *int, itemType string, countsAs []*string) (*model.Item, error) {
	// TODO: change name to UpdateItem (need to update client as well)
	idNum, _ := strconv.Atoi(id)

	_, deleteErr := db.Conn.Exec(context.Background(), `
			DELETE FROM item_counts_as
      WHERE specific_item_id = $1 
		`,
		idNum)

	if deleteErr != nil {
		fmt.Println("Error on delete:", deleteErr)
		return nil, deleteErr
	}

	updatedItem := model.Item{
		ID: id,
	}

	updateErr := db.Conn.QueryRow(context.Background(), `
			UPDATE item
      SET name = $2,
          category_id = $3,
          default_location_id = $4,
          default_shelflife = $5,
          item_type = $6
      WHERE id = $1
      RETURNING name, default_shelflife, item_type
	`,
		idNum,
		name,
		categoryID,
		defaultLocationID,
		defaultShelflife,
		itemType,
	).Scan(
		&updatedItem.Name,
		&updatedItem.DefaultShelflife,
		&updatedItem.ItemType,
	)

	if updateErr != nil {
		fmt.Println("Error on update:", updateErr)
		return nil, updateErr
	}

	for _, item := range countsAs {
		_, err := db.Conn.Exec(context.Background(), `
				WITH retrieved_item_id AS (
					SELECT item_id_for_insert($1)
				)
				INSERT INTO item_counts_as(specific_item_id, generic_item_id)
				SELECT 
					$2 as specific_item_id,
					(SELECT * FROM retrieved_item_id) AS generic_item_id
			`,
			item,
			idNum,
		)

		if err != nil {
			fmt.Println("Error on insert:", err)
			return nil, err
		}
	}

	return &updatedItem, nil
}

func (r *mutationResolver) AddDish(ctx context.Context, name string, tags []*string, isActive bool, ingredientSets []*model.IngredientSetInput) (*model.Dish, error) {
	panic(fmt.Errorf("AddDish not implemented"))
}

func (r *mutationResolver) UpdateDish(ctx context.Context, id string, name string, tags []*string, isActive bool, ingredientSets []*model.IngredientSetInput) (*model.Dish, error) {
	panic(fmt.Errorf("UpdateDish not implemented"))
}

func (r *mutationResolver) DeleteDish(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("DeleteDish not implemented"))
}

func (r *mutationResolver) AddDishDate(ctx context.Context, dishID string, date string) (*model.DishDate, error) {
	panic(fmt.Errorf("AddDishDate not implemented"))
}

func (r *mutationResolver) DeleteDishDate(ctx context.Context, id string) (*int, error) {
	panic(fmt.Errorf("DeleteDishDate not implemented"))
}

func (r *purchaseResolver) Location(ctx context.Context, obj *model.Purchase) (*model.PurchaseLocation, error) {
	var location model.PurchaseLocation
	var id int
	err := db.Conn.QueryRow(context.Background(), `
      SELECT purchase_location.id, name
      FROM purchase_location
      INNER JOIN purchase ON purchase.location_id = purchase_location.id
      WHERE purchase.id = $1
		`, obj.ID).Scan(&id, &location.Name)

	location.ID = strconv.Itoa(id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &location, nil
}

func (r *purchaseResolver) Items(ctx context.Context, obj *model.Purchase) ([]*model.PurchaseItem, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT 
        id, 
        price, 
        weight_amount,
        weight_unit,
        quantity_amount,
        quantity_unit
      FROM purchase_item
      WHERE purchase_id = $1
      ORDER BY id DESC
		`, obj.ID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*model.PurchaseItem{}

	for rows.Next() {
		var item model.PurchaseItem
		var tempID int

		err := rows.Scan(&tempID, &item.Price, &item.WeightAmount, &item.WeightUnit, &item.QuantityAmount, &item.QuantityUnit)
		if err != nil {
			return nil, err
		}

		item.ID = strconv.Itoa(tempID)
		items = append(items, &item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return items, nil
}

func (r *purchaseItemResolver) Item(ctx context.Context, obj *model.PurchaseItem) (*model.Item, error) {
	item := model.Item{}
	var tempID int

	err := db.Conn.QueryRow(context.Background(), `
      SELECT item.id, name, COALESCE(default_shelflife, 0) AS default_shelflife, item_type
      FROM item
			INNER JOIN purchase_item ON purchase_item.item_id = item.id
      WHERE purchase_item.id = $1 
		`, obj.ID).Scan(&tempID, &item.Name, &item.DefaultShelflife, &item.ItemType)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	item.ID = strconv.Itoa(tempID)

	return &item, nil
}

func (r *purchaseItemResolver) Purchase(ctx context.Context, obj *model.PurchaseItem) (*model.Purchase, error) {
	purchase := model.Purchase{}
	var tempID int

	err := db.Conn.QueryRow(context.Background(), `
			SELECT purchase.id, CAST(EXTRACT(epoch FROM date) * 1000 AS TEXT)
			FROM purchase
			INNER JOIN purchase_item on purchase_item.purchase_id = purchase.id
			WHERE purchase_item.id = $1
		`, obj.ID).Scan(&tempID, &purchase.Date)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	purchase.ID = strconv.Itoa(tempID)

	return &purchase, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT id, name, COALESCE(default_shelflife, 0), item_type
      FROM item 
      ORDER BY name
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*model.Item{}

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
			SELECT name, default_shelflife, item_type
			FROM item 
			WHERE id = $1
		`, idNum).Scan(&item.Name, &item.DefaultShelflife, &item.ItemType)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &item, nil
}

func (r *queryResolver) ItemByName(ctx context.Context, name string) (*model.Item, error) {
	item := model.Item{
		Name: name,
	}
	var itemID int

	err := db.Conn.QueryRow(context.Background(), `
      SELECT id, default_shelflife, item_type
      FROM item
      WHERE name = $1
		`, name).Scan(&itemID, &item.DefaultShelflife, &item.ItemType)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	item.ID = strconv.Itoa(itemID)

	return &item, nil
}

func (r *queryResolver) Dishes(ctx context.Context) ([]*model.Dish, error) {
	rows, err := db.Conn.Query(context.Background(), `
      SELECT id, name, is_active_dish
			FROM item
      WHERE item_type = 'dish'
      ORDER BY name
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dishes := []*model.Dish{}

	for rows.Next() {
		var dish model.Dish
		var dishID int

		err := rows.Scan(&dishID, &dish.Name, &dish.IsActiveDish)
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

func (r *queryResolver) Dish(ctx context.Context, id string) (*model.Dish, error) {
	dish := model.Dish{
		ID: id,
	}
	idNum, _ := strconv.Atoi(id)

	err := db.Conn.QueryRow(context.Background(), `
      SELECT name, is_active_dish
			FROM item
      WHERE id = $1
		`, idNum).Scan(&dish.Name, &dish.IsActiveDish)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dish, nil
}

func (r *queryResolver) InventoryItems(ctx context.Context) ([]*model.InventoryItem, error) {
	rows, err := db.Conn.Query(context.Background(), `
    SELECT id, expiration, add_date, amount
    FROM inventory_item 
    ORDER BY expiration ASC
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*model.InventoryItem{}

	for rows.Next() {
		var item model.InventoryItem
		var itemID int

		err := rows.Scan(&itemID, &item.Expiration, &item.AddDate, &item.Amount)
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

func (r *queryResolver) InventoryItem(ctx context.Context, id string) (*model.InventoryItem, error) {
	item := model.InventoryItem{
		ID: id,
	}
	idNum, _ := strconv.Atoi(id)

	err := db.Conn.QueryRow(context.Background(), `
      SELECT expiration, add_date, amount
      FROM inventory_item
      WHERE id = $1
		`, idNum).Scan(&item.Expiration, &item.AddDate, &item.Amount)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &item, nil
}

func (r *queryResolver) ItemLocations(ctx context.Context) ([]*model.ItemLocation, error) {
	rows, err := db.Conn.Query(context.Background(), `
		SELECT id, name
		FROM inventory_item_location
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locations := []*model.ItemLocation{}

	for rows.Next() {
		var location model.ItemLocation
		var locationID int

		err := rows.Scan(&locationID, &location.Name)
		if err != nil {
			return nil, err
		}

		location.ID = strconv.Itoa(locationID)
		locations = append(locations, &location)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return locations, nil
}

func (r *queryResolver) ItemCategories(ctx context.Context) ([]*model.ItemCategory, error) {
	rows, err := db.Conn.Query(context.Background(), `
		SELECT id, name
		FROM item_category
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []*model.ItemCategory{}

	for rows.Next() {
		var category model.ItemCategory
		var categoryID int

		err := rows.Scan(&categoryID, &category.Name)
		if err != nil {
			return nil, err
		}

		category.ID = strconv.Itoa(categoryID)
		categories = append(categories, &category)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return categories, nil
}

func (r *queryResolver) DishTags(ctx context.Context) ([]*model.DishTag, error) {
	rows, err := db.Conn.Query(context.Background(), `
		SELECT id, name
		FROM dish_tag
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []*model.DishTag{}

	for rows.Next() {
		var tag model.DishTag
		var tagID int

		err := rows.Scan(&tagID, &tag.Name)
		if err != nil {
			return nil, err
		}

		tag.ID = strconv.Itoa(tagID)
		tags = append(tags, &tag)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return tags, nil
}

func (r *queryResolver) Purchases(ctx context.Context) ([]*model.Purchase, error) {
	rows, err := db.Conn.Query(context.Background(), `
		SELECT id, CAST(EXTRACT(epoch FROM date) * 1000 AS TEXT)
		FROM purchase
		ORDER BY date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	purchases := []*model.Purchase{}

	for rows.Next() {
		var purchase model.Purchase
		var purchaseID int

		err := rows.Scan(&purchaseID, &purchase.Date)
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

func (r *queryResolver) Purchase(ctx context.Context, id string) (*model.Purchase, error) {
	purchase := model.Purchase{
		ID: id,
	}
	idNum, _ := strconv.Atoi(id)

	err := db.Conn.QueryRow(context.Background(), `
      SELECT CAST(EXTRACT(epoch FROM date) * 1000 AS TEXT) AS date
			FROM purchase
      WHERE id = $1
		`, idNum).Scan(&purchase.Date)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("DATE:\n", purchase.Date)

	return &purchase, nil
}

func (r *queryResolver) PurchaseLocations(ctx context.Context) ([]*model.PurchaseLocation, error) {
	rows, err := db.Conn.Query(context.Background(), `
		SELECT id, name
		FROM purchase_location
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locations := []*model.PurchaseLocation{}

	for rows.Next() {
		var location model.PurchaseLocation
		var locationID int

		err := rows.Scan(&locationID, &location.Name)
		if err != nil {
			return nil, err
		}

		location.ID = strconv.Itoa(locationID)
		locations = append(locations, &location)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return locations, nil
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
