package main

import (
	"testing"

	"github.com/monlor/feishu-sdk-golang/core/model/vo"
	"github.com/monlor/feishu-sdk-golang/sdk"
)

func TestMain(t *testing.T) {

	appId := "cli_a15ea2842178d00d"
  appSecret := "2fGPk5EFRGtM0fxYTWrxLdEofjrezVsn"

	app, err := sdk.BuildTenantInternal(appId, appSecret)

	if err != nil {
		panic(err)
	}

	t.Log(app)

	appToken := "bascnyJSvpEhlKIiAxpBxONMDEg"
  tableId := "tblAbvmLn5hTMIvR"
  var filters map[string]interface{}
	fieldNames := []string{"修复标题"}
	records, err := app.ListBitableRecords(appToken, tableId, filters, fieldNames, 1)

	if err != nil {
		panic(err)
	}
	if records.Code != 0 {
		panic(records.Msg)
	}

	t.Log(records.Data.Items)

	// 检索测试
	res, err := app.GetBitableRecord(appToken, tableId, "reckBZg79B")
	if err != nil {
		panic(err)
	}
	t.Log(res)

	//列出字段
	res2, err := app.ListBitableFields(appToken, tableId, 0)
	t.Log(res2)

	// 更新记录
	data := []vo.BitableRecords{
		vo.BitableRecords{
			// RecordId: "reckBZg79B",
			Fields: map[string]interface{}{"修复标题": "测试222", "模块": "base", "环境": "test", "发布日期": 1642057709000},
		},
	}
	body := vo.BatchUpdateBitableReq{
		Data: data,
	}
	t.Logf("请求体：%v", body)
	res3, err := app.BatchCreateBitable(appToken, tableId, body)
	if err != nil {
		panic(err)
	}
	if res3.Code != 0 {
		panic(res3.Msg)
	}
	t.Log(res3.Data)

	// 删除记录
	// deleteReq := vo.BatchDeleteBitableReq{
	// 	Data: []string{res.Data.Items[0].RecordId},
	// }
	// deleteRes, err := app.BatchDeleteBitable(appToken, tableId, deleteReq)
	// if err != nil {
	// 	panic(err)
	// }
	// if deleteRes.Code != 0 {
	// 	panic(deleteRes.Msg)
	// }
	// t.Log(deleteRes.Data)

}
