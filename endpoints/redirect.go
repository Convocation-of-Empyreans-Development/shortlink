package endpoints

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RouteRequest(redirects map[string]string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		path := params.ByName("path")
		if redirectUrl, ok := redirects[path]; ok {
			http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
