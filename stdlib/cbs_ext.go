package stdlib

import (
	"io/fs"

	"github.com/go-python/gpython/py"
)

func NewContextWithFS(opts py.ContextOpts, fs fs.FS) py.Context {
	ctx := &context{
		opts:    opts,
		done:    make(chan struct{}),
		closing: false,
		closed:  false,
		fs:      fs,
	}

	ctx.store = py.NewModuleStore()

	py.Import(ctx, "builtins", "sys")

	sys_mod := ctx.Store().MustGetModule("sys")
	sys_mod.Globals["argv"] = py.NewListFromStrings(opts.SysArgs)
	sys_mod.Globals["path"] = py.NewListFromStrings(opts.SysPaths)

	return ctx
}
