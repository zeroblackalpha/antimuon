package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httptest"

	"github.com/webview/webview"
)

//go:embed frontend/public
var frontend embed.FS

func httpServer() *httptest.Server {
	fsys, err := fs.Sub(frontend, "frontend/public")
	if err != nil {
		panic(err)
	}

	return httptest.NewServer(http.FileServer(http.FS(fsys)))
}

func main() {
	srv := httpServer()
	defer srv.Close()

	url := fmt.Sprintf("%s/index.html", srv.URL)
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
