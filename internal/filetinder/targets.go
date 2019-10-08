package filetinder

import (
	"net/http"
	"os"
	"path/filepath"
)

// Target is target files in FileTinder
type Target struct {
	ID   int64    `json:"id"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}

// TargetFileStats is target's file stats/metadatas
type TargetFileStats struct {
	Filename    string `json:"filename"`
	Filepath    string `json:"filepath"`
	ContentType string `json:"contentType"`
	Size        int64  `json:"size"`
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

// GetFile returns target's file reader
func (t *Target) GetFile() (reader *os.File, err error) {
	filepath := t.URL

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetStats return target's file stats
func (t *Target) GetStats() (filestats *TargetFileStats, err error) {
	f, err := t.GetFile()
	if err != nil {
		return nil, err
	}

	fstat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fmimetype, err := getContentType(f)
	if err != nil {
		return nil, err
	}

	_, filename := filepath.Split(t.URL)

	defer f.Close()

	// build stats
	stats := TargetFileStats{
		ContentType: fmimetype,
		Size:        fstat.Size(),
		Filename:    filename,
		Filepath:    t.URL,
	}

	return &stats, nil
}

// getContentType sniff file's content type
// Credits: https://golangcode.com/get-the-content-type-of-file/
func getContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
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
	ts.idIncr = 1

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
