package httpapi

import (
	"chargeguard/internal/charging"
	"net/http/httptest"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	charging.ResetInspectionHistory()
	a := httptest.NewRecorder()
	TaskHTTPHandler(a, httptest.NewRequest("GET", "/task?current=closed", nil))
	b := httptest.NewRecorder()
	TaskHTTPHandler(b, httptest.NewRequest("GET", "/task", nil))
	if b.Body.String() != "[\"checked\",\"rectified\",\"rechecked\"]\n" {
		t.Fatalf("history=%s", b.Body.String())
	}
}
