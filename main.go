package main

import (
	"log"
	"net"
	"bytes"
	"io"
	"net/http"
	"mime"
	"path/filepath"

	"github.com/zserge/webview"
)

// Counter is a simple example of automatic Go-to-JS data binding
type Counter struct {
	Value int `json:"value"`
}

// Add increases the value of a counter by n
func (c *Counter) Add(n int) {
	c.Value = c.Value + int(n)
}

// Reset sets the value of a counter back to zero
func (c *Counter) Reset() {
	c.Value = 0
}

// func runLocalHTTP() {
// 	ln, err := net.Listen("tcp", "127.0.0.1:0")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	go func() {
// 		defer ln.Close()
// 		http.Handle("/", http.FileServer(assetFS()))
// 		log.Fatal(http.Serve(ln, nil))
// 	}()
	
// 	url := "http://" + ln.Addr().String() + "/index.html"
// 	w := webview.New(webview.Settings{
// 		Title: "Loaded: " + url,
// 		URL:   url,
// 	})
// 	defer w.Exit()
// 	w.Run()
// }

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if len(path) > 0 && path[0] == '/' {
				path = path[1:]
			}
			if path == "" {
				path = "index.html"
			}
			if bs, err := Asset(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
				io.Copy(w, bytes.NewBuffer(bs))
			}
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

func main() {

	url := startServer()

	//---------------------------------
	// Webview
	//---------------------------------
	w := webview.New(webview.Settings{
		Width:  320,
		Height: 480,
		Title:  "Todo App",
		URL:    url,
	})
	defer w.Exit()
	go func() {
		w.Eval(string(MustAsset("index.html")))
	}()

	//---------------------------------
	// Lorca
	//---------------------------------
	// w, _ := lorca.New(url,"",1200,880)
	//defer w.Close()
	//<-w.Done()
}

//func main() {
	// w := webview.New(webview.Settings{
	// 	Title: "Nuxt Sample ",
	// })
	// defer w.Exit()

	// w.Dispatch(func() {
	// 	// Inject controller
	// 	w.Bind("counter", &Counter{})
	// 	w.Eval(string(MustAsset("index.html")))
	// })
	// w.Run()
//}
