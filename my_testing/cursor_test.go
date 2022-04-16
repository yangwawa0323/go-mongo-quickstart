package my_testing

import (
	"context"
	"fmt"
	"go-mongo-quickstart/utils"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const LIMIT int64 = 5

type Result struct {
	Title string `bson:"title"`
	Rated string `bson:"rated,omitempty"`
}

func Test_Cursor(t *testing.T) {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	dbc := &utils.DBContext{
		Context: context.TODO(),
	}

	// defer cancel()

	adOption := &options.ClientOptions{}

	client, err := dbc.ConnDB(adOption)
	if err != nil {
		t.Fatal(err)
	}

	defer client.Disconnect(dbc.Context)

	coll := client.Database("sample_mflix").Collection("movies")

	// SetProjection is controlled the output file,  by default the `_id` will visible.
	findOption := options.Find().SetLimit(LIMIT).SetProjection(bson.D{
		{"plot", 1}, {"rated", 1}, {"title", 1}, {"_id", 0},
	})

	// Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions)

	// method 1: filter by `bson.M`
	// filter := bson.M{
	// 	"title": bson.M{
	// 		"$regex": utils.Regex{Pattern: "civi", Options: "i"},
	// 	},
	// }

	// method 1: filter by `bson.D`

	filter := bson.D{
		{"title", bson.M{"$regex": utils.Regex{Pattern: "civi", Options: "i"}}},
	}

	cursor, err := coll.Find(dbc.Context, filter, findOption)
	if err != nil {
		t.Fatal(err)
	}

	defer cursor.Close(dbc.Context)

	// Result is a struct, mongo.Cursor can decode the data direct to the struct by `Decode()` function
	var res Result

	for cursor.Next(dbc.Context) {
		// var result bson.D
		if err := cursor.Decode(&res); err != nil {
			t.Fatal(err)
		}
		// bson.Unmarshal(result, res)
		fmt.Printf("Title: %s \nRated: %s\n", res.Title, res.Rated)
	}

	if err := cursor.Err(); err != nil {
		t.Fatal(err)
	}
}
