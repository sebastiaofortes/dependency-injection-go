package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

type Database interface {
	Connect()
}

type Mongo struct{}

func (db Mongo) Connect() {
	fmt.Println("Conectado ao Mongo")
}

func NewMongo() Database {
	return Mongo{}
}

func main() {
	var service1 Service1
	var service2 Service2
	var service3 Service3
	var service4 Service4
	var service5 Service5
	var service6 Service6
	var service7 Service7
	var service8 Service8
	var service9 Service9
	var service10 Service10

    app := fx.New(
        fx.Provide(NewMongo),
        fx.Invoke(service1.InjectDatabase1),
        fx.Invoke(service2.InjectDatabase2),
        fx.Invoke(service3.InjectDatabase3),
        fx.Invoke(service4.InjectDatabase4),
        fx.Invoke(service5.InjectDatabase5),
        fx.Invoke(service6.InjectDatabase6),
        fx.Invoke(service7.InjectDatabase7),
        fx.Invoke(service8.InjectDatabase8),
        fx.Invoke(service9.InjectDatabase9),
        fx.Invoke(service10.InjectDatabase10),
    )

    app.Start(context.Background())
    defer app.Stop(context.Background())
    /*
    // Global status
	db := NewMongo()

	service1.InjectDatabase1(db)
    service2.InjectDatabase2(db)
    service3.InjectDatabase3(db)
    service4.InjectDatabase4(db)
    service5.InjectDatabase5(db)
    service6.InjectDatabase6(db)
    service7.InjectDatabase7(db)
    service8.InjectDatabase8(db)
    service9.InjectDatabase9(db)
    service10.InjectDatabase10(db)

    // Invoke contructor
    service1.InjectDatabase1(NewMongo())
    service2.InjectDatabase2(NewMongo())
    service3.InjectDatabase3(NewMongo())
    service4.InjectDatabase4(NewMongo())
    service5.InjectDatabase5(NewMongo())
    service6.InjectDatabase6(NewMongo())
    service7.InjectDatabase7(NewMongo())
    service8.InjectDatabase8(NewMongo())
    service9.InjectDatabase9(NewMongo())
    service10.InjectDatabase10(NewMongo())
    */
}

type Service1 struct {
    db Database
}

func (s *Service1) InjectDatabase1(db Database) {
    fmt.Println("Dependencia injetada no serviço 1")
    s.db = db
}

type Service2 struct {
    db Database
}

func (s *Service2) InjectDatabase2(db Database) {
    fmt.Println("Dependencia injetada no serviço 2")
    s.db = db
}

type Service3 struct {
    db Database
}

func (s *Service3) InjectDatabase3(db Database) {
    fmt.Println("Dependencia injetada no serviço 3")
    s.db = db
}

type Service4 struct {
    db Database
}

func (s *Service4) InjectDatabase4(db Database) {
    fmt.Println("Dependencia injetada no serviço 4")
    s.db = db
}

type Service5 struct {
    db Database
}

func (s *Service5) InjectDatabase5(db Database) {
    fmt.Println("Dependencia injetada no serviço 5")
    s.db = db
}

type Service6 struct {
    db Database
}

func (s *Service6) InjectDatabase6(db Database) {
    fmt.Println("Dependencia injetada no serviço 6")
    s.db = db
}

type Service7 struct {
    db Database
}

func (s *Service7) InjectDatabase7(db Database) {
    fmt.Println("Dependencia injetada no serviço 7")
    s.db = db
}

type Service8 struct {
    db Database
}

func (s *Service8) InjectDatabase8(db Database) {
    fmt.Println("Dependencia injetada no serviço 8")
    s.db = db
}

type Service9 struct {
    db Database
}

func (s *Service9) InjectDatabase9(db Database) {
    fmt.Println("Dependencia injetada no serviço 9")
    s.db = db
}

type Service10 struct {
    db Database
}

func (s *Service10) InjectDatabase10(db Database) {
    fmt.Println("Dependencia injetada no serviço 10")
    s.db = db
}