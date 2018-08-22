// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// AdminStellarAsset is an object representing the database table.
type AdminStellarAsset struct {
	ID                int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IssuerPublicKeyID string    `boil:"issuer_public_key_id" json:"issuer_public_key_id" toml:"issuer_public_key_id" yaml:"issuer_public_key_id"`
	AssetCode         string    `boil:"asset_code" json:"asset_code" toml:"asset_code" yaml:"asset_code"`
	CreatedAt         time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt         time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy         string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`

	R *adminStellarAssetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminStellarAssetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminStellarAssetColumns = struct {
	ID                string
	IssuerPublicKeyID string
	AssetCode         string
	CreatedAt         string
	UpdatedAt         string
	UpdatedBy         string
}{
	ID:                "id",
	IssuerPublicKeyID: "issuer_public_key_id",
	AssetCode:         "asset_code",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	UpdatedBy:         "updated_by",
}

// AdminStellarAssetRels is where relationship names are stored.
var AdminStellarAssetRels = struct {
	IssuerPublicKey string
}{
	IssuerPublicKey: "IssuerPublicKey",
}

// adminStellarAssetR is where relationships are stored.
type adminStellarAssetR struct {
	IssuerPublicKey *AdminStellarAccount
}

// NewStruct creates a new relationship struct
func (*adminStellarAssetR) NewStruct() *adminStellarAssetR {
	return &adminStellarAssetR{}
}

// adminStellarAssetL is where Load methods for each relationship are stored.
type adminStellarAssetL struct{}

var (
	adminStellarAssetColumns               = []string{"id", "issuer_public_key_id", "asset_code", "created_at", "updated_at", "updated_by"}
	adminStellarAssetColumnsWithoutDefault = []string{"issuer_public_key_id", "asset_code", "updated_by"}
	adminStellarAssetColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	adminStellarAssetPrimaryKeyColumns     = []string{"id"}
)

type (
	// AdminStellarAssetSlice is an alias for a slice of pointers to AdminStellarAsset.
	// This should generally be used opposed to []AdminStellarAsset.
	AdminStellarAssetSlice []*AdminStellarAsset
	// AdminStellarAssetHook is the signature for custom AdminStellarAsset hook methods
	AdminStellarAssetHook func(boil.Executor, *AdminStellarAsset) error

	adminStellarAssetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminStellarAssetType                 = reflect.TypeOf(&AdminStellarAsset{})
	adminStellarAssetMapping              = queries.MakeStructMapping(adminStellarAssetType)
	adminStellarAssetPrimaryKeyMapping, _ = queries.BindMapping(adminStellarAssetType, adminStellarAssetMapping, adminStellarAssetPrimaryKeyColumns)
	adminStellarAssetInsertCacheMut       sync.RWMutex
	adminStellarAssetInsertCache          = make(map[string]insertCache)
	adminStellarAssetUpdateCacheMut       sync.RWMutex
	adminStellarAssetUpdateCache          = make(map[string]updateCache)
	adminStellarAssetUpsertCacheMut       sync.RWMutex
	adminStellarAssetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var adminStellarAssetBeforeInsertHooks []AdminStellarAssetHook
var adminStellarAssetBeforeUpdateHooks []AdminStellarAssetHook
var adminStellarAssetBeforeDeleteHooks []AdminStellarAssetHook
var adminStellarAssetBeforeUpsertHooks []AdminStellarAssetHook

var adminStellarAssetAfterInsertHooks []AdminStellarAssetHook
var adminStellarAssetAfterSelectHooks []AdminStellarAssetHook
var adminStellarAssetAfterUpdateHooks []AdminStellarAssetHook
var adminStellarAssetAfterDeleteHooks []AdminStellarAssetHook
var adminStellarAssetAfterUpsertHooks []AdminStellarAssetHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AdminStellarAsset) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AdminStellarAsset) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AdminStellarAsset) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AdminStellarAsset) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AdminStellarAsset) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AdminStellarAsset) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AdminStellarAsset) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AdminStellarAsset) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AdminStellarAsset) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarAssetAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAdminStellarAssetHook registers your hook function for all future operations.
func AddAdminStellarAssetHook(hookPoint boil.HookPoint, adminStellarAssetHook AdminStellarAssetHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		adminStellarAssetBeforeInsertHooks = append(adminStellarAssetBeforeInsertHooks, adminStellarAssetHook)
	case boil.BeforeUpdateHook:
		adminStellarAssetBeforeUpdateHooks = append(adminStellarAssetBeforeUpdateHooks, adminStellarAssetHook)
	case boil.BeforeDeleteHook:
		adminStellarAssetBeforeDeleteHooks = append(adminStellarAssetBeforeDeleteHooks, adminStellarAssetHook)
	case boil.BeforeUpsertHook:
		adminStellarAssetBeforeUpsertHooks = append(adminStellarAssetBeforeUpsertHooks, adminStellarAssetHook)
	case boil.AfterInsertHook:
		adminStellarAssetAfterInsertHooks = append(adminStellarAssetAfterInsertHooks, adminStellarAssetHook)
	case boil.AfterSelectHook:
		adminStellarAssetAfterSelectHooks = append(adminStellarAssetAfterSelectHooks, adminStellarAssetHook)
	case boil.AfterUpdateHook:
		adminStellarAssetAfterUpdateHooks = append(adminStellarAssetAfterUpdateHooks, adminStellarAssetHook)
	case boil.AfterDeleteHook:
		adminStellarAssetAfterDeleteHooks = append(adminStellarAssetAfterDeleteHooks, adminStellarAssetHook)
	case boil.AfterUpsertHook:
		adminStellarAssetAfterUpsertHooks = append(adminStellarAssetAfterUpsertHooks, adminStellarAssetHook)
	}
}

// OneG returns a single adminStellarAsset record from the query using the global executor.
func (q adminStellarAssetQuery) OneG() (*AdminStellarAsset, error) {
	return q.One(boil.GetDB())
}

// One returns a single adminStellarAsset record from the query.
func (q adminStellarAssetQuery) One(exec boil.Executor) (*AdminStellarAsset, error) {
	o := &AdminStellarAsset{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for admin_stellar_asset")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all AdminStellarAsset records from the query using the global executor.
func (q adminStellarAssetQuery) AllG() (AdminStellarAssetSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all AdminStellarAsset records from the query.
func (q adminStellarAssetQuery) All(exec boil.Executor) (AdminStellarAssetSlice, error) {
	var o []*AdminStellarAsset

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AdminStellarAsset slice")
	}

	if len(adminStellarAssetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all AdminStellarAsset records in the query, and panics on error.
func (q adminStellarAssetQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all AdminStellarAsset records in the query.
func (q adminStellarAssetQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count admin_stellar_asset rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q adminStellarAssetQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q adminStellarAssetQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if admin_stellar_asset exists")
	}

	return count > 0, nil
}

// IssuerPublicKey pointed to by the foreign key.
func (o *AdminStellarAsset) IssuerPublicKey(mods ...qm.QueryMod) adminStellarAccountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("public_key=?", o.IssuerPublicKeyID),
	}

	queryMods = append(queryMods, mods...)

	query := AdminStellarAccounts(queryMods...)
	queries.SetFrom(query.Query, "\"admin_stellar_account\"")

	return query
}

// LoadIssuerPublicKey allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (adminStellarAssetL) LoadIssuerPublicKey(e boil.Executor, singular bool, maybeAdminStellarAsset interface{}, mods queries.Applicator) error {
	var slice []*AdminStellarAsset
	var object *AdminStellarAsset

	if singular {
		object = maybeAdminStellarAsset.(*AdminStellarAsset)
	} else {
		slice = *maybeAdminStellarAsset.(*[]*AdminStellarAsset)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &adminStellarAssetR{}
		}
		args = append(args, object.IssuerPublicKeyID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &adminStellarAssetR{}
			}

			for _, a := range args {
				if a == obj.IssuerPublicKeyID {
					continue Outer
				}
			}

			args = append(args, obj.IssuerPublicKeyID)
		}
	}

	query := NewQuery(qm.From(`admin_stellar_account`), qm.WhereIn(`public_key in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AdminStellarAccount")
	}

	var resultSlice []*AdminStellarAccount
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AdminStellarAccount")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for admin_stellar_account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for admin_stellar_account")
	}

	if len(adminStellarAssetAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.IssuerPublicKey = foreign
		if foreign.R == nil {
			foreign.R = &adminStellarAccountR{}
		}
		foreign.R.IssuerPublicKeyAdminStellarAssets = append(foreign.R.IssuerPublicKeyAdminStellarAssets, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.IssuerPublicKeyID == foreign.PublicKey {
				local.R.IssuerPublicKey = foreign
				if foreign.R == nil {
					foreign.R = &adminStellarAccountR{}
				}
				foreign.R.IssuerPublicKeyAdminStellarAssets = append(foreign.R.IssuerPublicKeyAdminStellarAssets, local)
				break
			}
		}
	}

	return nil
}

