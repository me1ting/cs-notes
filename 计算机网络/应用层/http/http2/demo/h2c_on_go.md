```go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// dependency: go get golang.org/x/net@v0.8.0

// h2c updagrade from http1.1
// test: curl -v --http2 http://localhost:8080
func serverH2cUpgradeFromH1() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})
	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(handler, h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}

// h2c only http2
//test: curl -v --http2-prior-knowledge http://localhost:8080
func serverOnlyH2c() {
	server := http2.Server{}

	l, err := net.Listen("tcp", "0.0.0.0:8080")
	checkErr(err, "while listening")

	fmt.Printf("Listening [0.0.0.0:8080]...\n")
	for {
		conn, err := l.Accept()
		checkErr(err, "during accept")

		server.ServeConn(conn, &http2.ServeConnOpts{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
		})})
	}
}

func checkErr(err interface{}, msg string) {
	if err != nil {
		log.Fatal(err, msg)
	}
}

func main() {
	serverOnlyH2c()
}
```