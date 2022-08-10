package main

import (
	"fmt"
	"io"
	"net/http"
)

// func order(w http.ResponseWriter, req *http.Request) {

// 	fmt.Printf("这是订单")
// }

func readyBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ready body failed: %v", err)
		return
	}

	fmt.Fprintf(w, "ready the data: %s \n", body)

	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ready body failed: %v", err)
		return
	}

	fmt.Fprintf(w, "ready the data on more time:[%s] and read data length %d \n", string(body), len(body))

}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprint(w, "get body is nil \n")
	} else {
		fmt.Fprintf(w, "get body not nil")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v \n", values)
}

func main() {

	server := NewHttpServer("my-test-server")

	server.Route("POST", "/sign", SignUp)
	http.HandleFunc("/body", readyBodyOnce)
	http.HandleFunc("/get_body", getBodyIsNil)
	http.HandleFunc("/ ", queryParams)
	server.Start(":8090")
	//http.ListenAndServe(":8090", nil)
}
