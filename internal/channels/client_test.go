package channels

import "testing"

func TestAvailableCount(t *testing.T) {
	client := NewChannelsClient(0)
	if count := client.AvailableCount(); count != 0 {
		t.Fatalf("AvailableCount() = %d, want 0", count)
	}
	if err := client.Validate(); err == nil {
		t.Fatal("Validate() returned nil without connected clients")
	}

	client.ws_mu.Lock()
	client.ws_clients[&Client{}] = true
	client.ws_mu.Unlock()

	if count := client.AvailableCount(); count != 1 {
		t.Fatalf("AvailableCount() = %d, want 1", count)
	}
	if err := client.Validate(); err != nil {
		t.Fatalf("Validate() returned %v, want nil", err)
	}
}
