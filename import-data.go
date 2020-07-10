package main

import (
	"errors"
	"flag"
	"log"

	config "github.com/desafios-job/import-data/infraestructure/configs"
	"github.com/desafios-job/import-data/infraestructure/persistence"
	"github.com/desafios-job/import-data/service"
	_ "github.com/lib/pq"
)

func doProcess(filename string) {

	cfg := config.NewConfig()

	if len(filename) == 0 {
		log.Fatal(errors.New("Invalid input. "))
	}

	services, err := persistence.NewRepositories(*cfg.Db)

	if err != nil {
		panic(err)
	}

	defer services.Close()

	dw := service.NewNeowayDW(filename, *services)

	er := dw.Clean()
	if er != nil {
		log.Fatal(er)
	}

	rows, err := dw.Extract()
	if err != nil {
		log.Fatal(err)
	}

	shopping, inconsistencies := dw.Transform(rows)
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

	doProcess(fn)
}
