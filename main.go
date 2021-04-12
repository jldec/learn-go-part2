package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/jldec/learn-go-part2/counter"
	"github.com/jldec/learn-go-part2/sqrt"
	"github.com/jldec/learn-go-part2/static"
)

func main() {
	cnt := counter.New()
	http.Handle("/counter", cnt.Handler())
	http.Handle("/sqrt", sqrt.Handler(cnt))
	http.Handle("/", static.Handler(cnt, "/Users/jleschner/pub/fmc/generated/"))

	log.Println("Listening on :3000 (maxprocs: " + fmt.Sprint(runtime.GOMAXPROCS(0)) + ")")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
