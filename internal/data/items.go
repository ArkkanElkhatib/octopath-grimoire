package data

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Buy         int    `json:"buy_price"`
	Sell        int    `json:"sell_price"`
	Type        string `json:"type"`
}

type ItemModel struct {
	headings []string
	items    []Item
}

var emptyItem = Item{}

func (i Item) IsEmpty() bool {
	return i == emptyItem
}

func (m *ItemModel) GetHeadings() []string {
	return m.headings
}

func filterItems(items []Item, filter Filter) []Item {
	query := strings.ToLower(filter.Query)
	target := strings.ToLower(filter.QueryTarget)
	extension := strings.ToLower(filter.Extension)
	filteredItems := make([]Item, 0)

	var itemValue string
	var itemValueInt int
	for _, item := range items {
		switch target {
		// string values
		case "name":
			itemValue = strings.ToLower(item.Name)
		case "description":
			itemValue = strings.ToLower(item.Description)
		case "type":
			itemValue = strings.ToLower(item.Type)

		// int values
		case "buy":
			itemValueInt = item.Buy
		case "sell":
			itemValueInt = item.Sell

		default:
			itemValue = strings.ToLower(item.Name)
		}

		if itemValue != "" {
			if strings.Contains(itemValue, query) {
				filteredItems = append(filteredItems, item)
			}
		} else if itemValueInt != 0 {
			queryInt, err := strconv.Atoi(query)
			if err != nil {
				return filteredItems
			}

			if extension == "gt" {
				if itemValueInt >= queryInt {
					filteredItems = append(filteredItems, item)
				}
			} else {
				if itemValueInt <= queryInt {
					filteredItems = append(filteredItems, item)
				}
			}
		}
	}

	return filteredItems
}

func sortItems(items []Item, filter Filter) []Item {
	if len(filter.Sort) == 0 {
		return items
	}

	validSortFields := []string{"name", "description", "type", "buy", "sell"}
	field := strings.ToLower(filter.Sort)

	asc := true
	if []rune(filter.Sort)[0] == '-' {
		asc = false
		field = filter.Sort[1:] // Remove the '-' prefix
	}

	if !slices.Contains(validSortFields, field) {
		return items
	}

	// Copy items to sort in place
	// TODO:: TreeSet implementation rather than copy and sort with BinarySort
	itemsCopy := make([]Item, len(items))
	copy(itemsCopy, items)

	sort.Slice(itemsCopy, func(i, j int) bool {
		switch field {
		case "name":
			if asc {
				return itemsCopy[i].Name < itemsCopy[j].Name
			} else {
				return itemsCopy[i].Name > itemsCopy[j].Name
			}
		case "description":
			if asc {
				return itemsCopy[i].Description < itemsCopy[j].Description
			} else {
				return itemsCopy[i].Description > itemsCopy[j].Description
			}
		case "buy":
			if asc {
				return itemsCopy[i].Buy < itemsCopy[j].Buy
			} else {
				return itemsCopy[i].Buy > itemsCopy[j].Buy
			}
		case "sell":
			if asc {
				return itemsCopy[i].Sell < itemsCopy[j].Sell
			} else {
				return itemsCopy[i].Sell > itemsCopy[j].Sell
			}
		case "type":
			if asc {
				return itemsCopy[i].Type < itemsCopy[j].Type
			} else {
				return itemsCopy[i].Type > itemsCopy[j].Type
			}
		default:
			return false
		}
	})
	return itemsCopy
}

func paginateItems(items []Item, filter Filter) []Item {
	pageSize := filter.PageSize
	page := filter.Page

	// If starting index would be beyond final page
	if len(items) < (page-1)*pageSize {
		return []Item{}
	}

	if pageSize < len(items) {
		start := (page - 1) * pageSize
		if start+pageSize > len(items)-1 {
			items = items[start:]
		} else {
			items = items[start : start+pageSize]
		}
	}
	return items
}

func (m *ItemModel) GetItems(filter Filter) ([]Item, int) {
	filteredItems := filterItems(m.items, filter)
	numResults := len(filteredItems)

	return paginateItems(sortItems(filteredItems, filter), filter), numResults
}

func (m *ItemModel) GetItem(index int) Item {
	if 0 <= index && index <= len(m.items) {
		return m.items[index-1]
	}
	return Item{}
}

func NewItemModel(filepath string) (ItemModel, error) {
	var itemModel ItemModel

	records, err := ReadCSV(filepath)
	if err != nil {
		return ItemModel{}, nil
	}

	itemModel.headings = []string{"Name", "Description", "Buy", "Sell", "Type"}

	//itemModel.headings = records[0] // CSV headings do not correspond to desired headings
	itemModel.items = make([]Item, len(records)-1)
	for i, r := range records[1:] {
		intBuy, err := strconv.Atoi(r[2])
		if err != nil {
			return ItemModel{}, nil
		}

		intSell, err := strconv.Atoi(r[3])
		if err != nil {
			return ItemModel{}, nil
		}

		item := Item{
			Name:        r[0],
			Description: r[1],
			Buy:         intBuy,
			Sell:        intSell,
			Type:        r[4],
		}

		itemModel.items[i] = item
	}

	return itemModel, nil
}
