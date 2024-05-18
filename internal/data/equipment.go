package data

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Equipment struct {
	Name             string `json:"name"`
	HP               int    `json:"hp"`
	SP               int    `json:"sp"`
	PhysicalAttack   int    `json:"physical_attack"`
	ElementalAttack  int    `json:"elemental_attack"`
	PhysicalDefense  int    `json:"physical_defense"`
	ElementalDefense int    `json:"elemental_defense"`
	Accuracy         int    `json:"accuracy"`
	Speed            int    `json:"speed"`
	Critical         int    `json:"critical"`
	Evasion          int    `json:"evasion"`
	Effect           string `json:"effect"`
	Buy              int    `json:"buy_price"`
	Sell             int    `json:"sell_price"`
	Source           string `json:"sources"`
	Type             string `json:"type"`
}

type EquipmentModel struct {
	headings   []string
	equipments []Equipment
}

var emptyEquipment = Equipment{}

func (e Equipment) IsEmpty() bool {
	return e == emptyEquipment
}

func filterEquipment(equipments []Equipment, filter Filter) []Equipment {
	query := strings.ToLower(filter.Query)
	target := strings.ToLower(filter.QueryTarget)
	extension := strings.ToLower(filter.Extension)
	filteredEquipments := make([]Equipment, 0)

	var equipmentValue string
	var equipmentValueInt int
	for _, equipment := range equipments {
		switch target {
		// string values
		case "name":
			equipmentValue = strings.ToLower(equipment.Name)
		case "effect":
			equipmentValue = strings.ToLower(equipment.Effect)
		case "source":
			equipmentValue = strings.ToLower(equipment.Source)
		case "type":
			equipmentValue = strings.ToLower(equipment.Type)

		// int values
		case "hp":
			equipmentValueInt = equipment.HP
		case "sp":
			equipmentValueInt = equipment.SP
		case "physical attack":
			equipmentValueInt = equipment.PhysicalAttack
		case "elemental attack":
			equipmentValueInt = equipment.ElementalAttack
		case "physical defense":
			equipmentValueInt = equipment.PhysicalDefense
		case "elemental defense":
			equipmentValueInt = equipment.ElementalDefense
		case "accuracy":
			equipmentValueInt = equipment.Accuracy
		case "speed":
			equipmentValueInt = equipment.Speed
		case "critcal":
			equipmentValueInt = equipment.Critical
		case "evasion":
			equipmentValueInt = equipment.Evasion
		case "buy":
			equipmentValueInt = equipment.Buy
		case "sell":
			equipmentValueInt = equipment.Sell

		default:
			equipmentValue = strings.ToLower(equipment.Name)
		}

		if equipmentValue != "" {
			if strings.Contains(equipmentValue, query) {
				filteredEquipments = append(filteredEquipments, equipment)
			}
		} else if equipmentValueInt != 0 {
			queryInt, err := strconv.Atoi(query)
			if err != nil {
				return filteredEquipments
			}

			if extension == "gt" {
				if equipmentValueInt >= queryInt {
					filteredEquipments = append(filteredEquipments, equipment)
				}
			} else {
				if equipmentValueInt <= queryInt {
					filteredEquipments = append(filteredEquipments, equipment)
				}
			}
		}
	}

	return filteredEquipments
}

