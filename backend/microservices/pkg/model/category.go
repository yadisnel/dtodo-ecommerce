package model

type CategoryIn struct {
	LocalId       string        `bson:"local-id"`
	Name          string        `bson:"name"`
	NOrder        int64         `bson:"n-order"`
	SubCategories []SubCategory `bson:"sub-categories"`
	Created       int64         `bson:"created"`
	Modified      int64         `bson:"modified"`
	Deleted       bool          `bson:"deleted"`
}

type CategoryOut struct {
	Id            string        `bson:"_id,omitempty"`
	LocalId       string        `bson:"local-id"`
	Name          string        `bson:"name"`
	NOrder        int64         `bson:"n-order"`
	SubCategories []SubCategory `bson:"sub-categories"`
	Created       int64         `bson:"created"`
	Modified      int64         `bson:"modified"`
	Deleted       bool          `bson:"deleted"`
}
