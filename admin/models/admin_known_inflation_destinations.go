// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// AdminKnownInflationDestination is an object representing the database table.
type AdminKnownInflationDestination struct {
	ID               int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name             string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	IssuerPublicKey  string    `boil:"issuer_public_key" json:"issuer_public_key" toml:"issuer_public_key" yaml:"issuer_public_key"`
	ShortDescription string    `boil:"short_description" json:"short_description" toml:"short_description" yaml:"short_description"`
	LongDescription  string    `boil:"long_description" json:"long_description" toml:"long_description" yaml:"long_description"`
	OrderIndex       int       `boil:"order_index" json:"order_index" toml:"order_index" yaml:"order_index"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy        string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`

	R *adminKnownInflationDestinationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminKnownInflationDestinationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminKnownInflationDestinationColumns = struct {
	ID               string
	Name             string
	IssuerPublicKey  string
	ShortDescription string
	LongDescription  string
	OrderIndex       string
	CreatedAt        string
	UpdatedAt        string
	UpdatedBy        string
}{
	ID:               "id",
	Name:             "name",
	IssuerPublicKey:  "issuer_public_key",
	ShortDescription: "short_description",
	LongDescription:  "long_description",
	OrderIndex:       "order_index",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	UpdatedBy:        "updated_by",
}

// adminKnownInflationDestinationR is where relationships are stored.
type adminKnownInflationDestinationR struct {
}

// adminKnownInflationDestinationL is where Load methods for each relationship are stored.
type adminKnownInflationDestinationL struct{}

var (
	adminKnownInflationDestinationColumns               = []string{"id", "name", "issuer_public_key", "short_description", "long_description", "order_index", "created_at", "updated_at", "updated_by"}
	adminKnownInflationDestinationColumnsWithoutDefault = []string{"name", "issuer_public_key", "short_description", "long_description", "order_index", "updated_by"}
	adminKnownInflationDestinationColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	adminKnownInflationDestinationPrimaryKeyColumns     = []string{"id"}
)

type (
	// AdminKnownInflationDestinationSlice is an alias for a slice of pointers to AdminKnownInflationDestination.
	// This should generally be used opposed to []AdminKnownInflationDestination.
	AdminKnownInflationDestinationSlice []*AdminKnownInflationDestination
	// AdminKnownInflationDestinationHook is the signature for custom AdminKnownInflationDestination hook methods
	AdminKnownInflationDestinationHook func(boil.Executor, *AdminKnownInflationDestination) error

	adminKnownInflationDestinationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminKnownInflationDestinationType                 = reflect.TypeOf(&AdminKnownInflationDestination{})
	adminKnownInflationDestinationMapping              = queries.MakeStructMapping(adminKnownInflationDestinationType)
	adminKnownInflationDestinationPrimaryKeyMapping, _ = queries.BindMapping(adminKnownInflationDestinationType, adminKnownInflationDestinationMapping, adminKnownInflationDestinationPrimaryKeyColumns)
	adminKnownInflationDestinationInsertCacheMut       sync.RWMutex
	adminKnownInflationDestinationInsertCache          = make(map[string]insertCache)
	adminKnownInflationDestinationUpdateCacheMut       sync.RWMutex
	adminKnownInflationDestinationUpdateCache          = make(map[string]updateCache)
	adminKnownInflationDestinationUpsertCacheMut       sync.RWMutex
	adminKnownInflationDestinationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var adminKnownInflationDestinationBeforeInsertHooks []AdminKnownInflationDestinationHook
var adminKnownInflationDestinationBeforeUpdateHooks []AdminKnownInflationDestinationHook
var adminKnownInflationDestinationBeforeDeleteHooks []AdminKnownInflationDestinationHook
var adminKnownInflationDestinationBeforeUpsertHooks []AdminKnownInflationDestinationHook

var adminKnownInflationDestinationAfterInsertHooks []AdminKnownInflationDestinationHook
var adminKnownInflationDestinationAfterSelectHooks []AdminKnownInflationDestinationHook
var adminKnownInflationDestinationAfterUpdateHooks []AdminKnownInflationDestinationHook
var adminKnownInflationDestinationAfterDeleteHooks []AdminKnownInflationDestinationHook
var adminKnownInflationDestinationAfterUpsertHooks []AdminKnownInflationDestinationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AdminKnownInflationDestination) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AdminKnownInflationDestination) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AdminKnownInflationDestination) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AdminKnownInflationDestination) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AdminKnownInflationDestination) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AdminKnownInflationDestination) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AdminKnownInflationDestination) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AdminKnownInflationDestination) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AdminKnownInflationDestination) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminKnownInflationDestinationAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAdminKnownInflationDestinationHook registers your hook function for all future operations.
func AddAdminKnownInflationDestinationHook(hookPoint boil.HookPoint, adminKnownInflationDestinationHook AdminKnownInflationDestinationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		adminKnownInflationDestinationBeforeInsertHooks = append(adminKnownInflationDestinationBeforeInsertHooks, adminKnownInflationDestinationHook)
	case boil.BeforeUpdateHook:
		adminKnownInflationDestinationBeforeUpdateHooks = append(adminKnownInflationDestinationBeforeUpdateHooks, adminKnownInflationDestinationHook)
	case boil.BeforeDeleteHook:
		adminKnownInflationDestinationBeforeDeleteHooks = append(adminKnownInflationDestinationBeforeDeleteHooks, adminKnownInflationDestinationHook)
	case boil.BeforeUpsertHook:
		adminKnownInflationDestinationBeforeUpsertHooks = append(adminKnownInflationDestinationBeforeUpsertHooks, adminKnownInflationDestinationHook)
	case boil.AfterInsertHook:
		adminKnownInflationDestinationAfterInsertHooks = append(adminKnownInflationDestinationAfterInsertHooks, adminKnownInflationDestinationHook)
	case boil.AfterSelectHook:
		adminKnownInflationDestinationAfterSelectHooks = append(adminKnownInflationDestinationAfterSelectHooks, adminKnownInflationDestinationHook)
	case boil.AfterUpdateHook:
		adminKnownInflationDestinationAfterUpdateHooks = append(adminKnownInflationDestinationAfterUpdateHooks, adminKnownInflationDestinationHook)
	case boil.AfterDeleteHook:
		adminKnownInflationDestinationAfterDeleteHooks = append(adminKnownInflationDestinationAfterDeleteHooks, adminKnownInflationDestinationHook)
	case boil.AfterUpsertHook:
		adminKnownInflationDestinationAfterUpsertHooks = append(adminKnownInflationDestinationAfterUpsertHooks, adminKnownInflationDestinationHook)
	}
}