func sortEquipment(equipments []Equipment, filter Filter) []Equipment {
	if len(filter.Sort) == 0 {
		return equipments
	}

	validSortFields := []string{
		"name", "hp", "sp", // 0, 1, 2
		"physical attack ", "elemental attack", "elemental defense", // 3, 4, 5
		"elemental defense", "accuracy", "speed", // 6, 7, 8
		"critical", "evasion", "effect", // 9, 10, 11
		"buy", "sell", "source", // 12, 13, 14
		"type"} // 15
	field := strings.ToLower(filter.Sort)

	asc := true
	if []rune(filter.Sort)[0] == '-' {
		asc = false
		field = filter.Sort[1:] // Remove the '-' prefix
	}

	if !slices.Contains(validSortFields, field) {
		return equipments
	}

	// Copy items to sort in place
	// TODO:: TreeSet implementation rather than copy and sort with BinarySort
	equipmentsCopy := make([]Equipment, len(equipments))
	copy(equipmentsCopy, equipments)

	sort.Slice(equipmentsCopy, func(i, j int) bool {
		switch field {
		case "name":
			if asc {
				return equipmentsCopy[i].Name < equipmentsCopy[j].Name
			} else {
				return equipmentsCopy[i].Name > equipmentsCopy[j].Name
			}
		case "hp":
			if asc {
				return equipmentsCopy[i].HP < equipmentsCopy[j].HP
			} else {
				return equipmentsCopy[i].HP > equipmentsCopy[j].HP
			}
		case "sp":
			if asc {
				return equipmentsCopy[i].SP < equipmentsCopy[j].SP
			} else {
				return equipmentsCopy[i].SP > equipmentsCopy[j].SP
			}
		case "physical attack":
			if asc {
				return equipmentsCopy[i].PhysicalAttack < equipmentsCopy[j].PhysicalAttack
			} else {
				return equipmentsCopy[i].PhysicalAttack > equipmentsCopy[j].PhysicalAttack
			}
		case "elemental attack":
			if asc {
				return equipmentsCopy[i].ElementalAttack < equipmentsCopy[j].ElementalAttack
			} else {
				return equipmentsCopy[i].ElementalAttack > equipmentsCopy[j].ElementalAttack
			}
		case "physical defense":
			if asc {
				return equipmentsCopy[i].PhysicalDefense < equipmentsCopy[j].PhysicalDefense
			} else {
				return equipmentsCopy[i].PhysicalDefense > equipmentsCopy[j].PhysicalDefense
			}
		case "elemental defense":
			if asc {
				return equipmentsCopy[i].ElementalDefense < equipmentsCopy[j].ElementalDefense
			} else {
				return equipmentsCopy[i].ElementalDefense > equipmentsCopy[j].ElementalDefense
			}
		case "accuracy":
			if asc {
				return equipmentsCopy[i].Accuracy < equipmentsCopy[j].Accuracy
			} else {
				return equipmentsCopy[i].Accuracy > equipmentsCopy[j].Accuracy
			}
		case "speed":
			if asc {
				return equipmentsCopy[i].Speed < equipmentsCopy[j].Speed
			} else {
				return equipmentsCopy[i].Speed > equipmentsCopy[j].Speed
			}
		case "critical":
			if asc {
				return equipmentsCopy[i].Critical < equipmentsCopy[j].Critical
			} else {
				return equipmentsCopy[i].Critical > equipmentsCopy[j].Critical
			}
		case "evasion":
			if asc {
				return equipmentsCopy[i].Evasion < equipmentsCopy[j].Evasion
			} else {
				return equipmentsCopy[i].Evasion > equipmentsCopy[j].Evasion
			}
		case "effect":
			if asc {
				return equipmentsCopy[i].Effect < equipmentsCopy[j].Effect
			} else {
				return equipmentsCopy[i].Effect > equipmentsCopy[j].Effect
			}
		case "buy":
			if asc {
				return equipmentsCopy[i].Buy < equipmentsCopy[j].Buy
			} else {
				return equipmentsCopy[i].Buy > equipmentsCopy[j].Buy
			}
		case "sell":
			if asc {
				return equipmentsCopy[i].Sell < equipmentsCopy[j].Sell
			} else {
				return equipmentsCopy[i].Sell > equipmentsCopy[j].Sell
			}
		case "source":
			if asc {
				return equipmentsCopy[i].Source < equipmentsCopy[j].Source
			} else {
				return equipmentsCopy[i].Source > equipmentsCopy[j].Source
			}
		case "type":
			if asc {
				return equipmentsCopy[i].Type < equipmentsCopy[j].Type
			} else {
				return equipmentsCopy[i].Type > equipmentsCopy[j].Type
			}
		default:
			return false
		}
	})
	return equipmentsCopy
}

