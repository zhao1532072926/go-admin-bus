package bus

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFirstdddetailTable(ctx *context.Context) table.Table {

	firstDdDetail := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Int))

	info := firstDdDetail.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Sender", "sender", db.Varchar)
	info.AddField("Spu_id", "spu_id", db.Varchar)
	info.AddField("Spu_name", "spu_name", db.Text)
	info.AddField("Shop_id", "shop_id", db.Varchar)
	info.AddField("Shop_name", "shop_name", db.Text)
	info.AddField("Send_time", "send_time", db.Timestamp)
	info.AddField("Data_body", "data_body", db.Text)
	info.AddField("Remove", "remove", db.Int)
	info.AddField("Youhui", "youhui", db.Text)

	info.SetTable("first_dd_detail").SetTitle("Firstdddetail").SetDescription("Firstdddetail")

	formList := firstDdDetail.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Sender", "sender", db.Varchar, form.Text)
	formList.AddField("Spu_id", "spu_id", db.Varchar, form.Text)
	formList.AddField("Spu_name", "spu_name", db.Text, form.RichText)
	formList.AddField("Shop_id", "shop_id", db.Varchar, form.Text)
	formList.AddField("Shop_name", "shop_name", db.Text, form.RichText)
	formList.AddField("Send_time", "send_time", db.Timestamp, form.Datetime)
	formList.AddField("Data_body", "data_body", db.Text, form.RichText)
	formList.AddField("Remove", "remove", db.Int, form.Number)
	formList.AddField("Youhui", "youhui", db.Text, form.RichText)

	formList.SetTable("first_dd_detail").SetTitle("Firstdddetail").SetDescription("Firstdddetail")

	return firstDdDetail
}
