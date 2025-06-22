package apiserver

import (
	"LongTaskAPI/internal/core/http/handlers"
	"LongTaskAPI/internal/services"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	cfg     *Config
	router  *mux.Router
	logger  *logrus.Logger
	service *services.TaskService
}

func New(cfg *Config, service *services.TaskService) *APIServer {
	return &APIServer{
		cfg:     cfg,
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		service: service,
	}
}

func (s *APIServer) Run() error {
	s.configureRouter()

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Infof("Starting API server on %s", s.cfg.BindAddress)
	return http.ListenAndServe(s.cfg.BindAddress, s.router)
}

func (s *APIServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(s.cfg.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(lvl)
	return nil
}

func (s *APIServer) configureRouter() {
	handler := handlers.CreateHandler(s.logger, s.service)

	s.router.HandleFunc("/tasks/create", handler.CreateNewTasksHandler()).Methods("POST")
	s.router.HandleFunc("/tasks/", handler.GetAllTasksHandler()).Methods("GET")
	s.router.HandleFunc("/tasks/{id}", handler.GetTaskHandler()).Methods("GET")
	s.router.HandleFunc("/tasks/{id}", handler.DeleteTaskHandler()).Methods("DELETE")
}
