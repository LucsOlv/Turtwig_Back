// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/LucsOlv/Turtwing_Back/ent/historyproject"
	"github.com/LucsOlv/Turtwing_Back/ent/predicate"
)

// HistoryProjectDelete is the builder for deleting a HistoryProject entity.
type HistoryProjectDelete struct {
	config
	hooks    []Hook
	mutation *HistoryProjectMutation
}

// Where appends a list predicates to the HistoryProjectDelete builder.
func (hpd *HistoryProjectDelete) Where(ps ...predicate.HistoryProject) *HistoryProjectDelete {
	hpd.mutation.Where(ps...)
	return hpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hpd *HistoryProjectDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hpd.sqlExec, hpd.mutation, hpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hpd *HistoryProjectDelete) ExecX(ctx context.Context) int {
	n, err := hpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hpd *HistoryProjectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(historyproject.Table, sqlgraph.NewFieldSpec(historyproject.FieldID, field.TypeUUID))
	if ps := hpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hpd.mutation.done = true
	return affected, err
}

// HistoryProjectDeleteOne is the builder for deleting a single HistoryProject entity.
type HistoryProjectDeleteOne struct {
	hpd *HistoryProjectDelete
}

// Where appends a list predicates to the HistoryProjectDelete builder.
func (hpdo *HistoryProjectDeleteOne) Where(ps ...predicate.HistoryProject) *HistoryProjectDeleteOne {
	hpdo.hpd.mutation.Where(ps...)
	return hpdo
}

// Exec executes the deletion query.
func (hpdo *HistoryProjectDeleteOne) Exec(ctx context.Context) error {
	n, err := hpdo.hpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{historyproject.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hpdo *HistoryProjectDeleteOne) ExecX(ctx context.Context) {
	if err := hpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
