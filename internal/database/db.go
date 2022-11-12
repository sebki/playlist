package database

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const dbAddr = "localhost:9080"

type db struct {
	Client *dgo.Dgraph
	Closer CloseFunc
}

var Database db

// CloseFunc is used to pass the grpc.ClientConn Close() function out of newClient()
type CloseFunc func()

func NewClient() db {
	d, err := grpc.Dial(dbAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return db{
		Client: dgo.NewDgraphClient(api.NewDgraphClient(d)),
		Closer: func() {
			if err := d.Close(); err != nil {
				log.Printf("Error while closing connection:%v", err)
			}
		},
	}
}

func (db *db) Setup() {
	err := db.Client.Alter(context.Background(), &api.Operation{
		Schema: schema,
	})
	if err != nil {
		log.Fatal(err)
	}
}
