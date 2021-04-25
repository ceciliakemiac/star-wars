package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	addr   string
	router *mux.Router
	db     *mongo.Collection
}

func NewServer(addr string, db *mongo.Collection) (*Server, error) {
	s := &Server{
		addr:   addr,
		db:     db,
		router: mux.NewRouter(),
	}

	s.CreateRoutes(s.router, s.db)

	return s, nil
}

func (s *Server) Run() error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD"},
		AllowedHeaders: []string{"*"},
	})

	log.Println("Http Server starting to listen at", s.addr)
	err := http.ListenAndServe(s.addr, c.Handler(s.router))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) CreateRoutes(router *mux.Router, db *mongo.Collection) {
	basePath := router.PathPrefix("/api").Subrouter()

	basePath.Path("").HandlerFunc(s.Ping).Methods(http.MethodGet)
	basePath.Path("/planet").HandlerFunc(s.GetPlanets).Methods(http.MethodGet)
	basePath.Path("/planet/nome/{nome}").HandlerFunc(s.GetPlanetByName).Methods(http.MethodGet)
	basePath.Path("/planet/{id}").HandlerFunc(s.GetPlanetByID).Methods(http.MethodGet)
	basePath.Path("/planet").HandlerFunc(s.CreatePlanet).Methods(http.MethodPost)
	basePath.Path("/planet/{id}").HandlerFunc(s.DeleteBook).Methods(http.MethodDelete)
}

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("Hello from Star Wars backend server!")
	w.Write(res)
}
