package charging

var inspectionHistory = []string{"checked", "rectified", "rechecked"}

// InspectionHistory returns a defensive copy of the persisted inspection
// history. Display-only operations (appending the current status, reordering
// for presentation) must not mutate the saved records through slice aliasing,
// so callers receive a private copy rather than the underlying slice.
func InspectionHistory() []string {
	out := make([]string, len(inspectionHistory))
	copy(out, inspectionHistory)
	return out
}

func ResetInspectionHistory() {
	inspectionHistory = []string{"checked", "rectified", "rechecked"}
}
