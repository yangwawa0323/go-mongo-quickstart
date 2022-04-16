package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongo-quickstart/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	dbc := &utils.DBContext{
		context,
	}

	defer cancel()

	adOption := &options.ClientOptions{}

	client, err := dbc.ConnDB(adOption)

	if err != nil {
		panic(err)
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
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
