package main

import "fmt"

type Database interface {
    Connect()
}

type Mongo struct{}

func (db Mongo) Connect() {
    fmt.Println("Conectado ao Mongo")
}

func NewMongo() Database{
	return Mongo{}
}

type Service struct {
    db Database
}

func (s *Service) InjectDatabase(db Database) {
    s.db = db
}

func main() {
    var service Service

    db := NewMongo()

    service.InjectDatabase(db)
    service.db.Connect()
}