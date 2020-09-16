package model

type ConversationIn struct {
	fromUser      string `bson:"from-user"`
	withUser      string `bson:"with-user"`
	fromUserLeave bool   `bson:"from-user-leave"`
	withUserLeave bool   `bson:"from-user-leave"`
	created       int64  `bson:"created"`
	modified      int64  `bson:"modified"`
}

type ConversationOut struct {
	Id            string `bson:"_id,omitempty"`
	fromUser      string `bson:"from-user"`
	withUser      string `bson:"with-user"`
	fromUserLeave bool   `bson:"from-user-leave"`
	withUserLeave bool   `bson:"from-user-leave"`
	created       int64  `bson:"created"`
	modified      int64  `bson:"modified"`
}
