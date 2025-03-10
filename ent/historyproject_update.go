// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/LucsOlv/Turtwing_Back/ent/history"
	"github.com/LucsOlv/Turtwing_Back/ent/historyproject"
	"github.com/LucsOlv/Turtwing_Back/ent/predicate"
	"github.com/LucsOlv/Turtwing_Back/ent/project"
	"github.com/google/uuid"
)

// HistoryProjectUpdate is the builder for updating HistoryProject entities.
type HistoryProjectUpdate struct {
	config
	hooks    []Hook
	mutation *HistoryProjectMutation
}

// Where appends a list predicates to the HistoryProjectUpdate builder.
func (hpu *HistoryProjectUpdate) Where(ps ...predicate.HistoryProject) *HistoryProjectUpdate {
	hpu.mutation.Where(ps...)
	return hpu
}

// SetProjectID sets the "project_id" field.
func (hpu *HistoryProjectUpdate) SetProjectID(u uuid.UUID) *HistoryProjectUpdate {
	hpu.mutation.SetProjectID(u)
	return hpu
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (hpu *HistoryProjectUpdate) SetNillableProjectID(u *uuid.UUID) *HistoryProjectUpdate {
	if u != nil {
		hpu.SetProjectID(*u)
	}
	return hpu
}

// SetHistoryID sets the "history_id" field.
func (hpu *HistoryProjectUpdate) SetHistoryID(u uuid.UUID) *HistoryProjectUpdate {
	hpu.mutation.SetHistoryID(u)
	return hpu
}

// SetNillableHistoryID sets the "history_id" field if the given value is not nil.
func (hpu *HistoryProjectUpdate) SetNillableHistoryID(u *uuid.UUID) *HistoryProjectUpdate {
	if u != nil {
		hpu.SetHistoryID(*u)
	}
	return hpu
}

// SetChanges sets the "changes" field.
func (hpu *HistoryProjectUpdate) SetChanges(m map[string]interface{}) *HistoryProjectUpdate {
	hpu.mutation.SetChanges(m)
	return hpu
}

// SetStatus sets the "status" field.
func (hpu *HistoryProjectUpdate) SetStatus(h historyproject.Status) *HistoryProjectUpdate {
	hpu.mutation.SetStatus(h)
	return hpu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hpu *HistoryProjectUpdate) SetNillableStatus(h *historyproject.Status) *HistoryProjectUpdate {
	if h != nil {
		hpu.SetStatus(*h)
	}
	return hpu
}

// SetProject sets the "project" edge to the Project entity.
func (hpu *HistoryProjectUpdate) SetProject(p *Project) *HistoryProjectUpdate {
	return hpu.SetProjectID(p.ID)
}

// SetHistory sets the "history" edge to the History entity.
func (hpu *HistoryProjectUpdate) SetHistory(h *History) *HistoryProjectUpdate {
	return hpu.SetHistoryID(h.ID)
}

// Mutation returns the HistoryProjectMutation object of the builder.
func (hpu *HistoryProjectUpdate) Mutation() *HistoryProjectMutation {
	return hpu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (hpu *HistoryProjectUpdate) ClearProject() *HistoryProjectUpdate {
	hpu.mutation.ClearProject()
	return hpu
}

// ClearHistory clears the "history" edge to the History entity.
func (hpu *HistoryProjectUpdate) ClearHistory() *HistoryProjectUpdate {
	hpu.mutation.ClearHistory()
	return hpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hpu *HistoryProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hpu.sqlSave, hpu.mutation, hpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hpu *HistoryProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := hpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hpu *HistoryProjectUpdate) Exec(ctx context.Context) error {
	_, err := hpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hpu *HistoryProjectUpdate) ExecX(ctx context.Context) {
	if err := hpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hpu *HistoryProjectUpdate) check() error {
	if v, ok := hpu.mutation.Status(); ok {
		if err := historyproject.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HistoryProject.status": %w`, err)}
		}
	}
	if hpu.mutation.ProjectCleared() && len(hpu.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "HistoryProject.project"`)
	}
	if hpu.mutation.HistoryCleared() && len(hpu.mutation.HistoryIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "HistoryProject.history"`)
	}
	return nil
}

