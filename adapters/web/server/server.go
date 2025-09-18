package server

import (
	"github.com/kalinskilk/arquitetura-hexagonal/application"
	"github.com/kalinskilk/arquitetura-hexagonal/adapters/web/handler"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
    r := mux.NewRouter()

    n := negroni.New(
        negroni.NewLogger(),
    )

    // registra apenas as rotas no router
    handler.MakeProductHandlers(r, w.Service)

    // define router como handler do Negroni
    n.UseHandler(r)

    server := &http.Server{
        ReadHeaderTimeout: 10 * time.Second,
        WriteTimeout:      10 * time.Second,
        Addr:              ":9000",
        Handler:           n,
        ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
    }

    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}