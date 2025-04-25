package tm_no_ai_bots

import (
	"net/http"
)

func (t *TmNoAiBots) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for _, uaHeader := range req.Header.Values("User-Agent") {
		if t.Matcher.MatchString(uaHeader) {
			if req.URL.Path == "/robots.txt" {
				RobotsTxt(&res)
				return
			}
			BlockAgent(&res)
			return
		}
	}
	// Default -> Go on!
	t.next.ServeHTTP(res, req)
}
