package database

import (
    "context"
    "log"
    
    "github.com/LucsOlv/Turtwing_Back/ent"
    _ "github.com/mattn/go-sqlite3"
)

func NewClient() *ent.Client {
    client, err := ent.Open("sqlite3", "file:turtwig.db?cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }

    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    return client
}
