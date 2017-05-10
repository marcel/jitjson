// +build linux

package fixtures

type OnlyBuiltOnLinux struct {
	KernelRecompileNeeded bool `json:"kernel_recompile_needed"`
}
