package dependencyinjecttion

import "fmt"

type Database interface {
    Connect()
}

type PostgreSQL struct{}

func (db PostgreSQL) Connect() {
    fmt.Println("Conectado ao PostgreSQL")
}

type Service struct {
    db Database
}

func (s *Service) SetDatabase(db Database) {
    s.db = db
}

func main() {
    db := PostgreSQL{}
    var service Service
    service.SetDatabase(db)
    service.db.Connect()
}