// Code generated by entc, DO NOT EDIT.

package booking

import (
	"time"
)

const (
	// Label holds the string label denoting the booking type in the database.
	Label = "booking"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBOOKINGDATE holds the string denoting the booking_date field in the database.
	FieldBOOKINGDATE = "booking_date"
	// FieldTIMELEFT holds the string denoting the time_left field in the database.
	FieldTIMELEFT = "time_left"

	// EdgeUsedby holds the string denoting the usedby edge name in mutations.
	EdgeUsedby = "usedby"
	// EdgeBook holds the string denoting the book edge name in mutations.
	EdgeBook = "book"
	// EdgeUsing holds the string denoting the using edge name in mutations.
	EdgeUsing = "using"

	// Table holds the table name of the booking in the database.
	Table = "bookings"
	// UsedbyTable is the table the holds the usedby relation/edge.
	UsedbyTable = "bookings"
	// UsedbyInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsedbyInverseTable = "users"
	// UsedbyColumn is the table column denoting the usedby relation/edge.
	UsedbyColumn = "USER_ID"
	// BookTable is the table the holds the book relation/edge.
	BookTable = "bookings"
	// BookInverseTable is the table name for the Bookingtype entity.
	// It exists in this package in order to avoid circular dependency with the "bookingtype" package.
	BookInverseTable = "bookingtypes"
	// BookColumn is the table column denoting the book relation/edge.
	BookColumn = "BOOKINGTYPE_ID"
	// UsingTable is the table the holds the using relation/edge.
	UsingTable = "bookings"
	// UsingInverseTable is the table name for the ClientEntity entity.
	// It exists in this package in order to avoid circular dependency with the "cliententity" package.
	UsingInverseTable = "client_entities"
	// UsingColumn is the table column denoting the using relation/edge.
	UsingColumn = "CLIENT_ID"
)

// Columns holds all SQL columns for booking fields.
var Columns = []string{
	FieldID,
	FieldBOOKINGDATE,
	FieldTIMELEFT,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Booking type.
var ForeignKeys = []string{
	"BOOKINGTYPE_ID",
	"CLIENT_ID",
	"USER_ID",
}

var (
	// DefaultBOOKINGDATE holds the default value on creation for the BOOKING_DATE field.
	DefaultBOOKINGDATE func() time.Time
)
