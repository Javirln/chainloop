// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/robotaccount"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/google/uuid"
)

// RobotAccount is the model entity for the RobotAccount schema.
type RobotAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// RevokedAt holds the value of the "revoked_at" field.
	RevokedAt time.Time `json:"revoked_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RobotAccountQuery when eager-loading is set.
	Edges                  RobotAccountEdges `json:"edges"`
	workflow_robotaccounts *uuid.UUID
	selectValues           sql.SelectValues
}

// RobotAccountEdges holds the relations/edges for other nodes in the graph.
type RobotAccountEdges struct {
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RobotAccountEdges) WorkflowOrErr() (*Workflow, error) {
	if e.Workflow != nil {
		return e.Workflow, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: workflow.Label}
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RobotAccount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case robotaccount.FieldName:
			values[i] = new(sql.NullString)
		case robotaccount.FieldCreatedAt, robotaccount.FieldRevokedAt:
			values[i] = new(sql.NullTime)
		case robotaccount.FieldID:
			values[i] = new(uuid.UUID)
		case robotaccount.ForeignKeys[0]: // workflow_robotaccounts
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RobotAccount fields.
func (ra *RobotAccount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case robotaccount.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ra.ID = *value
			}
		case robotaccount.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ra.Name = value.String
			}
		case robotaccount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ra.CreatedAt = value.Time
			}
		case robotaccount.FieldRevokedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field revoked_at", values[i])
			} else if value.Valid {
				ra.RevokedAt = value.Time
			}
		case robotaccount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_robotaccounts", values[i])
			} else if value.Valid {
				ra.workflow_robotaccounts = new(uuid.UUID)
				*ra.workflow_robotaccounts = *value.S.(*uuid.UUID)
			}
		default:
			ra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RobotAccount.
// This includes values selected through modifiers, order, etc.
func (ra *RobotAccount) Value(name string) (ent.Value, error) {
	return ra.selectValues.Get(name)
}

// QueryWorkflow queries the "workflow" edge of the RobotAccount entity.
func (ra *RobotAccount) QueryWorkflow() *WorkflowQuery {
	return NewRobotAccountClient(ra.config).QueryWorkflow(ra)
}

// Update returns a builder for updating this RobotAccount.
// Note that you need to call RobotAccount.Unwrap() before calling this method if this RobotAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *RobotAccount) Update() *RobotAccountUpdateOne {
	return NewRobotAccountClient(ra.config).UpdateOne(ra)
}

// Unwrap unwraps the RobotAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *RobotAccount) Unwrap() *RobotAccount {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("ent: RobotAccount is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *RobotAccount) String() string {
	var builder strings.Builder
	builder.WriteString("RobotAccount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ra.ID))
	builder.WriteString("name=")
	builder.WriteString(ra.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ra.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("revoked_at=")
	builder.WriteString(ra.RevokedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RobotAccounts is a parsable slice of RobotAccount.
type RobotAccounts []*RobotAccount
