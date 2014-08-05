package main

import (
	"flag"
	"log"
	"net/http"
	"path"
	"text/template"

	"github.com/opentarock/frontend-user-management/server/api"
	"github.com/opentarock/frontend-user-management/server/logutil"
	"github.com/opentarock/frontend-user-management/server/service"
)

func main() {
	resourcesPath := flag.String("path", "../client/build/web/", "Path to resources to be served")
	listenOn := flag.String("listen", ":8080", "Hostname and port to listen on")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Lmicroseconds)

	log.Printf("Starting server on %s", *listenOn)
	log.Printf("Serving content from '%s'", *resourcesPath)

	staticPath := path.Join(*resourcesPath, "static")
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(staticPath))))

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		templatePath := path.Join(*resourcesPath, "register.html")
		// Change delims from default to not conflict with angular templates.
		template, err := template.New("register.html").Delims("[[", "]]").ParseFiles(templatePath)
		if err == nil {
			err = template.Execute(w, nil)
			if err == nil {
				return
			}
		}

		// Any error displaying a template is a bug or configuration error.
		logutil.Printf(r, "Error serving /register endpoint: %s", err)
		http.Error(w, "", http.StatusServiceUnavailable)
	})

	userService, err := service.NewUserServiceNanomsg()
	if err != nil {
		log.Fatalf("Error creating user service: %s", err)
	}
	defer userService.Close()

	http.Handle("/api/", http.StripPrefix("/api", api.New(userService)))

	http.ListenAndServe(*listenOn, nil)
}
