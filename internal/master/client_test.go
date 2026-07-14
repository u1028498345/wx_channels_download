package master

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterSendsForwardAddr(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/wework/v1/register" {
			t.Fatalf("path = %q, want register path", r.URL.Path)
		}
		if got := r.URL.Query().Get("forward_addr"); got != "https://node.example.com" {
			t.Fatalf("forward_addr = %q, want forwarded node address", got)
		}
		if got := r.URL.Query().Get("robot"); got != "node-1" {
			t.Fatalf("robot = %q, want node-1", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":200,"msg":"ok"}`))
	}))
	defer server.Close()

	client := NewClient(Config{
		BaseURL:     server.URL,
		ForwardAddr: "https://node.example.com",
		Robot:       "node-1",
		Name:        "节点一",
	}, nil)
	if err := client.Register(context.Background()); err != nil {
		t.Fatalf("Register() returned %v", err)
	}
}

func TestHeartbeatDoesNotSendForwardAddr(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/wework/v1/heartbeat" {
			t.Fatalf("path = %q, want heartbeat path", r.URL.Path)
		}
		if got := r.URL.Query().Get("forward_addr"); got != "" {
			t.Fatalf("forward_addr = %q, want empty on heartbeat", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":200,"msg":"ok"}`))
	}))
	defer server.Close()

	client := NewClient(Config{
		BaseURL:     server.URL,
		ForwardAddr: "https://node.example.com",
		Robot:       "node-1",
		Name:        "节点一",
	}, nil)
	if err := client.Heartbeat(context.Background()); err != nil {
		t.Fatalf("Heartbeat() returned %v", err)
	}
}
