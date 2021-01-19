package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := `https://api.coinbase.com/v2/prices/BTC-USD/spot`;

		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		bodyString := string(body)

		

		fmt.Println("STATUS: ", bodyString)
		fmt.Fprint(w, bodyString)
	})

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}