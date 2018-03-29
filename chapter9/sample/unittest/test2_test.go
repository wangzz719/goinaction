package unittest

import (
	"net/http/httptest"
	"net/http"
	"testing"
)

const checkMark2 = "\u2713"
const ballotX2 = "\u2717"

var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
</rss>`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(feed))
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload_2(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	t.Log("Given the need to test downloading different content.")
	t.Logf("When checking \"%s\" for status code \"%d\"", server.URL, statusCode)
	{
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatal("\t\t Should be able to get the url.", ballotX, err)
		}
		t.Log("\t\tShould be able to get the url.", checkMark)

		if resp.StatusCode == statusCode {
			t.Logf("\t\tShould have a \"%d\" status. %v", statusCode, checkMark)
		} else {
			t.Errorf("\t\tShould have a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
		}
		resp.Body.Close()
	}
}
