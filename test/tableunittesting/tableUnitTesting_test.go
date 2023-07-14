package tableunittesting_test

import (
	"net/http"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		url        string
		statusCode int
	}{
		{"https://www.google.com", http.StatusOK},
		{"http://rss.cnn.com/rss/cnn_topstories2.rss", http.StatusNotFound},
	}

	for _, test := range tests {
		resp, err := http.Get(test.url)
		if err != nil {
			t.Fatal("Error while download", test.url)
		}
		defer resp.Body.Close()

		if resp.StatusCode == test.statusCode {
			t.Logf("%s Deberiamos recibir a %d status code", succeed, test.statusCode)
		} else {
			t.Errorf("%s Deberiamos haber recibido a %d status code : %d", failed, test.statusCode, resp.StatusCode)
		}

	}

}
