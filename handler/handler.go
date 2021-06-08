package handler

import (
	"bakuvi/store"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)




type Service struct {
	S *store.Service
}

//func (h *Service) GetIDs(w http.ResponseWriter, r *http.Request){
//	keys, ok := r.URL.Query()["id"]
//
//	if !ok || len(keys[0]) < 1 {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte("Url Param 'id' is missing"))
//		return
//	}
//
//	dateFrom := keys[0]
//	dateTo:= keys[1]
//
//
//	us, err := h.S.Get(dateFrom, dateTo)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte(fmt.Sprintf("can't get elem: %v", err)))
//		return
//	}
//
//	resp, err := json.Marshal(us)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(fmt.Sprintf("can't marshal result: %v", err)))
//		return
//	}
//
//	w.Write(resp)
//	return
//}

func (h *Service) GetIDs(w http.ResponseWriter, r *http.Request) {

	v, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("can't read requuest data: %v", err)))
		return
	}

	var user store.GetUserDate
	if err := json.Unmarshal(v, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("can't unmarshal requuest data: %v", err)))
		return
	}
 dateFrom, err:=time.Parse("2006-01-02", user.DateFrom)
 if err!=nil{
	 w.WriteHeader(http.StatusBadRequest)
	 w.Write([]byte(fmt.Sprintf("can't unmarshal requuest data: %v", err)))
	 return
 }

 dateTo, err:=time.Parse("2006-01-02", user.DateTo)
	id, err := h.S.Get(dateFrom, dateTo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("can't add data: %v", err)))
		return
	}
	fmt.Println(id)
	w.Write([]byte(fmt.Sprintf("%v", id)))
	return
}
