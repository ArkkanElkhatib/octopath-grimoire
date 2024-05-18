package data

type Filter struct {
	Query       string `json:"query,omitempty"`
	QueryTarget string `json:"query_target,omitempty"`
	Extension   string `json:"extension,omitempty"`
	Sort        string `json:"sort,omitempty"`
	Page        int    `json:"page_number"`
	PageSize    int    `json:"page_size"`
}

type Metadata struct {
	Total    int `json:"total"`
	Returned int `json:"returned"`
}

type ModelsConfig struct {
	ItemsFilepath      string
	EquipmentsFilepath string
}

type Models struct {
	ItemModel      ItemModel
	EquipmentModel EquipmentModel
}

func LoadModels(cfg *ModelsConfig) *Models {
	itemModel, err := NewItemModel(cfg.ItemsFilepath)
	equipmentModel, err := NewEquipmentModel(cfg.EquipmentsFilepath)
	if err != nil {
		return &Models{}
	}

	return &Models{
		ItemModel:      itemModel,
		EquipmentModel: equipmentModel,
	}
}
