package my_testing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const URI string = "mongodb+srv://cluster0.ct3kn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

// const PLAIN_AUTH_URI string = "mongodb+srv://yangwawa:%40dmin2oo9ykun@cluster0.ct3kn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func Test_SetAuth(t *testing.T) {
	auth := &options.Credential{
		// AuthMechanism: "MONGODB-AWS",
		// AuthSource:    "",
		Username: "yangkun",
		Password: "yangkun_qq.com",
	}

	opts := options.Client().ApplyURI(URI).SetAuth(*auth)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		timeoutContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := client.Disconnect(timeoutContext)
		if err != nil {
			t.Fatalf("Disconnect Err: %v\n", err)
		}
	}()

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	coll := client.Database("sample_mflix").Collection("movies")

	filter := bson.D{}
	// coll := client.Database("sample_training").Collection("zips")
	// filter := bson.D{{"pop", bson.D{{"$lte", 500}}}}

	findOpt := options.Find().SetLimit(10).SetProjection(bson.D{
		{"title", 1}, {"type", 1}, {"_id", 0},
	})

	cursor, err := coll.Find(context.TODO(), filter, findOpt)

	if err != nil {
		t.Fatal(err)
	}

	var results []bson.D
	cursor.All(context.TODO(), &results)

	for _, result := range results {
		fmt.Printf("%v\n", result)
	}

}
