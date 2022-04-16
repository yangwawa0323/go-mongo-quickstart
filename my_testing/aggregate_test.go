package my_testing

import (
	"context"
	"fmt"
	"go-mongo-quickstart/utils"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Test_Aggregate_Match(t *testing.T) {

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

/*
db.orders.aggregate( [
   // Stage 1: Filter pizza order documents by date range
   {
      $match:
      {
         "date": { $gte: new ISODate( "2020-01-30" ), $lt: new ISODate( "2022-01-30" ) }
      }
   },
   // Stage 2: Group remaining documents by date and calculate results
   {
      $group:
      {
         _id: { $dateToString: { format: "%Y-%m-%d", date: "$date" } },
         totalOrderValue: { $sum: { $multiply: [ "$price", "$quantity" ] } },
         averageOrderQuantity: { $avg: "$quantity" }
      }
   },
   // Stage 3: Sort documents by totalOrderValue in descending order
   {
      $sort: { totalOrderValue: -1 }
   }
 ] )
*/

func Test_Aggregate_Order(t *testing.T) {
	dbc := &utils.DBContext{
		Context: context.TODO(),
	}

	adOption := &options.ClientOptions{}

	client, err := dbc.ConnDB(adOption)
	if err != nil {
		t.Fatal(err)
	}

	defer client.Disconnect(dbc.Context)

	coll := client.Database("myFirstDatabase").Collection("orders")

	since, _ := time.Parse("2006-01-02", "2022-01-01")
	util, _ := time.Parse("2006-01-02", "2022-01-30")

	matchStage := bson.D{
		{"$match", bson.D{
			{
				"date", bson.D{
					{"$gte", primitive.NewDateTimeFromTime(since)},
					{"$lt", primitive.NewDateTimeFromTime(util)},
				}},
		}}}

	groupStage := bson.D{
		{"$group", bson.D{
			{
				"_id", bson.D{
					{"$dateToString", bson.D{
						{"format", "%Y-%m-%d"},
						{"date", "$date"},
					}},
				}},
			{
				"totalOrderValue", bson.D{
					{"$sum", bson.D{
						{"$multiply", bson.A{"$price", "$quantity"}},
					},
					},
				}},
		}}}

	sortStage := bson.D{
		{"$sort", bson.D{{"totalOrderValue", -1}}},
	}

	cursor, err := coll.Aggregate(context.Background(),
		mongo.Pipeline{matchStage, groupStage, sortStage})

	if err != nil {
		t.Fatal(err)
	}

	var results []bson.M
	err = cursor.All(context.Background(), &results)
	if err != nil {
		t.Fatal(err)
	}

	for _, res := range results {
		fmt.Printf("%+v\n", res)
	}
}
