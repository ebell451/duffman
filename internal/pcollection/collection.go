package pcollection

// filtered out version
type Collection struct {
	Variables []KeyValue
	Requests  []Req
	Env       []KeyValue
	Schema    Schema
}

type Req struct {
	Method      string
	URL         string
	Headers     map[string]string
	Body        string
	ContentType string
	Parameters  Parameters
}

type Parameters struct {
	Get  map[string]string
	Post map[string]string
	Path map[string]string
}
