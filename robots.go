package tm_no_ai_bots

import "net/http"

const robotsTxt = `User-agent: *
Disallow: /
`

func RobotsTxt(res *http.ResponseWriter) {
	(*res).Header().Set("Content-Type", "text/plain; charset=utf-8")
	(*res).WriteHeader(http.StatusOK)
	_, _ = (*res).Write([]byte(robotsTxt))
}
