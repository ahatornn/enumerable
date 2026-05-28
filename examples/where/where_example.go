// go run examples/where/where_example.go
package main

import (
	"fmt"

	"github.com/ahatornn/enumerable"
)

type Order struct {
	ID     int
	Amount float64
	Status string // e.g., "Pending", "Shipped", "Cancelled"
}

func main() {
	orders := []Order{
		{ID: 1, Amount: 50.0, Status: "Pending"},
		{ID: 2, Amount: 150.0, Status: "Pending"},
		{ID: 3, Amount: 200.0, Status: "Shipped"},
		{ID: 4, Amount: 300.0, Status: "Shipped"},
		{ID: 5, Amount: 120.0, Status: "Cancelled"},
		{ID: 6, Amount: 180.0, Status: "Shipped"},
		{ID: 7, Amount: 250.0, Status: "Shipped"},
		{ID: 8, Amount: 110.0, Status: "Shipped"},
	}

	fmt.Println("All orders:")
	for _, o := range orders {
		fmt.Printf("  - ID: %d, Amount: %.2f, Status: %s\n", o.ID, o.Amount, o.Status)
	}

	idsOfShippedHighValueOrders := enumerable.FromSlice(orders).
		Where(func(o Order) bool {
			return o.Amount > 100.0 && o.Status == "Shipped"
		}).
		Take(3).
		ToSlice()

	var finalIds []int
	for _, o := range idsOfShippedHighValueOrders {
		finalIds = append(finalIds, o.ID)
	}

	fmt.Println("\n--- using enumerable ---")
	fmt.Println("ID first 3 orders with Amount > 100 and Status == 'Shipped':")
	fmt.Println(finalIds)
}
