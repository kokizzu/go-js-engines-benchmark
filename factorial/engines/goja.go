package engines

import "github.com/dop251/goja"

type GOJA struct {
	rt *goja.Runtime
}

func (g *GOJA) Name() string {
	return "GOJA"
}

func (g *GOJA) Init() error {
	g.rt = goja.New()
	return nil
}

func (g *GOJA) Close() error {
	g.rt = nil
	return nil
}

func (g *GOJA) Run(input string) error {
	_, err := g.rt.RunString(input)
	return err
}

var _ JSEngine = (*GOJA)(nil)
