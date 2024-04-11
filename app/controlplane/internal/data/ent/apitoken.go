// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/apitoken"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/google/uuid"
)

// APIToken is the model entity for the APIToken schema.
type APIToken struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// RevokedAt holds the value of the "revoked_at" field.
	RevokedAt time.Time `json:"revoked_at,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organization_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the APITokenQuery when eager-loading is set.
	Edges        APITokenEdges `json:"edges"`
	selectValues sql.SelectValues
}

// APITokenEdges holds the relations/edges for other nodes in the graph.
type APITokenEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APITokenEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*APIToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apitoken.FieldName, apitoken.FieldDescription:
			values[i] = new(sql.NullString)
		case apitoken.FieldCreatedAt, apitoken.FieldExpiresAt, apitoken.FieldRevokedAt:
			values[i] = new(sql.NullTime)
		case apitoken.FieldID, apitoken.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the APIToken fields.
func (at *APIToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apitoken.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				at.ID = *value
			}
		case apitoken.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				at.Name = value.String
			}
		case apitoken.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				at.Description = value.String
			}
		case apitoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				at.CreatedAt = value.Time
			}
		case apitoken.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				at.ExpiresAt = value.Time
			}
		case apitoken.FieldRevokedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field revoked_at", values[i])
			} else if value.Valid {
				at.RevokedAt = value.Time
			}
		case apitoken.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				at.OrganizationID = *value
			}
		default:
			at.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the APIToken.
// This includes values selected through modifiers, order, etc.
func (at *APIToken) Value(name string) (ent.Value, error) {
	return at.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the APIToken entity.
func (at *APIToken) QueryOrganization() *OrganizationQuery {
	return NewAPITokenClient(at.config).QueryOrganization(at)
}

// Update returns a builder for updating this APIToken.
// Note that you need to call APIToken.Unwrap() before calling this method if this APIToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *APIToken) Update() *APITokenUpdateOne {
	return NewAPITokenClient(at.config).UpdateOne(at)
}

// Unwrap unwraps the APIToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *APIToken) Unwrap() *APIToken {
	_tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: APIToken is not a transactional entity")
	}
	at.config.driver = _tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *APIToken) String() string {
	var builder strings.Builder
	builder.WriteString("APIToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", at.ID))
	builder.WriteString("name=")
	builder.WriteString(at.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(at.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(at.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(at.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("revoked_at=")
	builder.WriteString(at.RevokedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", at.OrganizationID))
	builder.WriteByte(')')
	return builder.String()
}

// APITokens is a parsable slice of APIToken.
type APITokens []*APIToken
