package exchangerate

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMonobankClient_EURUAHRate_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[
			{"currencyCodeA":840,"currencyCodeB":980,"date":1787605273,"rateBuy":44.43,"rateSell":44.831},
			{"currencyCodeA":978,"currencyCodeB":980,"date":1787640606,"rateBuy":51.8,"rateSell":52.499}
		]`))
	}))
	defer server.Close()

	client := &MonobankClient{baseURL: server.URL, httpClient: http.DefaultClient}

	rate, err := client.EURUAHRate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rate != 52.499 {
		t.Errorf("expected rate 52.499, got %v", rate)
	}
}

func TestMonobankClient_EURUAHRate_PairNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"currencyCodeA":840,"currencyCodeB":980,"rateBuy":44.43,"rateSell":44.831}]`))
	}))
	defer server.Close()

	client := &MonobankClient{baseURL: server.URL, httpClient: http.DefaultClient}

	_, err := client.EURUAHRate()
	if err == nil {
		t.Fatal("expected error when EUR/UAH pair is missing, got nil")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("expected 'not found' error, got: %v", err)
	}
}

func TestMonobankClient_EURUAHRate_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &MonobankClient{baseURL: server.URL, httpClient: http.DefaultClient}

	_, err := client.EURUAHRate()
	if err == nil {
		t.Fatal("expected error for non-200 response, got nil")
	}
	if !strings.Contains(err.Error(), "status code 500") {
		t.Errorf("expected status code error, got: %v", err)
	}
}

func TestMonobankClient_EURUAHRate_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`not json`))
	}))
	defer server.Close()

	client := &MonobankClient{baseURL: server.URL, httpClient: http.DefaultClient}

	_, err := client.EURUAHRate()
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestNewMonobankClient(t *testing.T) {
	client := NewMonobankClient()

	if client.baseURL != monobankBaseURL {
		t.Errorf("expected baseURL %q, got %q", monobankBaseURL, client.baseURL)
	}
	if client.httpClient == nil {
		t.Error("expected httpClient to be set")
	}
}
