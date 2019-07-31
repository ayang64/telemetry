package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	db *sql.DB
}

func WithDB(db *sql.DB) func(*Server) error {
	return func(s *Server) error {
		s.db = db
		return nil
	}
}

func New(opts ...func(*Server) error) (*Server, error) {
	rc := Server{}
	for _, opt := range opts {
		if err := opt(&rc); err != nil {
			return nil, err
		}
	}
	return &rc, nil
}

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func errorf(w http.ResponseWriter, r *http.Request, status int, f string, v ...interface{}) error {
	payload := struct {
		Error string `json:"error"`
	}{
		Error: fmt.Sprintf(f, v...),
	}

	w.WriteHeader(status)

	if err := send(w, r, payload); err != nil {
		return err
	}

	return nil
}

func send(w http.ResponseWriter, r *http.Request, v interface{}) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(v); err != nil {
		return err
	}
	return nil
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) (*User, error) {
	row := s.db.QueryRow(`select id, email from "user" where email=$1`, email)
	rc := User{}
	if err := row.Scan(&rc.Id, &rc.Email); err != nil {
		return nil, err
	}
	return &rc, nil
}

func (s *Server) CreateUser(email string) (*User, error) {
	row := s.db.QueryRow(`select * from CreateUser($1)`, email)
	rc := User{}
	if err := row.Scan(&rc.Id, &rc.Email); err != nil {
		return nil, err
	}
	return &rc, nil
}

func (s *Server) ApiUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		u, err := s.GetUser(r.URL.Query()["email"][0])
		if err != nil {
			errorf(w, r, http.StatusOK, "%v", err)
			return
		}
		send(w, r, u)
	case "POST":
	default:
	}
}

type Token struct {
	Token string
}

func (s *Server) GetDevice(token string, serial string) (*User, error) {
	return nil, nil
}

func (s *Server) ApiDevice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		u, err := s.GetUser(r.URL.Query()["email"][0])
		if err != nil {
			errorf(w, r, http.StatusOK, "%v", err)
			return
		}
		send(w, r, u)
	case "POST":
	default:
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api/v1/") {
		// return an error
	}
	switch strings.TrimPrefix(r.URL.Path, "/api/v1") {
	case "user":
	case "authkey":
	case "authtoken":
	case "device":
	}
}
