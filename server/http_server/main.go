package main

import (
	"fmt"
	"io"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	// http.ServeFile(w, req, "./form.html")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func simple_html(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, " This tag is compulsory for any HTML document.   ")

}

func static_file(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static.html")
}

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("printing since reached test")
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
}

func form1Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./form1.html")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fmt.Println("routine http server is serving on port 6897")
	http.HandleFunc("/test", test)

	http.HandleFunc("/trade", headers)
	http.HandleFunc("/simple_html", simple_html)
	http.HandleFunc("/static_file", static_file)
	http.HandleFunc("/form1", form1Handler)
	http.HandleFunc("/form", formHandler)
	http.ListenAndServe(":6897", nil)

}
