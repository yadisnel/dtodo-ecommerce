package model

type Shop struct {
	Name         string   `bson:"name"`
	Images       []string `bson:"images"`
	Location     Location `bson:"location"`
	ProvinceId   string   `bson:"province-id"`
	ProvinceName string   `bson:"province-name"`
}
