// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BookingsColumns holds the columns for the "bookings" table.
	BookingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "booking_date", Type: field.TypeTime},
		{Name: "time_left", Type: field.TypeTime},
		{Name: "BOOKINGTYPE_ID", Type: field.TypeInt, Nullable: true},
		{Name: "CLIENT_ID", Type: field.TypeInt, Nullable: true},
		{Name: "USER_ID", Type: field.TypeInt, Nullable: true},
	}
	// BookingsTable holds the schema information for the "bookings" table.
	BookingsTable = &schema.Table{
		Name:       "bookings",
		Columns:    BookingsColumns,
		PrimaryKey: []*schema.Column{BookingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bookings_bookingtypes_booktype",
				Columns: []*schema.Column{BookingsColumns[3]},

				RefColumns: []*schema.Column{BookingtypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookings_client_entities_booked",
				Columns: []*schema.Column{BookingsColumns[4]},

				RefColumns: []*schema.Column{ClientEntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookings_users_booking",
				Columns: []*schema.Column{BookingsColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BookingtypesColumns holds the columns for the "bookingtypes" table.
	BookingtypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "booktype_name", Type: field.TypeString, Unique: true},
	}
	// BookingtypesTable holds the schema information for the "bookingtypes" table.
	BookingtypesTable = &schema.Table{
		Name:        "bookingtypes",
		Columns:     BookingtypesColumns,
		PrimaryKey:  []*schema.Column{BookingtypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ClientEntitiesColumns holds the columns for the "client_entities" table.
	ClientEntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "client_name", Type: field.TypeString, Unique: true},
		{Name: "client_status", Type: field.TypeString},
	}
	// ClientEntitiesTable holds the schema information for the "client_entities" table.
	ClientEntitiesTable = &schema.Table{
		Name:        "client_entities",
		Columns:     ClientEntitiesColumns,
		PrimaryKey:  []*schema.Column{ClientEntitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_name", Type: field.TypeString, Unique: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_email", Type: field.TypeString, Unique: true},
		{Name: "user_name", Type: field.TypeString},
		{Name: "ROLE_ID", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_roles_role",
				Columns: []*schema.Column{UsersColumns[3]},

				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BookingsTable,
		BookingtypesTable,
		ClientEntitiesTable,
		RolesTable,
		UsersTable,
	}
)

func init() {
	BookingsTable.ForeignKeys[0].RefTable = BookingtypesTable
	BookingsTable.ForeignKeys[1].RefTable = ClientEntitiesTable
	BookingsTable.ForeignKeys[2].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = RolesTable
}