// SetIssuerPublicKeyG of the adminStellarAsset to the related item.
// Sets o.R.IssuerPublicKey to related.
// Adds o to related.R.IssuerPublicKeyAdminStellarAssets.
// Uses the global database handle.
func (o *AdminStellarAsset) SetIssuerPublicKeyG(insert bool, related *AdminStellarAccount) error {
	return o.SetIssuerPublicKey(boil.GetDB(), insert, related)
}

// SetIssuerPublicKey of the adminStellarAsset to the related item.
// Sets o.R.IssuerPublicKey to related.
// Adds o to related.R.IssuerPublicKeyAdminStellarAssets.
func (o *AdminStellarAsset) SetIssuerPublicKey(exec boil.Executor, insert bool, related *AdminStellarAccount) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"admin_stellar_asset\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"issuer_public_key_id"}),
		strmangle.WhereClause("\"", "\"", 2, adminStellarAssetPrimaryKeyColumns),
	)
	values := []interface{}{related.PublicKey, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.IssuerPublicKeyID = related.PublicKey
	if o.R == nil {
		o.R = &adminStellarAssetR{
			IssuerPublicKey: related,
		}
	} else {
		o.R.IssuerPublicKey = related
	}

	if related.R == nil {
		related.R = &adminStellarAccountR{
			IssuerPublicKeyAdminStellarAssets: AdminStellarAssetSlice{o},
		}
	} else {
		related.R.IssuerPublicKeyAdminStellarAssets = append(related.R.IssuerPublicKeyAdminStellarAssets, o)
	}

	return nil
}

