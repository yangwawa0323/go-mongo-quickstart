package my_testing

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongo-quickstart/utils"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_Connection(t *testing.T) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	dbc := &utils.DBContext{
		context,
	}

	defer cancel()

	client, err := dbc.ConnDB(nil)

	if err != nil {
		t.Fatal(err)
	}

	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"

	var result bson.M
	err = coll.FindOne(dbc.Context,
		bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		t.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", jsonData)
}
