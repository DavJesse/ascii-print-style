package Web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownloadArtHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/download-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DownloadArtHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check Content-Type
	if ct := rr.Header().Get("Content-Type"); ct != "text/plain" {
		t.Errorf("Content-Type header is wrong: got %v want %v", ct, "text/plain")
	}

	// Check Content-Length
	if cl := rr.Header().Get("Content-Length"); cl == "" {
		t.Error("Content-Length header is missing or empty")
	}

	// Check Content-Disposition
	if cd := rr.Header().Get("Content-Disposition"); cd != `attachment; filename="art.txt"` {
		t.Errorf("Content-Disposition header is wrong: got %v want %v", cd, `attachment; filename="art.txt"`)
	}
}

func TestDirectoryTraversal(t *testing.T) {
	// Attempt to access a file outside the intended directory
	req, err := http.NewRequest("GET", "/download-art?file=../etc/passwd", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DownloadArtHandler)

	handler.ServeHTTP(rr, req)

	// Check if the response status is 403 Forbidden or 400 Bad Request
	if status := rr.Code; status != http.StatusForbidden && status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v or %v",
			status, http.StatusForbidden, http.StatusBadRequest)
	}

	// Alternatively, check if the response does not contain sensitive data
	if body := rr.Body.String(); body != "Access Denied" { // Or another appropriate response message
		t.Errorf("Unexpected body content: got %v", body)
	}
}
