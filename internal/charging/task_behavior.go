package charging

var inspectionHistory = []string{"checked", "rectified", "rechecked"}

func InspectionHistory() []string { return inspectionHistory }
func ResetInspectionHistory()     { inspectionHistory = []string{"checked", "rectified", "rechecked"} }
