package rest

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func Handle(handler func(w http.ResponseWriter, r *http.Request) Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := handler(w, r)
		if response.Error != nil {
			err := response.Error

			logrus.WithFields(logrus.Fields{
				"type":        "request failed",
				"request":     r.RequestURI,
				"method":      r.Method,
				"status code": response.StatusCode,
			}).Error(err)
		}

		if h := response.Header(); h != nil {
			copyHeaders(w, h)
		}

		if err := sendJSON(w, response.Body, response.StatusCode); err != nil {
			logrus.WithFields(logrus.Fields{
				"type":    "request send json failed",
				"request": r.RequestURI,
				"method":  r.Method,
			}).Error(err)

			return
		}

		logrus.WithFields(logrus.Fields{
			"type":        "request success",
			"request":     r.RequestURI,
			"method":      r.Method,
			"status code": response.StatusCode,
		}).Info("success")
	}
}

func sendJSON(w http.ResponseWriter, payload any, statusCode int) error {
	if payload == nil {
		w.WriteHeader(statusCode)
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}

func copyHeaders(w http.ResponseWriter, h http.Header) {
	wh := w.Header()
	for header, values := range h {
		for _, value := range values {
			wh.Add(header, value)
		}
	}
}

func URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func QueryParams(r *http.Request) url.Values {
	return r.URL.Query()
}

func (r Response) WithError(err error) Response {
	r.Error = err
	return r
}

func (r Response) AddHeaders(headers map[string]string) Response {
	if len(headers) > 0 {
		for k, v := range headers {
			r = r.AddHeader(k, v)
		}
	}

	return r
}

func (r Response) AddHeader(name, value string) Response {
	if r.header == nil {
		r.header = http.Header{}
	}

	r.header.Add(name, value)

	return r
}

func (r Response) SetHeader(name, value string) Response {
	if r.header == nil {
		r.header = http.Header{}
	}

	r.header.Set(name, value)

	return r
}

func (r Response) Header() http.Header {
	return r.header
}
