package mockingwsresponse_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

type Item struct {
	Status  string `xml:"status"`
	Message string `xml:"message"`
}

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("ContentType", "application/xml")
		xmlResponse := "<item>\n    <status>OK</status>\n    <message>Tu solicitud ha sido procesada con éxito.</message>\n</item>"
		fmt.Fprintln(w, xmlResponse)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownloadMockServer(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("Error while download", server.URL)
	}
	defer resp.Body.Close()

	if resp.StatusCode == statusCode {
		t.Logf("%s Deberiamos recibir a %d status code", succeed, statusCode)
	}

	var item Item
	if err := xml.NewDecoder(resp.Body).Decode(&item); err != nil {
		t.Fatal("Error while decode item")
	}
	if item.Message != "" {
		t.Logf("%s Message reponse ok", succeed)
	} else {
		t.Errorf("%s Error en la respuesta %s", failed, item)
	}
}
