package models

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	gourl "net/url"
	"os"
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
	LongUrl      string             `json:"url" bson:"url"`
	ShortenedUrl string             `json:"shortened_url" bson:"shortened_url"`
	CreateAt     time.Time          `json:"create_at" bson:"create_at"`
	UpdateAt     time.Time          `json:"update_at" bson:"update_at"`
}

func CreateShortenedUrl(longUrl string) error {

	md5_hash := md5.Sum([]byte(longUrl))
	hash := hex.EncodeToString(md5_hash[:])

	// Get urls collections
	collection = client.Database("short-url").Collection("urls")

	var shortenedUrl ShortenedUrl
	// Else, create a new one
	s, err := gourl.JoinPath(os.Getenv("SERVER_NAME"), hash[:7])
	if err != nil {
		return err
	}
	log.Printf("%s: %s", os.Getenv("SERVER_NAME"), s)

	shortenedUrl = ShortenedUrl{
		LongUrl:      longUrl,
		ShortenedUrl: s, // Just do a simple slice for now
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
	}

	filter := bson.D{{Key: "shortened_url", Value: shortenedUrl.ShortenedUrl}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "update_at", Value: time.Now()}}}}
	opts := options.Update().SetUpsert(true)

	if err = collection.FindOne(context.TODO(), filter).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, err = collection.InsertOne(context.TODO(), shortenedUrl)
			return nil
		} else {
			return err
		}
	}

	_, err = collection.UpdateOne(ctx, filter, update, opts)

	return err
}

func GetShortenedUrl(url string) (ShortenedUrl, error) {
	collection = client.Database("short-url").Collection("urls")

	var shortenedUrl ShortenedUrl
	err := collection.FindOne(context.TODO(), bson.M{"url": url}).Decode(&shortenedUrl)

	return shortenedUrl, err
}

func GetOriginalUrl(url string) (ShortenedUrl, error) {

	collection = client.Database("short-url").Collection("urls")

	var shortenedUrl ShortenedUrl
	err := collection.FindOne(context.TODO(), bson.M{"shortened_url": url}).Decode(&shortenedUrl)
	return shortenedUrl, err
}
