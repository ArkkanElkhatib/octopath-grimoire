package data

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
