package model

type AccountIn struct {
	FacebookId string `bson:"facebook-id"`
	Token      string `bson:"token"`
	Name       string `bson:"name"`
	FirstName  string `bson:"first-name"`
	LastName   string `bson:"last-name"`
	Avatar     Image  `bson:"avatar"`
	Role       string `bson:"role"`
	Created    int64  `bson:"created"`
	Modified   int64  `bson:"modified"`
	Disabled   bool   `bson:"disabled"`
	Shop       Shop   `bson:"shop"`
}

type AccountOut struct {
	Id         string `bson:"_id,omitempty"`
	FacebookId string `bson:"facebook-id"`
	Token      string `bson:"token"`
	Name       string `bson:"name"`
	FirstName  string `bson:"first-name"`
	LastName   string `bson:"last-name"`
	Avatar     Image  `bson:"avatar"`
	Role       string `bson:"role"`
	Created    int64  `bson:"created"`
	Modified   int64  `bson:"modified"`
	Disabled   bool   `bson:"disabled"`
	Shop       Shop   `bson:"shop"`
}
