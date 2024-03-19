package paxi

import (
	"flag"
	"net/http"

	"github.com/steamgjk/paxi/log"
)

// Init setup paxi package
func Init() {
	flag.Parse()
	log.Setup()
	config.Load()
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
}
