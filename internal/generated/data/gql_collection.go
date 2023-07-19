// Code generated by ent, DO NOT EDIT.

package data

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/avptp/brain/internal/generated/data/authentication"
	"github.com/avptp/brain/internal/generated/data/person"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AuthenticationQuery) CollectFields(ctx context.Context, satisfies ...string) (*AuthenticationQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AuthenticationQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(authentication.Columns))
		selectedFields = []string{authentication.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "person":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PersonClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			a.withPerson = query
			if _, ok := fieldSeen[authentication.FieldPersonID]; !ok {
				selectedFields = append(selectedFields, authentication.FieldPersonID)
				fieldSeen[authentication.FieldPersonID] = struct{}{}
			}
		case "personID":
			if _, ok := fieldSeen[authentication.FieldPersonID]; !ok {
				selectedFields = append(selectedFields, authentication.FieldPersonID)
				fieldSeen[authentication.FieldPersonID] = struct{}{}
			}
		case "token":
			if _, ok := fieldSeen[authentication.FieldToken]; !ok {
				selectedFields = append(selectedFields, authentication.FieldToken)
				fieldSeen[authentication.FieldToken] = struct{}{}
			}
		case "createdIP":
			if _, ok := fieldSeen[authentication.FieldCreatedIP]; !ok {
				selectedFields = append(selectedFields, authentication.FieldCreatedIP)
				fieldSeen[authentication.FieldCreatedIP] = struct{}{}
			}
		case "lastUsedIP":
			if _, ok := fieldSeen[authentication.FieldLastUsedIP]; !ok {
				selectedFields = append(selectedFields, authentication.FieldLastUsedIP)
				fieldSeen[authentication.FieldLastUsedIP] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[authentication.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, authentication.FieldCreatedAt)
				fieldSeen[authentication.FieldCreatedAt] = struct{}{}
			}
		case "lastUsedAt":
			if _, ok := fieldSeen[authentication.FieldLastUsedAt]; !ok {
				selectedFields = append(selectedFields, authentication.FieldLastUsedAt)
				fieldSeen[authentication.FieldLastUsedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type authenticationPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AuthenticationPaginateOption
}

func newAuthenticationPaginateArgs(rv map[string]any) *authenticationPaginateArgs {
	args := &authenticationPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AuthenticationWhereInput); ok {
		args.opts = append(args.opts, WithAuthenticationFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PersonQuery) CollectFields(ctx context.Context, satisfies ...string) (*PersonQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pe, nil
	}
	if err := pe.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pe *PersonQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(person.Columns))
		selectedFields = []string{person.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "authentications":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AuthenticationClient{config: pe.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			pe.WithNamedAuthentications(alias, func(wq *AuthenticationQuery) {
				*wq = *query
			})
		case "email":
			if _, ok := fieldSeen[person.FieldEmail]; !ok {
				selectedFields = append(selectedFields, person.FieldEmail)
				fieldSeen[person.FieldEmail] = struct{}{}
			}
		case "emailVerifiedAt":
			if _, ok := fieldSeen[person.FieldEmailVerifiedAt]; !ok {
				selectedFields = append(selectedFields, person.FieldEmailVerifiedAt)
				fieldSeen[person.FieldEmailVerifiedAt] = struct{}{}
			}
		case "phone":
			if _, ok := fieldSeen[person.FieldPhone]; !ok {
				selectedFields = append(selectedFields, person.FieldPhone)
				fieldSeen[person.FieldPhone] = struct{}{}
			}
		case "taxID":
			if _, ok := fieldSeen[person.FieldTaxID]; !ok {
				selectedFields = append(selectedFields, person.FieldTaxID)
				fieldSeen[person.FieldTaxID] = struct{}{}
			}
		case "firstName":
			if _, ok := fieldSeen[person.FieldFirstName]; !ok {
				selectedFields = append(selectedFields, person.FieldFirstName)
				fieldSeen[person.FieldFirstName] = struct{}{}
			}
		case "lastName":
			if _, ok := fieldSeen[person.FieldLastName]; !ok {
				selectedFields = append(selectedFields, person.FieldLastName)
				fieldSeen[person.FieldLastName] = struct{}{}
			}
		case "language":
			if _, ok := fieldSeen[person.FieldLanguage]; !ok {
				selectedFields = append(selectedFields, person.FieldLanguage)
				fieldSeen[person.FieldLanguage] = struct{}{}
			}
		case "birthdate":
			if _, ok := fieldSeen[person.FieldBirthdate]; !ok {
				selectedFields = append(selectedFields, person.FieldBirthdate)
				fieldSeen[person.FieldBirthdate] = struct{}{}
			}
		case "gender":
			if _, ok := fieldSeen[person.FieldGender]; !ok {
				selectedFields = append(selectedFields, person.FieldGender)
				fieldSeen[person.FieldGender] = struct{}{}
			}
		case "address":
			if _, ok := fieldSeen[person.FieldAddress]; !ok {
				selectedFields = append(selectedFields, person.FieldAddress)
				fieldSeen[person.FieldAddress] = struct{}{}
			}
		case "postalCode":
			if _, ok := fieldSeen[person.FieldPostalCode]; !ok {
				selectedFields = append(selectedFields, person.FieldPostalCode)
				fieldSeen[person.FieldPostalCode] = struct{}{}
			}
		case "city":
			if _, ok := fieldSeen[person.FieldCity]; !ok {
				selectedFields = append(selectedFields, person.FieldCity)
				fieldSeen[person.FieldCity] = struct{}{}
			}
		case "country":
			if _, ok := fieldSeen[person.FieldCountry]; !ok {
				selectedFields = append(selectedFields, person.FieldCountry)
				fieldSeen[person.FieldCountry] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[person.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, person.FieldCreatedAt)
				fieldSeen[person.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[person.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, person.FieldUpdatedAt)
				fieldSeen[person.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pe.Select(selectedFields...)
	}
	return nil
}

type personPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PersonPaginateOption
}

func newPersonPaginateArgs(rv map[string]any) *personPaginateArgs {
	args := &personPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PersonWhereInput); ok {
		args.opts = append(args.opts, WithPersonFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
