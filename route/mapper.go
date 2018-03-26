package route

import (
	"fmt"
	"gogateway/continar"
	"net/http"
)

type Mapper struct{}

func (rh Mapper) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s%s\n", req.URL,
			continar.GetMsg())
	}
}
