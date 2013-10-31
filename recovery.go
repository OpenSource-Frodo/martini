package martini

import (
	"log"
	"net/http"
)

func Recovery(res http.ResponseWriter, c Context, logger *log.Logger) {
	defer func() {
		if err := recover(); err != nil {
			res.WriteHeader(500)
			logger.Printf("PANIC: %v\n", err)
		}
	}()

	c.Next()
}
