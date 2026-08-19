package charging

var inspectionHistory = []string{"checked", "rectified", "rechecked"}

func InspectionHistory() []string {
	out := make([]string, len(inspectionHistory))
	copy(out, inspectionHistory)
	return out
}
func ResetInspectionHistory() { inspectionHistory = []string{"checked", "rectified", "rechecked"} }
