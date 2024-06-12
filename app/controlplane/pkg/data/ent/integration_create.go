// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/integration"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/integrationattachment"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/google/uuid"
)

// IntegrationCreate is the builder for creating a Integration entity.
type IntegrationCreate struct {
	config
	mutation *IntegrationMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *IntegrationCreate) SetName(s string) *IntegrationCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetKind sets the "kind" field.
func (ic *IntegrationCreate) SetKind(s string) *IntegrationCreate {
	ic.mutation.SetKind(s)
	return ic
}

// SetDescription sets the "description" field.
func (ic *IntegrationCreate) SetDescription(s string) *IntegrationCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableDescription(s *string) *IntegrationCreate {
	if s != nil {
		ic.SetDescription(*s)
	}
	return ic
}

// SetSecretName sets the "secret_name" field.
func (ic *IntegrationCreate) SetSecretName(s string) *IntegrationCreate {
	ic.mutation.SetSecretName(s)
	return ic
}

// SetCreatedAt sets the "created_at" field.
func (ic *IntegrationCreate) SetCreatedAt(t time.Time) *IntegrationCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableCreatedAt(t *time.Time) *IntegrationCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetConfiguration sets the "configuration" field.
func (ic *IntegrationCreate) SetConfiguration(b []byte) *IntegrationCreate {
	ic.mutation.SetConfiguration(b)
	return ic
}

// SetDeletedAt sets the "deleted_at" field.
func (ic *IntegrationCreate) SetDeletedAt(t time.Time) *IntegrationCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableDeletedAt(t *time.Time) *IntegrationCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IntegrationCreate) SetID(u uuid.UUID) *IntegrationCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableID(u *uuid.UUID) *IntegrationCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// AddAttachmentIDs adds the "attachments" edge to the IntegrationAttachment entity by IDs.
func (ic *IntegrationCreate) AddAttachmentIDs(ids ...uuid.UUID) *IntegrationCreate {
	ic.mutation.AddAttachmentIDs(ids...)
	return ic
}

// AddAttachments adds the "attachments" edges to the IntegrationAttachment entity.
func (ic *IntegrationCreate) AddAttachments(i ...*IntegrationAttachment) *IntegrationCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddAttachmentIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ic *IntegrationCreate) SetOrganizationID(id uuid.UUID) *IntegrationCreate {
	ic.mutation.SetOrganizationID(id)
	return ic
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ic *IntegrationCreate) SetOrganization(o *Organization) *IntegrationCreate {
	return ic.SetOrganizationID(o.ID)
}

// Mutation returns the IntegrationMutation object of the builder.
func (ic *IntegrationCreate) Mutation() *IntegrationMutation {
	return ic.mutation
}

// Save creates the Integration in the database.
func (ic *IntegrationCreate) Save(ctx context.Context) (*Integration, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IntegrationCreate) SaveX(ctx context.Context) *Integration {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IntegrationCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IntegrationCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IntegrationCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := integration.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := integration.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IntegrationCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Integration.name"`)}
	}
	if _, ok := ic.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "Integration.kind"`)}
	}
	if _, ok := ic.mutation.SecretName(); !ok {
		return &ValidationError{Name: "secret_name", err: errors.New(`ent: missing required field "Integration.secret_name"`)}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Integration.created_at"`)}
	}
	if _, ok := ic.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "Integration.organization"`)}
	}
	return nil
}

func (ic *IntegrationCreate) sqlSave(ctx context.Context) (*Integration, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IntegrationCreate) createSpec() (*Integration, *sqlgraph.CreateSpec) {
	var (
		_node = &Integration{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(integration.Table, sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(integration.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Kind(); ok {
		_spec.SetField(integration.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(integration.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.SecretName(); ok {
		_spec.SetField(integration.FieldSecretName, field.TypeString, value)
		_node.SecretName = value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(integration.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.Configuration(); ok {
		_spec.SetField(integration.FieldConfiguration, field.TypeBytes, value)
		_node.Configuration = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(integration.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := ic.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   integration.AttachmentsTable,
			Columns: []string{integration.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OrganizationTable,
			Columns: []string{integration.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_integrations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IntegrationCreateBulk is the builder for creating many Integration entities in bulk.
type IntegrationCreateBulk struct {
	config
	builders []*IntegrationCreate
}

// Save creates the Integration entities in the database.
func (icb *IntegrationCreateBulk) Save(ctx context.Context) ([]*Integration, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Integration, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IntegrationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IntegrationCreateBulk) SaveX(ctx context.Context) []*Integration {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IntegrationCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IntegrationCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}