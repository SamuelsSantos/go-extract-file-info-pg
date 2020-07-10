package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/desafios-job/import-data/infraestructure/configs"
	"github.com/desafios-job/import-data/infraestructure/persistence"
	"github.com/desafios-job/import-data/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

// App handle rest
type App struct {
	InconsistencyApp service.InconsistencyAppInterface
	ShoppingApp      service.ShoppingAppInterface
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%+v\n", payload)
}

type handlers interface {
	GetShoppings(w http.ResponseWriter, r *http.Request)
	GetInconsistencies(w http.ResponseWriter, r *http.Request)
}

// GetShoppings handle
func (app *App) GetShoppings(w http.ResponseWriter, r *http.Request) {

	result, err := app.ShoppingApp.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondwithJSON(w, http.StatusOK, result)

}

// GetInconsistencies handle
func (app *App) GetInconsistencies(w http.ResponseWriter, r *http.Request) {

	result, err := app.InconsistencyApp.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondwithJSON(w, http.StatusOK, result)

}

func newRouter(app *App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/shoppings", app.GetShoppings)
	r.Get("/inconsistencies", app.GetInconsistencies)

	return r
}

func main() {

	cfg := config.NewConfig()

	services, err := persistence.NewRepositories(*cfg.Db)

	if err != nil {
		panic(err)
	}
	defer services.Close()

	app := &App{
		InconsistencyApp: service.NewInconsistencyService(services.Inconsistency),
		ShoppingApp:      service.NewShoppingService(services.Shopping),
	}

	address := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server 0.0.0.0%s\n", address)
	log.Println(cfg.Db.ToString())

	http.ListenAndServe(address, newRouter(app))
}
