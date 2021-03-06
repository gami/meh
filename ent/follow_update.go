// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"meh/ent/follow"
	"meh/ent/predicate"
	"meh/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FollowUpdate is the builder for updating Follow entities.
type FollowUpdate struct {
	config
	hooks    []Hook
	mutation *FollowMutation
}

// Where appends a list predicates to the FollowUpdate builder.
func (fu *FollowUpdate) Where(ps ...predicate.Follow) *FollowUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUserID sets the "user_id" field.
func (fu *FollowUpdate) SetUserID(u uint64) *FollowUpdate {
	fu.mutation.SetUserID(u)
	return fu
}

// SetFolloweeID sets the "followee_id" field.
func (fu *FollowUpdate) SetFolloweeID(u uint64) *FollowUpdate {
	fu.mutation.SetFolloweeID(u)
	return fu
}

// SetUser sets the "user" edge to the User entity.
func (fu *FollowUpdate) SetUser(u *User) *FollowUpdate {
	return fu.SetUserID(u.ID)
}

// SetFollowee sets the "followee" edge to the User entity.
func (fu *FollowUpdate) SetFollowee(u *User) *FollowUpdate {
	return fu.SetFolloweeID(u.ID)
}

// Mutation returns the FollowMutation object of the builder.
func (fu *FollowUpdate) Mutation() *FollowMutation {
	return fu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fu *FollowUpdate) ClearUser() *FollowUpdate {
	fu.mutation.ClearUser()
	return fu
}

// ClearFollowee clears the "followee" edge to the User entity.
func (fu *FollowUpdate) ClearFollowee() *FollowUpdate {
	fu.mutation.ClearFollowee()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FollowUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FollowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FollowUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FollowUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FollowUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FollowUpdate) check() error {
	if _, ok := fu.mutation.UserID(); fu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := fu.mutation.FolloweeID(); fu.mutation.FolloweeCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"followee\"")
	}
	return nil
}

func (fu *FollowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   follow.Table,
			Columns: follow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: follow.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.UserTable,
			Columns: []string{follow.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.UserTable,
			Columns: []string{follow.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.FolloweeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{follow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FollowUpdateOne is the builder for updating a single Follow entity.
type FollowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FollowMutation
}

// SetUserID sets the "user_id" field.
func (fuo *FollowUpdateOne) SetUserID(u uint64) *FollowUpdateOne {
	fuo.mutation.SetUserID(u)
	return fuo
}

// SetFolloweeID sets the "followee_id" field.
func (fuo *FollowUpdateOne) SetFolloweeID(u uint64) *FollowUpdateOne {
	fuo.mutation.SetFolloweeID(u)
	return fuo
}

// SetUser sets the "user" edge to the User entity.
func (fuo *FollowUpdateOne) SetUser(u *User) *FollowUpdateOne {
	return fuo.SetUserID(u.ID)
}

// SetFollowee sets the "followee" edge to the User entity.
func (fuo *FollowUpdateOne) SetFollowee(u *User) *FollowUpdateOne {
	return fuo.SetFolloweeID(u.ID)
}

// Mutation returns the FollowMutation object of the builder.
func (fuo *FollowUpdateOne) Mutation() *FollowMutation {
	return fuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fuo *FollowUpdateOne) ClearUser() *FollowUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// ClearFollowee clears the "followee" edge to the User entity.
func (fuo *FollowUpdateOne) ClearFollowee() *FollowUpdateOne {
	fuo.mutation.ClearFollowee()
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FollowUpdateOne) Select(field string, fields ...string) *FollowUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Follow entity.
func (fuo *FollowUpdateOne) Save(ctx context.Context) (*Follow, error) {
	var (
		err  error
		node *Follow
	)
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FollowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FollowUpdateOne) SaveX(ctx context.Context) *Follow {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FollowUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FollowUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FollowUpdateOne) check() error {
	if _, ok := fuo.mutation.UserID(); fuo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := fuo.mutation.FolloweeID(); fuo.mutation.FolloweeCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"followee\"")
	}
	return nil
}

func (fuo *FollowUpdateOne) sqlSave(ctx context.Context) (_node *Follow, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   follow.Table,
			Columns: follow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: follow.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Follow.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, follow.FieldID)
		for _, f := range fields {
			if !follow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != follow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.UserTable,
			Columns: []string{follow.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.UserTable,
			Columns: []string{follow.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.FolloweeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Follow{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{follow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
