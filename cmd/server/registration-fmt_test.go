package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/pbberlin/struc2frm"
	"github.com/zew/go-questionnaire/pkg/handlers"
)

func TestRegistrationFMTHandler(t *testing.T) {

	// valid form values
	form := url.Values{}
	form.Set("first_name", "Max")
	form.Set("last_name", "Mustermann")
	form.Set("unternehmen", "Beispiel GmbH")
	form.Set("abteilung", "IT")
	form.Set("position", "Entwickler")
	form.Set("plz", "12345")
	form.Set("ort", "Berlin")
	form.Set("strasse", "Musterstr. 1")
	form.Set("email", "max.mustermann@example.com")
	form.Set("telefon", "0123456789")
	form.Set("geburtsjahr", "1985")
	form.Set("einstieg", "2010")
	form.Set("terms", "true")

	s2f := struc2frm.New()
	tok := s2f.FormToken()
	form.Set("token", tok)

	// request with form data
	req := httptest.NewRequest(http.MethodPost, "/survey/registrationfmtde", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// record the response
	rr := httptest.NewRecorder()

	// execute http request
	handler := http.HandlerFunc(handlers.RegistrationFMTDeH)
	handler.ServeHTTP(rr, req)

	// Check response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Check response body
	expected1 := "Ihre Daten wurden gespeichert"
	expected2 := "Your data was saved"
	bdy := rr.Body.String()
	if !strings.Contains(bdy, expected1) && !strings.Contains(bdy, expected2) {
		t.Errorf("Expected body %q or %q - but got %q", expected1, expected2, bdy)
	}
}
