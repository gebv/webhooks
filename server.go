package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hooks/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-Version-ID", fmt.Sprintf("%v.%v", CfgVersion, CfgBuildStamp))

		re := regexp.MustCompile("/hooks/(.+)")
		route := re.FindStringSubmatch(r.URL.Path)

		if len(route) == 0 {
			log.Println("not valid route, path=" + r.URL.Path)

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token := strings.TrimSpace(route[1])

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
