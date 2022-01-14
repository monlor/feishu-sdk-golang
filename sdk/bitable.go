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
func (t Tenant) GetBitableRecords(appToken string, tableId string, filters map[string]string, pageSize int) (*vo.GetBitableRecordsResp, error) {
	filterStr := ""
	pageSizeStr := ""
	if filters != nil {
		filterStr = "AND("
		i := 0
		for k, v := range filters {
			i++
			filterStr += fmt.Sprintf("CurrentValue.[%s]=%s", k, v)
			if i != len(filters) {
				filterStr += ","
			}
		}
		filterStr += ")"
	}
	if pageSize != 0 {
		pageSizeStr = fmt.Sprintf("%d", pageSize)
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiBitableGetRecords + "?filter=%s&page_size=%v", appToken, tableId, encrypt.URLEncode(filterStr), pageSizeStr), nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetBitableRecordsResp{}
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
