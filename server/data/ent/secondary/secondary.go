// Code generated by ent, DO NOT EDIT.

package secondary

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the secondary type in the database.
	Label = "secondary"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSid holds the string denoting the sid field in the database.
	FieldSid = "sid"
	// FieldSteamID holds the string denoting the steam_id field in the database.
	FieldSteamID = "steam_id"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldQueryVersion holds the string denoting the query_version field in the database.
	FieldQueryVersion = "query_version"
	// EdgeServers holds the string denoting the servers edge name in mutations.
	EdgeServers = "servers"
	// Table holds the table name of the secondary in the database.
	Table = "secondaries"
	// ServersTable is the table that holds the servers relation/edge.
	ServersTable = "secondaries"
	// ServersInverseTable is the table name for the Server entity.
	// It exists in this package in order to avoid circular dependency with the "server" package.
	ServersInverseTable = "servers"
	// ServersColumn is the table column denoting the servers relation/edge.
	ServersColumn = "owner_id"
)

// Columns holds all SQL columns for secondary fields.
var Columns = []string{
	FieldID,
	FieldSid,
	FieldSteamID,
	FieldAddress,
	FieldPort,
	FieldOwnerID,
	FieldQueryVersion,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Secondary queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySid orders the results by the sid field.
func BySid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSid, opts...).ToFunc()
}

// BySteamID orders the results by the steam_id field.
func BySteamID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSteamID, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByQueryVersion orders the results by the query_version field.
func ByQueryVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQueryVersion, opts...).ToFunc()
}

// ByServersField orders the results by servers field.
func ByServersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServersStep(), sql.OrderByField(field, opts...))
	}
}
func newServersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServersTable, ServersColumn),
	)
}
