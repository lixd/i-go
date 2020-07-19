package filter

type StraightPipeline struct {
	Name    string
	Filters []IFilter
}

func NewStraightPipeline(name string, filters ...IFilter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: filters,
	}
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var (
		ret interface{}
		err error
	)
	for _, filter := range f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, nil
}
