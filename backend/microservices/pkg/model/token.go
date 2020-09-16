package model

type Token struct {
	AccessToken string `bson:"access-token"`
	TokenType      string `bson:"token-type"`
}