package setcronjob

type response struct {
	status  string
	code    int
	data    interface{}
	info    []string
	message string
}

func (r *response) GetData() interface{} {
	return r.data
}
