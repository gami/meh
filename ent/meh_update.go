// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"meh/ent/meh"
	"meh/ent/predicate"
	"meh/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MehUpdate is the builder for updating Meh entities.
type MehUpdate struct {
	config
	hooks    []Hook
	mutation *MehMutation
}

// Where appends a list predicates to the MehUpdate builder.
func (mu *MehUpdate) Where(ps ...predicate.Meh) *MehUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUserID sets the "user_id" field.
func (mu *MehUpdate) SetUserID(u uint64) *MehUpdate {
	mu.mutation.SetUserID(u)
	return mu
}

// SetText sets the "text" field.
func (mu *MehUpdate) SetText(s string) *MehUpdate {
	mu.mutation.SetText(s)
	return mu
}

// SetUser sets the "user" edge to the User entity.
func (mu *MehUpdate) SetUser(u *User) *MehUpdate {
	return mu.SetUserID(u.ID)
}

// Mutation returns the MehMutation object of the builder.
func (mu *MehUpdate) Mutation() *MehMutation {
	return mu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mu *MehUpdate) ClearUser() *MehUpdate {
	mu.mutation.ClearUser()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MehUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MehMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MehUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MehUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MehUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MehUpdate) check() error {
	if v, ok := mu.mutation.Text(); ok {
		if err := meh.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf("ent: validator failed for field \"text\": %w", err)}
		}
	}
	if _, ok := mu.mutation.UserID(); mu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (mu *MehUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   meh.Table,
			Columns: meh.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: meh.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: meh.FieldText,
		})
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   meh.UserTable,
			Columns: []string{meh.UserColumn},
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
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   meh.UserTable,
			Columns: []string{meh.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meh.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MehUpdateOne is the builder for updating a single Meh entity.
type MehUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MehMutation
}

// SetUserID sets the "user_id" field.
func (muo *MehUpdateOne) SetUserID(u uint64) *MehUpdateOne {
	muo.mutation.SetUserID(u)
	return muo
}

// SetText sets the "text" field.
func (muo *MehUpdateOne) SetText(s string) *MehUpdateOne {
	muo.mutation.SetText(s)
	return muo
}

// SetUser sets the "user" edge to the User entity.
func (muo *MehUpdateOne) SetUser(u *User) *MehUpdateOne {
	return muo.SetUserID(u.ID)
}

// Mutation returns the MehMutation object of the builder.
func (muo *MehUpdateOne) Mutation() *MehMutation {
	return muo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (muo *MehUpdateOne) ClearUser() *MehUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MehUpdateOne) Select(field string, fields ...string) *MehUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Meh entity.
func (muo *MehUpdateOne) Save(ctx context.Context) (*Meh, error) {
	var (
		err  error
		node *Meh
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MehMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MehUpdateOne) SaveX(ctx context.Context) *Meh {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MehUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MehUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MehUpdateOne) check() error {
	if v, ok := muo.mutation.Text(); ok {
		if err := meh.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf("ent: validator failed for field \"text\": %w", err)}
		}
	}
	if _, ok := muo.mutation.UserID(); muo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (muo *MehUpdateOne) sqlSave(ctx context.Context) (_node *Meh, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   meh.Table,
			Columns: meh.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: meh.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Meh.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meh.FieldID)
		for _, f := range fields {
			if !meh.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != meh.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: meh.FieldText,
		})
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   meh.UserTable,
			Columns: []string{meh.UserColumn},
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
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   meh.UserTable,
			Columns: []string{meh.UserColumn},
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
	_node = &Meh{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meh.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
