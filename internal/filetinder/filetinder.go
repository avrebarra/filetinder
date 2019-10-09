package filetinder

import (
	"github.com/shrotavre/filetinder/internal/config"
)

var (
	// TargetStoreInst is active target store in session
	TargetStoreInst *TargetStore

	// Config hold global configuration to be used in session
	Config *config.FileTinderConfig
)

func init() {
	TargetStoreInst = MakeTargetStore()
	Config = config.InitConfig()
}
