// Package server implements an example server.
package server

import (
	"fmt"
	"time"
)

// Server is an example server.
type Server struct {
	// now is a function returning a time that the server will consider the
	// current instant.
	now func() time.Time
}

// New constructs a new Server, which greets people based on the time returned
// by the given nowFunc. If nowFunc is nil, the Server will default to time.Now.
func New(nowFunc func() time.Time) *Server {
	return &Server{
		now: nowFunc,
	}
}

// Serve is an example function.
func (*Server) Serve() string {
	return "example serve command"
}

// Greet is an example fuzzable function.
// name is free form, but birthDate must be in YYYY-MM-DD format.
func (s *Server) Greet(name, birthDate string) (string, error) {
	b, err := time.ParseInLocation("2006-01-02", birthDate, time.Local)
	if err != nil {
		return "", fmt.Errorf("couldn't parse birthDate: %v", err)
	}
	now := s.now()
	if now.Before(b) {
		return "", fmt.Errorf("time travel detected")
	}
	// calculate time since birthday this year
	d := now.Sub(
		time.Date(now.Year(), b.Month(), b.Day(), 0, 0, 0, 0, time.Local))
	switch {
	case d < 0:
		// calculate time since birthday last year
		d = now.Sub(
			time.Date(now.Year()-1, b.Month(), b.Day(), 0, 0, 0, 0, time.Local))
		fallthrough
	case d > 0:
		return fmt.Sprintf("Hello %s, happy belated birthday for %d days ago.",
			name, int(d.Hours()/24)), nil
	default:
		return fmt.Sprintf("Hello %s, happy birthday!", name), nil
	}
}
