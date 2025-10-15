package engines

import (
	"fmt"
	"os"

	"github.com/fastschema/qjs"
)

type QJS struct {
	rt     *qjs.Runtime
	output [][]string
}

func (q *QJS) Name() string {
	return "QJS"
}

func (q *QJS) Init() (err error) {
	q.rt, err = qjs.New()
	if err != nil {
		return err
	}

	q.rt.Context().SetFunc("load", func(ctx *qjs.This) (*qjs.Value, error) {
		input := ctx.Args()[0].String()
		content, err := os.ReadFile("v8-v7/" + input)
		if err != nil {
			return nil, fmt.Errorf("Could not read %s: %v", input, err)
		}

		if _, err = q.rt.Eval(input, qjs.Code(string(content)), qjs.TypeGlobal()); err != nil {
			return nil, fmt.Errorf("Could not run script %s: %v", input, err)
		}

		return nil, nil
	})

	q.rt.Context().SetFunc("print", func(ctx *qjs.This) (*qjs.Value, error) {
		var line []string
		for i := 0; i < len(ctx.Args()); i++ {
			if i > 0 {
				fmt.Print(" ")
				line = append(line, " ")
			}

			arg := ctx.Args()[i].String()
			fmt.Print(arg)
			line = append(line, arg)
		}

		fmt.Println()
		q.output = append(q.output, line)
		return nil, nil
	})

	return nil
}

func (q *QJS) Close() error {
	q.output = nil
	if q.rt != nil {
		q.rt.Close()
	}

	q.rt = nil
	return nil
}

func (q *QJS) Run(inputFile string) ([][]string, error) {
	script, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", inputFile, err)
	}

	_, err = q.rt.Context().Eval("run.js", qjs.Code(string(script)), qjs.TypeGlobal())
	return q.output, err
}

var _ JSEngine = (*QJS)(nil)
