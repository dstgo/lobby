// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dstgo/lobby/server/data/ent/cronjob"
)

// CronJob is the model entity for the CronJob schema.
type CronJob struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Cron holds the value of the "cron" field.
	Cron string `json:"cron,omitempty"`
	// EntryID holds the value of the "entry_id" field.
	EntryID int `json:"entry_id,omitempty"`
	// Prev holds the value of the "prev" field.
	Prev int64 `json:"prev,omitempty"`
	// Next holds the value of the "next" field.
	Next         int64 `json:"next,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CronJob) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cronjob.FieldID, cronjob.FieldEntryID, cronjob.FieldPrev, cronjob.FieldNext:
			values[i] = new(sql.NullInt64)
		case cronjob.FieldName, cronjob.FieldCron:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CronJob fields.
func (cj *CronJob) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cronjob.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cj.ID = int(value.Int64)
		case cronjob.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cj.Name = value.String
			}
		case cronjob.FieldCron:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron", values[i])
			} else if value.Valid {
				cj.Cron = value.String
			}
		case cronjob.FieldEntryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entry_id", values[i])
			} else if value.Valid {
				cj.EntryID = int(value.Int64)
			}
		case cronjob.FieldPrev:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field prev", values[i])
			} else if value.Valid {
				cj.Prev = value.Int64
			}
		case cronjob.FieldNext:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field next", values[i])
			} else if value.Valid {
				cj.Next = value.Int64
			}
		default:
			cj.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CronJob.
// This includes values selected through modifiers, order, etc.
func (cj *CronJob) Value(name string) (ent.Value, error) {
	return cj.selectValues.Get(name)
}

// Update returns a builder for updating this CronJob.
// Note that you need to call CronJob.Unwrap() before calling this method if this CronJob
// was returned from a transaction, and the transaction was committed or rolled back.
func (cj *CronJob) Update() *CronJobUpdateOne {
	return NewCronJobClient(cj.config).UpdateOne(cj)
}

// Unwrap unwraps the CronJob entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cj *CronJob) Unwrap() *CronJob {
	_tx, ok := cj.config.driver.(*txDriver)
	if !ok {
		panic("ent: CronJob is not a transactional entity")
	}
	cj.config.driver = _tx.drv
	return cj
}

// String implements the fmt.Stringer.
func (cj *CronJob) String() string {
	var builder strings.Builder
	builder.WriteString("CronJob(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cj.ID))
	builder.WriteString("name=")
	builder.WriteString(cj.Name)
	builder.WriteString(", ")
	builder.WriteString("cron=")
	builder.WriteString(cj.Cron)
	builder.WriteString(", ")
	builder.WriteString("entry_id=")
	builder.WriteString(fmt.Sprintf("%v", cj.EntryID))
	builder.WriteString(", ")
	builder.WriteString("prev=")
	builder.WriteString(fmt.Sprintf("%v", cj.Prev))
	builder.WriteString(", ")
	builder.WriteString("next=")
	builder.WriteString(fmt.Sprintf("%v", cj.Next))
	builder.WriteByte(')')
	return builder.String()
}

// CronJobs is a parsable slice of CronJob.
type CronJobs []*CronJob
