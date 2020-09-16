package model

type RegionIn struct {
	Name     string `bson:"name"`
	NOrder   string `bson:"n-order"`
	created  string `bson:"created"`
	modified string `bson:"modified"`
	deleted  string `bson:"deleted"`
}

type RegionOut struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	NOrder   string `bson:"n-order"`
	created  string `bson:"created"`
	modified string `bson:"modified"`
	deleted  string `bson:"deleted"`
}
