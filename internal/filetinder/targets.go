package filetinder

// Target is target files in FileTinder
type Target struct {
	ID   int64    `json:"id"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}

// TargetsCollection is collection of targets
type TargetsCollection []*Target

var (
	// TargetStore is active target store in session
	TargetStore TargetsCollection

	// TargetIDIncrement is last incremented id for new target
	TargetIDIncrement int64
)

func init() {
	TargetStore = make([]*Target, 0)
	TargetIDIncrement = 1
}
