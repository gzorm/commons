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
	"github.com/gzorm/commons/ent/windict"
)

// WinDictQuery is the builder for querying WinDict entities.
type WinDictQuery struct {
	config
	ctx        *QueryContext
	order      []windict.OrderOption
	inters     []Interceptor
	predicates []predicate.WinDict
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WinDictQuery builder.
func (wdq *WinDictQuery) Where(ps ...predicate.WinDict) *WinDictQuery {
	wdq.predicates = append(wdq.predicates, ps...)
	return wdq
}

// Limit the number of records to be returned by this query.
func (wdq *WinDictQuery) Limit(limit int) *WinDictQuery {
	wdq.ctx.Limit = &limit
	return wdq
}

// Offset to start from.
func (wdq *WinDictQuery) Offset(offset int) *WinDictQuery {
	wdq.ctx.Offset = &offset
	return wdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wdq *WinDictQuery) Unique(unique bool) *WinDictQuery {
	wdq.ctx.Unique = &unique
	return wdq
}

// Order specifies how the records should be ordered.
func (wdq *WinDictQuery) Order(o ...windict.OrderOption) *WinDictQuery {
	wdq.order = append(wdq.order, o...)
	return wdq
}

// First returns the first WinDict entity from the query.
// Returns a *NotFoundError when no WinDict was found.
func (wdq *WinDictQuery) First(ctx context.Context) (*WinDict, error) {
	nodes, err := wdq.Limit(1).All(setContextOp(ctx, wdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{windict.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wdq *WinDictQuery) FirstX(ctx context.Context) *WinDict {
	node, err := wdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WinDict ID from the query.
// Returns a *NotFoundError when no WinDict ID was found.
func (wdq *WinDictQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = wdq.Limit(1).IDs(setContextOp(ctx, wdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{windict.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wdq *WinDictQuery) FirstIDX(ctx context.Context) int32 {
	id, err := wdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WinDict entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WinDict entity is found.
// Returns a *NotFoundError when no WinDict entities are found.
func (wdq *WinDictQuery) Only(ctx context.Context) (*WinDict, error) {
	nodes, err := wdq.Limit(2).All(setContextOp(ctx, wdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{windict.Label}
	default:
		return nil, &NotSingularError{windict.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wdq *WinDictQuery) OnlyX(ctx context.Context) *WinDict {
	node, err := wdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WinDict ID in the query.
// Returns a *NotSingularError when more than one WinDict ID is found.
// Returns a *NotFoundError when no entities are found.
func (wdq *WinDictQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = wdq.Limit(2).IDs(setContextOp(ctx, wdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{windict.Label}
	default:
		err = &NotSingularError{windict.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wdq *WinDictQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := wdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WinDicts.
func (wdq *WinDictQuery) All(ctx context.Context) ([]*WinDict, error) {
	ctx = setContextOp(ctx, wdq.ctx, "All")
	if err := wdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WinDict, *WinDictQuery]()
	return withInterceptors[[]*WinDict](ctx, wdq, qr, wdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wdq *WinDictQuery) AllX(ctx context.Context) []*WinDict {
	nodes, err := wdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WinDict IDs.
func (wdq *WinDictQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if wdq.ctx.Unique == nil && wdq.path != nil {
		wdq.Unique(true)
	}
	ctx = setContextOp(ctx, wdq.ctx, "IDs")
	if err = wdq.Select(windict.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wdq *WinDictQuery) IDsX(ctx context.Context) []int32 {
	ids, err := wdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wdq *WinDictQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wdq.ctx, "Count")
	if err := wdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wdq, querierCount[*WinDictQuery](), wdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wdq *WinDictQuery) CountX(ctx context.Context) int {
	count, err := wdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wdq *WinDictQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wdq.ctx, "Exist")
	switch _, err := wdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wdq *WinDictQuery) ExistX(ctx context.Context) bool {
	exist, err := wdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WinDictQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wdq *WinDictQuery) Clone() *WinDictQuery {
	if wdq == nil {
		return nil
	}
	return &WinDictQuery{
		config:     wdq.config,
		ctx:        wdq.ctx.Clone(),
		order:      append([]windict.OrderOption{}, wdq.order...),
		inters:     append([]Interceptor{}, wdq.inters...),
		predicates: append([]predicate.WinDict{}, wdq.predicates...),
		// clone intermediate query.
		sql:  wdq.sql.Clone(),
		path: wdq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WinDict.Query().
//		GroupBy(windict.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wdq *WinDictQuery) GroupBy(field string, fields ...string) *WinDictGroupBy {
	wdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WinDictGroupBy{build: wdq}
	grbuild.flds = &wdq.ctx.Fields
	grbuild.label = windict.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.WinDict.Query().
//		Select(windict.FieldTitle).
//		Scan(ctx, &v)
func (wdq *WinDictQuery) Select(fields ...string) *WinDictSelect {
	wdq.ctx.Fields = append(wdq.ctx.Fields, fields...)
	sbuild := &WinDictSelect{WinDictQuery: wdq}
	sbuild.label = windict.Label
	sbuild.flds, sbuild.scan = &wdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WinDictSelect configured with the given aggregations.
func (wdq *WinDictQuery) Aggregate(fns ...AggregateFunc) *WinDictSelect {
	return wdq.Select().Aggregate(fns...)
}

func (wdq *WinDictQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wdq); err != nil {
				return err
			}
		}
	}
	for _, f := range wdq.ctx.Fields {
		if !windict.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wdq.path != nil {
		prev, err := wdq.path(ctx)
		if err != nil {
			return err
		}
		wdq.sql = prev
	}
	return nil
}

func (wdq *WinDictQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WinDict, error) {
	var (
		nodes = []*WinDict{}
		_spec = wdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WinDict).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WinDict{config: wdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wdq *WinDictQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wdq.querySpec()
	_spec.Node.Columns = wdq.ctx.Fields
	if len(wdq.ctx.Fields) > 0 {
		_spec.Unique = wdq.ctx.Unique != nil && *wdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wdq.driver, _spec)
}

func (wdq *WinDictQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(windict.Table, windict.Columns, sqlgraph.NewFieldSpec(windict.FieldID, field.TypeInt32))
	_spec.From = wdq.sql
	if unique := wdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wdq.path != nil {
		_spec.Unique = true
	}
	if fields := wdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, windict.FieldID)
		for i := range fields {
			if fields[i] != windict.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wdq *WinDictQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wdq.driver.Dialect())
	t1 := builder.Table(windict.Table)
	columns := wdq.ctx.Fields
	if len(columns) == 0 {
		columns = windict.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wdq.sql != nil {
		selector = wdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wdq.ctx.Unique != nil && *wdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wdq.predicates {
		p(selector)
	}
	for _, p := range wdq.order {
		p(selector)
	}
	if offset := wdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (wdq *WinDictQuery) SqlQuery(ctx context.Context) (r1 string, r2 []any) {
	builder := sql.Dialect(wdq.driver.Dialect())
	t1 := builder.Table(windict.Table)
	columns := wdq.ctx.Fields
	if len(columns) == 0 {
		columns = windict.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wdq.sql != nil {
		selector = wdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wdq.ctx.Unique != nil && *wdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wdq.predicates {
		p(selector)
	}
	for _, p := range wdq.order {
		p(selector)
	}
	if offset := wdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	r1, r2 = selector.Query()
	//r3 = make([]*WinDict,0)

	return
}

// WinDictGroupBy is the group-by builder for WinDict entities.
type WinDictGroupBy struct {
	selector
	build *WinDictQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wdgb *WinDictGroupBy) Aggregate(fns ...AggregateFunc) *WinDictGroupBy {
	wdgb.fns = append(wdgb.fns, fns...)
	return wdgb
}

// Scan applies the selector query and scans the result into the given value.
func (wdgb *WinDictGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wdgb.build.ctx, "GroupBy")
	if err := wdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WinDictQuery, *WinDictGroupBy](ctx, wdgb.build, wdgb, wdgb.build.inters, v)
}

func (wdgb *WinDictGroupBy) sqlScan(ctx context.Context, root *WinDictQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wdgb.fns))
	for _, fn := range wdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wdgb.flds)+len(wdgb.fns))
		for _, f := range *wdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WinDictSelect is the builder for selecting fields of WinDict entities.
type WinDictSelect struct {
	*WinDictQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wds *WinDictSelect) Aggregate(fns ...AggregateFunc) *WinDictSelect {
	wds.fns = append(wds.fns, fns...)
	return wds
}

// Scan applies the selector query and scans the result into the given value.
func (wds *WinDictSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wds.ctx, "Select")
	if err := wds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WinDictQuery, *WinDictSelect](ctx, wds.WinDictQuery, wds, wds.inters, v)
}

func (wds *WinDictSelect) sqlScan(ctx context.Context, root *WinDictQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wds.fns))
	for _, fn := range wds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
