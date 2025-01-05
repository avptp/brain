// Code generated by ent, DO NOT EDIT.

package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/google/uuid"
)

// AuthorizationCreate is the builder for creating a Authorization entity.
type AuthorizationCreate struct {
	config
	mutation *AuthorizationMutation
	hooks    []Hook
}

// SetPersonID sets the "person_id" field.
func (ac *AuthorizationCreate) SetPersonID(u uuid.UUID) *AuthorizationCreate {
	ac.mutation.SetPersonID(u)
	return ac
}

// SetToken sets the "token" field.
func (ac *AuthorizationCreate) SetToken(b []byte) *AuthorizationCreate {
	ac.mutation.SetToken(b)
	return ac
}

// SetKind sets the "kind" field.
func (ac *AuthorizationCreate) SetKind(a authorization.Kind) *AuthorizationCreate {
	ac.mutation.SetKind(a)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AuthorizationCreate) SetCreatedAt(t time.Time) *AuthorizationCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AuthorizationCreate) SetNillableCreatedAt(t *time.Time) *AuthorizationCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AuthorizationCreate) SetID(u uuid.UUID) *AuthorizationCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetPerson sets the "person" edge to the Person entity.
func (ac *AuthorizationCreate) SetPerson(p *Person) *AuthorizationCreate {
	return ac.SetPersonID(p.ID)
}

// Mutation returns the AuthorizationMutation object of the builder.
func (ac *AuthorizationCreate) Mutation() *AuthorizationMutation {
	return ac.mutation
}

// Save creates the Authorization in the database.
func (ac *AuthorizationCreate) Save(ctx context.Context) (*Authorization, error) {
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthorizationCreate) SaveX(ctx context.Context) *Authorization {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuthorizationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuthorizationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AuthorizationCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if authorization.DefaultCreatedAt == nil {
			return fmt.Errorf("data: uninitialized authorization.DefaultCreatedAt (forgotten import data/runtime?)")
		}
		v := authorization.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthorizationCreate) check() error {
	if _, ok := ac.mutation.PersonID(); !ok {
		return &ValidationError{Name: "person_id", err: errors.New(`data: missing required field "Authorization.person_id"`)}
	}
	if _, ok := ac.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`data: missing required field "Authorization.token"`)}
	}
	if _, ok := ac.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`data: missing required field "Authorization.kind"`)}
	}
	if v, ok := ac.mutation.Kind(); ok {
		if err := authorization.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`data: validator failed for field "Authorization.kind": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`data: missing required field "Authorization.created_at"`)}
	}
	if len(ac.mutation.PersonIDs()) == 0 {
		return &ValidationError{Name: "person", err: errors.New(`data: missing required edge "Authorization.person"`)}
	}
	return nil
}

func (ac *AuthorizationCreate) sqlSave(ctx context.Context) (*Authorization, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AuthorizationCreate) createSpec() (*Authorization, *sqlgraph.CreateSpec) {
	var (
		_node = &Authorization{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(authorization.Table, sqlgraph.NewFieldSpec(authorization.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Token(); ok {
		_spec.SetField(authorization.FieldToken, field.TypeBytes, value)
		_node.Token = value
	}
	if value, ok := ac.mutation.Kind(); ok {
		_spec.SetField(authorization.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(authorization.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authorization.PersonTable,
			Columns: []string{authorization.PersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PersonID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AuthorizationCreateBulk is the builder for creating many Authorization entities in bulk.
type AuthorizationCreateBulk struct {
	config
	err      error
	builders []*AuthorizationCreate
}

// Save creates the Authorization entities in the database.
func (acb *AuthorizationCreateBulk) Save(ctx context.Context) ([]*Authorization, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Authorization, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthorizationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthorizationCreateBulk) SaveX(ctx context.Context) []*Authorization {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuthorizationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuthorizationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
