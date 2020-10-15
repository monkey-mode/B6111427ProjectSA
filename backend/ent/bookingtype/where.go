// Code generated by entc, DO NOT EDIT.

package bookingtype

import (
	"github.com/B6111427/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BOOKTYPENAME applies equality check predicate on the "BOOKTYPE_NAME" field. It's identical to BOOKTYPENAMEEQ.
func BOOKTYPENAME(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEEQ applies the EQ predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEEQ(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMENEQ applies the NEQ predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMENEQ(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEIn applies the In predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEIn(vs ...string) predicate.Bookingtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookingtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBOOKTYPENAME), v...))
	})
}

// BOOKTYPENAMENotIn applies the NotIn predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMENotIn(vs ...string) predicate.Bookingtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookingtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBOOKTYPENAME), v...))
	})
}

// BOOKTYPENAMEGT applies the GT predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEGT(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEGTE applies the GTE predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEGTE(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMELT applies the LT predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMELT(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMELTE applies the LTE predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMELTE(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEContains applies the Contains predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEContains(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEHasPrefix applies the HasPrefix predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEHasPrefix(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEHasSuffix applies the HasSuffix predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEHasSuffix(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEEqualFold applies the EqualFold predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEEqualFold(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBOOKTYPENAME), v))
	})
}

// BOOKTYPENAMEContainsFold applies the ContainsFold predicate on the "BOOKTYPE_NAME" field.
func BOOKTYPENAMEContainsFold(v string) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBOOKTYPENAME), v))
	})
}

// HasBooktype applies the HasEdge predicate on the "booktype" edge.
func HasBooktype() predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BooktypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BooktypeTable, BooktypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBooktypeWith applies the HasEdge predicate on the "booktype" edge with a given conditions (other predicates).
func HasBooktypeWith(preds ...predicate.Booking) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BooktypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BooktypeTable, BooktypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Bookingtype) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Bookingtype) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bookingtype) predicate.Bookingtype {
	return predicate.Bookingtype(func(s *sql.Selector) {
		p(s.Not())
	})
}