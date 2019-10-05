package filetinder

var (
	// TargetStoreInst is active target store in session
	TargetStoreInst *TargetStore
)

func init() {
	TargetStoreInst = MakeTargetStore()
}
