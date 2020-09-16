package model

type GlobWordIn struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
	LangIsoCode string `bson:"lang-iso-code"`
}

type GlobWordOut struct {
	Id          string `bson:"_id,omitempty"`
	Key         string `bson:"key"`
	Value       string `bson:"value"`
	LangIsoCode string `bson:"lang-iso-code"`
}