package cpprof


// #include <cpprof.h>
import "C"
import (
	"net/http"
)

func init() {
	http.HandleFunc("/debug/pprof/cheap", cHeapProfile)
}

func cHeapProfile(w http.ResponseWriter, r *http.Request) {
	C.useSystem()
}
