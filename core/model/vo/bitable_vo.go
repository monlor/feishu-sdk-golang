package vo

type BitableRecords struct {
	Fields   map[string]interface{} `json:"fields"`
	RecordId string                 `json:"record_id"`
}

type GetBitableRecordsResp struct {
	CommonVo
	Data GetBitableRecords `json:"data"`
}

type GetBitableRecords struct {
	PageToken string           `json:"page_token"`
	Total     int64            `json:"total"`
	HasMore   bool             `json:"has_more"`
	Items     []BitableRecords `json:"items"`
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
