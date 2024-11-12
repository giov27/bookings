package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/giov27/bookings/pkg/config"
	"github.com/giov27/bookings/pkg/handlers"
	"github.com/giov27/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":9001"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//change true when inProduction
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server")
	}
}
