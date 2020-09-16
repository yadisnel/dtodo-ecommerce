package model

type MqttPayload struct {
	Type string  `bson:"type"`
	Body string `bson:"Body"`
}