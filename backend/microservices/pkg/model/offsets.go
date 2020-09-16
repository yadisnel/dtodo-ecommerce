package model

type Offsets struct {
	Countries   int64 `bson:"countries"`
	Regions     int64 `bson:"regions"`
	Categories int64 `bson:"categories"`
	Words      int64 `bson:"words"`
}
