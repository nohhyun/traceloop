package tracemeta

type TraceMeta struct {
	TraceID      string `json:"traceid,omitempty"`
	ContainerID  string `json:"containerid,omitempty"`
	UID          string `json:"uid,omitempty"`
	Namespace    string `json:"namespace,omitempty"`
	Podname      string `json:"podname,omitempty"`
	Containeridx int    `json:"containeridx,omitempty"`
}

