package utils

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadURI() string {
	cwd, _ := os.Getwd()
	envFile := path.Join(cwd, ".env")

	if err := godotenv.Load(envFile); err != nil {
		log.Println("No .env file found.")
		return ""
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	return uri
}

type DBContext struct {
	Context context.Context
}

func (dbc *DBContext) ConnDB(adOption *options.ClientOptions) (*mongo.Client, error) {
	uri := LoadURI()

	clientOptions := options.Client().ApplyURI(uri)
	options.MergeClientOptions(clientOptions, adOption)
	// clientOptions.ApplyURI(uri)
	client, err := mongo.Connect(dbc.Context, clientOptions)

	if err != nil {
		return nil, err
	}
	return client, nil
}
