package model

type Item struct {
	ID     						string 					`json:"id"`
	Name   						string					`json:"name"`
	Category  				*ItemCategory  	`json:"category"`
	Dishes 						[]*Dish 				`json:"dishes"`
	DefaultLocation 	*ItemLocation 	`json:"defaultLocation"`
	DefaultShelflife 	int 						`json:"defaultShelflife"`
	ItemType 					string 					`json:"itemType"`
	Purchases 				[]*PurchaseItem `json:"purchases"`
	CountsAs 					[]*Item 				`json:"countsAs"`
}

type ItemCategory struct {
	ID 		string `json:"id"`
	Name	string `json:"name"`
}

type Dish struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Tags           []*DishTag       `json:"tags"`
	IsActiveDish   *bool            `json:"isActiveDish"`
	Dates          []*DishDate      `json:"dates"`
	IngredientSets []*IngredientSet `json:"ingredientSets"`
}

type DishDate struct {
	ID   string `json:"id"`
	Date string `json:"date"`
}

type DishTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Ingredient struct {
	ID            string `json:"id"`
	Item          *Item  `json:"item"`
	IsInInventory *bool  `json:"isInInventory"`
}

type IngredientInput struct {
	ID            string               `json:"id"`
	Item          *IngredientItemInput `json:"item"`
	IsInInventory *bool                `json:"isInInventory"`
}

type IngredientItemInput struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

type IngredientSet struct {
	ID          string        `json:"id"`
	IsOptional  *bool         `json:"isOptional"`
	Ingredients []*Ingredient `json:"ingredients"`
}

type IngredientSetInput struct {
	ID          string             `json:"id"`
	IsOptional  bool               `json:"isOptional"`
	Ingredients []*IngredientInput `json:"ingredients"`
}

type InventoryItem struct {
	ID         string        `json:"id"`
	Item       *Item         `json:"item"`
	Location   *ItemLocation `json:"location"`
	Expiration *string       `json:"expiration"`
	AddDate    *string       `json:"addDate"`
	Amount     *string       `json:"amount"`
}

type ItemLocation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Purchase struct {
	ID       string            `json:"id"`
	Date     string            `json:"date"`
	Location *PurchaseLocation `json:"location"`
	Items    []*PurchaseItem   `json:"items"`
}

type PurchaseItem struct {
	ID             string    `json:"id"`
	Item           *Item     `json:"item"`
	Price          *float64  `json:"price"`
	WeightAmount   *float64  `json:"weightAmount"`
	WeightUnit     *string   `json:"weightUnit"`
	QuantityAmount *float64  `json:"quantityAmount"`
	QuantityUnit   *string   `json:"quantityUnit"`
	PurchaseID     *int      `json:"purchaseId"`
	Purchase       *Purchase `json:"purchase"`
}

type PurchaseLocation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}