// OneP returns a single adminKnownInflationDestination record from the query, and panics on error.
func (q adminKnownInflationDestinationQuery) OneP() *AdminKnownInflationDestination {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single adminKnownInflationDestination record from the query.
func (q adminKnownInflationDestinationQuery) One() (*AdminKnownInflationDestination, error) {
	o := &AdminKnownInflationDestination{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for admin_known_inflation_destinations")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AdminKnownInflationDestination records from the query, and panics on error.
func (q adminKnownInflationDestinationQuery) AllP() AdminKnownInflationDestinationSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AdminKnownInflationDestination records from the query.
func (q adminKnownInflationDestinationQuery) All() (AdminKnownInflationDestinationSlice, error) {
	var o []*AdminKnownInflationDestination

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AdminKnownInflationDestination slice")
	}

	if len(adminKnownInflationDestinationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AdminKnownInflationDestination records in the query, and panics on error.
func (q adminKnownInflationDestinationQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AdminKnownInflationDestination records in the query.
func (q adminKnownInflationDestinationQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count admin_known_inflation_destinations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q adminKnownInflationDestinationQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q adminKnownInflationDestinationQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if admin_known_inflation_destinations exists")
	}

	return count > 0, nil
}

// AdminKnownInflationDestinationsG retrieves all records.
func AdminKnownInflationDestinationsG(mods ...qm.QueryMod) adminKnownInflationDestinationQuery {
	return AdminKnownInflationDestinations(boil.GetDB(), mods...)
}

// AdminKnownInflationDestinations retrieves all the records using an executor.
func AdminKnownInflationDestinations(exec boil.Executor, mods ...qm.QueryMod) adminKnownInflationDestinationQuery {
	mods = append(mods, qm.From("\"admin_known_inflation_destinations\""))
	return adminKnownInflationDestinationQuery{NewQuery(exec, mods...)}
}

// FindAdminKnownInflationDestinationG retrieves a single record by ID.
func FindAdminKnownInflationDestinationG(id int, selectCols ...string) (*AdminKnownInflationDestination, error) {
	return FindAdminKnownInflationDestination(boil.GetDB(), id, selectCols...)
}

// FindAdminKnownInflationDestinationGP retrieves a single record by ID, and panics on error.
func FindAdminKnownInflationDestinationGP(id int, selectCols ...string) *AdminKnownInflationDestination {
	retobj, err := FindAdminKnownInflationDestination(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAdminKnownInflationDestination retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdminKnownInflationDestination(exec boil.Executor, id int, selectCols ...string) (*AdminKnownInflationDestination, error) {
	adminKnownInflationDestinationObj := &AdminKnownInflationDestination{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"admin_known_inflation_destinations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(adminKnownInflationDestinationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from admin_known_inflation_destinations")
	}

	return adminKnownInflationDestinationObj, nil
}

// FindAdminKnownInflationDestinationP retrieves a single record by ID with an executor, and panics on error.
func FindAdminKnownInflationDestinationP(exec boil.Executor, id int, selectCols ...string) *AdminKnownInflationDestination {
	retobj, err := FindAdminKnownInflationDestination(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AdminKnownInflationDestination) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AdminKnownInflationDestination) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AdminKnownInflationDestination) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AdminKnownInflationDestination) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no admin_known_inflation_destinations provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminKnownInflationDestinationColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	adminKnownInflationDestinationInsertCacheMut.RLock()
	cache, cached := adminKnownInflationDestinationInsertCache[key]
	adminKnownInflationDestinationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			adminKnownInflationDestinationColumns,
			adminKnownInflationDestinationColumnsWithDefault,
			adminKnownInflationDestinationColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(adminKnownInflationDestinationType, adminKnownInflationDestinationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminKnownInflationDestinationType, adminKnownInflationDestinationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"admin_known_inflation_destinations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"admin_known_inflation_destinations\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into admin_known_inflation_destinations")
	}

	if !cached {
		adminKnownInflationDestinationInsertCacheMut.Lock()
		adminKnownInflationDestinationInsertCache[key] = cache
		adminKnownInflationDestinationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AdminKnownInflationDestination record. See Update for
// whitelist behavior description.
func (o *AdminKnownInflationDestination) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AdminKnownInflationDestination record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AdminKnownInflationDestination) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AdminKnownInflationDestination, and panics on error.
// See Update for whitelist behavior description.
func (o *AdminKnownInflationDestination) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AdminKnownInflationDestination.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AdminKnownInflationDestination) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	adminKnownInflationDestinationUpdateCacheMut.RLock()
	cache, cached := adminKnownInflationDestinationUpdateCache[key]
	adminKnownInflationDestinationUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			adminKnownInflationDestinationColumns,
			adminKnownInflationDestinationPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update admin_known_inflation_destinations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"admin_known_inflation_destinations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, adminKnownInflationDestinationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminKnownInflationDestinationType, adminKnownInflationDestinationMapping, append(wl, adminKnownInflationDestinationPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update admin_known_inflation_destinations row")
	}

	if !cached {
		adminKnownInflationDestinationUpdateCacheMut.Lock()
		adminKnownInflationDestinationUpdateCache[key] = cache
		adminKnownInflationDestinationUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q adminKnownInflationDestinationQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q adminKnownInflationDestinationQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for admin_known_inflation_destinations")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AdminKnownInflationDestinationSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AdminKnownInflationDestinationSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AdminKnownInflationDestinationSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminKnownInflationDestinationSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminKnownInflationDestinationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"admin_known_inflation_destinations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, adminKnownInflationDestinationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in adminKnownInflationDestination slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AdminKnownInflationDestination) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AdminKnownInflationDestination) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AdminKnownInflationDestination) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AdminKnownInflationDestination) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no admin_known_inflation_destinations provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminKnownInflationDestinationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()

	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	adminKnownInflationDestinationUpsertCacheMut.RLock()
	cache, cached := adminKnownInflationDestinationUpsertCache[key]
	adminKnownInflationDestinationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			adminKnownInflationDestinationColumns,
			adminKnownInflationDestinationColumnsWithDefault,
			adminKnownInflationDestinationColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			adminKnownInflationDestinationColumns,
			adminKnownInflationDestinationPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert admin_known_inflation_destinations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(adminKnownInflationDestinationPrimaryKeyColumns))
			copy(conflict, adminKnownInflationDestinationPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"admin_known_inflation_destinations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(adminKnownInflationDestinationType, adminKnownInflationDestinationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminKnownInflationDestinationType, adminKnownInflationDestinationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert admin_known_inflation_destinations")
	}

	if !cached {
		adminKnownInflationDestinationUpsertCacheMut.Lock()
		adminKnownInflationDestinationUpsertCache[key] = cache
		adminKnownInflationDestinationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AdminKnownInflationDestination record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AdminKnownInflationDestination) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AdminKnownInflationDestination record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AdminKnownInflationDestination) DeleteG() error {
	if o == nil {
		return errors.New("models: no AdminKnownInflationDestination provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AdminKnownInflationDestination record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AdminKnownInflationDestination) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AdminKnownInflationDestination record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AdminKnownInflationDestination) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AdminKnownInflationDestination provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminKnownInflationDestinationPrimaryKeyMapping)
	sql := "DELETE FROM \"admin_known_inflation_destinations\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from admin_known_inflation_destinations")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q adminKnownInflationDestinationQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q adminKnownInflationDestinationQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no adminKnownInflationDestinationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from admin_known_inflation_destinations")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AdminKnownInflationDestinationSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AdminKnownInflationDestinationSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AdminKnownInflationDestination slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AdminKnownInflationDestinationSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminKnownInflationDestinationSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AdminKnownInflationDestination slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(adminKnownInflationDestinationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminKnownInflationDestinationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"admin_known_inflation_destinations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminKnownInflationDestinationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from adminKnownInflationDestination slice")
	}

	if len(adminKnownInflationDestinationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AdminKnownInflationDestination) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AdminKnownInflationDestination) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AdminKnownInflationDestination) ReloadG() error {
	if o == nil {
		return errors.New("models: no AdminKnownInflationDestination provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AdminKnownInflationDestination) Reload(exec boil.Executor) error {
	ret, err := FindAdminKnownInflationDestination(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AdminKnownInflationDestinationSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AdminKnownInflationDestinationSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminKnownInflationDestinationSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AdminKnownInflationDestinationSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminKnownInflationDestinationSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	adminKnownInflationDestinations := AdminKnownInflationDestinationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminKnownInflationDestinationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"admin_known_inflation_destinations\".* FROM \"admin_known_inflation_destinations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminKnownInflationDestinationPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&adminKnownInflationDestinations)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AdminKnownInflationDestinationSlice")
	}

	*o = adminKnownInflationDestinations

	return nil
}

// AdminKnownInflationDestinationExists checks if the AdminKnownInflationDestination row exists.
func AdminKnownInflationDestinationExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"admin_known_inflation_destinations\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if admin_known_inflation_destinations exists")
	}

	return exists, nil
}

// AdminKnownInflationDestinationExistsG checks if the AdminKnownInflationDestination row exists.
func AdminKnownInflationDestinationExistsG(id int) (bool, error) {
	return AdminKnownInflationDestinationExists(boil.GetDB(), id)
}

// AdminKnownInflationDestinationExistsGP checks if the AdminKnownInflationDestination row exists. Panics on error.
func AdminKnownInflationDestinationExistsGP(id int) bool {
	e, err := AdminKnownInflationDestinationExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AdminKnownInflationDestinationExistsP checks if the AdminKnownInflationDestination row exists. Panics on error.
func AdminKnownInflationDestinationExistsP(exec boil.Executor, id int) bool {
	e, err := AdminKnownInflationDestinationExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}