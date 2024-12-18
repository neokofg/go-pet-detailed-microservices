// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NewsColumns holds the columns for the "news" table.
	NewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "image_url", Type: field.TypeString, Nullable: true, Size: 2048},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// NewsTable holds the schema information for the "news" table.
	NewsTable = &schema.Table{
		Name:       "news",
		Columns:    NewsColumns,
		PrimaryKey: []*schema.Column{NewsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "news_user_id",
				Unique:  false,
				Columns: []*schema.Column{NewsColumns[4]},
			},
			{
				Name:    "news_created_at",
				Unique:  false,
				Columns: []*schema.Column{NewsColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NewsTable,
	}
)

func init() {
}
