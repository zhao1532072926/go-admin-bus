package bus

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFirstusersTable(ctx *context.Context) table.Table {

	firstUsers := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Int))

	info := firstUsers.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Password", "password", db.Varchar)

	info.SetTable("first_users").SetTitle("Firstusers").SetDescription("Firstusers")

	formList := firstUsers.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)

	formList.SetTable("first_users").SetTitle("Firstusers").SetDescription("Firstusers")

	return firstUsers
}
