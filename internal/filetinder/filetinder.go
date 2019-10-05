package filetinder

// TargetsCollection is collection of targets
type TargetsCollection []*Target

var (
	// TargetColl is active target store in session
	TargetColl TargetsCollection

	// TargetStr is active target store in session
	TargetStr TargetStore

	// TargetIDIncrement is last incremented id for new target
	TargetIDIncrement int64
)

func init() {
	TargetStr = TargetStore{}.Init()
	TargetColl = make([]*Target, 0)
	TargetIDIncrement = 1
}
