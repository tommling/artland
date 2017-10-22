package artland

import (
	"fmt"
	"log"
	"net/http"
)

type Artland struct {
}

func NewArtLand() *Artland {
	return &Artland{}
}

func (al *Artland) Serve() error {
	http.HandleFunc("/verify", verifyFunc)
	http.HandleFunc("/login", loginFunc)
	http.HandleFunc("/register", registerFunc)
	http.HandleFunc("/post", postFunc)
	http.HandleFunc("/buy", buyFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

func verifyFunc(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello World!")); err != nil {
		fmt.Printf("Uh oh!")
	}
}

func loginFunc(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello World!")); err != nil {
		fmt.Printf("Uh oh!")
	}
}

func registerFunc(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello World!")); err != nil {
		fmt.Printf("Uh oh!")
	}
}

func postFunc(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello World!")); err != nil {
		fmt.Printf("Uh oh!")
	}
}

func buyFunc(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello World!")); err != nil {
		fmt.Printf("Uh oh!")
	}
}
