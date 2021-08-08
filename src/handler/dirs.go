package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bandozia/lolover/src/fileservice"
)

type dirReq struct {
	Path string `json:"path"`
}

func GetDir(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var dr dirReq
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		err = json.Unmarshal(body, &dr)
		if err != nil {
			WriteErr(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	if ch, err := fileservice.RenderTop(dr.Path); err == nil {
		dirs := []fileservice.FileEntity{}
		for fe := range ch {
			dirs = append(dirs, fe)
		}
		WriteJson(w, dirs)
	} else {
		WriteErr(w, http.StatusBadRequest, err.Error())
		return
	}
}
