package data

import (
	"strconv"
)

type EquipmentFilter struct {
	Query       string `json:"query,omitempty"`
	QueryTarget string `json:"query_target,omitempty"`
	Extension   string `json:"extension,omitempty"`
	Sort        string `json:"sort,omitempty"`
	Page        int    `json:"page_number"`
	PageSize    int    `json:"page_size"`
}

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

func (m *EquipmentModel) GetEquipments(filter EquipmentFilter) []Equipment {
	//retItems := paginateItems(sortItems(filterItems(m.equipments, filter), filter), filter)
	return m.equipments
}

func (m *EquipmentModel) GetEquipment(index int) Equipment {
	if index < len(m.equipments)-1 {
		return m.equipments[index+1]
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
		"Buy", "Sell", "Source(s)", // 12, 13, 14
		"Type"} // 15

	//itemModel.headings = records[0] // CSV headings do not correspond to desired headings
	equipmentModel.equipments = make([]Equipment, len(records)-1)
	for i, r := range records[1:] {
		name := r[0]
		hp, err := strconv.Atoi(r[1][:3])
		sp, err := strconv.Atoi(r[2][:3])
		physAtk, err := strconv.Atoi(r[3][:3])
		eleAtk, err := strconv.Atoi(r[4][:3])
		physDef, err := strconv.Atoi(r[5][:3])
		eleDef, err := strconv.Atoi(r[6][:3])
		accuracy, err := strconv.Atoi(r[7][:3])
		speed, err := strconv.Atoi(r[8][:3])
		critical, err := strconv.Atoi(r[9][:3])
		evasion, err := strconv.Atoi(r[10][:3])
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
