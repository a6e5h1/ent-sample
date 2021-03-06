// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "content", Type: field.TypeString, Size: 500},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime, Nullable: true},
		{Name: "post_id", Type: field.TypeUint64},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_posts_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "body", Type: field.TypeString, Size: 10000},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "user_name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		PostsTable,
		UsersTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = PostsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
}
