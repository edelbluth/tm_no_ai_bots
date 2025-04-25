package tm_no_ai_bots

import "net/http"

func BlockAgent(res *http.ResponseWriter) {
	(*res).Header().Set("Content-Type", "text/plain; charset=utf-8")
	(*res).WriteHeader(http.StatusForbidden)
	_, _ = (*res).Write([]byte("Access denied"))
}
