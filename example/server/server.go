package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"spark-url/utils/json"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func World(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "hello world")
	w.Write(buf.Bytes())
}

func Json(w http.ResponseWriter, r *http.Request) {
	var info = struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "xieWei",
		Age:  20,
	}
	bytes, _ := json.Marshal(&info)
	w.Write(bytes)

}

func String(w http.ResponseWriter, r *http.Request) {
	var buf strings.Builder
	buf.WriteString("真累啊")
	log.Println(buf.String())
	w.Write([]byte(buf.String()))

}
func Start() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/world", World)
	http.HandleFunc("/json", Json)
	http.HandleFunc("/string", String)
	var product Product
	product.Name = "Apple"
	product.Price = 10000
	http.HandleFunc("/product", product.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// method two

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	byte, _ := json.Marshal(&p)
	w.Write(byte)
}
