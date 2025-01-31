package server

import (
	"log"
	"net/http"
	"os"

	"github.com/YngviWarrior/microservice-exchange/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ServerInterface interface {
	InitServer(controller.ControllerInterface)
}

type server struct{}

func NewServer() ServerInterface {
	return &server{}
}

func (s *server) InitServer(controllers controller.ControllerInterface) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		Debug:            false,
	})

	r := mux.NewRouter()

	// r.HandleFunc("/exchange/balances/sync", controllersInterface.HandlerSyncBalances).Methods("GET")
	r.HandleFunc("/exchange", controllers.ListExchange).Methods("GET")
	// r.HandleFunc("/bot/health", controllersInterface.HandlerBotHealth).Methods("GET")
	// r.HandleFunc("/bot/sell-cotation", controllersInterface.HandlerBotHealth).Methods("GET")

	log.Printf("Running on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), c.Handler(r))

	if err != nil {
		log.Fatalf("%v", err)
	}
}
