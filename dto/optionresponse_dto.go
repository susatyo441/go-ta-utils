package dto

type InterfaceOptionDTO struct {
	Label string      `json:"label" bson:"label"`
	Value interface{} `json:"value" bson:"value"`
}

type OptionDTO[T any] struct {
	Label string `json:"label" bson:"label"`
	Value T      `json:"value" bson:"value"`
}
