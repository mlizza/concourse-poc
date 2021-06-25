package internal

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type Server struct {
	Port   string
	Router *httprouter.Router
}

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Infof("Requested: %s\n", r.URL.Path)
	_, _ = fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func (s Server) Init() *httprouter.Router {
	s.Router = httprouter.New()
	s.Router.GET("/", IndexHandler)
	return s.Router
}
