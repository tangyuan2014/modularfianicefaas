package errorhandeling

import "net/http"

type serviceHandeler func(writer http.ResponseWriter, request http.Request) error

func errWrapper(handeler serviceHandeler) func(writer http.ResponseWriter, request http.Request) {

return 	func (writer http.ResponseWriter request http.Request){
	err :=handeler(writer,request)
	if err!=nil{
		log.Println("Error Occured"+"Error: %s"+err.Error)
		if ue,ok :=err.(userError); ok{
			
		}
	}
}

}
