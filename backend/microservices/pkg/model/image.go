package model

type Image struct {
	Id             string `bson:"id"`
	OriginalKey    string `bson:"original-key"`
	OriginalWidth  int64  `bson:"original-with"`
	OriginalHeight int64  `bson:"original-height"`
	OriginalZize   int64  `bson:"original-size"`
	OriginalUrl    string `bson:"original-url"`
	ThumbKey       string `bson:"thumb-key"`
	ThumbWidth     int64  `bson:"thumb-width"`
	ThumbHeight    int64  `bson:"thumb-height"`
	ThumbSize      int64  `bson:"thumb-size"`
	ThumbUrl       string `bson:"thumb-url"`
}
