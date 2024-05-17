package data

import (
	"fmt"
	"strconv"
)

type Item struct {
	Name        string
	Description string
	Buy         int
	Sell        int
	Type        string
}

type ItemsModel struct {
	headings []string
	items    []Item
}

func (m *ItemsModel) GetItems() []Item {
	return m.items
}

func (m *ItemsModel) GetItem(index int) Item {
	if index < len(m.items)-1 {
		return m.items[index+1]
	}
	return Item{}
}

func NewItemsModel(filepath string) (ItemsModel, error) {
	var itemsModel ItemsModel
	fmt.Print("New Items Model\n")

	records, err := ReadCSV(filepath)
	if err != nil {
		return ItemsModel{}, nil
	}

	itemsModel.items = make([]Item, len(records)-1)
	for i, r := range records[1:] {
		intBuy, err := strconv.Atoi(r[2])
		if err != nil {
			return ItemsModel{}, nil
		}

		intSell, err := strconv.Atoi(r[3])
		if err != nil {
			return ItemsModel{}, nil
		}

		item := Item{
			Name:        r[0],
			Description: r[1],
			Buy:         intBuy,
			Sell:        intSell,
			Type:        r[4],
		}

		itemsModel.items[i] = item
	}

	return itemsModel, nil
}
