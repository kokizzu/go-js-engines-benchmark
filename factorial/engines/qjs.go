package engines

import "github.com/fastschema/qjs"

type QJS struct {
	rt *qjs.Runtime
}

func (q *QJS) Name() string {
	return "QJS"
}

func (q *QJS) Init() (err error) {
	q.rt, err = qjs.New()
	return err
}

func (q *QJS) Close() error {
	if q.rt != nil {
		q.rt.Close()
	}

	q.rt = nil
	return nil
}

func (q *QJS) Run(input string) error {
	_, err := q.rt.Context().Eval("<factorial>", qjs.Code(input))
	return err
}

var _ JSEngine = (*QJS)(nil)
