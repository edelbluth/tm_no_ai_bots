package tm_no_ai_bots_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edelbluth/tm_no_ai_bots"
)

type ResponderMethod func(*http.ResponseWriter)

type ResponderTestCase struct {
	RespondingMethod ResponderMethod
	RequestUrl       string
	ExpectedStatus   int
	ExpectedBody     string
}

func doResponderTest(testCase *ResponderTestCase, t *testing.T) {
	recorder := httptest.NewRecorder()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testCase.RespondingMethod(&w)
	})

	req, err := http.NewRequest("GET", testCase.RequestUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	h.ServeHTTP(recorder, req)

	if status := recorder.Code; status != testCase.ExpectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, testCase.ExpectedStatus)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if contentType := recorder.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	if recorder.Body.String() != testCase.ExpectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), testCase.ExpectedBody)
	}
}

func TestResponders(t *testing.T) {
	testCases := []ResponderTestCase{
		{
			RespondingMethod: tm_no_ai_bots.RobotsTxt,
			RequestUrl:       "/robots.txt",
			ExpectedStatus:   http.StatusOK,
			ExpectedBody: `User-agent: *
Disallow: /
`,
		},
		{
			RespondingMethod: tm_no_ai_bots.RobotsTxt,
			RequestUrl:       "/robots1.txt",
			ExpectedStatus:   http.StatusOK,
			ExpectedBody: `User-agent: *
Disallow: /
`,
		},
		{
			RespondingMethod: tm_no_ai_bots.BlockAgent,
			RequestUrl:       "/robots.txt",
			ExpectedStatus:   http.StatusForbidden,
			ExpectedBody:     "Access denied",
		},
		{
			RespondingMethod: tm_no_ai_bots.BlockAgent,
			RequestUrl:       "/",
			ExpectedStatus:   http.StatusForbidden,
			ExpectedBody:     "Access denied",
		},
	}
	for i, testCase := range testCases {
		testCase := testCase
		t.Run(fmt.Sprintf("case-%d", i), func(subTest *testing.T) {
			doResponderTest(&testCase, subTest)
		})
	}
}
