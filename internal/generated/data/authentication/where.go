// Code generated by ent, DO NOT EDIT.

package authentication

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/avptp/brain/internal/generated/data/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldLTE(FieldID, id))
}

// PersonID applies equality check predicate on the "person_id" field. It's identical to PersonIDEQ.
func PersonID(v uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldPersonID, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v []byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldToken, v))
}

// CreatedIP applies equality check predicate on the "created_ip" field. It's identical to CreatedIPEQ.
func CreatedIP(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldCreatedIP, v))
}

// LastUsedIP applies equality check predicate on the "last_used_ip" field. It's identical to LastUsedIPEQ.
func LastUsedIP(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldLastUsedIP, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldCreatedAt, v))
}

// LastUsedAt applies equality check predicate on the "last_used_at" field. It's identical to LastUsedAtEQ.
func LastUsedAt(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldLastUsedAt, v))
}

// PersonIDEQ applies the EQ predicate on the "person_id" field.
func PersonIDEQ(v uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldPersonID, v))
}

// PersonIDNEQ applies the NEQ predicate on the "person_id" field.
func PersonIDNEQ(v uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldNEQ(FieldPersonID, v))
}

// PersonIDIn applies the In predicate on the "person_id" field.
func PersonIDIn(vs ...uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldIn(FieldPersonID, vs...))
}

// PersonIDNotIn applies the NotIn predicate on the "person_id" field.
func PersonIDNotIn(vs ...uuid.UUID) predicate.Authentication {
	return predicate.Authentication(sql.FieldNotIn(FieldPersonID, vs...))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v []byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v []byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...[]byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...[]byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v []byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v []byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v []byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v []byte) predicate.Authentication {
	return predicate.Authentication(sql.FieldLTE(FieldToken, v))
}

// CreatedIPEQ applies the EQ predicate on the "created_ip" field.
func CreatedIPEQ(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldCreatedIP, v))
}

// CreatedIPNEQ applies the NEQ predicate on the "created_ip" field.
func CreatedIPNEQ(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldNEQ(FieldCreatedIP, v))
}

// CreatedIPIn applies the In predicate on the "created_ip" field.
func CreatedIPIn(vs ...string) predicate.Authentication {
	return predicate.Authentication(sql.FieldIn(FieldCreatedIP, vs...))
}

// CreatedIPNotIn applies the NotIn predicate on the "created_ip" field.
func CreatedIPNotIn(vs ...string) predicate.Authentication {
	return predicate.Authentication(sql.FieldNotIn(FieldCreatedIP, vs...))
}

// CreatedIPGT applies the GT predicate on the "created_ip" field.
func CreatedIPGT(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldGT(FieldCreatedIP, v))
}

// CreatedIPGTE applies the GTE predicate on the "created_ip" field.
func CreatedIPGTE(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldGTE(FieldCreatedIP, v))
}

// CreatedIPLT applies the LT predicate on the "created_ip" field.
func CreatedIPLT(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldLT(FieldCreatedIP, v))
}

// CreatedIPLTE applies the LTE predicate on the "created_ip" field.
func CreatedIPLTE(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldLTE(FieldCreatedIP, v))
}

// CreatedIPContains applies the Contains predicate on the "created_ip" field.
func CreatedIPContains(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldContains(FieldCreatedIP, v))
}

// CreatedIPHasPrefix applies the HasPrefix predicate on the "created_ip" field.
func CreatedIPHasPrefix(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldHasPrefix(FieldCreatedIP, v))
}

// CreatedIPHasSuffix applies the HasSuffix predicate on the "created_ip" field.
func CreatedIPHasSuffix(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldHasSuffix(FieldCreatedIP, v))
}

// CreatedIPEqualFold applies the EqualFold predicate on the "created_ip" field.
func CreatedIPEqualFold(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldEqualFold(FieldCreatedIP, v))
}

// CreatedIPContainsFold applies the ContainsFold predicate on the "created_ip" field.
func CreatedIPContainsFold(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldContainsFold(FieldCreatedIP, v))
}

// LastUsedIPEQ applies the EQ predicate on the "last_used_ip" field.
func LastUsedIPEQ(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldLastUsedIP, v))
}

