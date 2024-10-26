package main

import (
	"embed"
	// "fmt"
	// "io/ioutil"
	"io/fs"
	"net/http"

	"github.com/webui-dev/go-webui/v2"
)

func Close(_ webui.Event) any {
	println("Exit.")

	webui.Exit()

	return nil

}

//go:embed ui/dist
var embeddedFiles embed.FS

func main() {

	fsys, err := fs.Sub(embeddedFiles, "ui/dist")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(":8050", nil)

	w := webui.NewWindow()

	w.SetRootFolder("ui")

	w.Bind("__close-btn", Close)

	if err := w.ShowBrowser("index.html", webui.ChromiumBased); err != nil {
		println("Warning: Install a Chromium-based web browser for an optimized experience.")
		w.Show("index.html")
	}

	webui.Wait()
}
