package unittest

import (
	"testing"
	"net/http"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	var urls = [] struct {
		url        string
		statusCode int
	}{
		{
			url:        "http://www.goinggo.net/feeds/posts/default?alt=rss",
			statusCode: http.StatusOK,
		},
		{
			url:        "http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			statusCode: http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("When checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\t Should be able to get the url.", ballotX, err)
				}
				t.Log("\t\tShould be able to get the url.", checkMark)

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status. %v %v", u.statusCode, ballotX, resp.StatusCode)
				}
				resp.Body.Close()
			}
		}
	}
}
