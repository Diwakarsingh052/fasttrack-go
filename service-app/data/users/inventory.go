package users

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

// CreateInventory method add inventory for a particular user
func (db *DbService) CreateInventory(ctx context.Context, ni NewShirtInventory, userId string, now time.Time) (ShirtInventory, error) {

	// assigning values provide by the handler func
	inv := ShirtInventory{

		UserId:      userId,
		ItemName:    ni.ItemName,
		Quantity:    ni.Quantity,
		DateCreated: now,
		DateUpdated: now,
	}

	//inserting the data in db for specific user
	const q = `INSERT INTO inventory
		(user_id, item_name, quantity, date_created, date_updated)
		VALUES ( $1, $2, $3, $4, $5)
		Returning id`

	var id int
	//exec the query // QueryRowContext is used as we are expecting one row back in the result
	row := db.QueryRowContext(ctx, q, userId, inv.ItemName, inv.Quantity, inv.DateCreated, inv.DateUpdated)
	err := row.Scan(&id)

	if err != nil {
		return ShirtInventory{}, fmt.Errorf("inserting inventory %w", err)
	}

	inv.ID = strconv.Itoa(id)

	// returning inventory var if everything works fine
	return inv, nil
}

func (db *DbService) ViewAll(ctx context.Context, userId string) ([]ShirtInventory, error) {
	var inv []ShirtInventory

	const q = `Select id,user_id, item_name,quantity,date_created,date_updated FROM inventory where user_id = $1`
	//QueryContext is used when we are expecting multiple rows to be returned
	rows, err := db.QueryContext(ctx, q, userId)

	defer rows.Close()

	// rows.Next() returns true if more rows are left to read otherwise false which will stop the loop
	for rows.Next() {
		var newInv ShirtInventory
		err = rows.Scan(&newInv.ID, &newInv.UserId, &newInv.ItemName, &newInv.Quantity, &newInv.DateCreated, &newInv.DateUpdated)
		if err != nil {
			return nil, err
		}
		//appending the struct into the slice to maintain a whole list
		inv = append(inv, newInv)
	}

	return inv, nil
}
