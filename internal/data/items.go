package data

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type ItemFilter struct {
	Query       string
	QueryTarget string
	Extension   string
	Sort        string
	Page        int
	PageSize    int
}

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

func (m *ItemsModel) GetHeadings() []string {
	return m.headings
}

func filterItems(items []Item, filter ItemFilter) []Item {
	query := strings.ToLower(filter.Query)
	target := strings.ToLower(filter.QueryTarget)
	extension := strings.ToLower(filter.Extension)
	filteredItems := make([]Item, 0)
	fmt.Printf("filter items called: query=%s target=%s\n", query, target)

	var itemValue string
	var itemValueInt int
	for _, item := range items {
		switch target {
		case "name":
			itemValue = strings.ToLower(item.Name)
		case "description":
			itemValue = strings.ToLower(item.Description)
		case "type":
			itemValue = strings.ToLower(item.Type)
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

			switch extension {
			case "gt":
				if itemValueInt >= queryInt {
					filteredItems = append(filteredItems, item)
				}
			default:
				if itemValueInt <= queryInt {
					filteredItems = append(filteredItems, item)
				}
			}
		}
	}

	return filteredItems
}

func sortItems(items []Item, filter ItemFilter) []Item {
	if len(filter.Sort) == 0 {
		return items
	}

	validSortFields := []string{"name", "description", "type", "buy", "sell"}
	field := strings.ToLower(filter.Sort)

	asc := true
	if []rune(filter.Sort)[0] == '-' {
		asc = false
		field = filter.Sort[1:]
	}

	if !slices.Contains(validSortFields, field) {
		return items
	}

	itemsCopy := make([]Item, len(items))
	copy(itemsCopy, items)

	sort.Slice(itemsCopy, func(i, j int) bool {
		switch field {
		case "name":
			fmt.Print("sorting by name\n")
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
		case "type":
			if asc {
				return itemsCopy[i].Type < itemsCopy[j].Type
			} else {
				return itemsCopy[i].Type > itemsCopy[j].Type
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
		}

		return false
	})
	return itemsCopy
}

func (m *ItemsModel) GetItems(filter ItemFilter) []Item {

	retItems := sortItems(filterItems(m.items, filter), filter)

	pageSize := filter.PageSize
	page := filter.Page

	// If starting index would be beyond final page
	if len(retItems) < (page-1)*pageSize {
		return []Item{}
	}

	if pageSize < len(retItems) {
		start := (page - 1) * pageSize
		fmt.Printf("start: %d\n", start)
		if start+pageSize > len(retItems)-1 {
			fmt.Printf("end: end\n")
			retItems = retItems[start:]
		} else {
			fmt.Printf("end: %d\n", start+pageSize)
			retItems = retItems[start : start+pageSize]
		}
	}
	return retItems
}

func (m *ItemsModel) GetItem(index int) Item {
	if index < len(m.items)-1 {
		return m.items[index+1]
	}
	return Item{}
}

func NewItemsModel(filepath string) (ItemsModel, error) {
	var itemsModel ItemsModel

	records, err := ReadCSV(filepath)
	if err != nil {
		return ItemsModel{}, nil
	}

	itemsModel.headings = []string{"Name", "Description", "Buy", "Sell", "Type"}

	//itemsModel.headings = records[0] // CSV headings do not correspond to desired headings
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
