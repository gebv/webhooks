package main

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-Version-ID", fmt.Sprintf("%v.%v", CfgVersion, CfgBuildStamp))

		token := r.URL.Query().Get("webhook_key")

		if err := Execute(token, r.Body); err != nil {

			if err != ErrNotFound {
				log.Println("error execute, err=" + err.Error() + ", token=" + token)
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Println("Listening... " + flagAddr)
	http.ListenAndServe(flagAddr, mux)
}