// AdminStellarAssets retrieves all the records using an executor.
func AdminStellarAssets(mods ...qm.QueryMod) adminStellarAssetQuery {
	mods = append(mods, qm.From("\"admin_stellar_asset\""))
	return adminStellarAssetQuery{NewQuery(mods...)}
}

// FindAdminStellarAssetG retrieves a single record by ID.
func FindAdminStellarAssetG(iD int, selectCols ...string) (*AdminStellarAsset, error) {
	return FindAdminStellarAsset(boil.GetDB(), iD, selectCols...)
}

// FindAdminStellarAsset retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdminStellarAsset(exec boil.Executor, iD int, selectCols ...string) (*AdminStellarAsset, error) {
	adminStellarAssetObj := &AdminStellarAsset{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"admin_stellar_asset\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, adminStellarAssetObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from admin_stellar_asset")
	}

	return adminStellarAssetObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AdminStellarAsset) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AdminStellarAsset) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_stellar_asset provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(adminStellarAssetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	adminStellarAssetInsertCacheMut.RLock()
	cache, cached := adminStellarAssetInsertCache[key]
	adminStellarAssetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			adminStellarAssetColumns,
			adminStellarAssetColumnsWithDefault,
			adminStellarAssetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(adminStellarAssetType, adminStellarAssetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminStellarAssetType, adminStellarAssetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"admin_stellar_asset\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"admin_stellar_asset\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into admin_stellar_asset")
	}

	if !cached {
		adminStellarAssetInsertCacheMut.Lock()
		adminStellarAssetInsertCache[key] = cache
		adminStellarAssetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AdminStellarAsset record using the global executor.
// See Update for more documentation.
func (o *AdminStellarAsset) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the AdminStellarAsset.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AdminStellarAsset) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	adminStellarAssetUpdateCacheMut.RLock()
	cache, cached := adminStellarAssetUpdateCache[key]
	adminStellarAssetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			adminStellarAssetColumns,
			adminStellarAssetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update admin_stellar_asset, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"admin_stellar_asset\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, adminStellarAssetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminStellarAssetType, adminStellarAssetMapping, append(wl, adminStellarAssetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update admin_stellar_asset row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for admin_stellar_asset")
	}

	if !cached {
		adminStellarAssetUpdateCacheMut.Lock()
		adminStellarAssetUpdateCache[key] = cache
		adminStellarAssetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q adminStellarAssetQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for admin_stellar_asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for admin_stellar_asset")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AdminStellarAssetSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminStellarAssetSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminStellarAssetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"admin_stellar_asset\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, adminStellarAssetPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in adminStellarAsset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all adminStellarAsset")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AdminStellarAsset) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AdminStellarAsset) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_stellar_asset provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminStellarAssetColumnsWithDefault, o)

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

	adminStellarAssetUpsertCacheMut.RLock()
	cache, cached := adminStellarAssetUpsertCache[key]
	adminStellarAssetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			adminStellarAssetColumns,
			adminStellarAssetColumnsWithDefault,
			adminStellarAssetColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			adminStellarAssetColumns,
			adminStellarAssetPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert admin_stellar_asset, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(adminStellarAssetPrimaryKeyColumns))
			copy(conflict, adminStellarAssetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"admin_stellar_asset\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(adminStellarAssetType, adminStellarAssetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminStellarAssetType, adminStellarAssetMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert admin_stellar_asset")
	}

	if !cached {
		adminStellarAssetUpsertCacheMut.Lock()
		adminStellarAssetUpsertCache[key] = cache
		adminStellarAssetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single AdminStellarAsset record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AdminStellarAsset) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single AdminStellarAsset record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AdminStellarAsset) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AdminStellarAsset provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminStellarAssetPrimaryKeyMapping)
	sql := "DELETE FROM \"admin_stellar_asset\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from admin_stellar_asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for admin_stellar_asset")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q adminStellarAssetQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no adminStellarAssetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from admin_stellar_asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_stellar_asset")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AdminStellarAssetSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminStellarAssetSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AdminStellarAsset slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(adminStellarAssetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminStellarAssetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"admin_stellar_asset\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminStellarAssetPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from adminStellarAsset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_stellar_asset")
	}

	if len(adminStellarAssetAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AdminStellarAsset) ReloadG() error {
	if o == nil {
		return errors.New("models: no AdminStellarAsset provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AdminStellarAsset) Reload(exec boil.Executor) error {
	ret, err := FindAdminStellarAsset(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminStellarAssetSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AdminStellarAssetSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminStellarAssetSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AdminStellarAssetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminStellarAssetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"admin_stellar_asset\".* FROM \"admin_stellar_asset\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminStellarAssetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AdminStellarAssetSlice")
	}

	*o = slice

	return nil
}

// AdminStellarAssetExistsG checks if the AdminStellarAsset row exists.
func AdminStellarAssetExistsG(iD int) (bool, error) {
	return AdminStellarAssetExists(boil.GetDB(), iD)
}

// AdminStellarAssetExists checks if the AdminStellarAsset row exists.
func AdminStellarAssetExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"admin_stellar_asset\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if admin_stellar_asset exists")
	}

	return exists, nil
}
