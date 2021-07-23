package pipeline

// Pipeline represents a stream processing pipeline
type Pipeline struct {
	dangling chan interface{} // reference to last unwired output channel
	stages   []*stage
}

// New is the pipeline constructor
func New(ingest func(chan interface{})) *Pipeline {
	p := &Pipeline{
		dangling: make(chan interface{}),
		stages:   []*stage{},
	}

	go ingest(p.dangling)

	return p
}

// AddStage adds a stage to the processing pipeline
func (p *Pipeline) AddStage(name string, action action) {
	p.stages = append(p.stages, newStage(name, action, p.dangling))
	p.dangling = p.stages[len(p.stages)-1].out
}

// Run kicks-off pipeline stage threads (in reverse order)
// then blocks until the dangling output channel is closed
func (p *Pipeline) Run() {
	for i := len(p.stages) - 1; i >= 0; i-- {
		// TODO: make each stage multithreaded
		go p.stages[i].run()
	}

	// drain dangling output
	for {
		if _, ok := <-p.dangling; !ok {
			break
		}
	}
}
