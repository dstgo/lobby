// Code generated by ent, DO NOT EDIT.

package node

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dstgo/lobby/server/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldID, id))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldName, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldAddress, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldNote, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUpdatedAt, v))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldUID, vs...))
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldUID, v))
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldUID, v))
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldUID, v))
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldUID, v))
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldUID, v))
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldUID, v))
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldUID, v))
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldUID, v))
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldUID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldName, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldAddress, v))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldNote, v))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasContainers applies the HasEdge predicate on the "containers" edge.
func HasContainers() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ContainersTable, ContainersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContainersWith applies the HasEdge predicate on the "containers" edge with a given conditions (other predicates).
func HasContainersWith(preds ...predicate.Container) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newContainersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(sql.NotPredicates(p))
}
