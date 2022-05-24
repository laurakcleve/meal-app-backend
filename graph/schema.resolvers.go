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

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	var item model.Item
	idNum, _ := strconv.Atoi(id)
	err := db.Conn.QueryRow(context.Background(), `SELECT name FROM item WHERE id = $1`, idNum).Scan(&item.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &item, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	rows, err := db.Conn.Query(context.Background(),`SELECT id, name FROM item`)
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }