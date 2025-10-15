package engines

import (
	"fmt"
	"os"

	"github.com/dop251/goja"
)

type GOJA struct {
	rt     *goja.Runtime
	output [][]string
}

func (g *GOJA) Name() string {
	return "GOJA"
}

func (g *GOJA) Init() error {
	g.rt = goja.New()
	if err := g.rt.Set("load", func(call goja.FunctionCall) goja.Value {
		input := call.Argument(0).String()
		content, err := os.ReadFile("v8-v7/" + input)
		if err != nil {
			return g.rt.ToValue(fmt.Errorf("Could not read %s: %v", input, err))
		}

		if _, err = g.rt.RunScript(input, string(content)); err != nil {
			return g.rt.ToValue(fmt.Errorf("Could not run script %s: %v", input, err))
		}

		return nil
	}); err != nil {
		return err
	}

	if err := g.rt.Set("print", func(call goja.FunctionCall) goja.Value {
		var line []string
		for i := 0; i < len(call.Arguments); i++ {
			if i > 0 {
				fmt.Print(" ")
				line = append(line, " ")
			}

			arg := call.Argument(i).String()
			fmt.Print(arg)
			line = append(line, arg)
		}

		fmt.Println()
		g.output = append(g.output, line)
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (g *GOJA) Close() error {
	g.output = nil
	g.rt = nil
	return nil
}

func (g *GOJA) Run(inputFile string) ([][]string, error) {
	script, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", inputFile, err)
	}

	_, err = g.rt.RunString(string(script))
	return g.output, err
}

var _ JSEngine = (*GOJA)(nil)
