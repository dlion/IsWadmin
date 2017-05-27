package IsWadmin

import "testing"

func TestIsWadmin(t *testing.T) {
	if IsWadmin() {
		t.Log("This process is running as administrator on Windows")
	}
}
