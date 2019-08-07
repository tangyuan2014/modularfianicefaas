package errorhandeling

import "net/http"

type serviceHandeler func(writer http.ResponseWriter, request http.Request) error

func errWrapper(handeler serviceHandeler) func(writer http.ResponseWriter, request http.Request) {

}
