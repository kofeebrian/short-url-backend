package models

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/kofeebrian/short-url-server/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client = db.DBConnect()

var collection *mongo.Collection
var ctx = context.TODO()

type ShortenedUrl struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Url          string             `json:"url" bson:"url"`
	ShortenedUrl string             `json:"shortened_url" bson:"shortened_url"`
	CreateAt     time.Time          `json:"create_at" bson:"create_at"`
	UpdateAt     time.Time          `json:"update_at" bson:"update_at"`
}

func CreateShortenedUrl(url string) error {

	md5_hash := md5.Sum([]byte(url))
	hash := hex.EncodeToString(md5_hash[:])

	// Get urls collections
	collection = client.Database("short-url").Collection("urls")

	var shortenedUrl ShortenedUrl
	// Else, create a new one
	shortenedUrl = ShortenedUrl{
		Url:          url,
		ShortenedUrl: "http://localhost:8080/" + hash[:7], // Just do a simple slice for now
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
	}

	filter := bson.D{{Key: "shortened_url", Value: shortenedUrl.ShortenedUrl}}
	update := bson.D{{Key: "$set", Value: shortenedUrl}}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)

	return err
}

func GetShortenedUrl(url string) (ShortenedUrl, error) {
	var shortenedUrl ShortenedUrl

	err := collection.FindOne(context.TODO(), bson.M{"url": url}).Decode(&shortenedUrl)

	return shortenedUrl, err
}
