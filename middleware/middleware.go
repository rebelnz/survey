package middleware

import (
	// "fmt"
	"net/http"	
)

func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _,m := range mid {
		handler = m(handler)
	}
	return handler
}




















