package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrearcaina/wudb/kvstore"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	r chi.Router
	s *kvstore.Store
}

func NewServer(s *kvstore.Store) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/set", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.Set(req.Key, req.Value)
		w.WriteHeader(http.StatusCreated)

		err := json.NewEncoder(w).Encode(map[string]string{
			"key":   req.Key,
			"value": req.Value,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/get/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")

		value, ok := s.Get(key)
		if !ok {
			http.Error(w, "key not found", http.StatusNotFound)
			return
		}

		err := json.NewEncoder(w).Encode(map[string]string{"value": value})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Delete("/del/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		if err := s.Del(key); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	return &Server{
		r: r,
		s: s,
	}
}

func (s *Server) Start(addr string) error {
	fmt.Printf("Starting server on %s\n", addr)
	return http.ListenAndServe(addr, s.r)
}
