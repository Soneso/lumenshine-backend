// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package horizon

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// HistoryOperationParticipant is an object representing the database table.
type HistoryOperationParticipant struct {
	ID                 int   `boil:"id" json:"id" toml:"id" yaml:"id"`
	HistoryOperationID int64 `boil:"history_operation_id" json:"history_operation_id" toml:"history_operation_id" yaml:"history_operation_id"`
	HistoryAccountID   int64 `boil:"history_account_id" json:"history_account_id" toml:"history_account_id" yaml:"history_account_id"`

	R *historyOperationParticipantR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L historyOperationParticipantL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HistoryOperationParticipantColumns = struct {
	ID                 string
	HistoryOperationID string
	HistoryAccountID   string
}{
	ID:                 "id",
	HistoryOperationID: "history_operation_id",
	HistoryAccountID:   "history_account_id",
}

// HistoryOperationParticipantRels is where relationship names are stored.
var HistoryOperationParticipantRels = struct {
}{}

// historyOperationParticipantR is where relationships are stored.
type historyOperationParticipantR struct {
}

// NewStruct creates a new relationship struct
func (*historyOperationParticipantR) NewStruct() *historyOperationParticipantR {
	return &historyOperationParticipantR{}
}

// historyOperationParticipantL is where Load methods for each relationship are stored.
type historyOperationParticipantL struct{}

var (
	historyOperationParticipantColumns               = []string{"id", "history_operation_id", "history_account_id"}
	historyOperationParticipantColumnsWithoutDefault = []string{"history_operation_id", "history_account_id"}
	historyOperationParticipantColumnsWithDefault    = []string{"id"}
	historyOperationParticipantPrimaryKeyColumns     = []string{"id"}
)

type (
	// HistoryOperationParticipantSlice is an alias for a slice of pointers to HistoryOperationParticipant.
	// This should generally be used opposed to []HistoryOperationParticipant.
	HistoryOperationParticipantSlice []*HistoryOperationParticipant
	// HistoryOperationParticipantHook is the signature for custom HistoryOperationParticipant hook methods
	HistoryOperationParticipantHook func(boil.Executor, *HistoryOperationParticipant) error

	historyOperationParticipantQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	historyOperationParticipantType                 = reflect.TypeOf(&HistoryOperationParticipant{})
	historyOperationParticipantMapping              = queries.MakeStructMapping(historyOperationParticipantType)
	historyOperationParticipantPrimaryKeyMapping, _ = queries.BindMapping(historyOperationParticipantType, historyOperationParticipantMapping, historyOperationParticipantPrimaryKeyColumns)
	historyOperationParticipantInsertCacheMut       sync.RWMutex
	historyOperationParticipantInsertCache          = make(map[string]insertCache)
	historyOperationParticipantUpdateCacheMut       sync.RWMutex
	historyOperationParticipantUpdateCache          = make(map[string]updateCache)
	historyOperationParticipantUpsertCacheMut       sync.RWMutex
	historyOperationParticipantUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var historyOperationParticipantBeforeInsertHooks []HistoryOperationParticipantHook
var historyOperationParticipantBeforeUpdateHooks []HistoryOperationParticipantHook
var historyOperationParticipantBeforeDeleteHooks []HistoryOperationParticipantHook
var historyOperationParticipantBeforeUpsertHooks []HistoryOperationParticipantHook

var historyOperationParticipantAfterInsertHooks []HistoryOperationParticipantHook
var historyOperationParticipantAfterSelectHooks []HistoryOperationParticipantHook
var historyOperationParticipantAfterUpdateHooks []HistoryOperationParticipantHook
var historyOperationParticipantAfterDeleteHooks []HistoryOperationParticipantHook
var historyOperationParticipantAfterUpsertHooks []HistoryOperationParticipantHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HistoryOperationParticipant) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HistoryOperationParticipant) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HistoryOperationParticipant) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HistoryOperationParticipant) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HistoryOperationParticipant) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HistoryOperationParticipant) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HistoryOperationParticipant) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HistoryOperationParticipant) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HistoryOperationParticipant) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range historyOperationParticipantAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHistoryOperationParticipantHook registers your hook function for all future operations.
func AddHistoryOperationParticipantHook(hookPoint boil.HookPoint, historyOperationParticipantHook HistoryOperationParticipantHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		historyOperationParticipantBeforeInsertHooks = append(historyOperationParticipantBeforeInsertHooks, historyOperationParticipantHook)
	case boil.BeforeUpdateHook:
		historyOperationParticipantBeforeUpdateHooks = append(historyOperationParticipantBeforeUpdateHooks, historyOperationParticipantHook)
	case boil.BeforeDeleteHook:
		historyOperationParticipantBeforeDeleteHooks = append(historyOperationParticipantBeforeDeleteHooks, historyOperationParticipantHook)
	case boil.BeforeUpsertHook:
		historyOperationParticipantBeforeUpsertHooks = append(historyOperationParticipantBeforeUpsertHooks, historyOperationParticipantHook)
	case boil.AfterInsertHook:
		historyOperationParticipantAfterInsertHooks = append(historyOperationParticipantAfterInsertHooks, historyOperationParticipantHook)
	case boil.AfterSelectHook:
		historyOperationParticipantAfterSelectHooks = append(historyOperationParticipantAfterSelectHooks, historyOperationParticipantHook)
	case boil.AfterUpdateHook:
		historyOperationParticipantAfterUpdateHooks = append(historyOperationParticipantAfterUpdateHooks, historyOperationParticipantHook)
	case boil.AfterDeleteHook:
		historyOperationParticipantAfterDeleteHooks = append(historyOperationParticipantAfterDeleteHooks, historyOperationParticipantHook)
	case boil.AfterUpsertHook:
		historyOperationParticipantAfterUpsertHooks = append(historyOperationParticipantAfterUpsertHooks, historyOperationParticipantHook)
	}
}

