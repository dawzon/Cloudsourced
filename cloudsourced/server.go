package main

import "net/http"

func init() {

	http.ListenAndServeTLS(":443", "your_certfile", "your_key", nil)
}
