package filetinder

// Target is target files in FileTinder
type Target struct {
	ID   int64    `json:"id"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}

// TargetsCollection is collection of targets
type TargetsCollection []*Target
