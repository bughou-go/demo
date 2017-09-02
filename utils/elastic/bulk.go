package elastic

type BulkResult struct {
	Errors bool                                `json:"errors"`
	Items  []map[string]map[string]interface{} `json:"items"`
}

func (es *ES) BulkCreate(path string, data []map[string]interface{}) error {
	if len(data) <= 0 {
		return nil
	}
	body, err := MakeBulkCreate(data)
	if err != nil {
		return err
	}
	return es.BulkDo(path, body, `create`, data)
}

func (es *ES) BulkUpdate(path string, data []map[string]interface{}) error {
	if len(data) <= 0 {
		return nil
	}
	return es.BulkDo(path, MakeBulkUpdate(data), `update`, data)
}

func (es *ES) BulkDo(path string, body, typ string, data []map[string]interface{}) error {
	result := BulkResult{}
	if err := es.client.PostJson(es.Uri(path+`/_bulk`), nil, body, &result); err != nil {
		return err
	}
	if !result.Errors {
		return nil
	}
	return bulkError{typ: typ, inputs: data, results: result.Items}
}
