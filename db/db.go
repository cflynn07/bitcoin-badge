// Page 28 of The Little Go Book
package db

type Item struct {
	Price int32 // US Cents
}

func LoadItem(id int) *Item {
	return &Item{
		Price: 100,
	}
}