// LastUsedIPNEQ applies the NEQ predicate on the "last_used_ip" field.
func LastUsedIPNEQ(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldNEQ(FieldLastUsedIP, v))
}

// LastUsedIPIn applies the In predicate on the "last_used_ip" field.
func LastUsedIPIn(vs ...string) predicate.Authentication {
	return predicate.Authentication(sql.FieldIn(FieldLastUsedIP, vs...))
}

// LastUsedIPNotIn applies the NotIn predicate on the "last_used_ip" field.
func LastUsedIPNotIn(vs ...string) predicate.Authentication {
	return predicate.Authentication(sql.FieldNotIn(FieldLastUsedIP, vs...))
}

// LastUsedIPGT applies the GT predicate on the "last_used_ip" field.
func LastUsedIPGT(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldGT(FieldLastUsedIP, v))
}

// LastUsedIPGTE applies the GTE predicate on the "last_used_ip" field.
func LastUsedIPGTE(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldGTE(FieldLastUsedIP, v))
}

// LastUsedIPLT applies the LT predicate on the "last_used_ip" field.
func LastUsedIPLT(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldLT(FieldLastUsedIP, v))
}

// LastUsedIPLTE applies the LTE predicate on the "last_used_ip" field.
func LastUsedIPLTE(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldLTE(FieldLastUsedIP, v))
}

// LastUsedIPContains applies the Contains predicate on the "last_used_ip" field.
func LastUsedIPContains(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldContains(FieldLastUsedIP, v))
}

// LastUsedIPHasPrefix applies the HasPrefix predicate on the "last_used_ip" field.
func LastUsedIPHasPrefix(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldHasPrefix(FieldLastUsedIP, v))
}

// LastUsedIPHasSuffix applies the HasSuffix predicate on the "last_used_ip" field.
func LastUsedIPHasSuffix(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldHasSuffix(FieldLastUsedIP, v))
}

// LastUsedIPEqualFold applies the EqualFold predicate on the "last_used_ip" field.
func LastUsedIPEqualFold(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldEqualFold(FieldLastUsedIP, v))
}

// LastUsedIPContainsFold applies the ContainsFold predicate on the "last_used_ip" field.
func LastUsedIPContainsFold(v string) predicate.Authentication {
	return predicate.Authentication(sql.FieldContainsFold(FieldLastUsedIP, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldLTE(FieldCreatedAt, v))
}

// LastUsedAtEQ applies the EQ predicate on the "last_used_at" field.
func LastUsedAtEQ(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldEQ(FieldLastUsedAt, v))
}

// LastUsedAtNEQ applies the NEQ predicate on the "last_used_at" field.
func LastUsedAtNEQ(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldNEQ(FieldLastUsedAt, v))
}

// LastUsedAtIn applies the In predicate on the "last_used_at" field.
func LastUsedAtIn(vs ...time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldIn(FieldLastUsedAt, vs...))
}

// LastUsedAtNotIn applies the NotIn predicate on the "last_used_at" field.
func LastUsedAtNotIn(vs ...time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldNotIn(FieldLastUsedAt, vs...))
}

// LastUsedAtGT applies the GT predicate on the "last_used_at" field.
func LastUsedAtGT(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldGT(FieldLastUsedAt, v))
}

// LastUsedAtGTE applies the GTE predicate on the "last_used_at" field.
func LastUsedAtGTE(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldGTE(FieldLastUsedAt, v))
}

// LastUsedAtLT applies the LT predicate on the "last_used_at" field.
func LastUsedAtLT(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldLT(FieldLastUsedAt, v))
}

// LastUsedAtLTE applies the LTE predicate on the "last_used_at" field.
func LastUsedAtLTE(v time.Time) predicate.Authentication {
	return predicate.Authentication(sql.FieldLTE(FieldLastUsedAt, v))
}

// HasPerson applies the HasEdge predicate on the "person" edge.
func HasPerson() predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonTable, PersonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonWith applies the HasEdge predicate on the "person" edge with a given conditions (other predicates).
func HasPersonWith(preds ...predicate.Person) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		step := newPersonStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Authentication) predicate.Authentication {
	return predicate.Authentication(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Authentication) predicate.Authentication {
	return predicate.Authentication(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Authentication) predicate.Authentication {
	return predicate.Authentication(sql.NotPredicates(p))
}