func (hpu *HistoryProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(historyproject.Table, historyproject.Columns, sqlgraph.NewFieldSpec(historyproject.FieldID, field.TypeUUID))
	if ps := hpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hpu.mutation.Changes(); ok {
		_spec.SetField(historyproject.FieldChanges, field.TypeJSON, value)
	}
	if value, ok := hpu.mutation.Status(); ok {
		_spec.SetField(historyproject.FieldStatus, field.TypeEnum, value)
	}
	if hpu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.ProjectTable,
			Columns: []string{historyproject.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hpu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.ProjectTable,
			Columns: []string{historyproject.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hpu.mutation.HistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.HistoryTable,
			Columns: []string{historyproject.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hpu.mutation.HistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.HistoryTable,
			Columns: []string{historyproject.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{historyproject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hpu.mutation.done = true
	return n, nil
}

// HistoryProjectUpdateOne is the builder for updating a single HistoryProject entity.
type HistoryProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HistoryProjectMutation
}

// SetProjectID sets the "project_id" field.
func (hpuo *HistoryProjectUpdateOne) SetProjectID(u uuid.UUID) *HistoryProjectUpdateOne {
	hpuo.mutation.SetProjectID(u)
	return hpuo
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (hpuo *HistoryProjectUpdateOne) SetNillableProjectID(u *uuid.UUID) *HistoryProjectUpdateOne {
	if u != nil {
		hpuo.SetProjectID(*u)
	}
	return hpuo
}

// SetHistoryID sets the "history_id" field.
func (hpuo *HistoryProjectUpdateOne) SetHistoryID(u uuid.UUID) *HistoryProjectUpdateOne {
	hpuo.mutation.SetHistoryID(u)
	return hpuo
}

// SetNillableHistoryID sets the "history_id" field if the given value is not nil.
func (hpuo *HistoryProjectUpdateOne) SetNillableHistoryID(u *uuid.UUID) *HistoryProjectUpdateOne {
	if u != nil {
		hpuo.SetHistoryID(*u)
	}
	return hpuo
}

// SetChanges sets the "changes" field.
func (hpuo *HistoryProjectUpdateOne) SetChanges(m map[string]interface{}) *HistoryProjectUpdateOne {
	hpuo.mutation.SetChanges(m)
	return hpuo
}

// SetStatus sets the "status" field.
func (hpuo *HistoryProjectUpdateOne) SetStatus(h historyproject.Status) *HistoryProjectUpdateOne {
	hpuo.mutation.SetStatus(h)
	return hpuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hpuo *HistoryProjectUpdateOne) SetNillableStatus(h *historyproject.Status) *HistoryProjectUpdateOne {
	if h != nil {
		hpuo.SetStatus(*h)
	}
	return hpuo
}

// SetProject sets the "project" edge to the Project entity.
func (hpuo *HistoryProjectUpdateOne) SetProject(p *Project) *HistoryProjectUpdateOne {
	return hpuo.SetProjectID(p.ID)
}

// SetHistory sets the "history" edge to the History entity.
func (hpuo *HistoryProjectUpdateOne) SetHistory(h *History) *HistoryProjectUpdateOne {
	return hpuo.SetHistoryID(h.ID)
}

// Mutation returns the HistoryProjectMutation object of the builder.
func (hpuo *HistoryProjectUpdateOne) Mutation() *HistoryProjectMutation {
	return hpuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (hpuo *HistoryProjectUpdateOne) ClearProject() *HistoryProjectUpdateOne {
	hpuo.mutation.ClearProject()
	return hpuo
}

// ClearHistory clears the "history" edge to the History entity.
func (hpuo *HistoryProjectUpdateOne) ClearHistory() *HistoryProjectUpdateOne {
	hpuo.mutation.ClearHistory()
	return hpuo
}

// Where appends a list predicates to the HistoryProjectUpdate builder.
func (hpuo *HistoryProjectUpdateOne) Where(ps ...predicate.HistoryProject) *HistoryProjectUpdateOne {
	hpuo.mutation.Where(ps...)
	return hpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hpuo *HistoryProjectUpdateOne) Select(field string, fields ...string) *HistoryProjectUpdateOne {
	hpuo.fields = append([]string{field}, fields...)
	return hpuo
}

// Save executes the query and returns the updated HistoryProject entity.
func (hpuo *HistoryProjectUpdateOne) Save(ctx context.Context) (*HistoryProject, error) {
	return withHooks(ctx, hpuo.sqlSave, hpuo.mutation, hpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hpuo *HistoryProjectUpdateOne) SaveX(ctx context.Context) *HistoryProject {
	node, err := hpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hpuo *HistoryProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := hpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hpuo *HistoryProjectUpdateOne) ExecX(ctx context.Context) {
	if err := hpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hpuo *HistoryProjectUpdateOne) check() error {
	if v, ok := hpuo.mutation.Status(); ok {
		if err := historyproject.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HistoryProject.status": %w`, err)}
		}
	}
	if hpuo.mutation.ProjectCleared() && len(hpuo.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "HistoryProject.project"`)
	}
	if hpuo.mutation.HistoryCleared() && len(hpuo.mutation.HistoryIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "HistoryProject.history"`)
	}
	return nil
}

func (hpuo *HistoryProjectUpdateOne) sqlSave(ctx context.Context) (_node *HistoryProject, err error) {
	if err := hpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(historyproject.Table, historyproject.Columns, sqlgraph.NewFieldSpec(historyproject.FieldID, field.TypeUUID))
	id, ok := hpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HistoryProject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, historyproject.FieldID)
		for _, f := range fields {
			if !historyproject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != historyproject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hpuo.mutation.Changes(); ok {
		_spec.SetField(historyproject.FieldChanges, field.TypeJSON, value)
	}
	if value, ok := hpuo.mutation.Status(); ok {
		_spec.SetField(historyproject.FieldStatus, field.TypeEnum, value)
	}
	if hpuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.ProjectTable,
			Columns: []string{historyproject.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hpuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.ProjectTable,
			Columns: []string{historyproject.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hpuo.mutation.HistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.HistoryTable,
			Columns: []string{historyproject.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hpuo.mutation.HistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historyproject.HistoryTable,
			Columns: []string{historyproject.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HistoryProject{config: hpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{historyproject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hpuo.mutation.done = true
	return _node, nil
}