// OneG returns a single historyOperationParticipant record from the query using the global executor.
func (q historyOperationParticipantQuery) OneG() (*HistoryOperationParticipant, error) {
	return q.One(boil.GetDB())
}

// One returns a single historyOperationParticipant record from the query.
func (q historyOperationParticipantQuery) One(exec boil.Executor) (*HistoryOperationParticipant, error) {
	o := &HistoryOperationParticipant{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "horizon: failed to execute a one query for history_operation_participants")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all HistoryOperationParticipant records from the query using the global executor.
func (q historyOperationParticipantQuery) AllG() (HistoryOperationParticipantSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all HistoryOperationParticipant records from the query.
func (q historyOperationParticipantQuery) All(exec boil.Executor) (HistoryOperationParticipantSlice, error) {
	var o []*HistoryOperationParticipant

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "horizon: failed to assign all query results to HistoryOperationParticipant slice")
	}

	if len(historyOperationParticipantAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all HistoryOperationParticipant records in the query, and panics on error.
func (q historyOperationParticipantQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all HistoryOperationParticipant records in the query.
func (q historyOperationParticipantQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "horizon: failed to count history_operation_participants rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q historyOperationParticipantQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q historyOperationParticipantQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "horizon: failed to check if history_operation_participants exists")
	}

	return count > 0, nil
}

// HistoryOperationParticipants retrieves all the records using an executor.
func HistoryOperationParticipants(mods ...qm.QueryMod) historyOperationParticipantQuery {
	mods = append(mods, qm.From("\"history_operation_participants\""))
	return historyOperationParticipantQuery{NewQuery(mods...)}
}

// FindHistoryOperationParticipantG retrieves a single record by ID.
func FindHistoryOperationParticipantG(iD int, selectCols ...string) (*HistoryOperationParticipant, error) {
	return FindHistoryOperationParticipant(boil.GetDB(), iD, selectCols...)
}

// FindHistoryOperationParticipant retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHistoryOperationParticipant(exec boil.Executor, iD int, selectCols ...string) (*HistoryOperationParticipant, error) {
	historyOperationParticipantObj := &HistoryOperationParticipant{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"history_operation_participants\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, historyOperationParticipantObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "horizon: unable to select from history_operation_participants")
	}

	return historyOperationParticipantObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *HistoryOperationParticipant) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HistoryOperationParticipant) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("horizon: no history_operation_participants provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(historyOperationParticipantColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	historyOperationParticipantInsertCacheMut.RLock()
	cache, cached := historyOperationParticipantInsertCache[key]
	historyOperationParticipantInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			historyOperationParticipantColumns,
			historyOperationParticipantColumnsWithDefault,
			historyOperationParticipantColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(historyOperationParticipantType, historyOperationParticipantMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(historyOperationParticipantType, historyOperationParticipantMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"history_operation_participants\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"history_operation_participants\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
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
		return errors.Wrap(err, "horizon: unable to insert into history_operation_participants")
	}

	if !cached {
		historyOperationParticipantInsertCacheMut.Lock()
		historyOperationParticipantInsertCache[key] = cache
		historyOperationParticipantInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single HistoryOperationParticipant record using the global executor.
// See Update for more documentation.
func (o *HistoryOperationParticipant) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the HistoryOperationParticipant.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HistoryOperationParticipant) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	historyOperationParticipantUpdateCacheMut.RLock()
	cache, cached := historyOperationParticipantUpdateCache[key]
	historyOperationParticipantUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			historyOperationParticipantColumns,
			historyOperationParticipantPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("horizon: unable to update history_operation_participants, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"history_operation_participants\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, historyOperationParticipantPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(historyOperationParticipantType, historyOperationParticipantMapping, append(wl, historyOperationParticipantPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to update history_operation_participants row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "horizon: failed to get rows affected by update for history_operation_participants")
	}

	if !cached {
		historyOperationParticipantUpdateCacheMut.Lock()
		historyOperationParticipantUpdateCache[key] = cache
		historyOperationParticipantUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q historyOperationParticipantQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to update all for history_operation_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to retrieve rows affected for history_operation_participants")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o HistoryOperationParticipantSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HistoryOperationParticipantSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("horizon: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), historyOperationParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"history_operation_participants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, historyOperationParticipantPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to update all in historyOperationParticipant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to retrieve rows affected all in update all historyOperationParticipant")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *HistoryOperationParticipant) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HistoryOperationParticipant) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("horizon: no history_operation_participants provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(historyOperationParticipantColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	historyOperationParticipantUpsertCacheMut.RLock()
	cache, cached := historyOperationParticipantUpsertCache[key]
	historyOperationParticipantUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			historyOperationParticipantColumns,
			historyOperationParticipantColumnsWithDefault,
			historyOperationParticipantColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			historyOperationParticipantColumns,
			historyOperationParticipantPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("horizon: unable to upsert history_operation_participants, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(historyOperationParticipantPrimaryKeyColumns))
			copy(conflict, historyOperationParticipantPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"history_operation_participants\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(historyOperationParticipantType, historyOperationParticipantMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(historyOperationParticipantType, historyOperationParticipantMapping, ret)
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
		return errors.Wrap(err, "horizon: unable to upsert history_operation_participants")
	}

	if !cached {
		historyOperationParticipantUpsertCacheMut.Lock()
		historyOperationParticipantUpsertCache[key] = cache
		historyOperationParticipantUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single HistoryOperationParticipant record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *HistoryOperationParticipant) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single HistoryOperationParticipant record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HistoryOperationParticipant) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("horizon: no HistoryOperationParticipant provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), historyOperationParticipantPrimaryKeyMapping)
	sql := "DELETE FROM \"history_operation_participants\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to delete from history_operation_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "horizon: failed to get rows affected by delete for history_operation_participants")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q historyOperationParticipantQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("horizon: no historyOperationParticipantQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to delete all from history_operation_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "horizon: failed to get rows affected by deleteall for history_operation_participants")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o HistoryOperationParticipantSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HistoryOperationParticipantSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("horizon: no HistoryOperationParticipant slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(historyOperationParticipantBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), historyOperationParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"history_operation_participants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, historyOperationParticipantPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "horizon: unable to delete all from historyOperationParticipant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "horizon: failed to get rows affected by deleteall for history_operation_participants")
	}

	if len(historyOperationParticipantAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *HistoryOperationParticipant) ReloadG() error {
	if o == nil {
		return errors.New("horizon: no HistoryOperationParticipant provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *HistoryOperationParticipant) Reload(exec boil.Executor) error {
	ret, err := FindHistoryOperationParticipant(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HistoryOperationParticipantSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("horizon: empty HistoryOperationParticipantSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HistoryOperationParticipantSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HistoryOperationParticipantSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), historyOperationParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"history_operation_participants\".* FROM \"history_operation_participants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, historyOperationParticipantPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "horizon: unable to reload all in HistoryOperationParticipantSlice")
	}

	*o = slice

	return nil
}

// HistoryOperationParticipantExistsG checks if the HistoryOperationParticipant row exists.
func HistoryOperationParticipantExistsG(iD int) (bool, error) {
	return HistoryOperationParticipantExists(boil.GetDB(), iD)
}

// HistoryOperationParticipantExists checks if the HistoryOperationParticipant row exists.
func HistoryOperationParticipantExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"history_operation_participants\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "horizon: unable to check if history_operation_participants exists")
	}

	return exists, nil
}