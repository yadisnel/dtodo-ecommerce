package model

type MessageIn struct {
	UserOwnerId string  `bson:"user-owner-id"`
	ConversationId string `bson:"conversation-id"`
	FromUserId string `bson:"from-user-id"`
	FromUserName string `bson:"from-user-name"`
	ToUserId string `bson:"to-user-id"`
	ToUserName string `bson:"to-user-name"`
	Body string `bson:"body"`
	Created int64 `bson:"created"`
	Modified int64 `bson:"modified"`
	Deleted bool `bson:"deleted"`
}

type MessageOut struct {
	Id         string `bson:"_id,omitempty"`
	UserOwnerId string  `bson:"user-owner-id"`
	ConversationId string `bson:"conversation-id"`
	FromUserId string `bson:"from-user-id"`
	FromUserName string `bson:"from-user-name"`
	ToUserId string `bson:"to-user-id"`
	ToUserName string `bson:"to-user-name"`
	Body string `bson:"body"`
	Created int64 `bson:"created"`
	Modified int64 `bson:"modified"`
	Deleted bool `bson:"deleted"`
}