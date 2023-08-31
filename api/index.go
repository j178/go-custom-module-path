package api

import (
	"fmt"
	"net/http"
	"strings"
)

const tmpl = `
<html>
	<head>
		<meta name="go-import" content="go.j178.dev/%s git https://github.com/j178/%[1]s">
	</head>
	<body>
		<a href="https://github.com/j178/%[1]s">go.j178.dev/%[1]s</a>
	</body>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	module := strings.Split(r.URL.Path, "/")[1]
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(tmpl, module)))
}
