package saga

import (
	"errors"
	"reflect"
)

type Coordinator struct {
	saga 	Saga
	ended 	bool
}

func NewCoordinator(s Saga) *Coordinator {
	return &Coordinator{
		saga: s,
	}
}

func (c *Coordinator) Commit() error {
	if c.ended == true {
		return nil
	}

	for i := 0; i < len(c.saga.steps); i++ {
		err := c.execStep(i)
		if err != nil {
			return err
		}
	}

	c.ended = true
	return nil
}

func (c *Coordinator) execStep(i int) error {
	f := c.saga.steps[i].Func
	v := reflect.ValueOf(f)

	if v.Kind() != reflect.Func {
		return errors.New("func must be function")
	}

	r := v.Call([]reflect.Value{})
	err := c.isError(r)
	if err != nil {
		c.abort(i)

		return err
	}

	return nil
}

func (c *Coordinator) abort(from int) {
	for i := from; i >= 0; i-- {
		err := c.abortStep(i)

		if err != nil {
			continue
		}
	}
}

func (c *Coordinator) abortStep(step int) error {
	f := c.saga.steps[step].Compensation
	v := reflect.ValueOf(f)

	if v.Kind() != reflect.Func {
		return errors.New("compensation must be function")
	}

	v.Call([]reflect.Value{})

	return nil
}

func (c *Coordinator) isError(r []reflect.Value) error {
	if len(r) > 0 && !r[len(r)-1].IsNil() {
		return r[len(r)-1].Interface().(error)
	}

	return nil
}