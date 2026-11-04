package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// A page that reads from the back of a room: the app's name and version, large.
var page = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Name}} {{.Version}}</title>
<style>
  body { margin: 0; min-height: 100vh; display: grid; place-content: center;
         text-align: center; font-family: system-ui, sans-serif;
         background: #0f172a; color: #f8fafc; }
  h1 { margin: 0; font-size: clamp(3rem, 10vw, 8rem); }
  p  { margin: 0.5em 0 0; font-size: clamp(1.5rem, 5vw, 4rem); color: #38bdf8; }
</style>
</head>
<body>
<h1>{{.Name}}</h1>
<p>{{.Version}}</p>
</body>
</html>
`))

func main() {
	data := struct{ Name, Version string }{os.Getenv("APP_NAME"), os.Getenv("APP_VERSION")}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := page.Execute(w, data); err != nil {
			log.Print(err)
		}
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
