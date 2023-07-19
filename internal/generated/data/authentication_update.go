// Code generated by ent, DO NOT EDIT.

package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/avptp/brain/internal/generated/data/authentication"
	"github.com/avptp/brain/internal/generated/data/predicate"
)

// AuthenticationUpdate is the builder for updating Authentication entities.
type AuthenticationUpdate struct {
	config
	hooks    []Hook
	mutation *AuthenticationMutation
}

// Where appends a list predicates to the AuthenticationUpdate builder.
func (au *AuthenticationUpdate) Where(ps ...predicate.Authentication) *AuthenticationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedIP sets the "created_ip" field.
func (au *AuthenticationUpdate) SetCreatedIP(s string) *AuthenticationUpdate {
	au.mutation.SetCreatedIP(s)
	return au
}

// SetLastUsedIP sets the "last_used_ip" field.
func (au *AuthenticationUpdate) SetLastUsedIP(s string) *AuthenticationUpdate {
	au.mutation.SetLastUsedIP(s)
	return au
}

// SetLastUsedAt sets the "last_used_at" field.
func (au *AuthenticationUpdate) SetLastUsedAt(t time.Time) *AuthenticationUpdate {
	au.mutation.SetLastUsedAt(t)
	return au
}

// Mutation returns the AuthenticationMutation object of the builder.
func (au *AuthenticationUpdate) Mutation() *AuthenticationMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuthenticationUpdate) Save(ctx context.Context) (int, error) {
	if err := au.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthenticationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthenticationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthenticationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AuthenticationUpdate) defaults() error {
	if _, ok := au.mutation.LastUsedAt(); !ok {
		if authentication.UpdateDefaultLastUsedAt == nil {
			return fmt.Errorf("data: uninitialized authentication.UpdateDefaultLastUsedAt (forgotten import data/runtime?)")
		}
		v := authentication.UpdateDefaultLastUsedAt()
		au.mutation.SetLastUsedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (au *AuthenticationUpdate) check() error {
	if _, ok := au.mutation.PersonID(); au.mutation.PersonCleared() && !ok {
		return errors.New(`data: clearing a required unique edge "Authentication.person"`)
	}
	return nil
}

func (au *AuthenticationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(authentication.Table, authentication.Columns, sqlgraph.NewFieldSpec(authentication.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedIP(); ok {
		_spec.SetField(authentication.FieldCreatedIP, field.TypeString, value)
	}
	if value, ok := au.mutation.LastUsedIP(); ok {
		_spec.SetField(authentication.FieldLastUsedIP, field.TypeString, value)
	}
	if value, ok := au.mutation.LastUsedAt(); ok {
		_spec.SetField(authentication.FieldLastUsedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authentication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AuthenticationUpdateOne is the builder for updating a single Authentication entity.
type AuthenticationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthenticationMutation
}

// SetCreatedIP sets the "created_ip" field.
func (auo *AuthenticationUpdateOne) SetCreatedIP(s string) *AuthenticationUpdateOne {
	auo.mutation.SetCreatedIP(s)
	return auo
}

// SetLastUsedIP sets the "last_used_ip" field.
func (auo *AuthenticationUpdateOne) SetLastUsedIP(s string) *AuthenticationUpdateOne {
	auo.mutation.SetLastUsedIP(s)
	return auo
}

// SetLastUsedAt sets the "last_used_at" field.
func (auo *AuthenticationUpdateOne) SetLastUsedAt(t time.Time) *AuthenticationUpdateOne {
	auo.mutation.SetLastUsedAt(t)
	return auo
}

// Mutation returns the AuthenticationMutation object of the builder.
func (auo *AuthenticationUpdateOne) Mutation() *AuthenticationMutation {
	return auo.mutation
}

// Where appends a list predicates to the AuthenticationUpdate builder.
func (auo *AuthenticationUpdateOne) Where(ps ...predicate.Authentication) *AuthenticationUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuthenticationUpdateOne) Select(field string, fields ...string) *AuthenticationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Authentication entity.
func (auo *AuthenticationUpdateOne) Save(ctx context.Context) (*Authentication, error) {
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthenticationUpdateOne) SaveX(ctx context.Context) *Authentication {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuthenticationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthenticationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AuthenticationUpdateOne) defaults() error {
	if _, ok := auo.mutation.LastUsedAt(); !ok {
		if authentication.UpdateDefaultLastUsedAt == nil {
			return fmt.Errorf("data: uninitialized authentication.UpdateDefaultLastUsedAt (forgotten import data/runtime?)")
		}
		v := authentication.UpdateDefaultLastUsedAt()
		auo.mutation.SetLastUsedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (auo *AuthenticationUpdateOne) check() error {
	if _, ok := auo.mutation.PersonID(); auo.mutation.PersonCleared() && !ok {
		return errors.New(`data: clearing a required unique edge "Authentication.person"`)
	}
	return nil
}

func (auo *AuthenticationUpdateOne) sqlSave(ctx context.Context) (_node *Authentication, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(authentication.Table, authentication.Columns, sqlgraph.NewFieldSpec(authentication.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`data: missing "Authentication.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authentication.FieldID)
		for _, f := range fields {
			if !authentication.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("data: invalid field %q for query", f)}
			}
			if f != authentication.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedIP(); ok {
		_spec.SetField(authentication.FieldCreatedIP, field.TypeString, value)
	}
	if value, ok := auo.mutation.LastUsedIP(); ok {
		_spec.SetField(authentication.FieldLastUsedIP, field.TypeString, value)
	}
	if value, ok := auo.mutation.LastUsedAt(); ok {
		_spec.SetField(authentication.FieldLastUsedAt, field.TypeTime, value)
	}
	_node = &Authentication{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authentication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
