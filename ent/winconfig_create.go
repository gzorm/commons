// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/commons/ent/winconfig"
)

// WinConfigCreate is the builder for creating a WinConfig entity.
type WinConfigCreate struct {
	config
	mutation *WinConfigMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (wcc *WinConfigCreate) SetTitle(s string) *WinConfigCreate {
	wcc.mutation.SetTitle(s)
	return wcc
}

// SetTitleZh sets the "title_zh" field.
func (wcc *WinConfigCreate) SetTitleZh(s string) *WinConfigCreate {
	wcc.mutation.SetTitleZh(s)
	return wcc
}

// SetValue sets the "value" field.
func (wcc *WinConfigCreate) SetValue(s string) *WinConfigCreate {
	wcc.mutation.SetValue(s)
	return wcc
}

// SetShowApp sets the "show_app" field.
func (wcc *WinConfigCreate) SetShowApp(i int8) *WinConfigCreate {
	wcc.mutation.SetShowApp(i)
	return wcc
}

// SetCanModify sets the "can_modify" field.
func (wcc *WinConfigCreate) SetCanModify(i int8) *WinConfigCreate {
	wcc.mutation.SetCanModify(i)
	return wcc
}

// SetStatus sets the "status" field.
func (wcc *WinConfigCreate) SetStatus(i int8) *WinConfigCreate {
	wcc.mutation.SetStatus(i)
	return wcc
}

// SetCreatedAt sets the "created_at" field.
func (wcc *WinConfigCreate) SetCreatedAt(i int32) *WinConfigCreate {
	wcc.mutation.SetCreatedAt(i)
	return wcc
}

// SetUpdatedAt sets the "updated_at" field.
func (wcc *WinConfigCreate) SetUpdatedAt(i int32) *WinConfigCreate {
	wcc.mutation.SetUpdatedAt(i)
	return wcc
}

// SetID sets the "id" field.
func (wcc *WinConfigCreate) SetID(i int32) *WinConfigCreate {
	wcc.mutation.SetID(i)
	return wcc
}

// Mutation returns the WinConfigMutation object of the builder.
func (wcc *WinConfigCreate) Mutation() *WinConfigMutation {
	return wcc.mutation
}

// Save creates the WinConfig in the database.
func (wcc *WinConfigCreate) Save(ctx context.Context) (*WinConfig, error) {
	return withHooks(ctx, wcc.sqlSave, wcc.mutation, wcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wcc *WinConfigCreate) SaveX(ctx context.Context) *WinConfig {
	v, err := wcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcc *WinConfigCreate) Exec(ctx context.Context) error {
	_, err := wcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcc *WinConfigCreate) ExecX(ctx context.Context) {
	if err := wcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcc *WinConfigCreate) check() error {
	if _, ok := wcc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "WinConfig.title"`)}
	}
	if _, ok := wcc.mutation.TitleZh(); !ok {
		return &ValidationError{Name: "title_zh", err: errors.New(`ent: missing required field "WinConfig.title_zh"`)}
	}
	if _, ok := wcc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "WinConfig.value"`)}
	}
	if _, ok := wcc.mutation.ShowApp(); !ok {
		return &ValidationError{Name: "show_app", err: errors.New(`ent: missing required field "WinConfig.show_app"`)}
	}
	if _, ok := wcc.mutation.CanModify(); !ok {
		return &ValidationError{Name: "can_modify", err: errors.New(`ent: missing required field "WinConfig.can_modify"`)}
	}
	if _, ok := wcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "WinConfig.status"`)}
	}
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WinConfig.created_at"`)}
	}
	if _, ok := wcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WinConfig.updated_at"`)}
	}
	return nil
}

func (wcc *WinConfigCreate) sqlSave(ctx context.Context) (*WinConfig, error) {
	if err := wcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	wcc.mutation.id = &_node.ID
	wcc.mutation.done = true
	return _node, nil
}

func (wcc *WinConfigCreate) createSpec() (*WinConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &WinConfig{config: wcc.config}
		_spec = sqlgraph.NewCreateSpec(winconfig.Table, sqlgraph.NewFieldSpec(winconfig.FieldID, field.TypeInt32))
	)
	if id, ok := wcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wcc.mutation.Title(); ok {
		_spec.SetField(winconfig.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := wcc.mutation.TitleZh(); ok {
		_spec.SetField(winconfig.FieldTitleZh, field.TypeString, value)
		_node.TitleZh = value
	}
	if value, ok := wcc.mutation.Value(); ok {
		_spec.SetField(winconfig.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := wcc.mutation.ShowApp(); ok {
		_spec.SetField(winconfig.FieldShowApp, field.TypeInt8, value)
		_node.ShowApp = value
	}
	if value, ok := wcc.mutation.CanModify(); ok {
		_spec.SetField(winconfig.FieldCanModify, field.TypeInt8, value)
		_node.CanModify = value
	}
	if value, ok := wcc.mutation.Status(); ok {
		_spec.SetField(winconfig.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := wcc.mutation.CreatedAt(); ok {
		_spec.SetField(winconfig.FieldCreatedAt, field.TypeInt32, value)
		_node.CreatedAt = value
	}
	if value, ok := wcc.mutation.UpdatedAt(); ok {
		_spec.SetField(winconfig.FieldUpdatedAt, field.TypeInt32, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WinConfigCreateBulk is the builder for creating many WinConfig entities in bulk.
type WinConfigCreateBulk struct {
	config
	builders []*WinConfigCreate
}

// Save creates the WinConfig entities in the database.
func (wccb *WinConfigCreateBulk) Save(ctx context.Context) ([]*WinConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wccb.builders))
	nodes := make([]*WinConfig, len(wccb.builders))
	mutators := make([]Mutator, len(wccb.builders))
	for i := range wccb.builders {
		func(i int, root context.Context) {
			builder := wccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, wccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, wccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wccb *WinConfigCreateBulk) SaveX(ctx context.Context) []*WinConfig {
	v, err := wccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wccb *WinConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := wccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wccb *WinConfigCreateBulk) ExecX(ctx context.Context) {
	if err := wccb.Exec(ctx); err != nil {
		panic(err)
	}
}
