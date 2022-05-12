package saga

type Saga struct {
	name string
	steps []Step
}

func (s *Saga) SetName(name string) {
	s.name = name
}

func (s *Saga) AddStep(step Step) {
	s.steps = append(s.steps, step)
}

type Step struct {
	Name 			string
	Func 			interface{}
	Compensation 	interface{}
}