func paginateEquipment(equipments []Equipment, filter Filter) []Equipment {
	pageSize := filter.PageSize
	page := filter.Page

	// If starting index would be beyond final page
	if len(equipments) < (page-1)*pageSize {
		return []Equipment{}
	}

	/* if pageSize > len(equipments) {
		if page == 1 {
			return equipments
		}
		return []Equipment{}
	} */

	if pageSize < len(equipments) {
		start := (page - 1) * pageSize
		if start+pageSize > len(equipments)-1 {
			equipments = equipments[start:]
		} else {
			equipments = equipments[start : start+pageSize]
		}
	}
	return equipments
}

func (m *EquipmentModel) GetEquipments(filter Filter) ([]Equipment, int) {
	filteredEquipment := filterEquipment(m.equipments, filter)
	numResults := len(filteredEquipment)

	return paginateEquipment(sortEquipment(filteredEquipment, filter), filter), numResults
}

func (m *EquipmentModel) GetEquipment(index int) Equipment {
	if index >= 0 && index <= len(m.equipments) {
		return m.equipments[index-1]
	}
	return Equipment{}
}

func NewEquipmentModel(filepath string) (EquipmentModel, error) {
	var equipmentModel EquipmentModel

	records, err := ReadCSV(filepath)
	if err != nil {
		return EquipmentModel{}, nil
	}

	equipmentModel.headings = []string{
		"Name", "Maximum HP", "Maximum SP", // 0, 1, 2
		"Physical Attack", "Elemental Attack", "Physical Defense", // 3, 4, 5
		"Elemental Defense", "Accuracy", "Speed", // 6, 7, 8
		"Critical", "Evasion", "Effect", // 9, 10, 11
		"Buy", "Sell", "Source", // 12, 13, 14
		"Type"} // 15

	intFromFloatStringRX := regexp.MustCompile(`^[0-9\-]+`)
	//itemModel.headings = records[0] // CSV headings do not correspond to desired headings
	fmt.Printf("Num Records: %d\n", len(records))
	equipmentModel.equipments = make([]Equipment, len(records)-1)
	for i, r := range records[1:] {
		name := r[0]
		hp, err := strconv.Atoi(r[1])
		sp, err := strconv.Atoi(r[2])
		physAtk, err := strconv.Atoi(intFromFloatStringRX.FindString(r[3]))
		eleAtk, err := strconv.Atoi(intFromFloatStringRX.FindString(r[4]))
		physDef, err := strconv.Atoi(intFromFloatStringRX.FindString(r[5]))
		eleDef, err := strconv.Atoi(intFromFloatStringRX.FindString(r[6]))
		accuracy, err := strconv.Atoi(intFromFloatStringRX.FindString(r[7]))
		speed, err := strconv.Atoi(intFromFloatStringRX.FindString(r[8]))
		critical, err := strconv.Atoi(intFromFloatStringRX.FindString(r[9]))
		evasion, err := strconv.Atoi(intFromFloatStringRX.FindString(r[10]))
		effect := r[11]
		buy, err := strconv.Atoi(r[12])
		sell, err := strconv.Atoi(r[13])
		source := r[14]
		equipmentType := r[15]
		if err != nil {
			return EquipmentModel{}, nil
		}

		equipment := Equipment{
			Name:             name,
			HP:               hp,
			SP:               sp,
			PhysicalAttack:   physAtk,
			ElementalAttack:  eleAtk,
			PhysicalDefense:  physDef,
			ElementalDefense: eleDef,
			Accuracy:         accuracy,
			Speed:            speed,
			Critical:         critical,
			Evasion:          evasion,
			Effect:           effect,
			Buy:              buy,
			Sell:             sell,
			Source:           source,
			Type:             equipmentType,
		}

		equipmentModel.equipments[i] = equipment
	}

	return equipmentModel, nil
}
