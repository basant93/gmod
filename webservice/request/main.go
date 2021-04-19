package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(f.Message))
}

type Weight struct {
	Unit     string
	Quantity float32
}

type Money struct {
	Amount   float32
	Currency string
}

type PaintBucket struct {
	ColorCode int    `json:"ColorCode,omitempty"`
	ColorName string `json:"ColorName,omitempty"`
	Weight    Weight
	Cost      Money
}

func main() {
	http.Handle("/ping", &fooHandler{Message: "Hello Server"})

	bar := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello Go!!"))
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		case http.MethodPost:
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)

				return
			}
			var productname string
			data := json.Unmarshal(bodyBytes, &productname)
			fmt.Println(data)
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			return

		}

	}
	http.HandleFunc("/bar", bar)

	color := &PaintBucket{ColorCode: 2699, ColorName: "Light Pink", Weight: Weight{Quantity: 20, Unit: "Kg"}, Cost: Money{Amount: 1400, Currency: "Ruppees"}}
	//field names must be exported
	data, _ := json.Marshal(&color)
	fmt.Println(string(data))

	type temp struct {
		Message string `json:"message"`
	}

	str := `{
		"message": "Hello Apple"
	}`
	var t temp
	err := json.Unmarshal([]byte(str), &t)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(t)

	//Mostly services are after load balancer which handles the request encryption and downgrade the request to http.
	//err := http.ListenAndServeTLS(":3000", nil)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

//servemux uses the pattern to route the traffic to specific handler. But if you want the same handler to
//accept multiple request, we can use pattern along with http methods. for this we need to implement new handlers.

//Middleware are the functions that are called before and after the handlers are called. It is mainly used for authentication, session
//management, logging and monitoring.

func middlewares(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Method", "POST, GET, OPTIONS")

		handler.ServeHTTP(w, r)
		end := time.Now()
		diff := end.Sub(start)
		fmt.Println(diff)

	})
}
