// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package ref

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type newscategoriesTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *newscategoriesTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("newscategories").
func (v *newscategoriesTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *newscategoriesTableType) Columns() []string {
	return []string{
		"id",
		"news_id",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *newscategoriesTableType) NewStruct() reform.Struct {
	return new(Newscategories)
}

// NewRecord makes a new record for that table.
func (v *newscategoriesTableType) NewRecord() reform.Record {
	return new(Newscategories)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *newscategoriesTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// NewscategoriesTable represents newscategories view or table in SQL database.
var NewscategoriesTable = &newscategoriesTableType{
	s: parse.StructInfo{
		Type:    "Newscategories",
		SQLName: "newscategories",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int32", Column: "id"},
			{Name: "NewsID", Type: "int32", Column: "news_id"},
		},
		PKFieldIndex: 0,
	},
	z: new(Newscategories).Values(),
}

// String returns a string representation of this struct or record.
func (s Newscategories) String() string {
	res := make([]string, 2)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "NewsID: " + reform.Inspect(s.NewsID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Newscategories) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.NewsID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Newscategories) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.NewsID,
	}
}

// View returns View object for that struct.
func (s *Newscategories) View() reform.View {
	return NewscategoriesTable
}

// Table returns Table object for that record.
func (s *Newscategories) Table() reform.Table {
	return NewscategoriesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Newscategories) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Newscategories) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Newscategories) HasPK() bool {
	return s.ID != NewscategoriesTable.z[NewscategoriesTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Newscategories) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = NewscategoriesTable
	_ reform.Struct = (*Newscategories)(nil)
	_ reform.Table  = NewscategoriesTable
	_ reform.Record = (*Newscategories)(nil)
	_ fmt.Stringer  = (*Newscategories)(nil)
)

func init() {
	parse.AssertUpToDate(&NewscategoriesTable.s, new(Newscategories))
}
