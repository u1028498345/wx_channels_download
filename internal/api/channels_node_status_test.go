package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"wx_channel/internal/channels"
)

func TestChannelsNodeStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	client := &APIClient{
		channels: channels.NewChannelsClient(0),
		engine:   gin.New(),
	}
	client.engine.GET("/api/channels/node_status", client.handleChannelsNodeStatus)

	resp := httptest.NewRecorder()
	client.ServeHTTP(resp, httptest.NewRequest(http.MethodGet, "/api/channels/node_status", nil))
	if resp.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", resp.Code, http.StatusOK)
	}

	var body struct {
		Code int `json:"code"`
		Data struct {
			Available bool `json:"available"`
			Channels  int  `json:"channels"`
		} `json:"data"`
	}
	if err := json.Unmarshal(resp.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Code != 0 {
		t.Fatalf("code = %d, want 0", body.Code)
	}
	if body.Data.Available {
		t.Fatal("available = true, want false without connected clients")
	}
	if body.Data.Channels != 0 {
		t.Fatalf("channels = %d, want 0", body.Data.Channels)
	}
}
