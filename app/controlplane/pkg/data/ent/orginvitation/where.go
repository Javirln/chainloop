// Code generated by ent, DO NOT EDIT.

package orginvitation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/authz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLTE(FieldID, id))
}

// ReceiverEmail applies equality check predicate on the "receiver_email" field. It's identical to ReceiverEmailEQ.
func ReceiverEmail(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldReceiverEmail, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldCreatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldDeletedAt, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldOrganizationID, v))
}

// SenderID applies equality check predicate on the "sender_id" field. It's identical to SenderIDEQ.
func SenderID(v uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldSenderID, v))
}

// ReceiverEmailEQ applies the EQ predicate on the "receiver_email" field.
func ReceiverEmailEQ(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldReceiverEmail, v))
}

// ReceiverEmailNEQ applies the NEQ predicate on the "receiver_email" field.
func ReceiverEmailNEQ(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNEQ(FieldReceiverEmail, v))
}

// ReceiverEmailIn applies the In predicate on the "receiver_email" field.
func ReceiverEmailIn(vs ...string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIn(FieldReceiverEmail, vs...))
}

// ReceiverEmailNotIn applies the NotIn predicate on the "receiver_email" field.
func ReceiverEmailNotIn(vs ...string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotIn(FieldReceiverEmail, vs...))
}

// ReceiverEmailGT applies the GT predicate on the "receiver_email" field.
func ReceiverEmailGT(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGT(FieldReceiverEmail, v))
}

// ReceiverEmailGTE applies the GTE predicate on the "receiver_email" field.
func ReceiverEmailGTE(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGTE(FieldReceiverEmail, v))
}

// ReceiverEmailLT applies the LT predicate on the "receiver_email" field.
func ReceiverEmailLT(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLT(FieldReceiverEmail, v))
}

// ReceiverEmailLTE applies the LTE predicate on the "receiver_email" field.
func ReceiverEmailLTE(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLTE(FieldReceiverEmail, v))
}

// ReceiverEmailContains applies the Contains predicate on the "receiver_email" field.
func ReceiverEmailContains(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldContains(FieldReceiverEmail, v))
}

// ReceiverEmailHasPrefix applies the HasPrefix predicate on the "receiver_email" field.
func ReceiverEmailHasPrefix(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldHasPrefix(FieldReceiverEmail, v))
}

// ReceiverEmailHasSuffix applies the HasSuffix predicate on the "receiver_email" field.
func ReceiverEmailHasSuffix(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldHasSuffix(FieldReceiverEmail, v))
}

// ReceiverEmailEqualFold applies the EqualFold predicate on the "receiver_email" field.
func ReceiverEmailEqualFold(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEqualFold(FieldReceiverEmail, v))
}

// ReceiverEmailContainsFold applies the ContainsFold predicate on the "receiver_email" field.
func ReceiverEmailContainsFold(v string) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldContainsFold(FieldReceiverEmail, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v biz.OrgInvitationStatus) predicate.OrgInvitation {
	vc := v
	return predicate.OrgInvitation(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v biz.OrgInvitationStatus) predicate.OrgInvitation {
	vc := v
	return predicate.OrgInvitation(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...biz.OrgInvitationStatus) predicate.OrgInvitation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrgInvitation(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...biz.OrgInvitationStatus) predicate.OrgInvitation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrgInvitation(sql.FieldNotIn(FieldStatus, v...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLTE(FieldCreatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotNull(FieldDeletedAt))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// SenderIDEQ applies the EQ predicate on the "sender_id" field.
func SenderIDEQ(v uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldEQ(FieldSenderID, v))
}

// SenderIDNEQ applies the NEQ predicate on the "sender_id" field.
func SenderIDNEQ(v uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNEQ(FieldSenderID, v))
}

// SenderIDIn applies the In predicate on the "sender_id" field.
func SenderIDIn(vs ...uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIn(FieldSenderID, vs...))
}

// SenderIDNotIn applies the NotIn predicate on the "sender_id" field.
func SenderIDNotIn(vs ...uuid.UUID) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotIn(FieldSenderID, vs...))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v authz.Role) predicate.OrgInvitation {
	vc := v
	return predicate.OrgInvitation(sql.FieldEQ(FieldRole, vc))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v authz.Role) predicate.OrgInvitation {
	vc := v
	return predicate.OrgInvitation(sql.FieldNEQ(FieldRole, vc))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...authz.Role) predicate.OrgInvitation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrgInvitation(sql.FieldIn(FieldRole, v...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...authz.Role) predicate.OrgInvitation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrgInvitation(sql.FieldNotIn(FieldRole, v...))
}

// RoleIsNil applies the IsNil predicate on the "role" field.
func RoleIsNil() predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldIsNull(FieldRole))
}

// RoleNotNil applies the NotNil predicate on the "role" field.
func RoleNotNil() predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.FieldNotNull(FieldRole))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.OrgInvitation {
	return predicate.OrgInvitation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.OrgInvitation {
	return predicate.OrgInvitation(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSender applies the HasEdge predicate on the "sender" edge.
func HasSender() predicate.OrgInvitation {
	return predicate.OrgInvitation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SenderTable, SenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSenderWith applies the HasEdge predicate on the "sender" edge with a given conditions (other predicates).
func HasSenderWith(preds ...predicate.User) predicate.OrgInvitation {
	return predicate.OrgInvitation(func(s *sql.Selector) {
		step := newSenderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrgInvitation) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrgInvitation) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrgInvitation) predicate.OrgInvitation {
	return predicate.OrgInvitation(sql.NotPredicates(p))
}
