// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/B6111427/app/ent/booking"
	"github.com/B6111427/app/ent/cliententity"
	"github.com/B6111427/app/ent/predicate"
	"github.com/B6111427/app/ent/status"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClientEntityQuery is the builder for querying ClientEntity entities.
type ClientEntityQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.ClientEntity
	// eager-loading edges.
	withBooked *BookingQuery
	withState  *StatusQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ceq *ClientEntityQuery) Where(ps ...predicate.ClientEntity) *ClientEntityQuery {
	ceq.predicates = append(ceq.predicates, ps...)
	return ceq
}

// Limit adds a limit step to the query.
func (ceq *ClientEntityQuery) Limit(limit int) *ClientEntityQuery {
	ceq.limit = &limit
	return ceq
}

// Offset adds an offset step to the query.
func (ceq *ClientEntityQuery) Offset(offset int) *ClientEntityQuery {
	ceq.offset = &offset
	return ceq
}

// Order adds an order step to the query.
func (ceq *ClientEntityQuery) Order(o ...OrderFunc) *ClientEntityQuery {
	ceq.order = append(ceq.order, o...)
	return ceq
}

// QueryBooked chains the current query on the booked edge.
func (ceq *ClientEntityQuery) QueryBooked() *BookingQuery {
	query := &BookingQuery{config: ceq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cliententity.Table, cliententity.FieldID, ceq.sqlQuery()),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cliententity.BookedTable, cliententity.BookedColumn),
		)
		fromU = sqlgraph.SetNeighbors(ceq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryState chains the current query on the state edge.
func (ceq *ClientEntityQuery) QueryState() *StatusQuery {
	query := &StatusQuery{config: ceq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cliententity.Table, cliententity.FieldID, ceq.sqlQuery()),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cliententity.StateTable, cliententity.StateColumn),
		)
		fromU = sqlgraph.SetNeighbors(ceq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClientEntity entity in the query. Returns *NotFoundError when no cliententity was found.
func (ceq *ClientEntityQuery) First(ctx context.Context) (*ClientEntity, error) {
	ces, err := ceq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ces) == 0 {
		return nil, &NotFoundError{cliententity.Label}
	}
	return ces[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ceq *ClientEntityQuery) FirstX(ctx context.Context) *ClientEntity {
	ce, err := ceq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ce
}

// FirstID returns the first ClientEntity id in the query. Returns *NotFoundError when no id was found.
func (ceq *ClientEntityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ceq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cliententity.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ceq *ClientEntityQuery) FirstXID(ctx context.Context) int {
	id, err := ceq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ClientEntity entity in the query, returns an error if not exactly one entity was returned.
func (ceq *ClientEntityQuery) Only(ctx context.Context) (*ClientEntity, error) {
	ces, err := ceq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ces) {
	case 1:
		return ces[0], nil
	case 0:
		return nil, &NotFoundError{cliententity.Label}
	default:
		return nil, &NotSingularError{cliententity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ceq *ClientEntityQuery) OnlyX(ctx context.Context) *ClientEntity {
	ce, err := ceq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ce
}

// OnlyID returns the only ClientEntity id in the query, returns an error if not exactly one id was returned.
func (ceq *ClientEntityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ceq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = &NotSingularError{cliententity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ceq *ClientEntityQuery) OnlyIDX(ctx context.Context) int {
	id, err := ceq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClientEntities.
func (ceq *ClientEntityQuery) All(ctx context.Context) ([]*ClientEntity, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ceq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ceq *ClientEntityQuery) AllX(ctx context.Context) []*ClientEntity {
	ces, err := ceq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ces
}

// IDs executes the query and returns a list of ClientEntity ids.
func (ceq *ClientEntityQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ceq.Select(cliententity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ceq *ClientEntityQuery) IDsX(ctx context.Context) []int {
	ids, err := ceq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ceq *ClientEntityQuery) Count(ctx context.Context) (int, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ceq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ceq *ClientEntityQuery) CountX(ctx context.Context) int {
	count, err := ceq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ceq *ClientEntityQuery) Exist(ctx context.Context) (bool, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ceq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ceq *ClientEntityQuery) ExistX(ctx context.Context) bool {
	exist, err := ceq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ceq *ClientEntityQuery) Clone() *ClientEntityQuery {
	return &ClientEntityQuery{
		config:     ceq.config,
		limit:      ceq.limit,
		offset:     ceq.offset,
		order:      append([]OrderFunc{}, ceq.order...),
		unique:     append([]string{}, ceq.unique...),
		predicates: append([]predicate.ClientEntity{}, ceq.predicates...),
		// clone intermediate query.
		sql:  ceq.sql.Clone(),
		path: ceq.path,
	}
}

//  WithBooked tells the query-builder to eager-loads the nodes that are connected to
// the "booked" edge. The optional arguments used to configure the query builder of the edge.
func (ceq *ClientEntityQuery) WithBooked(opts ...func(*BookingQuery)) *ClientEntityQuery {
	query := &BookingQuery{config: ceq.config}
	for _, opt := range opts {
		opt(query)
	}
	ceq.withBooked = query
	return ceq
}

//  WithState tells the query-builder to eager-loads the nodes that are connected to
// the "state" edge. The optional arguments used to configure the query builder of the edge.
func (ceq *ClientEntityQuery) WithState(opts ...func(*StatusQuery)) *ClientEntityQuery {
	query := &StatusQuery{config: ceq.config}
	for _, opt := range opts {
		opt(query)
	}
	ceq.withState = query
	return ceq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CLIENTNAME string `json:"CLIENT_NAME,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClientEntity.Query().
//		GroupBy(cliententity.FieldCLIENTNAME).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ceq *ClientEntityQuery) GroupBy(field string, fields ...string) *ClientEntityGroupBy {
	group := &ClientEntityGroupBy{config: ceq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ceq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CLIENTNAME string `json:"CLIENT_NAME,omitempty"`
//	}
//
//	client.ClientEntity.Query().
//		Select(cliententity.FieldCLIENTNAME).
//		Scan(ctx, &v)
//
func (ceq *ClientEntityQuery) Select(field string, fields ...string) *ClientEntitySelect {
	selector := &ClientEntitySelect{config: ceq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ceq.sqlQuery(), nil
	}
	return selector
}

func (ceq *ClientEntityQuery) prepareQuery(ctx context.Context) error {
	if ceq.path != nil {
		prev, err := ceq.path(ctx)
		if err != nil {
			return err
		}
		ceq.sql = prev
	}
	return nil
}

func (ceq *ClientEntityQuery) sqlAll(ctx context.Context) ([]*ClientEntity, error) {
	var (
		nodes       = []*ClientEntity{}
		withFKs     = ceq.withFKs
		_spec       = ceq.querySpec()
		loadedTypes = [2]bool{
			ceq.withBooked != nil,
			ceq.withState != nil,
		}
	)
	if ceq.withState != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cliententity.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &ClientEntity{config: ceq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ceq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ceq.withBooked; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ClientEntity)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Booking(func(s *sql.Selector) {
			s.Where(sql.InValues(cliententity.BookedColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.CLIENT_ID
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "CLIENT_ID" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "CLIENT_ID" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Booked = append(node.Edges.Booked, n)
		}
	}

	if query := ceq.withState; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ClientEntity)
		for i := range nodes {
			if fk := nodes[i].Status_ID; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(status.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "Status_ID" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.State = n
			}
		}
	}

	return nodes, nil
}

func (ceq *ClientEntityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ceq.querySpec()
	return sqlgraph.CountNodes(ctx, ceq.driver, _spec)
}

func (ceq *ClientEntityQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ceq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ceq *ClientEntityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cliententity.Table,
			Columns: cliententity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cliententity.FieldID,
			},
		},
		From:   ceq.sql,
		Unique: true,
	}
	if ps := ceq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ceq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ceq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ceq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ceq *ClientEntityQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ceq.driver.Dialect())
	t1 := builder.Table(cliententity.Table)
	selector := builder.Select(t1.Columns(cliententity.Columns...)...).From(t1)
	if ceq.sql != nil {
		selector = ceq.sql
		selector.Select(selector.Columns(cliententity.Columns...)...)
	}
	for _, p := range ceq.predicates {
		p(selector)
	}
	for _, p := range ceq.order {
		p(selector)
	}
	if offset := ceq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ceq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClientEntityGroupBy is the builder for group-by ClientEntity entities.
type ClientEntityGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cegb *ClientEntityGroupBy) Aggregate(fns ...AggregateFunc) *ClientEntityGroupBy {
	cegb.fns = append(cegb.fns, fns...)
	return cegb
}

// Scan applies the group-by query and scan the result into the given value.
func (cegb *ClientEntityGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cegb.path(ctx)
	if err != nil {
		return err
	}
	cegb.sql = query
	return cegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ClientEntityGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) StringsX(ctx context.Context) []string {
	v, err := cegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntityGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) StringX(ctx context.Context) string {
	v, err := cegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ClientEntityGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) IntsX(ctx context.Context) []int {
	v, err := cegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntityGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) IntX(ctx context.Context) int {
	v, err := cegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ClientEntityGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntityGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ClientEntityGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (cegb *ClientEntityGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntityGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cegb *ClientEntityGroupBy) BoolX(ctx context.Context) bool {
	v, err := cegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cegb *ClientEntityGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cegb.sqlQuery().Query()
	if err := cegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cegb *ClientEntityGroupBy) sqlQuery() *sql.Selector {
	selector := cegb.sql
	columns := make([]string, 0, len(cegb.fields)+len(cegb.fns))
	columns = append(columns, cegb.fields...)
	for _, fn := range cegb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(cegb.fields...)
}

// ClientEntitySelect is the builder for select fields of ClientEntity entities.
type ClientEntitySelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ces *ClientEntitySelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ces.path(ctx)
	if err != nil {
		return err
	}
	ces.sql = query
	return ces.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ces *ClientEntitySelect) ScanX(ctx context.Context, v interface{}) {
	if err := ces.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) Strings(ctx context.Context) ([]string, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ClientEntitySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ces *ClientEntitySelect) StringsX(ctx context.Context) []string {
	v, err := ces.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ces.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntitySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ces *ClientEntitySelect) StringX(ctx context.Context) string {
	v, err := ces.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) Ints(ctx context.Context) ([]int, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ClientEntitySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ces *ClientEntitySelect) IntsX(ctx context.Context) []int {
	v, err := ces.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ces.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntitySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ces *ClientEntitySelect) IntX(ctx context.Context) int {
	v, err := ces.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ClientEntitySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ces *ClientEntitySelect) Float64sX(ctx context.Context) []float64 {
	v, err := ces.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ces.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntitySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ces *ClientEntitySelect) Float64X(ctx context.Context) float64 {
	v, err := ces.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ClientEntitySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ces *ClientEntitySelect) BoolsX(ctx context.Context) []bool {
	v, err := ces.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ces *ClientEntitySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ces.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cliententity.Label}
	default:
		err = fmt.Errorf("ent: ClientEntitySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ces *ClientEntitySelect) BoolX(ctx context.Context) bool {
	v, err := ces.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ces *ClientEntitySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ces.sqlQuery().Query()
	if err := ces.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ces *ClientEntitySelect) sqlQuery() sql.Querier {
	selector := ces.sql
	selector.Select(selector.Columns(ces.fields...)...)
	return selector
}
