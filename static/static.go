package static

import (
	"log"
	"net/http"
	"regexp"

	"github.com/jldec/learn-go-part2/counter"
)

// Static HTML middleware
// Assumes all paths start with /.
// Appends .html when filename is missing extension.
// 404s if any part of the path starts with.
// Ignores RawPath, does not make copy of Request or URL.
func Handler(cnt *counter.Counter, path string) http.Handler {
	fileServer := http.FileServer(http.Dir(path))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path

		cnt.Inc() // increment thread-safe counter
		log.Println(cnt.Get(), p)

		if dotPath, _ := regexp.MatchString(`/\.`, p); dotPath {
			http.NotFound(w, r)
			return
		}
		if noExtension, _ := regexp.MatchString(`/[^./]+$`, p); noExtension {
			r.URL.Path = p + ".html"
		}
		fileServer.ServeHTTP(w, r)
	})
}
