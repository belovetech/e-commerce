// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql"
	"time"
)

type Order struct {
	ID        int32
	UserID    int32
	Total     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	ID        int32
	OrderID   int32
	ProductID int32
	Quantity  int32
	Price     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          int32
	Name        string
	Description sql.NullString
	Price       string
	Stock       int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SeedingHistory struct {
	ID         int32
	SeederName string
	ExecutedAt time.Time
}

type User struct {
	ID        int32
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}