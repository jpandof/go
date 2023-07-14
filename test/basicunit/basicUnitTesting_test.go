package basicunit_test

import (
	"net/http"
	"testing"
)

const (
	Succeed = "\u2713"
	Failed  = "\u2717"
)

func TestBasic(t *testing.T) {
	url := "https://www.google.com"
	statusCode := 200
	t.Log("Download contest")
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal("Error while download", url)
	}

	defer resp.Body.Close()

	if resp.StatusCode == statusCode {
		t.Logf("%s Deberiamos recibir a %d status code", Succeed, statusCode)
	} else {
		t.Errorf("%s Deberiamos haber recivido a %d status code : %d", Failed, statusCode, resp.StatusCode)
	}

}
