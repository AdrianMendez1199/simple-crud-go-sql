package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// WriteLog Middleware
func WriteLog(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		t := time.Now()

		fmt.Println(t.Year())
		formatted := fmt.Sprintf("%d/%d/%d",
			t.Year(), t.Month(), t.Day(),
		)

		fmt.Println(r.Header)

		logPath := "logs/" + formatted
		err := os.MkdirAll(logPath, 0755)
		check(err)

		f, err := os.OpenFile(logPath+"/mid.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		check(err)

		_, err = f.WriteString(
			"" + " Date" + t.String() +
				"User Agrent: " + r.Header.Get("User-Agent") + " Host: " + r.Header.Get("Host") + "\n")

		check(err)

		next(w, r)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
