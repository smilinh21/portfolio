// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Portfolio/ent/formdata"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// FormData is the model entity for the FormData schema.
type FormData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FormData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case formdata.FieldID:
			values[i] = new(sql.NullInt64)
		case formdata.FieldName, formdata.FieldEmail, formdata.FieldMessage:
			values[i] = new(sql.NullString)
		case formdata.FieldCreatedAt, formdata.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FormData fields.
func (fd *FormData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case formdata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fd.ID = int(value.Int64)
		case formdata.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fd.Name = value.String
			}
		case formdata.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				fd.Email = value.String
			}
		case formdata.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				fd.Message = value.String
			}
		case formdata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fd.CreatedAt = value.Time
			}
		case formdata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fd.UpdatedAt = value.Time
			}
		default:
			fd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FormData.
// This includes values selected through modifiers, order, etc.
func (fd *FormData) Value(name string) (ent.Value, error) {
	return fd.selectValues.Get(name)
}

// Update returns a builder for updating this FormData.
// Note that you need to call FormData.Unwrap() before calling this method if this FormData
// was returned from a transaction, and the transaction was committed or rolled back.
func (fd *FormData) Update() *FormDataUpdateOne {
	return NewFormDataClient(fd.config).UpdateOne(fd)
}

// Unwrap unwraps the FormData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fd *FormData) Unwrap() *FormData {
	_tx, ok := fd.config.driver.(*txDriver)
	if !ok {
		panic("ent: FormData is not a transactional entity")
	}
	fd.config.driver = _tx.drv
	return fd
}

// String implements the fmt.Stringer.
func (fd *FormData) String() string {
	var builder strings.Builder
	builder.WriteString("FormData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fd.ID))
	builder.WriteString("name=")
	builder.WriteString(fd.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(fd.Email)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(fd.Message)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fd.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FormDataSlice is a parsable slice of FormData.
type FormDataSlice []*FormData
