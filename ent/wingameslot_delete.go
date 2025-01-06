// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/common/ent/predicate"
	"github.com/gzorm/common/ent/wingameslot"
)

// WinGameSlotDelete is the builder for deleting a WinGameSlot entity.
type WinGameSlotDelete struct {
	config
	hooks    []Hook
	mutation *WinGameSlotMutation
}

// Where appends a list predicates to the WinGameSlotDelete builder.
func (wgsd *WinGameSlotDelete) Where(ps ...predicate.WinGameSlot) *WinGameSlotDelete {
	wgsd.mutation.Where(ps...)
	return wgsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wgsd *WinGameSlotDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wgsd.sqlExec, wgsd.mutation, wgsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wgsd *WinGameSlotDelete) ExecX(ctx context.Context) int {
	n, err := wgsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wgsd *WinGameSlotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wingameslot.Table, sqlgraph.NewFieldSpec(wingameslot.FieldID, field.TypeString))
	if ps := wgsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wgsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wgsd.mutation.done = true
	return affected, err
}

// WinGameSlotDeleteOne is the builder for deleting a single WinGameSlot entity.
type WinGameSlotDeleteOne struct {
	wgsd *WinGameSlotDelete
}

// Where appends a list predicates to the WinGameSlotDelete builder.
func (wgsdo *WinGameSlotDeleteOne) Where(ps ...predicate.WinGameSlot) *WinGameSlotDeleteOne {
	wgsdo.wgsd.mutation.Where(ps...)
	return wgsdo
}

// Exec executes the deletion query.
func (wgsdo *WinGameSlotDeleteOne) Exec(ctx context.Context) error {
	n, err := wgsdo.wgsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wingameslot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wgsdo *WinGameSlotDeleteOne) ExecX(ctx context.Context) {
	if err := wgsdo.Exec(ctx); err != nil {
		panic(err)
	}
}