// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"meh/ent/meh"
	"meh/ent/timeline"
	"meh/ent/user"

	"entgo.io/ent/dialect/sql"
)

// Timeline is the model entity for the Timeline schema.
type Timeline struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// MehID holds the value of the "meh_id" field.
	MehID uint64 `json:"meh_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimelineQuery when eager-loading is set.
	Edges TimelineEdges `json:"edges"`
}

// TimelineEdges holds the relations/edges for other nodes in the graph.
type TimelineEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Meh holds the value of the meh edge.
	Meh *Meh `json:"meh,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TimelineEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MehOrErr returns the Meh value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TimelineEdges) MehOrErr() (*Meh, error) {
	if e.loadedTypes[1] {
		if e.Meh == nil {
			// The edge meh was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: meh.Label}
		}
		return e.Meh, nil
	}
	return nil, &NotLoadedError{edge: "meh"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Timeline) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case timeline.FieldID, timeline.FieldUserID, timeline.FieldMehID:
			values[i] = new(sql.NullInt64)
		case timeline.FieldCreatedAt, timeline.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Timeline", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Timeline fields.
func (t *Timeline) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timeline.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = uint64(value.Int64)
		case timeline.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case timeline.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case timeline.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				t.UserID = uint64(value.Int64)
			}
		case timeline.FieldMehID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field meh_id", values[i])
			} else if value.Valid {
				t.MehID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Timeline entity.
func (t *Timeline) QueryUser() *UserQuery {
	return (&TimelineClient{config: t.config}).QueryUser(t)
}

// QueryMeh queries the "meh" edge of the Timeline entity.
func (t *Timeline) QueryMeh() *MehQuery {
	return (&TimelineClient{config: t.config}).QueryMeh(t)
}

// Update returns a builder for updating this Timeline.
// Note that you need to call Timeline.Unwrap() before calling this method if this Timeline
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Timeline) Update() *TimelineUpdateOne {
	return (&TimelineClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Timeline entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Timeline) Unwrap() *Timeline {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Timeline is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Timeline) String() string {
	var builder strings.Builder
	builder.WriteString("Timeline(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", t.UserID))
	builder.WriteString(", meh_id=")
	builder.WriteString(fmt.Sprintf("%v", t.MehID))
	builder.WriteByte(')')
	return builder.String()
}

// Timelines is a parsable slice of Timeline.
type Timelines []*Timeline

func (t Timelines) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
