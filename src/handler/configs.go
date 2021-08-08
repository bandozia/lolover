package handler

import "net/http"

func GetConfigs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("soon"))
}
