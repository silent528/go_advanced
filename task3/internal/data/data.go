package data

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"go_advanced/task3/internal/data/ent"
	"go_advanced/task3/internal/data/ent/migrate"
	"log"
)

var ProviderSet = wire.NewSet(NewData, NewEntClient, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
}

func NewEntClient() *ent.Client {
	client, err := ent.Open(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/test?parseTime=true&loc=Local",
		ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewData(entClient *ent.Client) (*Data, func(), error) {
	d := &Data{
		db: entClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Println(err)
		}
	}, nil
}
