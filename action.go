package pipeline

type ingest func() (interface{}, error)
type transform func(interface{}) (interface{}, error)
type publish func(interface{}) error
