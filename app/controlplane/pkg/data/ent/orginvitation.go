// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/authz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/orginvitation"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/user"
	"github.com/google/uuid"
)

// OrgInvitation is the model entity for the OrgInvitation schema.
type OrgInvitation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ReceiverEmail holds the value of the "receiver_email" field.
	ReceiverEmail string `json:"receiver_email,omitempty"`
	// Status holds the value of the "status" field.
	Status biz.OrgInvitationStatus `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organization_id,omitempty"`
	// SenderID holds the value of the "sender_id" field.
	SenderID uuid.UUID `json:"sender_id,omitempty"`
	// Role holds the value of the "role" field.
	Role authz.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgInvitationQuery when eager-loading is set.
	Edges        OrgInvitationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrgInvitationEdges holds the relations/edges for other nodes in the graph.
type OrgInvitationEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Sender holds the value of the sender edge.
	Sender *User `json:"sender,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgInvitationEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgInvitationEdges) SenderOrErr() (*User, error) {
	if e.Sender != nil {
		return e.Sender, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgInvitation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orginvitation.FieldReceiverEmail, orginvitation.FieldStatus, orginvitation.FieldRole:
			values[i] = new(sql.NullString)
		case orginvitation.FieldCreatedAt, orginvitation.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case orginvitation.FieldID, orginvitation.FieldOrganizationID, orginvitation.FieldSenderID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgInvitation fields.
func (oi *OrgInvitation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orginvitation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				oi.ID = *value
			}
		case orginvitation.FieldReceiverEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_email", values[i])
			} else if value.Valid {
				oi.ReceiverEmail = value.String
			}
		case orginvitation.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				oi.Status = biz.OrgInvitationStatus(value.String)
			}
		case orginvitation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oi.CreatedAt = value.Time
			}
		case orginvitation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				oi.DeletedAt = value.Time
			}
		case orginvitation.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				oi.OrganizationID = *value
			}
		case orginvitation.FieldSenderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field sender_id", values[i])
			} else if value != nil {
				oi.SenderID = *value
			}
		case orginvitation.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				oi.Role = authz.Role(value.String)
			}
		default:
			oi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrgInvitation.
// This includes values selected through modifiers, order, etc.
func (oi *OrgInvitation) Value(name string) (ent.Value, error) {
	return oi.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the OrgInvitation entity.
func (oi *OrgInvitation) QueryOrganization() *OrganizationQuery {
	return NewOrgInvitationClient(oi.config).QueryOrganization(oi)
}

// QuerySender queries the "sender" edge of the OrgInvitation entity.
func (oi *OrgInvitation) QuerySender() *UserQuery {
	return NewOrgInvitationClient(oi.config).QuerySender(oi)
}

// Update returns a builder for updating this OrgInvitation.
// Note that you need to call OrgInvitation.Unwrap() before calling this method if this OrgInvitation
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrgInvitation) Update() *OrgInvitationUpdateOne {
	return NewOrgInvitationClient(oi.config).UpdateOne(oi)
}

// Unwrap unwraps the OrgInvitation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrgInvitation) Unwrap() *OrgInvitation {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgInvitation is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrgInvitation) String() string {
	var builder strings.Builder
	builder.WriteString("OrgInvitation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("receiver_email=")
	builder.WriteString(oi.ReceiverEmail)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", oi.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(oi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(oi.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("sender_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.SenderID))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", oi.Role))
	builder.WriteByte(')')
	return builder.String()
}

// OrgInvitations is a parsable slice of OrgInvitation.
type OrgInvitations []*OrgInvitation
