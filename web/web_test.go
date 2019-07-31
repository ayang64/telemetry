package web

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func NewServer() (*Server, error) {
	db, err := sql.Open("postgres", "user=ayan database=telemetry host=/tmp")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return New(WithDB(db))
}

func TestUser(t *testing.T) {
	s, err := NewServer()
	if err != nil {
		t.Fatal(err)
	}

	s.db.Exec(`delete from "user" where email='ayan@ayan.net';`)
	t.Run("Create", func(t *testing.T) {
		u, err := s.CreateUser("ayan@ayan.net")
		if err != nil {
			t.Fatal(err)
		}
		if expected, got := "ayan@ayan.net", u.Email; got != expected {
			t.Fatalf("expected returned email to be %q; got %q", expected, got)
		}
	})

	t.Run("Get", func(t *testing.T) {
		u, err := s.GetUser("ayan@ayan.net")
		if err != nil {
			t.Fatal(err)
		}
		if expected, got := "ayan@ayan.net", u.Email; got != expected {
			t.Fatalf("expected returned email to be %q; got %q", expected, got)
		}
	})

}
