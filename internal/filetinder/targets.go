package filetinder

// Target is target files in FileTinder
type Target struct {
	ID   int64    `json:"id"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}

// TargetStore is storage that maintains active targets
type TargetStore struct {
	targets []*Target
	idIncr  int64
}

// Init setup and initialize new TargetStore
func (tm TargetStore) Init() (targetManager TargetStore) {
	tm.targets = make([]*Target, 0)
	tm.idIncr = 0

	return tm
}

// Add adds new target to store
func (tm TargetStore) Add(t Target) (target Target) {
	t.ID = tm.idIncr
	tm.targets = append(tm.targets, &t)

	tm.idIncr = tm.idIncr + 1

	return t
}

// List lists all stored targets
func (tm TargetStore) List() (targets []*Target) {
	return tm.targets
}
