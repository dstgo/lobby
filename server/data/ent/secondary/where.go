// Code generated by ent, DO NOT EDIT.

package secondary

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dstgo/lobby/server/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Secondary {
	return predicate.Secondary(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Secondary {
	return predicate.Secondary(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Secondary {
	return predicate.Secondary(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Secondary {
	return predicate.Secondary(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Secondary {
	return predicate.Secondary(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Secondary {
	return predicate.Secondary(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Secondary {
	return predicate.Secondary(sql.FieldLTE(FieldID, id))
}

// Sid applies equality check predicate on the "sid" field. It's identical to SidEQ.
func Sid(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldSid, v))
}

// SteamID applies equality check predicate on the "steam_id" field. It's identical to SteamIDEQ.
func SteamID(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldSteamID, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldAddress, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldPort, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldOwnerID, v))
}

// QueryVersion applies equality check predicate on the "query_version" field. It's identical to QueryVersionEQ.
func QueryVersion(v int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldQueryVersion, v))
}

// SidEQ applies the EQ predicate on the "sid" field.
func SidEQ(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldSid, v))
}

// SidNEQ applies the NEQ predicate on the "sid" field.
func SidNEQ(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldNEQ(FieldSid, v))
}

// SidIn applies the In predicate on the "sid" field.
func SidIn(vs ...string) predicate.Secondary {
	return predicate.Secondary(sql.FieldIn(FieldSid, vs...))
}

// SidNotIn applies the NotIn predicate on the "sid" field.
func SidNotIn(vs ...string) predicate.Secondary {
	return predicate.Secondary(sql.FieldNotIn(FieldSid, vs...))
}

// SidGT applies the GT predicate on the "sid" field.
func SidGT(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldGT(FieldSid, v))
}

// SidGTE applies the GTE predicate on the "sid" field.
func SidGTE(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldGTE(FieldSid, v))
}

// SidLT applies the LT predicate on the "sid" field.
func SidLT(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldLT(FieldSid, v))
}

// SidLTE applies the LTE predicate on the "sid" field.
func SidLTE(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldLTE(FieldSid, v))
}

// SidContains applies the Contains predicate on the "sid" field.
func SidContains(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldContains(FieldSid, v))
}

// SidHasPrefix applies the HasPrefix predicate on the "sid" field.
func SidHasPrefix(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldHasPrefix(FieldSid, v))
}

// SidHasSuffix applies the HasSuffix predicate on the "sid" field.
func SidHasSuffix(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldHasSuffix(FieldSid, v))
}

// SidEqualFold applies the EqualFold predicate on the "sid" field.
func SidEqualFold(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEqualFold(FieldSid, v))
}

// SidContainsFold applies the ContainsFold predicate on the "sid" field.
func SidContainsFold(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldContainsFold(FieldSid, v))
}

// SteamIDEQ applies the EQ predicate on the "steam_id" field.
func SteamIDEQ(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldSteamID, v))
}

// SteamIDNEQ applies the NEQ predicate on the "steam_id" field.
func SteamIDNEQ(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldNEQ(FieldSteamID, v))
}

// SteamIDIn applies the In predicate on the "steam_id" field.
func SteamIDIn(vs ...string) predicate.Secondary {
	return predicate.Secondary(sql.FieldIn(FieldSteamID, vs...))
}

// SteamIDNotIn applies the NotIn predicate on the "steam_id" field.
func SteamIDNotIn(vs ...string) predicate.Secondary {
	return predicate.Secondary(sql.FieldNotIn(FieldSteamID, vs...))
}

// SteamIDGT applies the GT predicate on the "steam_id" field.
func SteamIDGT(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldGT(FieldSteamID, v))
}

// SteamIDGTE applies the GTE predicate on the "steam_id" field.
func SteamIDGTE(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldGTE(FieldSteamID, v))
}

// SteamIDLT applies the LT predicate on the "steam_id" field.
func SteamIDLT(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldLT(FieldSteamID, v))
}

// SteamIDLTE applies the LTE predicate on the "steam_id" field.
func SteamIDLTE(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldLTE(FieldSteamID, v))
}

// SteamIDContains applies the Contains predicate on the "steam_id" field.
func SteamIDContains(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldContains(FieldSteamID, v))
}

// SteamIDHasPrefix applies the HasPrefix predicate on the "steam_id" field.
func SteamIDHasPrefix(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldHasPrefix(FieldSteamID, v))
}

// SteamIDHasSuffix applies the HasSuffix predicate on the "steam_id" field.
func SteamIDHasSuffix(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldHasSuffix(FieldSteamID, v))
}

// SteamIDEqualFold applies the EqualFold predicate on the "steam_id" field.
func SteamIDEqualFold(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEqualFold(FieldSteamID, v))
}

// SteamIDContainsFold applies the ContainsFold predicate on the "steam_id" field.
func SteamIDContainsFold(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldContainsFold(FieldSteamID, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Secondary {
	return predicate.Secondary(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Secondary {
	return predicate.Secondary(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Secondary {
	return predicate.Secondary(sql.FieldContainsFold(FieldAddress, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...int) predicate.Secondary {
	return predicate.Secondary(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...int) predicate.Secondary {
	return predicate.Secondary(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldLTE(FieldPort, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v int) predicate.Secondary {
	return predicate.Secondary(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...int) predicate.Secondary {
	return predicate.Secondary(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...int) predicate.Secondary {
	return predicate.Secondary(sql.FieldNotIn(FieldOwnerID, vs...))
}

// QueryVersionEQ applies the EQ predicate on the "query_version" field.
func QueryVersionEQ(v int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldEQ(FieldQueryVersion, v))
}

// QueryVersionNEQ applies the NEQ predicate on the "query_version" field.
func QueryVersionNEQ(v int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldNEQ(FieldQueryVersion, v))
}

// QueryVersionIn applies the In predicate on the "query_version" field.
func QueryVersionIn(vs ...int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldIn(FieldQueryVersion, vs...))
}

// QueryVersionNotIn applies the NotIn predicate on the "query_version" field.
func QueryVersionNotIn(vs ...int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldNotIn(FieldQueryVersion, vs...))
}

// QueryVersionGT applies the GT predicate on the "query_version" field.
func QueryVersionGT(v int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldGT(FieldQueryVersion, v))
}

// QueryVersionGTE applies the GTE predicate on the "query_version" field.
func QueryVersionGTE(v int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldGTE(FieldQueryVersion, v))
}

// QueryVersionLT applies the LT predicate on the "query_version" field.
func QueryVersionLT(v int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldLT(FieldQueryVersion, v))
}

// QueryVersionLTE applies the LTE predicate on the "query_version" field.
func QueryVersionLTE(v int64) predicate.Secondary {
	return predicate.Secondary(sql.FieldLTE(FieldQueryVersion, v))
}

// HasServers applies the HasEdge predicate on the "servers" edge.
func HasServers() predicate.Secondary {
	return predicate.Secondary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ServersTable, ServersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServersWith applies the HasEdge predicate on the "servers" edge with a given conditions (other predicates).
func HasServersWith(preds ...predicate.Server) predicate.Secondary {
	return predicate.Secondary(func(s *sql.Selector) {
		step := newServersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Secondary) predicate.Secondary {
	return predicate.Secondary(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Secondary) predicate.Secondary {
	return predicate.Secondary(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Secondary) predicate.Secondary {
	return predicate.Secondary(sql.NotPredicates(p))
}
