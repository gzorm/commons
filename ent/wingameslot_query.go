// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/commons/ent/predicate"
	"github.com/gzorm/commons/ent/wingameslot"
)

// WinGameSlotQuery is the builder for querying WinGameSlot entities.
type WinGameSlotQuery struct {
	config
	ctx        *QueryContext
	order      []wingameslot.OrderOption
	inters     []Interceptor
	predicates []predicate.WinGameSlot
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WinGameSlotQuery builder.
func (wgsq *WinGameSlotQuery) Where(ps ...predicate.WinGameSlot) *WinGameSlotQuery {
	wgsq.predicates = append(wgsq.predicates, ps...)
	return wgsq
}

// Limit the number of records to be returned by this query.
func (wgsq *WinGameSlotQuery) Limit(limit int) *WinGameSlotQuery {
	wgsq.ctx.Limit = &limit
	return wgsq
}

// Offset to start from.
func (wgsq *WinGameSlotQuery) Offset(offset int) *WinGameSlotQuery {
	wgsq.ctx.Offset = &offset
	return wgsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wgsq *WinGameSlotQuery) Unique(unique bool) *WinGameSlotQuery {
	wgsq.ctx.Unique = &unique
	return wgsq
}

// Order specifies how the records should be ordered.
func (wgsq *WinGameSlotQuery) Order(o ...wingameslot.OrderOption) *WinGameSlotQuery {
	wgsq.order = append(wgsq.order, o...)
	return wgsq
}

// First returns the first WinGameSlot entity from the query.
// Returns a *NotFoundError when no WinGameSlot was found.
func (wgsq *WinGameSlotQuery) First(ctx context.Context) (*WinGameSlot, error) {
	nodes, err := wgsq.Limit(1).All(setContextOp(ctx, wgsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wingameslot.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) FirstX(ctx context.Context) *WinGameSlot {
	node, err := wgsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WinGameSlot ID from the query.
// Returns a *NotFoundError when no WinGameSlot ID was found.
func (wgsq *WinGameSlotQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wgsq.Limit(1).IDs(setContextOp(ctx, wgsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wingameslot.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) FirstIDX(ctx context.Context) string {
	id, err := wgsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WinGameSlot entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WinGameSlot entity is found.
// Returns a *NotFoundError when no WinGameSlot entities are found.
func (wgsq *WinGameSlotQuery) Only(ctx context.Context) (*WinGameSlot, error) {
	nodes, err := wgsq.Limit(2).All(setContextOp(ctx, wgsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wingameslot.Label}
	default:
		return nil, &NotSingularError{wingameslot.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) OnlyX(ctx context.Context) *WinGameSlot {
	node, err := wgsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WinGameSlot ID in the query.
// Returns a *NotSingularError when more than one WinGameSlot ID is found.
// Returns a *NotFoundError when no entities are found.
func (wgsq *WinGameSlotQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wgsq.Limit(2).IDs(setContextOp(ctx, wgsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wingameslot.Label}
	default:
		err = &NotSingularError{wingameslot.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) OnlyIDX(ctx context.Context) string {
	id, err := wgsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WinGameSlots.
func (wgsq *WinGameSlotQuery) All(ctx context.Context) ([]*WinGameSlot, error) {
	ctx = setContextOp(ctx, wgsq.ctx, "All")
	if err := wgsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WinGameSlot, *WinGameSlotQuery]()
	return withInterceptors[[]*WinGameSlot](ctx, wgsq, qr, wgsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) AllX(ctx context.Context) []*WinGameSlot {
	nodes, err := wgsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WinGameSlot IDs.
func (wgsq *WinGameSlotQuery) IDs(ctx context.Context) (ids []string, err error) {
	if wgsq.ctx.Unique == nil && wgsq.path != nil {
		wgsq.Unique(true)
	}
	ctx = setContextOp(ctx, wgsq.ctx, "IDs")
	if err = wgsq.Select(wingameslot.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) IDsX(ctx context.Context) []string {
	ids, err := wgsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wgsq *WinGameSlotQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wgsq.ctx, "Count")
	if err := wgsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wgsq, querierCount[*WinGameSlotQuery](), wgsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) CountX(ctx context.Context) int {
	count, err := wgsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wgsq *WinGameSlotQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wgsq.ctx, "Exist")
	switch _, err := wgsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wgsq *WinGameSlotQuery) ExistX(ctx context.Context) bool {
	exist, err := wgsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WinGameSlotQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wgsq *WinGameSlotQuery) Clone() *WinGameSlotQuery {
	if wgsq == nil {
		return nil
	}
	return &WinGameSlotQuery{
		config:     wgsq.config,
		ctx:        wgsq.ctx.Clone(),
		order:      append([]wingameslot.OrderOption{}, wgsq.order...),
		inters:     append([]Interceptor{}, wgsq.inters...),
		predicates: append([]predicate.WinGameSlot{}, wgsq.predicates...),
		// clone intermediate query.
		sql:  wgsq.sql.Clone(),
		path: wgsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GameID int32 `json:"game_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WinGameSlot.Query().
//		GroupBy(wingameslot.FieldGameID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wgsq *WinGameSlotQuery) GroupBy(field string, fields ...string) *WinGameSlotGroupBy {
	wgsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WinGameSlotGroupBy{build: wgsq}
	grbuild.flds = &wgsq.ctx.Fields
	grbuild.label = wingameslot.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GameID int32 `json:"game_id,omitempty"`
//	}
//
//	client.WinGameSlot.Query().
//		Select(wingameslot.FieldGameID).
//		Scan(ctx, &v)
func (wgsq *WinGameSlotQuery) Select(fields ...string) *WinGameSlotSelect {
	wgsq.ctx.Fields = append(wgsq.ctx.Fields, fields...)
	sbuild := &WinGameSlotSelect{WinGameSlotQuery: wgsq}
	sbuild.label = wingameslot.Label
	sbuild.flds, sbuild.scan = &wgsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WinGameSlotSelect configured with the given aggregations.
func (wgsq *WinGameSlotQuery) Aggregate(fns ...AggregateFunc) *WinGameSlotSelect {
	return wgsq.Select().Aggregate(fns...)
}

func (wgsq *WinGameSlotQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wgsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wgsq); err != nil {
				return err
			}
		}
	}
	for _, f := range wgsq.ctx.Fields {
		if !wingameslot.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wgsq.path != nil {
		prev, err := wgsq.path(ctx)
		if err != nil {
			return err
		}
		wgsq.sql = prev
	}
	return nil
}

func (wgsq *WinGameSlotQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WinGameSlot, error) {
	var (
		nodes = []*WinGameSlot{}
		_spec = wgsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WinGameSlot).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WinGameSlot{config: wgsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wgsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wgsq *WinGameSlotQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wgsq.querySpec()
	_spec.Node.Columns = wgsq.ctx.Fields
	if len(wgsq.ctx.Fields) > 0 {
		_spec.Unique = wgsq.ctx.Unique != nil && *wgsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wgsq.driver, _spec)
}

func (wgsq *WinGameSlotQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wingameslot.Table, wingameslot.Columns, sqlgraph.NewFieldSpec(wingameslot.FieldID, field.TypeString))
	_spec.From = wgsq.sql
	if unique := wgsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wgsq.path != nil {
		_spec.Unique = true
	}
	if fields := wgsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wingameslot.FieldID)
		for i := range fields {
			if fields[i] != wingameslot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wgsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wgsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wgsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wgsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wgsq *WinGameSlotQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wgsq.driver.Dialect())
	t1 := builder.Table(wingameslot.Table)
	columns := wgsq.ctx.Fields
	if len(columns) == 0 {
		columns = wingameslot.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wgsq.sql != nil {
		selector = wgsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wgsq.ctx.Unique != nil && *wgsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wgsq.predicates {
		p(selector)
	}
	for _, p := range wgsq.order {
		p(selector)
	}
	if offset := wgsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wgsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (wgsq *WinGameSlotQuery) SqlQuery(ctx context.Context) (r1 string, r2 []any) {
	builder := sql.Dialect(wgsq.driver.Dialect())
	t1 := builder.Table(wingameslot.Table)
	columns := wgsq.ctx.Fields
	if len(columns) == 0 {
		columns = wingameslot.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wgsq.sql != nil {
		selector = wgsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wgsq.ctx.Unique != nil && *wgsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wgsq.predicates {
		p(selector)
	}
	for _, p := range wgsq.order {
		p(selector)
	}
	if offset := wgsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wgsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	r1, r2 = selector.Query()
	//r3 = make([]*WinGameSlot,0)

	return
}

// WinGameSlotGroupBy is the group-by builder for WinGameSlot entities.
type WinGameSlotGroupBy struct {
	selector
	build *WinGameSlotQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgsgb *WinGameSlotGroupBy) Aggregate(fns ...AggregateFunc) *WinGameSlotGroupBy {
	wgsgb.fns = append(wgsgb.fns, fns...)
	return wgsgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgsgb *WinGameSlotGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgsgb.build.ctx, "GroupBy")
	if err := wgsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WinGameSlotQuery, *WinGameSlotGroupBy](ctx, wgsgb.build, wgsgb, wgsgb.build.inters, v)
}

func (wgsgb *WinGameSlotGroupBy) sqlScan(ctx context.Context, root *WinGameSlotQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgsgb.fns))
	for _, fn := range wgsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgsgb.flds)+len(wgsgb.fns))
		for _, f := range *wgsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WinGameSlotSelect is the builder for selecting fields of WinGameSlot entities.
type WinGameSlotSelect struct {
	*WinGameSlotQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wgss *WinGameSlotSelect) Aggregate(fns ...AggregateFunc) *WinGameSlotSelect {
	wgss.fns = append(wgss.fns, fns...)
	return wgss
}

// Scan applies the selector query and scans the result into the given value.
func (wgss *WinGameSlotSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgss.ctx, "Select")
	if err := wgss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WinGameSlotQuery, *WinGameSlotSelect](ctx, wgss.WinGameSlotQuery, wgss, wgss.inters, v)
}

func (wgss *WinGameSlotSelect) sqlScan(ctx context.Context, root *WinGameSlotQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wgss.fns))
	for _, fn := range wgss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wgss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
