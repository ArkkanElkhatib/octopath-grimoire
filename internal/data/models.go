package data

type ModelsConfig struct {
	ItemsFilepath string
}

type Models struct {
	ItemsModel ItemsModel
}

func LoadModels(cfg *ModelsConfig) *Models {
	itemsModel, err := NewItemsModel(cfg.ItemsFilepath)
	if err != nil {
		return &Models{}
	}

	return &Models{
		ItemsModel: itemsModel,
	}
}
