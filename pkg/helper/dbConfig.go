package helper

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

var Database *memdb.MemDB

func InitDb() {
	// Create a new data base
	db, err := memdb.NewMemDB(createSchema())
	if err != nil {
		panic(err)
	}
	Database = db
	populateDbWithDefaultValues()
}

func GetDatabase() *memdb.MemDB {
	return Database
}

func createSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": {
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UUIDFieldIndex{Field: "Id"},
					},
					"firstName": {
						Name:    "firstName",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "FirstName"},
					},
					"lastName": {
						Name:    "lastName",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "LastName"},
					},
					"email": {
						Name:    "email",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"userTickets": {
						Name:    "userTickets",
						Unique:  false,
						Indexer: &memdb.UintFieldIndex{Field: "UserTickets"},
					},
				},
			},
		},
	}
}

func populateDbWithDefaultValues() {
	bookings := []*User{
		{uuid.NewString(), "Szym", "Tru", "szym@tru", 2},
	}

	transaction := GetDatabase().Txn(true)
	for _, p := range bookings {
		if err := transaction.Insert("user", p); err != nil {
			panic(err)
		}
	}
	transaction.Commit()
}
