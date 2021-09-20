package app

import (
	"TestTask/app/entities"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const maxBodyLength = 20

func (s *Server) MyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	switch r.Method {
	case http.MethodPost:
		listUrl := entities.ListUrlRequest{}

		err := json.NewDecoder(r.Body).Decode(&listUrl)
		if err != nil {
			s.log.Error("invalid data format")
			http.Error(w, "invalid data format", http.StatusBadRequest)
			return
		}

		if len(listUrl) > maxBodyLength {
			s.log.Error("request body too long", )
			http.Error(w, "request body too long", http.StatusBadRequest)
			return
		}

		result, err := s.httpClient.ProcessUrls(ctx, listUrl)
		if err != nil {
			s.log.Error("failed while processing url")
			http.Error(w, fmt.Sprintf("error while processing url :'%s'", err.Error()),
				http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "result : %v", result)
	default:
		fmt.Fprintf(w, "%s - method not allowed", r.Method)

	}
}
