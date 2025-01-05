// Code generated by ent, DO NOT EDIT.

package data

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/avptp/brain/internal/api/types"
	"github.com/avptp/brain/internal/generated/data/authentication"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/avptp/brain/internal/generated/data/predicate"
)

// PersonQuery is the builder for querying Person entities.
type PersonQuery struct {
	config
	ctx                      *QueryContext
	order                    []person.OrderOption
	inters                   []Interceptor
	predicates               []predicate.Person
	withAuthentications      *AuthenticationQuery
	withAuthorizations       *AuthorizationQuery
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*Person) error
	withNamedAuthentications map[string]*AuthenticationQuery
	withNamedAuthorizations  map[string]*AuthorizationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PersonQuery builder.
func (pq *PersonQuery) Where(ps ...predicate.Person) *PersonQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PersonQuery) Limit(limit int) *PersonQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PersonQuery) Offset(offset int) *PersonQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PersonQuery) Unique(unique bool) *PersonQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PersonQuery) Order(o ...person.OrderOption) *PersonQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryAuthentications chains the current query on the "authentications" edge.
func (pq *PersonQuery) QueryAuthentications() *AuthenticationQuery {
	query := (&AuthenticationClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(authentication.Table, authentication.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, person.AuthenticationsTable, person.AuthenticationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthorizations chains the current query on the "authorizations" edge.
func (pq *PersonQuery) QueryAuthorizations() *AuthorizationQuery {
	query := (&AuthorizationClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(authorization.Table, authorization.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, person.AuthorizationsTable, person.AuthorizationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Person entity from the query.
// Returns a *NotFoundError when no Person was found.
func (pq *PersonQuery) First(ctx context.Context) (*Person, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{person.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PersonQuery) FirstX(ctx context.Context) *Person {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Person ID from the query.
// Returns a *NotFoundError when no Person ID was found.
func (pq *PersonQuery) FirstID(ctx context.Context) (id types.ID, err error) {
	var ids []types.ID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{person.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PersonQuery) FirstIDX(ctx context.Context) types.ID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Person entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Person entity is found.
// Returns a *NotFoundError when no Person entities are found.
func (pq *PersonQuery) Only(ctx context.Context) (*Person, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{person.Label}
	default:
		return nil, &NotSingularError{person.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PersonQuery) OnlyX(ctx context.Context) *Person {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Person ID in the query.
// Returns a *NotSingularError when more than one Person ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PersonQuery) OnlyID(ctx context.Context) (id types.ID, err error) {
	var ids []types.ID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{person.Label}
	default:
		err = &NotSingularError{person.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PersonQuery) OnlyIDX(ctx context.Context) types.ID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Persons.
func (pq *PersonQuery) All(ctx context.Context) ([]*Person, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Person, *PersonQuery]()
	return withInterceptors[[]*Person](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PersonQuery) AllX(ctx context.Context) []*Person {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Person IDs.
func (pq *PersonQuery) IDs(ctx context.Context) (ids []types.ID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(person.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PersonQuery) IDsX(ctx context.Context) []types.ID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PersonQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PersonQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PersonQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PersonQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("data: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PersonQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PersonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PersonQuery) Clone() *PersonQuery {
	if pq == nil {
		return nil
	}
	return &PersonQuery{
		config:              pq.config,
		ctx:                 pq.ctx.Clone(),
		order:               append([]person.OrderOption{}, pq.order...),
		inters:              append([]Interceptor{}, pq.inters...),
		predicates:          append([]predicate.Person{}, pq.predicates...),
		withAuthentications: pq.withAuthentications.Clone(),
		withAuthorizations:  pq.withAuthorizations.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithAuthentications tells the query-builder to eager-load the nodes that are connected to
// the "authentications" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithAuthentications(opts ...func(*AuthenticationQuery)) *PersonQuery {
	query := (&AuthenticationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withAuthentications = query
	return pq
}

// WithAuthorizations tells the query-builder to eager-load the nodes that are connected to
// the "authorizations" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithAuthorizations(opts ...func(*AuthorizationQuery)) *PersonQuery {
	query := (&AuthorizationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withAuthorizations = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StripeID string `json:"stripe_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Person.Query().
//		GroupBy(person.FieldStripeID).
//		Aggregate(data.Count()).
//		Scan(ctx, &v)
func (pq *PersonQuery) GroupBy(field string, fields ...string) *PersonGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PersonGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = person.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StripeID string `json:"stripe_id,omitempty"`
//	}
//
//	client.Person.Query().
//		Select(person.FieldStripeID).
//		Scan(ctx, &v)
func (pq *PersonQuery) Select(fields ...string) *PersonSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PersonSelect{PersonQuery: pq}
	sbuild.label = person.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PersonSelect configured with the given aggregations.
func (pq *PersonQuery) Aggregate(fns ...AggregateFunc) *PersonSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PersonQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("data: uninitialized interceptor (forgotten import data/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !person.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("data: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	if person.Policy == nil {
		return errors.New("data: uninitialized person.Policy (forgotten import data/runtime?)")
	}
	if err := person.Policy.EvalQuery(ctx, pq); err != nil {
		return err
	}
	return nil
}

func (pq *PersonQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Person, error) {
	var (
		nodes       = []*Person{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withAuthentications != nil,
			pq.withAuthorizations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Person).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Person{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withAuthentications; query != nil {
		if err := pq.loadAuthentications(ctx, query, nodes,
			func(n *Person) { n.Edges.Authentications = []*Authentication{} },
			func(n *Person, e *Authentication) { n.Edges.Authentications = append(n.Edges.Authentications, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withAuthorizations; query != nil {
		if err := pq.loadAuthorizations(ctx, query, nodes,
			func(n *Person) { n.Edges.Authorizations = []*Authorization{} },
			func(n *Person, e *Authorization) { n.Edges.Authorizations = append(n.Edges.Authorizations, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedAuthentications {
		if err := pq.loadAuthentications(ctx, query, nodes,
			func(n *Person) { n.appendNamedAuthentications(name) },
			func(n *Person, e *Authentication) { n.appendNamedAuthentications(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedAuthorizations {
		if err := pq.loadAuthorizations(ctx, query, nodes,
			func(n *Person) { n.appendNamedAuthorizations(name) },
			func(n *Person, e *Authorization) { n.appendNamedAuthorizations(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PersonQuery) loadAuthentications(ctx context.Context, query *AuthenticationQuery, nodes []*Person, init func(*Person), assign func(*Person, *Authentication)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[types.ID]*Person)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(authentication.FieldPersonID)
	}
	query.Where(predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(person.AuthenticationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PersonID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "person_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PersonQuery) loadAuthorizations(ctx context.Context, query *AuthorizationQuery, nodes []*Person, init func(*Person), assign func(*Person, *Authorization)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[types.ID]*Person)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(authorization.FieldPersonID)
	}
	query.Where(predicate.Authorization(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(person.AuthorizationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PersonID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "person_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PersonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PersonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for i := range fields {
			if fields[i] != person.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PersonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(person.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = person.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAuthentications tells the query-builder to eager-load the nodes that are connected to the "authentications"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithNamedAuthentications(name string, opts ...func(*AuthenticationQuery)) *PersonQuery {
	query := (&AuthenticationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedAuthentications == nil {
		pq.withNamedAuthentications = make(map[string]*AuthenticationQuery)
	}
	pq.withNamedAuthentications[name] = query
	return pq
}

// WithNamedAuthorizations tells the query-builder to eager-load the nodes that are connected to the "authorizations"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithNamedAuthorizations(name string, opts ...func(*AuthorizationQuery)) *PersonQuery {
	query := (&AuthorizationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedAuthorizations == nil {
		pq.withNamedAuthorizations = make(map[string]*AuthorizationQuery)
	}
	pq.withNamedAuthorizations[name] = query
	return pq
}

// PersonGroupBy is the group-by builder for Person entities.
type PersonGroupBy struct {
	selector
	build *PersonQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PersonGroupBy) Aggregate(fns ...AggregateFunc) *PersonGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PersonGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonQuery, *PersonGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PersonGroupBy) sqlScan(ctx context.Context, root *PersonQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PersonSelect is the builder for selecting fields of Person entities.
type PersonSelect struct {
	*PersonQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PersonSelect) Aggregate(fns ...AggregateFunc) *PersonSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PersonSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonQuery, *PersonSelect](ctx, ps.PersonQuery, ps, ps.inters, v)
}

func (ps *PersonSelect) sqlScan(ctx context.Context, root *PersonQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
