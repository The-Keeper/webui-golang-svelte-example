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
	go http.ListenAndServe(":8050", nil)

	w := webui.NewWindow()
	w.SetPort(8051)
	
	w.Bind("__close-btn", Close)

	w.Show("http://localhost:8050")

	webui.Wait()
}
