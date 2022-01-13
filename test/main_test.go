package main

import (
	"testing"

	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/sdk"
)

func TestMain(t *testing.T) {

	appId := ""
  appSecret := ""

	app, err := sdk.BuildTenantInternal(appId, appSecret)

	if err != nil {
		panic(err)
	}

	t.Log(app)

	appToken := ""
  tableId := ""
	filters := map[string]string{"发布日期": "TODAY()"}
	records, err := app.GetBitableRecords(appToken, tableId, filters, 1)

	if err != nil {
		panic(err)
	}
	if records.Code != 0 {
		panic(records.Msg)
	}

	t.Log(records.Data.Items)

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
	res, err := app.BatchCreateBitable(appToken, tableId, body)
	if err != nil {
		panic(err)
	}
	if res.Code != 0 {
		panic(res.Msg)
	}
	t.Log(res.Data)

	// 删除记录
	deleteReq := vo.BatchDeleteBitableReq{
		Data: []string{res.Data.Items[0].RecordId},
	}
	deleteRes, err := app.BatchDeleteBitable(appToken, tableId, deleteReq)
	if err != nil {
		panic(err)
	}
	if deleteRes.Code != 0 {
		panic(deleteRes.Msg)
	}
	t.Log(deleteRes.Data)

}
