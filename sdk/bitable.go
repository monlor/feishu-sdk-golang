package sdk

import (
	"fmt"

	"github.com/monlor/feishu-sdk-golang/core/consts"
	"github.com/monlor/feishu-sdk-golang/core/model/vo"
	"github.com/monlor/feishu-sdk-golang/core/util/encrypt"
	"github.com/monlor/feishu-sdk-golang/core/util/http"
	"github.com/monlor/feishu-sdk-golang/core/util/json"
	"github.com/monlor/feishu-sdk-golang/core/util/log"
)

//多维表格 列出记录 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list
func (t Tenant) ListBitableRecords(appToken string, tableId string, filters map[string]interface{}, fieldNames []string, pageSize int) (*vo.ListBitableRecordsResp, error) {
	filterStr := ""
	fieldNameStr := ""
	pageSizeStr := ""
	if filters != nil {
		if len(filters) == 0 {
			filterStr = ""
		} else {
			filterStr = "AND("
			i := 0
			for k, v := range filters {
				i++
				filterStr += fmt.Sprintf("CurrentValue.[%s]=%v", k, v)
				if i != len(filters) {
					filterStr += ","
				}
			}
			filterStr += ")"
		}
	}
	if fieldNames != nil {
		if len(fieldNames) == 0 {
			fieldNameStr = ""
		} else {
			fieldNameStr = "["
			for i, v := range fieldNames {
				fieldNameStr += "\"" + v + "\""
				if i+1 != len(fieldNames) {
					fieldNameStr += ","
				}
			}
			fieldNameStr += "]"
		}
	}
	if pageSize != 0 {
		pageSizeStr = fmt.Sprintf("%d", pageSize)
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiBitableListRecords + "?filter=%s&page_size=%v&field_names=%s", 
		appToken, tableId, encrypt.URLEncode(filterStr), pageSizeStr, encrypt.URLEncode(fieldNameStr)), 
		nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.ListBitableRecordsResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//多维表格 检索记录 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/get
func (t Tenant) GetBitableRecord(appToken string, tableId string, recordId string) (*vo.GetBitableRecordResp, error) {
	respBody, err := http.Get(fmt.Sprintf(consts.ApiBitableGetRecord, appToken, tableId, recordId), 
		nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetBitableRecordResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//多维表格 更新多条记录 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_update
func (t Tenant) BatchUpdateBitable(appToken string, tableId string, bodyParams vo.BatchUpdateBitableReq) (*vo.BatchUpdateBitableResp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiBitableBatchUpdate, appToken, tableId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.BatchUpdateBitableResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//多维表格 新增多条记录 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_create
func (t Tenant) BatchCreateBitable(appToken string, tableId string, bodyParams vo.BatchUpdateBitableReq) (*vo.BatchUpdateBitableResp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiBitableBatchCreate, appToken, tableId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.BatchUpdateBitableResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//多维表格 新增多条记录 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_delete
func (t Tenant) BatchDeleteBitable(appToken string, tableId string, bodyParams vo.BatchDeleteBitableReq) (*vo.BatchDeleteBitableResp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiBitableBatchDelete, appToken, tableId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.BatchDeleteBitableResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//多维表格 列出字段 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list
func (t Tenant) ListBitableFields(appToken string, tableId string, pageSize int) (*vo.ListBitableFieldsResp, error) {
	pageSizeStr := ""
	if (pageSize != 0) {
		pageSizeStr = fmt.Sprintf("%d", pageSize)
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiBitableListFields + "?page_size=%s", appToken, tableId, pageSizeStr), 
		nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.ListBitableFieldsResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}