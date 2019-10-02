package filetinder

// Target is target files in FileTinder
type Target struct {
	ID  int64  `json:"id"`
	URL string `json:"url" form:"url"`
	Tag string `json:"tag"`
}

// TargetsCollection is collection of targets
type TargetsCollection []*Target
