// Code generated by ent, DO NOT EDIT.

package data

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/avptp/brain/internal/api/types"
	"github.com/avptp/brain/internal/generated/data/authentication"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[types.ID]
	PageInfo       = entgql.PageInfo[types.ID]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// AuthenticationEdge is the edge representation of Authentication.
type AuthenticationEdge struct {
	Node   *Authentication `json:"node"`
	Cursor Cursor          `json:"cursor"`
}

// AuthenticationConnection is the connection containing edges to Authentication.
type AuthenticationConnection struct {
	Edges      []*AuthenticationEdge `json:"edges"`
	PageInfo   PageInfo              `json:"pageInfo"`
	TotalCount int                   `json:"totalCount"`
}

func (c *AuthenticationConnection) build(nodes []*Authentication, pager *authenticationPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Authentication
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Authentication {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Authentication {
			return nodes[i]
		}
	}
	c.Edges = make([]*AuthenticationEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &AuthenticationEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// AuthenticationPaginateOption enables pagination customization.
type AuthenticationPaginateOption func(*authenticationPager) error

// WithAuthenticationOrder configures pagination ordering.
func WithAuthenticationOrder(order *AuthenticationOrder) AuthenticationPaginateOption {
	if order == nil {
		order = DefaultAuthenticationOrder
	}
	o := *order
	return func(pager *authenticationPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultAuthenticationOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithAuthenticationFilter configures pagination filter.
func WithAuthenticationFilter(filter func(*AuthenticationQuery) (*AuthenticationQuery, error)) AuthenticationPaginateOption {
	return func(pager *authenticationPager) error {
		if filter == nil {
			return errors.New("AuthenticationQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type authenticationPager struct {
	reverse bool
	order   *AuthenticationOrder
	filter  func(*AuthenticationQuery) (*AuthenticationQuery, error)
}

func newAuthenticationPager(opts []AuthenticationPaginateOption, reverse bool) (*authenticationPager, error) {
	pager := &authenticationPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultAuthenticationOrder
	}
	return pager, nil
}

func (p *authenticationPager) applyFilter(query *AuthenticationQuery) (*AuthenticationQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *authenticationPager) toCursor(a *Authentication) Cursor {
	return p.order.Field.toCursor(a)
}

func (p *authenticationPager) applyCursors(query *AuthenticationQuery, after, before *Cursor) (*AuthenticationQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultAuthenticationOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *authenticationPager) applyOrder(query *AuthenticationQuery) *AuthenticationQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultAuthenticationOrder.Field {
		query = query.Order(DefaultAuthenticationOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *authenticationPager) orderExpr(query *AuthenticationQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultAuthenticationOrder.Field {
			b.Comma().Ident(DefaultAuthenticationOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Authentication.
func (a *AuthenticationQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...AuthenticationPaginateOption,
) (*AuthenticationConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newAuthenticationPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if a, err = pager.applyFilter(a); err != nil {
		return nil, err
	}
	conn := &AuthenticationConnection{Edges: []*AuthenticationEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := a.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if a, err = pager.applyCursors(a, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		a.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := a.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	a = pager.applyOrder(a)
	nodes, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// AuthenticationOrderField defines the ordering field of Authentication.
type AuthenticationOrderField struct {
	// Value extracts the ordering value from the given Authentication.
	Value    func(*Authentication) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) authentication.OrderOption
	toCursor func(*Authentication) Cursor
}

// AuthenticationOrder defines the ordering of Authentication.
type AuthenticationOrder struct {
	Direction OrderDirection            `json:"direction"`
	Field     *AuthenticationOrderField `json:"field"`
}

// DefaultAuthenticationOrder is the default ordering of Authentication.
var DefaultAuthenticationOrder = &AuthenticationOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &AuthenticationOrderField{
		Value: func(a *Authentication) (ent.Value, error) {
			return a.ID, nil
		},
		column: authentication.FieldID,
		toTerm: authentication.ByID,
		toCursor: func(a *Authentication) Cursor {
			return Cursor{ID: a.ID}
		},
	},
}

// ToEdge converts Authentication into AuthenticationEdge.
func (a *Authentication) ToEdge(order *AuthenticationOrder) *AuthenticationEdge {
	if order == nil {
		order = DefaultAuthenticationOrder
	}
	return &AuthenticationEdge{
		Node:   a,
		Cursor: order.Field.toCursor(a),
	}
}

// AuthorizationEdge is the edge representation of Authorization.
type AuthorizationEdge struct {
	Node   *Authorization `json:"node"`
	Cursor Cursor         `json:"cursor"`
}

// AuthorizationConnection is the connection containing edges to Authorization.
type AuthorizationConnection struct {
	Edges      []*AuthorizationEdge `json:"edges"`
	PageInfo   PageInfo             `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

func (c *AuthorizationConnection) build(nodes []*Authorization, pager *authorizationPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Authorization
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Authorization {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Authorization {
			return nodes[i]
		}
	}
	c.Edges = make([]*AuthorizationEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &AuthorizationEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// AuthorizationPaginateOption enables pagination customization.
type AuthorizationPaginateOption func(*authorizationPager) error

// WithAuthorizationOrder configures pagination ordering.
func WithAuthorizationOrder(order *AuthorizationOrder) AuthorizationPaginateOption {
	if order == nil {
		order = DefaultAuthorizationOrder
	}
	o := *order
	return func(pager *authorizationPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultAuthorizationOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithAuthorizationFilter configures pagination filter.
func WithAuthorizationFilter(filter func(*AuthorizationQuery) (*AuthorizationQuery, error)) AuthorizationPaginateOption {
	return func(pager *authorizationPager) error {
		if filter == nil {
			return errors.New("AuthorizationQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type authorizationPager struct {
	reverse bool
	order   *AuthorizationOrder
	filter  func(*AuthorizationQuery) (*AuthorizationQuery, error)
}

func newAuthorizationPager(opts []AuthorizationPaginateOption, reverse bool) (*authorizationPager, error) {
	pager := &authorizationPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultAuthorizationOrder
	}
	return pager, nil
}

func (p *authorizationPager) applyFilter(query *AuthorizationQuery) (*AuthorizationQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *authorizationPager) toCursor(a *Authorization) Cursor {
	return p.order.Field.toCursor(a)
}

func (p *authorizationPager) applyCursors(query *AuthorizationQuery, after, before *Cursor) (*AuthorizationQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultAuthorizationOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *authorizationPager) applyOrder(query *AuthorizationQuery) *AuthorizationQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultAuthorizationOrder.Field {
		query = query.Order(DefaultAuthorizationOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *authorizationPager) orderExpr(query *AuthorizationQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultAuthorizationOrder.Field {
			b.Comma().Ident(DefaultAuthorizationOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Authorization.
func (a *AuthorizationQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...AuthorizationPaginateOption,
) (*AuthorizationConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newAuthorizationPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if a, err = pager.applyFilter(a); err != nil {
		return nil, err
	}
	conn := &AuthorizationConnection{Edges: []*AuthorizationEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := a.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if a, err = pager.applyCursors(a, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		a.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := a.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	a = pager.applyOrder(a)
	nodes, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// AuthorizationOrderField defines the ordering field of Authorization.
type AuthorizationOrderField struct {
	// Value extracts the ordering value from the given Authorization.
	Value    func(*Authorization) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) authorization.OrderOption
	toCursor func(*Authorization) Cursor
}

// AuthorizationOrder defines the ordering of Authorization.
type AuthorizationOrder struct {
	Direction OrderDirection           `json:"direction"`
	Field     *AuthorizationOrderField `json:"field"`
}

// DefaultAuthorizationOrder is the default ordering of Authorization.
var DefaultAuthorizationOrder = &AuthorizationOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &AuthorizationOrderField{
		Value: func(a *Authorization) (ent.Value, error) {
			return a.ID, nil
		},
		column: authorization.FieldID,
		toTerm: authorization.ByID,
		toCursor: func(a *Authorization) Cursor {
			return Cursor{ID: a.ID}
		},
	},
}

// ToEdge converts Authorization into AuthorizationEdge.
func (a *Authorization) ToEdge(order *AuthorizationOrder) *AuthorizationEdge {
	if order == nil {
		order = DefaultAuthorizationOrder
	}
	return &AuthorizationEdge{
		Node:   a,
		Cursor: order.Field.toCursor(a),
	}
}

// PersonEdge is the edge representation of Person.
type PersonEdge struct {
	Node   *Person `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// PersonConnection is the connection containing edges to Person.
type PersonConnection struct {
	Edges      []*PersonEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *PersonConnection) build(nodes []*Person, pager *personPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Person
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Person {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Person {
			return nodes[i]
		}
	}
	c.Edges = make([]*PersonEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &PersonEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// PersonPaginateOption enables pagination customization.
type PersonPaginateOption func(*personPager) error

// WithPersonOrder configures pagination ordering.
func WithPersonOrder(order *PersonOrder) PersonPaginateOption {
	if order == nil {
		order = DefaultPersonOrder
	}
	o := *order
	return func(pager *personPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultPersonOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithPersonFilter configures pagination filter.
func WithPersonFilter(filter func(*PersonQuery) (*PersonQuery, error)) PersonPaginateOption {
	return func(pager *personPager) error {
		if filter == nil {
			return errors.New("PersonQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type personPager struct {
	reverse bool
	order   *PersonOrder
	filter  func(*PersonQuery) (*PersonQuery, error)
}

func newPersonPager(opts []PersonPaginateOption, reverse bool) (*personPager, error) {
	pager := &personPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultPersonOrder
	}
	return pager, nil
}

func (p *personPager) applyFilter(query *PersonQuery) (*PersonQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *personPager) toCursor(pe *Person) Cursor {
	return p.order.Field.toCursor(pe)
}

func (p *personPager) applyCursors(query *PersonQuery, after, before *Cursor) (*PersonQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultPersonOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *personPager) applyOrder(query *PersonQuery) *PersonQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultPersonOrder.Field {
		query = query.Order(DefaultPersonOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *personPager) orderExpr(query *PersonQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultPersonOrder.Field {
			b.Comma().Ident(DefaultPersonOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Person.
func (pe *PersonQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...PersonPaginateOption,
) (*PersonConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPersonPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if pe, err = pager.applyFilter(pe); err != nil {
		return nil, err
	}
	conn := &PersonConnection{Edges: []*PersonEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := pe.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if pe, err = pager.applyCursors(pe, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		pe.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := pe.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	pe = pager.applyOrder(pe)
	nodes, err := pe.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// PersonOrderField defines the ordering field of Person.
type PersonOrderField struct {
	// Value extracts the ordering value from the given Person.
	Value    func(*Person) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) person.OrderOption
	toCursor func(*Person) Cursor
}

// PersonOrder defines the ordering of Person.
type PersonOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *PersonOrderField `json:"field"`
}

// DefaultPersonOrder is the default ordering of Person.
var DefaultPersonOrder = &PersonOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &PersonOrderField{
		Value: func(pe *Person) (ent.Value, error) {
			return pe.ID, nil
		},
		column: person.FieldID,
		toTerm: person.ByID,
		toCursor: func(pe *Person) Cursor {
			return Cursor{ID: pe.ID}
		},
	},
}

// ToEdge converts Person into PersonEdge.
func (pe *Person) ToEdge(order *PersonOrder) *PersonEdge {
	if order == nil {
		order = DefaultPersonOrder
	}
	return &PersonEdge{
		Node:   pe,
		Cursor: order.Field.toCursor(pe),
	}
}
