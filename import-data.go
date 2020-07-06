package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/desafios-job/import-data/infraestructure/persistence"
	"github.com/desafios-job/import-data/service"
	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	host     = "postgresdb"
	port     = "5432"
	user     = "postgres"
	password = "neoway"
	dbname   = "import-data"
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

	optFile := flag.String("file", "", "Source file")
	flag.Parse()
	fn := *optFile
	if fn == "" {
		log.Fatal(errors.New("File parameter was not informed! Exemplo: -file=./resource/file.txt"))
	}

	fmt.Println(fn)

	doProcess(fn)
}
