package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// RenderJSON writes an http response using the value passed in v as JSON.
func RenderJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("render_error:%s", errString(err))
	} else {
		w.WriteHeader(code)
	}
	_, err := w.Write(b.Bytes())
	if err != nil{
		log.Printf("render_error:%s", errString(err))
	}
}

func RenderNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

type ErrResponse struct {
	Status  string `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
	ErrorID string `json:"error_id,omitempty"`
}

func RenderErrNotFound(w http.ResponseWriter) {
	RenderJSON(w, http.StatusNotFound, ErrResponse{Status: "not found", Error: "not found"})
}

func RenderErrResourceNotFound(w http.ResponseWriter, resource string) {
	RenderJSON(w, http.StatusNotFound, ErrResponse{Status: resource + " not found", Error: resource + " not found"})
}

func RenderErrUnauthorized(w http.ResponseWriter) {
	RenderJSON(w, http.StatusUnauthorized, ErrResponse{Status: "not authorized", Error: "not authorized"})
}

func RenderErrInvalidRequest(w http.ResponseWriter, err error) {
	RenderJSON(w, http.StatusBadRequest, ErrResponse{Status: "invalid request", Error: errString(err)})
}

func RenderErrInternal(w http.ResponseWriter, err error) {
	RenderJSON(w, http.StatusInternalServerError, ErrResponse{Status: "internal error", Error: errString(err)})
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func DecodeJSON(r io.Reader, v interface{}) error {
	defer func(dst io.Writer, src io.Reader) {
		_, err := io.Copy(dst, src)
		if err != nil {
			log.Printf("render_error:%s", errString(err))
		}
	}(ioutil.Discard, r)
	return json.NewDecoder(r).Decode(v)
}
