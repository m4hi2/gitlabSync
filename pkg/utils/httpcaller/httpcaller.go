package httpcaller

import "net/http"

type HTTPCaller interface {
	Do(req *http.Request) (*http.Response, error)
}
