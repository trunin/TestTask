package app

import (
	httpClient "TestTask/app/http"
	"TestTask/app/logger"
	"log"
	"net/http"
)

type Server struct {
	httpClient *httpClient.HttpClient
	log        *logger.Logger
}

func NewServer() Server {
	newLogger := logger.NewLogger()
	client := httpClient.NewHttpClient(&newLogger)
	return Server{
		httpClient: &client,
		log:        &newLogger,
	}
}

func (s *Server) Serve() {
	handler := http.HandlerFunc(s.MyHandler)
	http.Handle("/", s.requestLimits(handler))
	s.log.Info("listen and serve")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
