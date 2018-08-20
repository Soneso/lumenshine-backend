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

// JWTKey is an object representing the database table.
type JWTKey struct {
	ID             int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	KeyName        string    `boil:"key_name" json:"key_name" toml:"key_name" yaml:"key_name"`
	KeyValue1      string    `boil:"key_value1" json:"key_value1" toml:"key_value1" yaml:"key_value1"`
	KeyValue2      string    `boil:"key_value2" json:"key_value2" toml:"key_value2" yaml:"key_value2"`
	KeyDescription string    `boil:"key_description" json:"key_description" toml:"key_description" yaml:"key_description"`
	Valid1To       time.Time `boil:"valid1_to" json:"valid1_to" toml:"valid1_to" yaml:"valid1_to"`
	Valid2To       time.Time `boil:"valid2_to" json:"valid2_to" toml:"valid2_to" yaml:"valid2_to"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *jwtKeyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jwtKeyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JWTKeyColumns = struct {
	ID             string
	KeyName        string
	KeyValue1      string
	KeyValue2      string
	KeyDescription string
	Valid1To       string
	Valid2To       string
	UpdatedAt      string
}{
	ID:             "id",
	KeyName:        "key_name",
	KeyValue1:      "key_value1",
	KeyValue2:      "key_value2",
	KeyDescription: "key_description",
	Valid1To:       "valid1_to",
	Valid2To:       "valid2_to",
	UpdatedAt:      "updated_at",
}

// jwtKeyR is where relationships are stored.
type jwtKeyR struct {
}

// jwtKeyL is where Load methods for each relationship are stored.
type jwtKeyL struct{}

var (
	jwtKeyColumns               = []string{"id", "key_name", "key_value1", "key_value2", "key_description", "valid1_to", "valid2_to", "updated_at"}
	jwtKeyColumnsWithoutDefault = []string{"key_name", "key_value1", "key_value2", "key_description", "valid1_to", "valid2_to"}
	jwtKeyColumnsWithDefault    = []string{"id", "updated_at"}
	jwtKeyPrimaryKeyColumns     = []string{"id"}
)

type (
	// JWTKeySlice is an alias for a slice of pointers to JWTKey.
	// This should generally be used opposed to []JWTKey.
	JWTKeySlice []*JWTKey
	// JWTKeyHook is the signature for custom JWTKey hook methods
	JWTKeyHook func(boil.Executor, *JWTKey) error

	jwtKeyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jwtKeyType                 = reflect.TypeOf(&JWTKey{})
	jwtKeyMapping              = queries.MakeStructMapping(jwtKeyType)
	jwtKeyPrimaryKeyMapping, _ = queries.BindMapping(jwtKeyType, jwtKeyMapping, jwtKeyPrimaryKeyColumns)
	jwtKeyInsertCacheMut       sync.RWMutex
	jwtKeyInsertCache          = make(map[string]insertCache)
	jwtKeyUpdateCacheMut       sync.RWMutex
	jwtKeyUpdateCache          = make(map[string]updateCache)
	jwtKeyUpsertCacheMut       sync.RWMutex
	jwtKeyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var jwtKeyBeforeInsertHooks []JWTKeyHook
var jwtKeyBeforeUpdateHooks []JWTKeyHook
var jwtKeyBeforeDeleteHooks []JWTKeyHook
var jwtKeyBeforeUpsertHooks []JWTKeyHook

var jwtKeyAfterInsertHooks []JWTKeyHook
var jwtKeyAfterSelectHooks []JWTKeyHook
var jwtKeyAfterUpdateHooks []JWTKeyHook
var jwtKeyAfterDeleteHooks []JWTKeyHook
var jwtKeyAfterUpsertHooks []JWTKeyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *JWTKey) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *JWTKey) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *JWTKey) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *JWTKey) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *JWTKey) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *JWTKey) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *JWTKey) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *JWTKey) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *JWTKey) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jwtKeyAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddJWTKeyHook registers your hook function for all future operations.
func AddJWTKeyHook(hookPoint boil.HookPoint, jwtKeyHook JWTKeyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		jwtKeyBeforeInsertHooks = append(jwtKeyBeforeInsertHooks, jwtKeyHook)
	case boil.BeforeUpdateHook:
		jwtKeyBeforeUpdateHooks = append(jwtKeyBeforeUpdateHooks, jwtKeyHook)
	case boil.BeforeDeleteHook:
		jwtKeyBeforeDeleteHooks = append(jwtKeyBeforeDeleteHooks, jwtKeyHook)
	case boil.BeforeUpsertHook:
		jwtKeyBeforeUpsertHooks = append(jwtKeyBeforeUpsertHooks, jwtKeyHook)
	case boil.AfterInsertHook:
		jwtKeyAfterInsertHooks = append(jwtKeyAfterInsertHooks, jwtKeyHook)
	case boil.AfterSelectHook:
		jwtKeyAfterSelectHooks = append(jwtKeyAfterSelectHooks, jwtKeyHook)
	case boil.AfterUpdateHook:
		jwtKeyAfterUpdateHooks = append(jwtKeyAfterUpdateHooks, jwtKeyHook)
	case boil.AfterDeleteHook:
		jwtKeyAfterDeleteHooks = append(jwtKeyAfterDeleteHooks, jwtKeyHook)
	case boil.AfterUpsertHook:
		jwtKeyAfterUpsertHooks = append(jwtKeyAfterUpsertHooks, jwtKeyHook)
	}
}

// OneP returns a single jwtKey record from the query, and panics on error.
func (q jwtKeyQuery) OneP() *JWTKey {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single jwtKey record from the query.
func (q jwtKeyQuery) One() (*JWTKey, error) {
	o := &JWTKey{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for jwt_key")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all JWTKey records from the query, and panics on error.
func (q jwtKeyQuery) AllP() JWTKeySlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all JWTKey records from the query.
func (q jwtKeyQuery) All() (JWTKeySlice, error) {
	var o []*JWTKey

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to JWTKey slice")
	}

	if len(jwtKeyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all JWTKey records in the query, and panics on error.
func (q jwtKeyQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all JWTKey records in the query.
func (q jwtKeyQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count jwt_key rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q jwtKeyQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q jwtKeyQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if jwt_key exists")
	}

	return count > 0, nil
}

// JWTKeysG retrieves all records.
func JWTKeysG(mods ...qm.QueryMod) jwtKeyQuery {
	return JWTKeys(boil.GetDB(), mods...)
}

// JWTKeys retrieves all the records using an executor.
func JWTKeys(exec boil.Executor, mods ...qm.QueryMod) jwtKeyQuery {
	mods = append(mods, qm.From("\"jwt_key\""))
	return jwtKeyQuery{NewQuery(exec, mods...)}
}

// FindJWTKeyG retrieves a single record by ID.
func FindJWTKeyG(id int, selectCols ...string) (*JWTKey, error) {
	return FindJWTKey(boil.GetDB(), id, selectCols...)
}

// FindJWTKeyGP retrieves a single record by ID, and panics on error.
func FindJWTKeyGP(id int, selectCols ...string) *JWTKey {
	retobj, err := FindJWTKey(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindJWTKey retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJWTKey(exec boil.Executor, id int, selectCols ...string) (*JWTKey, error) {
	jwtKeyObj := &JWTKey{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"jwt_key\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(jwtKeyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from jwt_key")
	}

	return jwtKeyObj, nil
}

// FindJWTKeyP retrieves a single record by ID with an executor, and panics on error.
func FindJWTKeyP(exec boil.Executor, id int, selectCols ...string) *JWTKey {
	retobj, err := FindJWTKey(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *JWTKey) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *JWTKey) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *JWTKey) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *JWTKey) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no jwt_key provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jwtKeyColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	jwtKeyInsertCacheMut.RLock()
	cache, cached := jwtKeyInsertCache[key]
	jwtKeyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			jwtKeyColumns,
			jwtKeyColumnsWithDefault,
			jwtKeyColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(jwtKeyType, jwtKeyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jwtKeyType, jwtKeyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"jwt_key\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"jwt_key\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into jwt_key")
	}

	if !cached {
		jwtKeyInsertCacheMut.Lock()
		jwtKeyInsertCache[key] = cache
		jwtKeyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single JWTKey record. See Update for
// whitelist behavior description.
func (o *JWTKey) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single JWTKey record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *JWTKey) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the JWTKey, and panics on error.
// See Update for whitelist behavior description.
func (o *JWTKey) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the JWTKey.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *JWTKey) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	jwtKeyUpdateCacheMut.RLock()
	cache, cached := jwtKeyUpdateCache[key]
	jwtKeyUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			jwtKeyColumns,
			jwtKeyPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update jwt_key, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"jwt_key\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jwtKeyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jwtKeyType, jwtKeyMapping, append(wl, jwtKeyPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update jwt_key row")
	}

	if !cached {
		jwtKeyUpdateCacheMut.Lock()
		jwtKeyUpdateCache[key] = cache
		jwtKeyUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q jwtKeyQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q jwtKeyQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for jwt_key")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o JWTKeySlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o JWTKeySlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o JWTKeySlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JWTKeySlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jwtKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"jwt_key\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, jwtKeyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in jwtKey slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *JWTKey) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *JWTKey) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *JWTKey) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *JWTKey) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no jwt_key provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jwtKeyColumnsWithDefault, o)

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

	jwtKeyUpsertCacheMut.RLock()
	cache, cached := jwtKeyUpsertCache[key]
	jwtKeyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			jwtKeyColumns,
			jwtKeyColumnsWithDefault,
			jwtKeyColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			jwtKeyColumns,
			jwtKeyPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert jwt_key, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jwtKeyPrimaryKeyColumns))
			copy(conflict, jwtKeyPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"jwt_key\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(jwtKeyType, jwtKeyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jwtKeyType, jwtKeyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert jwt_key")
	}

	if !cached {
		jwtKeyUpsertCacheMut.Lock()
		jwtKeyUpsertCache[key] = cache
		jwtKeyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single JWTKey record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JWTKey) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single JWTKey record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *JWTKey) DeleteG() error {
	if o == nil {
		return errors.New("models: no JWTKey provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single JWTKey record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JWTKey) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single JWTKey record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *JWTKey) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no JWTKey provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jwtKeyPrimaryKeyMapping)
	sql := "DELETE FROM \"jwt_key\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from jwt_key")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q jwtKeyQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q jwtKeyQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no jwtKeyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from jwt_key")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o JWTKeySlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o JWTKeySlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no JWTKey slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o JWTKeySlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JWTKeySlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no JWTKey slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(jwtKeyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jwtKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"jwt_key\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jwtKeyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from jwtKey slice")
	}

	if len(jwtKeyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *JWTKey) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *JWTKey) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *JWTKey) ReloadG() error {
	if o == nil {
		return errors.New("models: no JWTKey provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *JWTKey) Reload(exec boil.Executor) error {
	ret, err := FindJWTKey(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JWTKeySlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JWTKeySlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JWTKeySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty JWTKeySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JWTKeySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	jwtKeys := JWTKeySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jwtKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"jwt_key\".* FROM \"jwt_key\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jwtKeyPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&jwtKeys)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JWTKeySlice")
	}

	*o = jwtKeys

	return nil
}

// JWTKeyExists checks if the JWTKey row exists.
func JWTKeyExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"jwt_key\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if jwt_key exists")
	}

	return exists, nil
}

// JWTKeyExistsG checks if the JWTKey row exists.
func JWTKeyExistsG(id int) (bool, error) {
	return JWTKeyExists(boil.GetDB(), id)
}

// JWTKeyExistsGP checks if the JWTKey row exists. Panics on error.
func JWTKeyExistsGP(id int) bool {
	e, err := JWTKeyExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// JWTKeyExistsP checks if the JWTKey row exists. Panics on error.
func JWTKeyExistsP(exec boil.Executor, id int) bool {
	e, err := JWTKeyExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}