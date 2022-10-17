package database

import (
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const dbAddr = "localhost:9080"

// CloseFunc is used to pass the grpc.ClientConn Close() function out of newClient()
type CloseFunc func()

func NewClient() (*dgo.Dgraph, CloseFunc) {
	d, err := grpc.Dial(dbAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
			api.NewDgraphClient(d),
		), func() {
			if err := d.Close(); err != nil {
				log.Printf("Error while closing connection:%v", err)
			}
		}
}

// func setup(c *dgo.Dgraph) {
// 	err := c.Alter(context.Background(), &api.Operation{
// 		Schema: `
// 			username: string @index(term) @lang @upsert .
// 			email: string @index(exact) @upsert .
// 			password: password .
// 		`,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
