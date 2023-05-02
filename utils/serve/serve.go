package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fileSystem := os.DirFS(".")

	fileServer := http.FileServer(http.FS(fileSystem))

	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// default to "index.html"
			if r.URL.Path == "/" {
				f, err := fileSystem.Open("index.html")
				// if we found index.html serve it to the user
				if err == nil {
					io.Copy(w, f)
					return
				}
			}

			fileServer.ServeHTTP(w, r)
		}),
	}

	fmt.Println("serving static files on :8080")
	fmt.Println(server.ListenAndServe())
}
