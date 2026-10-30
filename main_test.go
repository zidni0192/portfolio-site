package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeRendersPortfolioContent(t *testing.T) {
	content := Portfolio{Name: "Casey", Role: "Engineer", Introduction: "Building useful tools.", Projects: []Project{{Name: "Alpha", Description: "A project", URL: "https://example.com", Tags: []string{"Go"}}}}
	recorder := httptest.NewRecorder()
	newHandler(content).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if got := response.Header.Get("Content-Type"); got != "text/html; charset=utf-8" {
		t.Fatalf("Content-Type = %q", got)
	}
	body := recorder.Body.String()
	for _, want := range []string{"Hello, I’m Casey.", "Building useful tools.", "Alpha", "A project", ">Go<", "aria-labelledby=\"projects-heading\"", "focus-visible"} {
		if !strings.Contains(body, want) {
			t.Errorf("response body missing %q", want)
		}
	}
}

func TestHomeHandlesEmptyProjectList(t *testing.T) {
	recorder := httptest.NewRecorder()
	newHandler(Portfolio{Name: "Casey"}).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))
	if !strings.Contains(recorder.Body.String(), "No projects to show yet.") {
		t.Fatal("empty portfolio fallback not rendered")
	}
}

func TestHomeEscapesContentAndSortsProjects(t *testing.T) {
	content := Portfolio{Name: "<script>alert(1)</script>", Projects: []Project{{Name: "Zulu"}, {Name: "Alpha"}}}
	recorder := httptest.NewRecorder()
	newHandler(content).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))
	body := recorder.Body.String()
	if strings.Contains(body, "<script>alert(1)") {
		t.Fatal("unescaped user content in response")
	}
	if strings.Index(body, "Alpha") > strings.Index(body, "Zulu") {
		t.Fatal("projects are not sorted by name")
	}
}

func TestHomeNotFoundForUnknownPath(t *testing.T) {
	recorder := httptest.NewRecorder()
	newHandler(portfolio).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/missing", nil))
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusNotFound)
	}
}
