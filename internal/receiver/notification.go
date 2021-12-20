package receiver

type InvolvedObject struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	Uid             string `json:"uid"`
	ApiVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
}
type Notification struct {
	Object InvolvedObject `json:"involvedObject"`
	// Severity e.g info
	Severity  string `json:"severity"`
	Timestamp string `json:"timestamp"`
	// A string explainin the event e.g `ArtifactFailed`
	Message string `json:"message"`
	// Seems to be inconsistent, `ArtifactFailed` or `ReconciliationSucceeded` for Kustomizations, but `info` for
	// GitRepository.
	Reason string `json:"reason"`
	// Present for Kustomization resource but not GitRepository
	Metadata            map[string]string `json:"metadata"`
	ReportingController string            `json:"reportingController"`
	ReportingInstance   string            `json:"reportingInstance"`
}
