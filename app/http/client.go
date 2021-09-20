package http

import (
	"TestTask/app/logger"
	"TestTask/app/workers"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const maxProcess = 4

type HttpClient struct {
	log *logger.Logger
}

func NewHttpClient(log *logger.Logger) HttpClient {
	return HttpClient{log: log}
}

func (h *HttpClient) ProcessUrls(ctx context.Context, urls []string) ([]string, error) {
	var result []string
	resultCh := make(chan []byte)
	errCh := make(chan error)
	pool := workers.NewPool(maxProcess, h.getUrl)
	for _, url := range urls {
		pool.Do(ctx, url, resultCh, errCh)
	}

	for range urls {
		select {
		case <-ctx.Done():
			h.log.Warn("context was canceled")
			return nil, fmt.Errorf("context was canceled")
		case r := <-resultCh:
			result = append(result, string(r))
		case err := <-errCh:
			h.log.Error("failed while processing url")
			return nil, err
		}
	}

	return result, nil

}

func (h *HttpClient) getUrl(ctx context.Context, url string, outCh chan []byte, errCh chan error) {
	var decoded map[string]interface{}
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	r, err := client.Get(url)
	if err != nil {
		errCh <- err
		h.log.Error("error while request")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&decoded); err != nil {
		h.log.Error("error unmarshal request")
		errCh <- err
		return
	}
	result, err := json.MarshalIndent(decoded, "", "  ");
	if err != nil {
		h.log.Error("error indent request")
		errCh <- err
		return
	}

	select {
	case <-ctx.Done():
		h.log.Warn("context was canceled")
		return
	case outCh <- result:
		h.log.Info("response received")
		return
	}
}
