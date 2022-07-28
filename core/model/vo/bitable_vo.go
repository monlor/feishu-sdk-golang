package vo

type BitableRecords struct {
	Fields   map[string]interface{} `json:"fields"`
	RecordId string                 `json:"record_id"`
}

type ListBitableRecordsResp struct {
	CommonVo
	Data ListBitableRecords `json:"data"`
}

type ListBitableRecords struct {
	PageToken string           `json:"page_token"`
	Total     int64            `json:"total"`
	HasMore   bool             `json:"has_more"`
	Items     []BitableRecords `json:"items"`
}

type GetBitableRecordResp struct {
	CommonVo
	Data GetBitableRecord `json:"data"`
}

type GetBitableRecord struct {
	Record BitableRecords `json:"record"`
}

type BatchUpdateBitableResp struct {
	CommonVo
	Data BatchUpdateBitable `json:"data"`
}

type BatchUpdateBitable struct {
	Items []BitableRecords `json:"records"`
}

type BatchUpdateBitableReq struct {
	Data []BitableRecords `json:"records"`
}

type BatchDeleteBitableReq struct {
	Data []string `json:"records"`
}

type BatchDeleteBitableResp struct {
	CommonVo
	Data BatchDeleteBitableData `json:"data"`
}

type BatchDeleteBitableData struct {
	Records []BatchDeleteBitable `json:"records"`
}

type BatchDeleteBitable struct {
	Deleted  bool   `json:"deleted"`
	RecordId string `json:"record_id"`
}

type ListBitableFieldsResp struct {
	CommonVo 
	Data ListBitableFieldsData
}

type ListBitableFieldsData struct {
	PageToken string           `json:"page_token"`
	Total     int64            `json:"total"`
	HasMore   bool             `json:"has_more"`
	Items     []ListBitableFields `json:"items"`
}

type ListBitableFields struct {
	FieldId string `json:"field_id"`
	FieldName string `json:"field_name"`
	Type int `json:"type"`
	Property map[string]interface{} `json:"property"`
}

type ListBitableFieldsProperty struct {
	Options []ListBitableFieldsOptions `json:"options"`
	Formatter string `json:"formatter"`
	DateFormatter string `json:"date_formatter"`
	AutoFill string `json:"auto_fill"`
	Multiple string `json:"multiple"`
	TableId string `json:"table_id"`
	TableName string `json:"table_name"`
	BackFieldName string `json:"back_field_name"`
}

type ListBitableFieldsOptions struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Color string `json:"color"`
}