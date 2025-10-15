package engines

import (
	"fmt"
	"os"

	"modernc.org/quickjs"
)

type ModerncQuickJS struct {
	vm     *quickjs.VM
	output [][]string
}

func (q *ModerncQuickJS) Name() string {
	return "ModerncQuickJS"
}

func (q *ModerncQuickJS) Init() (err error) {
	if q.vm, err = quickjs.NewVM(); err != nil {
		return err
	}

	global := q.vm.GlobalObject()
	defer global.Free()

	if err := q.vm.RegisterFunc("load", func(input string) error {
		content, err := os.ReadFile("v8-v7/" + input)
		if err != nil {
			return fmt.Errorf("Could not read %s: %v", input, err)
		}

		if _, err = q.vm.Eval(string(content), quickjs.EvalGlobal); err != nil {
			return fmt.Errorf("Could not run script %s: %v", input, err)
		}

		return nil
	}, false); err != nil {
		return err
	}

	if err := q.vm.RegisterFunc("print", func(inputs ...any) error {
		var line []string
		for i := range inputs {
			if i > 0 {
				fmt.Print(" ")
				line = append(line, " ")
			}

			fmt.Print(inputs[i])
			line = append(line, fmt.Sprint(inputs[i]))
		}
		fmt.Println()
		q.output = append(q.output, line)
		return nil
	}, false); err != nil {
		return err
	}

	return nil
}

func (q *ModerncQuickJS) Close() error {
	q.output = nil
	if q.vm != nil {
		defer func() { q.vm = nil }()
		return q.vm.Close()
	}

	return nil
}

func (q *ModerncQuickJS) Run(inputFile string) ([][]string, error) {
	script, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", inputFile, err)
	}

	_, err = q.vm.Eval(string(script), quickjs.EvalGlobal)
	return q.output, err
}

var _ JSEngine = (*ModerncQuickJS)(nil)
