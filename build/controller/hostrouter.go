package controller

import (
	"net/http"
	"strings"
)

func (h *BaseHandler) HostRouter(w http.ResponseWriter, r *http.Request) {

	//log.Println(r.URL.Path)
	http.FileServer(http.Dir("www/"+strings.TrimLeft(r.Host, "www."))).ServeHTTP(w, r)
	//http.ServeFile(w, r, "www")
}
