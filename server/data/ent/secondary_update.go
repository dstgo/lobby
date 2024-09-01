// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dstgo/lobby/server/data/ent/predicate"
	"github.com/dstgo/lobby/server/data/ent/secondary"
	"github.com/dstgo/lobby/server/data/ent/server"
)

// SecondaryUpdate is the builder for updating Secondary entities.
type SecondaryUpdate struct {
	config
	hooks     []Hook
	mutation  *SecondaryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SecondaryUpdate builder.
func (su *SecondaryUpdate) Where(ps ...predicate.Secondary) *SecondaryUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSid sets the "sid" field.
func (su *SecondaryUpdate) SetSid(s string) *SecondaryUpdate {
	su.mutation.SetSid(s)
	return su
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (su *SecondaryUpdate) SetNillableSid(s *string) *SecondaryUpdate {
	if s != nil {
		su.SetSid(*s)
	}
	return su
}

// SetSteamID sets the "steam_id" field.
func (su *SecondaryUpdate) SetSteamID(s string) *SecondaryUpdate {
	su.mutation.SetSteamID(s)
	return su
}

// SetNillableSteamID sets the "steam_id" field if the given value is not nil.
func (su *SecondaryUpdate) SetNillableSteamID(s *string) *SecondaryUpdate {
	if s != nil {
		su.SetSteamID(*s)
	}
	return su
}

// SetAddress sets the "address" field.
func (su *SecondaryUpdate) SetAddress(s string) *SecondaryUpdate {
	su.mutation.SetAddress(s)
	return su
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (su *SecondaryUpdate) SetNillableAddress(s *string) *SecondaryUpdate {
	if s != nil {
		su.SetAddress(*s)
	}
	return su
}

// SetPort sets the "port" field.
func (su *SecondaryUpdate) SetPort(i int) *SecondaryUpdate {
	su.mutation.ResetPort()
	su.mutation.SetPort(i)
	return su
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (su *SecondaryUpdate) SetNillablePort(i *int) *SecondaryUpdate {
	if i != nil {
		su.SetPort(*i)
	}
	return su
}

// AddPort adds i to the "port" field.
func (su *SecondaryUpdate) AddPort(i int) *SecondaryUpdate {
	su.mutation.AddPort(i)
	return su
}

// SetOwnerID sets the "owner_id" field.
func (su *SecondaryUpdate) SetOwnerID(i int) *SecondaryUpdate {
	su.mutation.SetOwnerID(i)
	return su
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (su *SecondaryUpdate) SetNillableOwnerID(i *int) *SecondaryUpdate {
	if i != nil {
		su.SetOwnerID(*i)
	}
	return su
}

// SetServersID sets the "servers" edge to the Server entity by ID.
func (su *SecondaryUpdate) SetServersID(id int) *SecondaryUpdate {
	su.mutation.SetServersID(id)
	return su
}

// SetServers sets the "servers" edge to the Server entity.
func (su *SecondaryUpdate) SetServers(s *Server) *SecondaryUpdate {
	return su.SetServersID(s.ID)
}

// Mutation returns the SecondaryMutation object of the builder.
func (su *SecondaryUpdate) Mutation() *SecondaryMutation {
	return su.mutation
}

// ClearServers clears the "servers" edge to the Server entity.
func (su *SecondaryUpdate) ClearServers() *SecondaryUpdate {
	su.mutation.ClearServers()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SecondaryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SecondaryUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SecondaryUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SecondaryUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SecondaryUpdate) check() error {
	if su.mutation.ServersCleared() && len(su.mutation.ServersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Secondary.servers"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SecondaryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SecondaryUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SecondaryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(secondary.Table, secondary.Columns, sqlgraph.NewFieldSpec(secondary.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Sid(); ok {
		_spec.SetField(secondary.FieldSid, field.TypeString, value)
	}
	if value, ok := su.mutation.SteamID(); ok {
		_spec.SetField(secondary.FieldSteamID, field.TypeString, value)
	}
	if value, ok := su.mutation.Address(); ok {
		_spec.SetField(secondary.FieldAddress, field.TypeString, value)
	}
	if value, ok := su.mutation.Port(); ok {
		_spec.SetField(secondary.FieldPort, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedPort(); ok {
		_spec.AddField(secondary.FieldPort, field.TypeInt, value)
	}
	if su.mutation.ServersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   secondary.ServersTable,
			Columns: []string{secondary.ServersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ServersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   secondary.ServersTable,
			Columns: []string{secondary.ServersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{secondary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SecondaryUpdateOne is the builder for updating a single Secondary entity.
type SecondaryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SecondaryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetSid sets the "sid" field.
func (suo *SecondaryUpdateOne) SetSid(s string) *SecondaryUpdateOne {
	suo.mutation.SetSid(s)
	return suo
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (suo *SecondaryUpdateOne) SetNillableSid(s *string) *SecondaryUpdateOne {
	if s != nil {
		suo.SetSid(*s)
	}
	return suo
}

// SetSteamID sets the "steam_id" field.
func (suo *SecondaryUpdateOne) SetSteamID(s string) *SecondaryUpdateOne {
	suo.mutation.SetSteamID(s)
	return suo
}

// SetNillableSteamID sets the "steam_id" field if the given value is not nil.
func (suo *SecondaryUpdateOne) SetNillableSteamID(s *string) *SecondaryUpdateOne {
	if s != nil {
		suo.SetSteamID(*s)
	}
	return suo
}

// SetAddress sets the "address" field.
func (suo *SecondaryUpdateOne) SetAddress(s string) *SecondaryUpdateOne {
	suo.mutation.SetAddress(s)
	return suo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (suo *SecondaryUpdateOne) SetNillableAddress(s *string) *SecondaryUpdateOne {
	if s != nil {
		suo.SetAddress(*s)
	}
	return suo
}

// SetPort sets the "port" field.
func (suo *SecondaryUpdateOne) SetPort(i int) *SecondaryUpdateOne {
	suo.mutation.ResetPort()
	suo.mutation.SetPort(i)
	return suo
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (suo *SecondaryUpdateOne) SetNillablePort(i *int) *SecondaryUpdateOne {
	if i != nil {
		suo.SetPort(*i)
	}
	return suo
}

// AddPort adds i to the "port" field.
func (suo *SecondaryUpdateOne) AddPort(i int) *SecondaryUpdateOne {
	suo.mutation.AddPort(i)
	return suo
}

// SetOwnerID sets the "owner_id" field.
func (suo *SecondaryUpdateOne) SetOwnerID(i int) *SecondaryUpdateOne {
	suo.mutation.SetOwnerID(i)
	return suo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (suo *SecondaryUpdateOne) SetNillableOwnerID(i *int) *SecondaryUpdateOne {
	if i != nil {
		suo.SetOwnerID(*i)
	}
	return suo
}

// SetServersID sets the "servers" edge to the Server entity by ID.
func (suo *SecondaryUpdateOne) SetServersID(id int) *SecondaryUpdateOne {
	suo.mutation.SetServersID(id)
	return suo
}

// SetServers sets the "servers" edge to the Server entity.
func (suo *SecondaryUpdateOne) SetServers(s *Server) *SecondaryUpdateOne {
	return suo.SetServersID(s.ID)
}

// Mutation returns the SecondaryMutation object of the builder.
func (suo *SecondaryUpdateOne) Mutation() *SecondaryMutation {
	return suo.mutation
}

// ClearServers clears the "servers" edge to the Server entity.
func (suo *SecondaryUpdateOne) ClearServers() *SecondaryUpdateOne {
	suo.mutation.ClearServers()
	return suo
}

// Where appends a list predicates to the SecondaryUpdate builder.
func (suo *SecondaryUpdateOne) Where(ps ...predicate.Secondary) *SecondaryUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SecondaryUpdateOne) Select(field string, fields ...string) *SecondaryUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Secondary entity.
func (suo *SecondaryUpdateOne) Save(ctx context.Context) (*Secondary, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SecondaryUpdateOne) SaveX(ctx context.Context) *Secondary {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SecondaryUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SecondaryUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SecondaryUpdateOne) check() error {
	if suo.mutation.ServersCleared() && len(suo.mutation.ServersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Secondary.servers"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SecondaryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SecondaryUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SecondaryUpdateOne) sqlSave(ctx context.Context) (_node *Secondary, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(secondary.Table, secondary.Columns, sqlgraph.NewFieldSpec(secondary.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Secondary.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, secondary.FieldID)
		for _, f := range fields {
			if !secondary.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != secondary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Sid(); ok {
		_spec.SetField(secondary.FieldSid, field.TypeString, value)
	}
	if value, ok := suo.mutation.SteamID(); ok {
		_spec.SetField(secondary.FieldSteamID, field.TypeString, value)
	}
	if value, ok := suo.mutation.Address(); ok {
		_spec.SetField(secondary.FieldAddress, field.TypeString, value)
	}
	if value, ok := suo.mutation.Port(); ok {
		_spec.SetField(secondary.FieldPort, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedPort(); ok {
		_spec.AddField(secondary.FieldPort, field.TypeInt, value)
	}
	if suo.mutation.ServersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   secondary.ServersTable,
			Columns: []string{secondary.ServersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ServersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   secondary.ServersTable,
			Columns: []string{secondary.ServersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Secondary{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{secondary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
