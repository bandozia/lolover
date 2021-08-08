package handler

import (
	"net/http"

	"github.com/bandozia/lolover/src/fileservice"
)

func GetDir(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("not allowed"))
	}

	if ch, err := fileservice.RenderTop("."); err == nil {
		for fe := range ch {
			w.Write([]byte(fe.Name))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

}
