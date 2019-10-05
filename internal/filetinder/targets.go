package filetinder

// Target is target files in FileTinder
type Target struct {
	ID   int64    `json:"id"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}

// HasTag check if target has specific tag
func (t *Target) HasTag(s string) bool {
	for _, t := range t.Tags {
		if t == s {
			return true
		}
	}

	return false
}

// TargetStore is storage that maintains active targets
type TargetStore struct {
	targets []*Target
	idIncr  int64
}

// MakeTargetStore setup and initialize new TargetStore
func MakeTargetStore() (targetStore *TargetStore) {
	ts := TargetStore{}

	ts.targets = make([]*Target, 0)
	ts.idIncr = 0

	return &ts
}

func (tm *TargetStore) findTarget(t *Target) (index int, target *Target) {
	for i, st := range tm.targets {
		if st.ID == int64(t.ID) {
			return i, st
		}
	}

	return -1, nil
}

// Add adds new target to store
func (tm *TargetStore) Add(t *Target) (target *Target) {
	t.ID = tm.idIncr
	tm.targets = append(tm.targets, t)

	tm.idIncr = tm.idIncr + 1

	return t
}

// Del removes target from store
func (tm *TargetStore) Del(t *Target) (target *Target) {
	i, st := tm.findTarget(t)
	tm.targets = append(tm.targets[:i], tm.targets[i+1:]...)

	return st
}

// List lists all stored targets
func (tm *TargetStore) List() (targets []*Target) {
	return tm.targets
}

// FindByID find target by ID from stored targets
func (tm *TargetStore) FindByID(id int) (target *Target) {
	_, t := tm.findTarget(&Target{ID: int64(id)})
	return t
}
