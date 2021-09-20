// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"meh/ent/meh"
	"meh/ent/predicate"
	"meh/ent/timeline"
	"meh/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimelineUpdate is the builder for updating Timeline entities.
type TimelineUpdate struct {
	config
	hooks    []Hook
	mutation *TimelineMutation
}

// Where appends a list predicates to the TimelineUpdate builder.
func (tu *TimelineUpdate) Where(ps ...predicate.Timeline) *TimelineUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TimelineUpdate) SetUserID(u uint64) *TimelineUpdate {
	tu.mutation.SetUserID(u)
	return tu
}

// SetMehID sets the "meh_id" field.
func (tu *TimelineUpdate) SetMehID(u uint64) *TimelineUpdate {
	tu.mutation.SetMehID(u)
	return tu
}

// SetUser sets the "user" edge to the User entity.
func (tu *TimelineUpdate) SetUser(u *User) *TimelineUpdate {
	return tu.SetUserID(u.ID)
}

// SetMeh sets the "meh" edge to the Meh entity.
func (tu *TimelineUpdate) SetMeh(m *Meh) *TimelineUpdate {
	return tu.SetMehID(m.ID)
}

// Mutation returns the TimelineMutation object of the builder.
func (tu *TimelineUpdate) Mutation() *TimelineMutation {
	return tu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TimelineUpdate) ClearUser() *TimelineUpdate {
	tu.mutation.ClearUser()
	return tu
}

// ClearMeh clears the "meh" edge to the Meh entity.
func (tu *TimelineUpdate) ClearMeh() *TimelineUpdate {
	tu.mutation.ClearMeh()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TimelineUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimelineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TimelineUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TimelineUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TimelineUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TimelineUpdate) check() error {
	if _, ok := tu.mutation.UserID(); tu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := tu.mutation.MehID(); tu.mutation.MehCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"meh\"")
	}
	return nil
}

func (tu *TimelineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   timeline.Table,
			Columns: timeline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: timeline.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.UserTable,
			Columns: []string{timeline.UserColumn},
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
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.UserTable,
			Columns: []string{timeline.UserColumn},
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
	if tu.mutation.MehCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.MehTable,
			Columns: []string{timeline.MehColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: meh.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.MehIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.MehTable,
			Columns: []string{timeline.MehColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: meh.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timeline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TimelineUpdateOne is the builder for updating a single Timeline entity.
type TimelineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TimelineMutation
}

// SetUserID sets the "user_id" field.
func (tuo *TimelineUpdateOne) SetUserID(u uint64) *TimelineUpdateOne {
	tuo.mutation.SetUserID(u)
	return tuo
}

// SetMehID sets the "meh_id" field.
func (tuo *TimelineUpdateOne) SetMehID(u uint64) *TimelineUpdateOne {
	tuo.mutation.SetMehID(u)
	return tuo
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TimelineUpdateOne) SetUser(u *User) *TimelineUpdateOne {
	return tuo.SetUserID(u.ID)
}

// SetMeh sets the "meh" edge to the Meh entity.
func (tuo *TimelineUpdateOne) SetMeh(m *Meh) *TimelineUpdateOne {
	return tuo.SetMehID(m.ID)
}

// Mutation returns the TimelineMutation object of the builder.
func (tuo *TimelineUpdateOne) Mutation() *TimelineMutation {
	return tuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TimelineUpdateOne) ClearUser() *TimelineUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// ClearMeh clears the "meh" edge to the Meh entity.
func (tuo *TimelineUpdateOne) ClearMeh() *TimelineUpdateOne {
	tuo.mutation.ClearMeh()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TimelineUpdateOne) Select(field string, fields ...string) *TimelineUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Timeline entity.
func (tuo *TimelineUpdateOne) Save(ctx context.Context) (*Timeline, error) {
	var (
		err  error
		node *Timeline
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimelineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TimelineUpdateOne) SaveX(ctx context.Context) *Timeline {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TimelineUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TimelineUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TimelineUpdateOne) check() error {
	if _, ok := tuo.mutation.UserID(); tuo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	if _, ok := tuo.mutation.MehID(); tuo.mutation.MehCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"meh\"")
	}
	return nil
}

func (tuo *TimelineUpdateOne) sqlSave(ctx context.Context) (_node *Timeline, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   timeline.Table,
			Columns: timeline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: timeline.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Timeline.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timeline.FieldID)
		for _, f := range fields {
			if !timeline.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != timeline.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.UserTable,
			Columns: []string{timeline.UserColumn},
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
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.UserTable,
			Columns: []string{timeline.UserColumn},
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
	if tuo.mutation.MehCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.MehTable,
			Columns: []string{timeline.MehColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: meh.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.MehIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   timeline.MehTable,
			Columns: []string{timeline.MehColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: meh.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Timeline{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timeline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}