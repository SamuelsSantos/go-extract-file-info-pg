package main

import (
	"errors"
	"flag"
	"log"

	"github.com/desafios-job/import-data/infraestructure/config"
	"github.com/desafios-job/import-data/infraestructure/persistence"
	"github.com/desafios-job/import-data/service"
	_ "github.com/lib/pq"
)

func doProcess(filename string) {

	conf := config.GetConf()

	if len(filename) == 0 {
		log.Fatal(errors.New("Invalid input. "))
	}

	services, err := persistence.NewRepositories(
		conf.DbDriver,
		conf.DbUser,
		conf.DbPassword,
		conf.DbPort,
		conf.DbHost,
		conf.DbName,
	)

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

	doProcess(fn)
}
