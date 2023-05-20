package validator

type DataCollector interface {
	Set(key string, value any)
	Get(key string) (value any)
	Has(key string) bool
}

func NewMapDataCollector() mapDataCollector {
	return mapDataCollector{}
}

type mapDataCollector map[string]any

func (c mapDataCollector) Set(key string, value any) {
	c[key] = value
}

func (c mapDataCollector) Get(key string) any {
	return c[key]
}

func (c mapDataCollector) Has(key string) bool {
	_, exists := c[key]

	return exists
}
