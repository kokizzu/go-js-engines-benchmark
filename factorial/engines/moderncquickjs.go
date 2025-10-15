package engines

import (
	"modernc.org/quickjs"
)

type ModerncQuickJS struct {
	vm *quickjs.VM
}

func (g *ModerncQuickJS) Name() string {
	return "ModerncQuickJS"
}

func (g *ModerncQuickJS) Init() error {
	var err error
	g.vm, err = quickjs.NewVM()
	return err
}

func (q *ModerncQuickJS) Close() error {
	if q.vm != nil {
		defer func() { q.vm = nil }()
		return q.vm.Close()
	}

	return nil
}

func (g *ModerncQuickJS) Run(input string) error {
	_, err := g.vm.Eval(input, quickjs.EvalGlobal)
	return err
}

var _ JSEngine = (*ModerncQuickJS)(nil)
