package logger

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)
var logger zerolog.Logger

func init () {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

    output.FormatLevel = func(i interface{}) string {
    return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
    return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
    return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	logger = zerolog.New(output).With().Timestamp().Logger()
}


func Logger(next http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.Info().Msgf("%s\t%s\t%s\t%s",
					  r.Method,
					  r.RequestURI,
					  name,
					  time.Since(start))
	})
}