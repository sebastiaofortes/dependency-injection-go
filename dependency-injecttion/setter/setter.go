package dependencyinjecttion

import "fmt"

type PostgreSQL struct{}

func (db PostgreSQL) Connect() {
    fmt.Println("Conectado ao PostgreSQL")
}

type Service struct {
    db PostgreSQL
}

func (s *Service) SetDatabase(db PostgreSQL) {
    s.db = db
}

func main() {
    db := PostgreSQL{}
    var service Service
    service.SetDatabase(db)
    service.db.Connect()
}