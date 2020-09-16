package model

type ProductIn struct {
	LocalId string `bson:"local-id"`
	UserId string  `bson:"user-id"`
	name string  `bson:"name"`
	description string  `bson:"description"`
	CategoryId string  `bson:"category-id"`
	Categoryame string  `bson:"category-name"`
	SubCategoryId string  `bson:"sub-category-id"`
	SubCategoryName string `bson:"sub-category-name"`
	RegionId string  `bson:"province-id"`
	RegionName string  `bson:"province-name"`
	Price float64  `bson:"price"`
	Images [] Image   `bson:"images"`
	IsNew bool  `bson:"is-new"`
	Currency string  `bson:"currency"`
	Created int64  `bson:"created"`
	Modified int64  `bson:"modified"`
	Location Location  `bson:"location"`
	Promoted bool  `bson:"promoted"`
	Likes int64  `bson:"likes"`
	Views int64  `bson:"views"`
	Score float64  `bson:"score"`
	Deleted bool  `bson:"deleted"`
}

type ProductOut struct {
	Id         string `bson:"_id,omitempty"`
	LocalId string `bson:"local-id"`
	UserId string  `bson:"user-id"`
	name string  `bson:"name"`
	description string  `bson:"description"`
	CategoryId string  `bson:"category-id"`
	Categoryame string  `bson:"category-name"`
	SubCategoryId string  `bson:"sub-category-id"`
	SubCategoryName string `bson:"sub-category-name"`
	RegionId string  `bson:"province-id"`
	RegionName string  `bson:"province-name"`
	Price float64  `bson:"price"`
	Images [] Image   `bson:"images"`
	IsNew bool  `bson:"is-new"`
	Currency string  `bson:"currency"`
	Created int64  `bson:"created"`
	Modified int64  `bson:"modified"`
	Location Location  `bson:"location"`
	Promoted bool  `bson:"promoted"`
	Likes int64  `bson:"likes"`
	Views int64  `bson:"views"`
	Score float64  `bson:"score"`
	Deleted bool  `bson:"deleted"`
}
