package my_testing

import (
	"context"
	"fmt"
	"go-mongo-quickstart/utils"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBCollection struct {
	Connect *mongo.Client
	Context context.Context
	T       *testing.T
}

var dbclient *DBCollection

func init() {

	ctx := context.TODO()
	dbc := &utils.DBContext{
		ctx,
	}

	adOption := &options.ClientOptions{}
	client, _ := dbc.ConnDB(adOption)

	dbclient = &DBCollection{
		Connect: client,
		Context: ctx,
	}
}

func Test_FindOne(t *testing.T) {

	var title string = "The Room"

	dbclient.T = t

	err := dbclient.FindTitle(title)

	defer dbclient.Connect.Disconnect(dbclient.Context)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			t.Logf("No found `%s\n`", title)
		}
		t.Fatal(err)
	}

}

func Test_Not_Found(t *testing.T) {

	var title string = "The room"

	dbclient.T = t

	err := dbclient.FindTitle(title)

	defer dbclient.Connect.Disconnect(dbclient.Context)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			t.Logf("No found `%s\n`", title)
		}
		t.Fatal(err)
	}
}

func (c *DBCollection) FindTitle(title string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	dbc := &utils.DBContext{
		Context: ctx,
	}

	defer cancel()

	client, err := dbc.ConnDB(&options.ClientOptions{})

	if err != nil {
		c.T.Fatal(err)
	}
	coll := client.Database("sample_mflix").Collection("movies")

	var result bson.M
	// The `R`oom
	err = coll.FindOne(ctx, bson.D{{"title", title}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.T.Log("No document finded.\n")
		}
		c.T.Fatal(err)
	}

	fmt.Printf("%#v\n", result)
	return nil
}
