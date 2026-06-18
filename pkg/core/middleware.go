package core

import "net/http"

type HandlerFunc func(*Context) error

func (s *Service) Guest(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{w: w, r: r}
		if err := fn(ctx); err != nil {
			s.log.Error("Failed to handle request", err)
		}
	}
}

func (s *Service) Auth(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{w: w, r: r}
		user, err := s.auth(ctx)
		if err != nil {
			s.log.Error("Failed to authenticate request", err)
			return
		}
		ctx.u = user
		err = fn(ctx)
		if err != nil {
			s.log.Error("Failed to handle request", err)
		}
	}
}
