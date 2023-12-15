package logger

import (
	"encoding/json"
	"fmt"

	"github.com/wardonne/gopi/web/middleware"
)

func stdWriter(info Information) {
	bytes, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// StdLogger write log to std output
func StdLogger() middleware.IMiddleware {
	return New(stdWriter)
}
