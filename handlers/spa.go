package handlers

import (
	"net/http"
	"os"

	"path/filepath"
)

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

func (spa *SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path, err := filepath.Abs(r.URL.Path)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(spa.StaticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(spa.StaticPath, spa.IndexPath))
		return
	} else if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(spa.StaticPath)).ServeHTTP(w, r)
}
