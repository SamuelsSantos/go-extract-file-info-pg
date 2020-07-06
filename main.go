package main

import (
	"errors"
	"log"
	"os"

	"github.com/desafios-job/import-data/infraestructure/persistence"
	"github.com/desafios-job/import-data/service"
	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "navita"
	dbname   = "neoway"
	limit    = 65535 //Postgress limit parameters
)

func doProcess(filename string) {

	if len(filename) == 0 {
		log.Fatal(errors.New("Invalid input. "))
	}

	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()

	dw := service.DWNeoway{
		FileName:         filename,
		InconsistencyApp: service.NewInconsistencyApp(services.Inconsistency),
		ShoppingApp:      service.NewShoppingApp(services.Shopping),
		LayoutFile:       *service.NeowayLayout(),
	}

	er := dw.Clean()
	if er != nil {
		log.Fatal(er)
	}

	rows, err := dw.Extract()
	if err != nil {
		log.Fatal(err)
	}

	shopping, inconsistencies := dw.Transform(*rows)
	dw.SaveShoppings(shopping)
	dw.SaveInconsistencies(inconsistencies)
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal(errors.New("File parameter was not informed! Exemplo: ./resource/file.txt"))
	}

	doProcess(os.Args[1])
}
