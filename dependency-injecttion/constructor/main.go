package main

import "fmt"

type MySQL struct{}

func (db MySQL) Connect() {
    fmt.Println("Conectado ao MySQL")
}

type Service struct {
    db MySQL
}

func NewService(db MySQL) *Service {
    return &Service{
        db: db,
    }
}

func main() {
    db := MySQL{}
    service := NewService(db)
    service.db.Connect()
}