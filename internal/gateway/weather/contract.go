package weather

import "net/http"

type httpClient interface {
	Do(request *http.Request) (*http.Response, error)
}
