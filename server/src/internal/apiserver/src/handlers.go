package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *server) GetGoods() http.HandlerFunc {

	type request struct {
		ID int `json:"id"`
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
		for i := 0; i < len(s.orders); i++ {
			if s.orders[i].id == id {
				str = i
				break
			}
		}
		if str == -1 {
			str = len(s.orders)
			s.orders = append(s.orders, order{item: []good{}, ready: false})
			s.orders[str].id = id
		}

		if !s.orders[str].ready {
			req := &request{}
			istr := -1
			sstr := -1

			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				s.Err(w, r, http.StatusBadRequest, err)
				return
			}

			for i := 0; i < len(s.store); i++ {
				if s.store[i].id == req.ID {
					sstr = i
					break
				}
			}
			if sstr == -1 {
				s.respond(w, r, http.StatusBadRequest, nil)
				return
			}

			for i := 0; i < len(s.orders[str].item); i++ {
				if s.orders[str].item[i].id == req.ID {
					istr = i
					break
				}
			}

			if istr == -1 {
				istr = len(s.orders[str].item)
				s.orders[str].item = append(s.orders[str].item, s.store[sstr])
				s.orders[str].item[istr].amount = 0
			}

			s.orders[str].item[istr].amount += 1
		}

		s.Logger.Info(s.orders[str])

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
		for i := 0; i < len(s.orders); i++ {
			if s.orders[i].id == id {
				str = i
				break
			}
		}
		if str != -1 {
			s.orders[str].ready = true
		}

		s.Logger.Info(s.orders[str])

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
		for i := 0; i < len(s.orders); i++ {
			if s.orders[i].id == id {
				str = i
				break
			}
		}
		if str != -1 {
			if s.orders[str].ready {
				for i := 0; i < len(s.orders[str].item); i++ {
					req = Alg(req, s.orders[str].item[i].amount, s.orders[str].item[i].id)
				}

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

		for i := 0; i < len(s.orders); i++ {
			if s.orders[i].ready {
				id = append(id, s.orders[i].id)
			}
		}

		s.respond(w, r, http.StatusFound, id)
	})
}
