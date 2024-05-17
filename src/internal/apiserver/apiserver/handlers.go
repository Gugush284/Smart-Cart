package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *server) GetGoods() http.HandlerFunc {

	type request struct {
		First  int `json:"first"`
		Second int `json:"second"`
		Third  int `json:"third"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := strings.ReplaceAll(r.URL.Path, "/tg/", "")

		id, err := strconv.Atoi(key)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		str := -1
		for i := 0; i < len(cid); i++ {
			if cid[i].id == id {
				str = i
				break
			}
		}
		if str == -1 {
			str = len(cid)
			cid = append(cid, goods{0, 0, 0, false, 0})
			cid[str].id = id
		}

		if !cid[str].ready {
			req := &request{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				s.Err(w, r, http.StatusBadRequest, err)
				return
			}

			cid[str].first += req.First
			cid[str].second += req.Second
			cid[str].third += req.Third
		}

		s.Logger.Info(cid[str])

		s.respond(w, r, http.StatusAccepted, nil)
	})
}

func (s *server) ReadyTG() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := strings.ReplaceAll(r.URL.Path, "/ready/tg/", "")

		id, err := strconv.Atoi(key)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		str := -1
		for i := 0; i < len(cid); i++ {
			if cid[i].id == id {
				str = i
				break
			}
		}
		if str != -1 {
			cid[str].ready = true
		}

		s.Logger.Info(cid[str])

		s.respond(w, r, http.StatusAccepted, nil)
	})
}

func (s *server) Ready() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req [][]int

		key := strings.ReplaceAll(r.URL.Path, "/ready/", "")

		id, err := strconv.Atoi(key)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		str := -1
		for i := 0; i < len(cid); i++ {
			if cid[i].id == id {
				str = i
				break
			}
		}
		if str != -1 {
			if cid[str].ready {
				req = Alg(req, 1, cid[str].first)
				req = Alg(req, 2, cid[str].second)
				req = Alg(req, 3, cid[str].third)

				s.respond(w, r, http.StatusOK, req)

				//FIX delete order
			}
		}

		s.respond(w, r, http.StatusFound, nil)
	})
}

func (s *server) GetUsers() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id []int

		for i := 0; i < len(cid); i++ {
			if cid[i].ready {
				id = append(id, cid[i].id)
			}
		}

		s.respond(w, r, http.StatusFound, id)
	})
}
