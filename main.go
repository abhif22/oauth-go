package main

import (
	"encoding/json"
	"net/http"
)

type person struct {
	Name string // Must be caps to be properly marshalled/unmarshalled
	Age  int
}

func main() {
	http.HandleFunc("/encode", foo)
	http.ListenAndServe(":3030", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	p := person{
		Name: "James Bond",
		Age:  35,
	}
	b, err := json.Marshal(&p)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`Some error occured in marshaling`))
		return
	}
	w.WriteHeader(200)
	w.Write(b)
}