package apiserver

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
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

func Random(cin int) int {
	var pack int
	n, err := rand.Int(rand.Reader, big.NewInt(int64(cin)))
	if err != nil {
		pack = 0
	} else {
		pack = int(n.Int64())
	}

	return pack
}

func Alg(req [][]int, good int, amount int) [][]int {
	pack := (Random(1000) % 2) + 1

	req = append(req, []int{pack, good, amount})

	return req
}
