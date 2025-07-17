package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/compermane/gontabilizador/render"
	"github.com/compermane/gontabilizador/service/ensaio"
	"github.com/compermane/gontabilizador/service/presenca"
	"github.com/compermane/gontabilizador/service/ritmista"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	
	ritmistaStore   := ritmista.NewStore(s.db)
	ritmistaHandler := ritmista.NewHandler(ritmistaStore)
	ritmistaHandler.RegistroRoutes(subrouter)

	ensaioStore     := ensaio.NewStore(s.db)
	ensaioHandler   := ensaio.NewHandler(ensaioStore)
	ensaioHandler.RegistroRoutes(subrouter)

	presencaStore   := presenca.NewStore(s.db)
	presencaHandler := presenca.NewHandler(presencaStore)
	presencaHandler.PresencaRoutes(subrouter)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	renderHandler := render.NewPageHandler(ritmistaStore, ensaioStore, presencaStore)
	router.HandleFunc("/", renderHandler.Home).Methods("GET")
	router.HandleFunc("/ensaios", renderHandler.Ensaios).Methods("GET")
	router.HandleFunc("/presencas", renderHandler.Presencas).Methods("GET")

	log.Println("[APIServer] Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}