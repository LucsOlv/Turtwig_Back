// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/LucsOlv/Turtwing_Back/ent/predicate"
	"github.com/LucsOlv/Turtwing_Back/ent/project"
	"github.com/LucsOlv/Turtwing_Back/ent/projectmember"
	"github.com/LucsOlv/Turtwing_Back/ent/user"
	"github.com/google/uuid"
)

// ProjectMemberUpdate is the builder for updating ProjectMember entities.
type ProjectMemberUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMemberMutation
}

// Where appends a list predicates to the ProjectMemberUpdate builder.
func (pmu *ProjectMemberUpdate) Where(ps ...predicate.ProjectMember) *ProjectMemberUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetProjectID sets the "project_id" field.
func (pmu *ProjectMemberUpdate) SetProjectID(u uuid.UUID) *ProjectMemberUpdate {
	pmu.mutation.SetProjectID(u)
	return pmu
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (pmu *ProjectMemberUpdate) SetNillableProjectID(u *uuid.UUID) *ProjectMemberUpdate {
	if u != nil {
		pmu.SetProjectID(*u)
	}
	return pmu
}

// SetUserID sets the "user_id" field.
func (pmu *ProjectMemberUpdate) SetUserID(u uuid.UUID) *ProjectMemberUpdate {
	pmu.mutation.SetUserID(u)
	return pmu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pmu *ProjectMemberUpdate) SetNillableUserID(u *uuid.UUID) *ProjectMemberUpdate {
	if u != nil {
		pmu.SetUserID(*u)
	}
	return pmu
}

// SetUpdatedAt sets the "updated_at" field.
func (pmu *ProjectMemberUpdate) SetUpdatedAt(t time.Time) *ProjectMemberUpdate {
	pmu.mutation.SetUpdatedAt(t)
	return pmu
}

// SetRole sets the "role" field.
func (pmu *ProjectMemberUpdate) SetRole(pr projectmember.Role) *ProjectMemberUpdate {
	pmu.mutation.SetRole(pr)
	return pmu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (pmu *ProjectMemberUpdate) SetNillableRole(pr *projectmember.Role) *ProjectMemberUpdate {
	if pr != nil {
		pmu.SetRole(*pr)
	}
	return pmu
}

// SetStatus sets the "status" field.
func (pmu *ProjectMemberUpdate) SetStatus(pr projectmember.Status) *ProjectMemberUpdate {
	pmu.mutation.SetStatus(pr)
	return pmu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pmu *ProjectMemberUpdate) SetNillableStatus(pr *projectmember.Status) *ProjectMemberUpdate {
	if pr != nil {
		pmu.SetStatus(*pr)
	}
	return pmu
}

// SetProject sets the "project" edge to the Project entity.
func (pmu *ProjectMemberUpdate) SetProject(p *Project) *ProjectMemberUpdate {
	return pmu.SetProjectID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pmu *ProjectMemberUpdate) SetUser(u *User) *ProjectMemberUpdate {
	return pmu.SetUserID(u.ID)
}

// Mutation returns the ProjectMemberMutation object of the builder.
func (pmu *ProjectMemberUpdate) Mutation() *ProjectMemberMutation {
	return pmu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (pmu *ProjectMemberUpdate) ClearProject() *ProjectMemberUpdate {
	pmu.mutation.ClearProject()
	return pmu
}

// ClearUser clears the "user" edge to the User entity.
func (pmu *ProjectMemberUpdate) ClearUser() *ProjectMemberUpdate {
	pmu.mutation.ClearUser()
	return pmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *ProjectMemberUpdate) Save(ctx context.Context) (int, error) {
	pmu.defaults()
	return withHooks(ctx, pmu.sqlSave, pmu.mutation, pmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *ProjectMemberUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *ProjectMemberUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *ProjectMemberUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmu *ProjectMemberUpdate) defaults() {
	if _, ok := pmu.mutation.UpdatedAt(); !ok {
		v := projectmember.UpdateDefaultUpdatedAt()
		pmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmu *ProjectMemberUpdate) check() error {
	if v, ok := pmu.mutation.Role(); ok {
		if err := projectmember.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "ProjectMember.role": %w`, err)}
		}
	}
	if v, ok := pmu.mutation.Status(); ok {
		if err := projectmember.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProjectMember.status": %w`, err)}
		}
	}
	if pmu.mutation.ProjectCleared() && len(pmu.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProjectMember.project"`)
	}
	if pmu.mutation.UserCleared() && len(pmu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProjectMember.user"`)
	}
	return nil
}

func (pmu *ProjectMemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectmember.Table, projectmember.Columns, sqlgraph.NewFieldSpec(projectmember.FieldID, field.TypeUUID))
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.UpdatedAt(); ok {
		_spec.SetField(projectmember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pmu.mutation.Role(); ok {
		_spec.SetField(projectmember.FieldRole, field.TypeEnum, value)
	}
	if value, ok := pmu.mutation.Status(); ok {
		_spec.SetField(projectmember.FieldStatus, field.TypeEnum, value)
	}
	if pmu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.ProjectTable,
			Columns: []string{projectmember.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.ProjectTable,
			Columns: []string{projectmember.ProjectColumn},
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
	if pmu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.UserTable,
			Columns: []string{projectmember.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.UserTable,
			Columns: []string{projectmember.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pmu.mutation.done = true
	return n, nil
}

// ProjectMemberUpdateOne is the builder for updating a single ProjectMember entity.
type ProjectMemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMemberMutation
}

// SetProjectID sets the "project_id" field.
func (pmuo *ProjectMemberUpdateOne) SetProjectID(u uuid.UUID) *ProjectMemberUpdateOne {
	pmuo.mutation.SetProjectID(u)
	return pmuo
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (pmuo *ProjectMemberUpdateOne) SetNillableProjectID(u *uuid.UUID) *ProjectMemberUpdateOne {
	if u != nil {
		pmuo.SetProjectID(*u)
	}
	return pmuo
}

// SetUserID sets the "user_id" field.
func (pmuo *ProjectMemberUpdateOne) SetUserID(u uuid.UUID) *ProjectMemberUpdateOne {
	pmuo.mutation.SetUserID(u)
	return pmuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pmuo *ProjectMemberUpdateOne) SetNillableUserID(u *uuid.UUID) *ProjectMemberUpdateOne {
	if u != nil {
		pmuo.SetUserID(*u)
	}
	return pmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pmuo *ProjectMemberUpdateOne) SetUpdatedAt(t time.Time) *ProjectMemberUpdateOne {
	pmuo.mutation.SetUpdatedAt(t)
	return pmuo
}

// SetRole sets the "role" field.
func (pmuo *ProjectMemberUpdateOne) SetRole(pr projectmember.Role) *ProjectMemberUpdateOne {
	pmuo.mutation.SetRole(pr)
	return pmuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (pmuo *ProjectMemberUpdateOne) SetNillableRole(pr *projectmember.Role) *ProjectMemberUpdateOne {
	if pr != nil {
		pmuo.SetRole(*pr)
	}
	return pmuo
}

// SetStatus sets the "status" field.
func (pmuo *ProjectMemberUpdateOne) SetStatus(pr projectmember.Status) *ProjectMemberUpdateOne {
	pmuo.mutation.SetStatus(pr)
	return pmuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pmuo *ProjectMemberUpdateOne) SetNillableStatus(pr *projectmember.Status) *ProjectMemberUpdateOne {
	if pr != nil {
		pmuo.SetStatus(*pr)
	}
	return pmuo
}

// SetProject sets the "project" edge to the Project entity.
func (pmuo *ProjectMemberUpdateOne) SetProject(p *Project) *ProjectMemberUpdateOne {
	return pmuo.SetProjectID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pmuo *ProjectMemberUpdateOne) SetUser(u *User) *ProjectMemberUpdateOne {
	return pmuo.SetUserID(u.ID)
}

// Mutation returns the ProjectMemberMutation object of the builder.
func (pmuo *ProjectMemberUpdateOne) Mutation() *ProjectMemberMutation {
	return pmuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (pmuo *ProjectMemberUpdateOne) ClearProject() *ProjectMemberUpdateOne {
	pmuo.mutation.ClearProject()
	return pmuo
}

// ClearUser clears the "user" edge to the User entity.
func (pmuo *ProjectMemberUpdateOne) ClearUser() *ProjectMemberUpdateOne {
	pmuo.mutation.ClearUser()
	return pmuo
}

// Where appends a list predicates to the ProjectMemberUpdate builder.
func (pmuo *ProjectMemberUpdateOne) Where(ps ...predicate.ProjectMember) *ProjectMemberUpdateOne {
	pmuo.mutation.Where(ps...)
	return pmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *ProjectMemberUpdateOne) Select(field string, fields ...string) *ProjectMemberUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated ProjectMember entity.
func (pmuo *ProjectMemberUpdateOne) Save(ctx context.Context) (*ProjectMember, error) {
	pmuo.defaults()
	return withHooks(ctx, pmuo.sqlSave, pmuo.mutation, pmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *ProjectMemberUpdateOne) SaveX(ctx context.Context) *ProjectMember {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *ProjectMemberUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *ProjectMemberUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmuo *ProjectMemberUpdateOne) defaults() {
	if _, ok := pmuo.mutation.UpdatedAt(); !ok {
		v := projectmember.UpdateDefaultUpdatedAt()
		pmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmuo *ProjectMemberUpdateOne) check() error {
	if v, ok := pmuo.mutation.Role(); ok {
		if err := projectmember.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "ProjectMember.role": %w`, err)}
		}
	}
	if v, ok := pmuo.mutation.Status(); ok {
		if err := projectmember.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProjectMember.status": %w`, err)}
		}
	}
	if pmuo.mutation.ProjectCleared() && len(pmuo.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProjectMember.project"`)
	}
	if pmuo.mutation.UserCleared() && len(pmuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProjectMember.user"`)
	}
	return nil
}

func (pmuo *ProjectMemberUpdateOne) sqlSave(ctx context.Context) (_node *ProjectMember, err error) {
	if err := pmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectmember.Table, projectmember.Columns, sqlgraph.NewFieldSpec(projectmember.FieldID, field.TypeUUID))
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectMember.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectmember.FieldID)
		for _, f := range fields {
			if !projectmember.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(projectmember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pmuo.mutation.Role(); ok {
		_spec.SetField(projectmember.FieldRole, field.TypeEnum, value)
	}
	if value, ok := pmuo.mutation.Status(); ok {
		_spec.SetField(projectmember.FieldStatus, field.TypeEnum, value)
	}
	if pmuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.ProjectTable,
			Columns: []string{projectmember.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.ProjectTable,
			Columns: []string{projectmember.ProjectColumn},
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
	if pmuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.UserTable,
			Columns: []string{projectmember.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectmember.UserTable,
			Columns: []string{projectmember.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectMember{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pmuo.mutation.done = true
	return _node, nil
}
