// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/robotaccount"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/google/uuid"
)

// RobotAccountCreate is the builder for creating a RobotAccount entity.
type RobotAccountCreate struct {
	config
	mutation *RobotAccountMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rac *RobotAccountCreate) SetName(s string) *RobotAccountCreate {
	rac.mutation.SetName(s)
	return rac
}

// SetCreatedAt sets the "created_at" field.
func (rac *RobotAccountCreate) SetCreatedAt(t time.Time) *RobotAccountCreate {
	rac.mutation.SetCreatedAt(t)
	return rac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rac *RobotAccountCreate) SetNillableCreatedAt(t *time.Time) *RobotAccountCreate {
	if t != nil {
		rac.SetCreatedAt(*t)
	}
	return rac
}

// SetRevokedAt sets the "revoked_at" field.
func (rac *RobotAccountCreate) SetRevokedAt(t time.Time) *RobotAccountCreate {
	rac.mutation.SetRevokedAt(t)
	return rac
}

// SetNillableRevokedAt sets the "revoked_at" field if the given value is not nil.
func (rac *RobotAccountCreate) SetNillableRevokedAt(t *time.Time) *RobotAccountCreate {
	if t != nil {
		rac.SetRevokedAt(*t)
	}
	return rac
}

// SetID sets the "id" field.
func (rac *RobotAccountCreate) SetID(u uuid.UUID) *RobotAccountCreate {
	rac.mutation.SetID(u)
	return rac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rac *RobotAccountCreate) SetNillableID(u *uuid.UUID) *RobotAccountCreate {
	if u != nil {
		rac.SetID(*u)
	}
	return rac
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (rac *RobotAccountCreate) SetWorkflowID(id uuid.UUID) *RobotAccountCreate {
	rac.mutation.SetWorkflowID(id)
	return rac
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (rac *RobotAccountCreate) SetNillableWorkflowID(id *uuid.UUID) *RobotAccountCreate {
	if id != nil {
		rac = rac.SetWorkflowID(*id)
	}
	return rac
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (rac *RobotAccountCreate) SetWorkflow(w *Workflow) *RobotAccountCreate {
	return rac.SetWorkflowID(w.ID)
}

// Mutation returns the RobotAccountMutation object of the builder.
func (rac *RobotAccountCreate) Mutation() *RobotAccountMutation {
	return rac.mutation
}

// Save creates the RobotAccount in the database.
func (rac *RobotAccountCreate) Save(ctx context.Context) (*RobotAccount, error) {
	rac.defaults()
	return withHooks(ctx, rac.sqlSave, rac.mutation, rac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rac *RobotAccountCreate) SaveX(ctx context.Context) *RobotAccount {
	v, err := rac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rac *RobotAccountCreate) Exec(ctx context.Context) error {
	_, err := rac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rac *RobotAccountCreate) ExecX(ctx context.Context) {
	if err := rac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rac *RobotAccountCreate) defaults() {
	if _, ok := rac.mutation.CreatedAt(); !ok {
		v := robotaccount.DefaultCreatedAt()
		rac.mutation.SetCreatedAt(v)
	}
	if _, ok := rac.mutation.ID(); !ok {
		v := robotaccount.DefaultID()
		rac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rac *RobotAccountCreate) check() error {
	if _, ok := rac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "RobotAccount.name"`)}
	}
	if _, ok := rac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RobotAccount.created_at"`)}
	}
	return nil
}

func (rac *RobotAccountCreate) sqlSave(ctx context.Context) (*RobotAccount, error) {
	if err := rac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rac.driver, _spec); err != nil {
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
	rac.mutation.id = &_node.ID
	rac.mutation.done = true
	return _node, nil
}

func (rac *RobotAccountCreate) createSpec() (*RobotAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &RobotAccount{config: rac.config}
		_spec = sqlgraph.NewCreateSpec(robotaccount.Table, sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID))
	)
	if id, ok := rac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rac.mutation.Name(); ok {
		_spec.SetField(robotaccount.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rac.mutation.CreatedAt(); ok {
		_spec.SetField(robotaccount.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rac.mutation.RevokedAt(); ok {
		_spec.SetField(robotaccount.FieldRevokedAt, field.TypeTime, value)
		_node.RevokedAt = value
	}
	if nodes := rac.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   robotaccount.WorkflowTable,
			Columns: []string{robotaccount.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_robotaccounts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RobotAccountCreateBulk is the builder for creating many RobotAccount entities in bulk.
type RobotAccountCreateBulk struct {
	config
	err      error
	builders []*RobotAccountCreate
}

// Save creates the RobotAccount entities in the database.
func (racb *RobotAccountCreateBulk) Save(ctx context.Context) ([]*RobotAccount, error) {
	if racb.err != nil {
		return nil, racb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(racb.builders))
	nodes := make([]*RobotAccount, len(racb.builders))
	mutators := make([]Mutator, len(racb.builders))
	for i := range racb.builders {
		func(i int, root context.Context) {
			builder := racb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RobotAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, racb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, racb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, racb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (racb *RobotAccountCreateBulk) SaveX(ctx context.Context) []*RobotAccount {
	v, err := racb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (racb *RobotAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := racb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (racb *RobotAccountCreateBulk) ExecX(ctx context.Context) {
	if err := racb.Exec(ctx); err != nil {
		panic(err)
	}
}
