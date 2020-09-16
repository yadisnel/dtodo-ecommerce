package model

type GlobCountryIn struct {
	Name        string `bson:"name"`
	LangIsoCode string `bson:"lang-iso-code"`
}

type GlobCountryOut struct {
	Id          string `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	LangIsoCode string `bson:"lang-iso-code"`
}
