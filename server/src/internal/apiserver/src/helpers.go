package apiserver

import (
	"encoding/json"
	"net/http"
)

func (s *server) Err(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func Alg(req [][]int, g []good, p []pack, s *server) [][]int {

	for j := 0; j < len(g)-1; j++ {
		var swap good
		f := 0

		for i := 0; i < len(g)-j-1; i++ {
			if g[i].weight > g[i+1].weight {
				swap = g[i]
				g[i] = g[i+1]
				g[i+1] = swap

				f = 1
			}
		}

		if f == 0 {
			break
		}
	}

	s.Logger.Debug(g)

	for j := 0; j < len(p); j++ {
		for i := len(g) - 1; i > -1; i-- {
			for z := g[i].amount - j; z > 0; z -= len(p) {
				req = append(req, []int{p[j].id, g[i].id})
			}
		}
	}

	return req
}
