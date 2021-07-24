package pipeline

import "log"

// Pipeline represents a stream processing pipeline
type Pipeline struct {
	source *source
	stages []*stage
	sink   *sink
}

// New is the pipeline constructor
func New() *Pipeline {
	return &Pipeline{stages: []*stage{}}
}

// AddStage adds a stage to the processing pipeline
func (p *Pipeline) AddStage(name string, threads int, transform transform) {
	if p.sink != nil {
		log.Fatal("cannot add stages after sink is set")
	}

	inputChan := p.source.out
	if len(p.stages) > 0 {
		inputChan = p.stages[len(p.stages)-1].out
	}

	p.stages = append(p.stages, newStage(name, threads, transform, inputChan))
}

// SetSource sets data ingestion source in the pipeline
func (p *Pipeline) SetSource(name string, ingest ingest) {
	p.source = newSource(name, ingest)
}

// SetSink sets data sink in the pipeline
func (p *Pipeline) SetSink(name string, commit publish) {
	if len(p.stages) == 0 {
		p.sink = newSink(name, commit, p.source.out)
	}
	p.sink = newSink(name, commit, p.stages[len(p.stages)-1].out)
}

// Run runs the pipeline and blocks until done
func (p *Pipeline) Run() {
	go p.sink.run()
	for i := len(p.stages) - 1; i >= 0; i-- {
		go p.stages[i].run()
	}
	go p.source.run()

	<-p.sink.done
}
