package model

type Pagination struct {
	Pages    string `bson:"pages"`
	Elements string `bson:"elements"`
	Total    string `bson:"total"`
}
