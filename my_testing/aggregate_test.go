package my_testing

import (
	"context"
	"fmt"
	"go-mongo-quickstart/utils"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
IMPORTANT: The field name of the struct if it is lowcase will not access,
FOR EXAMPLE: I been tried many time to troubleshooting the mongo.Cursor.Decode( &AggregateResult{})
*/
type AggregateResult struct {
	Id            string `bson:"_id"`
	TotalQuantity uint   `bson:"totalQuantity"`
}

func Test_Aggregate(t *testing.T) {

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

	coll := client.Database("myFirstDatabase").Collection("orders")

	matchStage := bson.D{
		{"$match", bson.D{
			{"size", "medium"},
		},
		},
	}

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$name"},
			{"totalQuantity", bson.D{{"$sum", "$quantity"}}},
		},
		},
	}

	// addFieldStage := bson.D{
	// 	{"$addFields", bson.D{
	// 		{"totalQuantity", bson.D{{"$sum", "$quantity"}}},
	// 	},
	// 	},
	// }

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage /*addFieldStage*/})
	if err != nil {
		t.Fatal(err)
	}

	// var results []AggregateResult = make([]AggregateResult, 0)
	// var results []bson.M
	// if err := cursor.All(context.TODO(), &results); err != nil {
	// 	t.Fatal(err)
	// }

	// for _, result := range results {

	// 	fmt.Printf("%+v\n", result)
	// }

	for cursor.Next(context.Background()) {
		result := &AggregateResult{}

		if err := cursor.Decode(result); err != nil {
			t.Fatal(err)
		}
		// append(results, result)
		fmt.Printf("%+v\n", *result)
	}
}
