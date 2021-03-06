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

// UserSecurity is an object representing the database table.
type UserSecurity struct {
	ID                int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID            int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	KDFSalt           string    `boil:"kdf_salt" json:"kdf_salt" toml:"kdf_salt" yaml:"kdf_salt"`
	MnemonicMasterKey string    `boil:"mnemonic_master_key" json:"mnemonic_master_key" toml:"mnemonic_master_key" yaml:"mnemonic_master_key"`
	MnemonicMasterIv  string    `boil:"mnemonic_master_iv" json:"mnemonic_master_iv" toml:"mnemonic_master_iv" yaml:"mnemonic_master_iv"`
	WordlistMasterKey string    `boil:"wordlist_master_key" json:"wordlist_master_key" toml:"wordlist_master_key" yaml:"wordlist_master_key"`
	WordlistMasterIv  string    `boil:"wordlist_master_iv" json:"wordlist_master_iv" toml:"wordlist_master_iv" yaml:"wordlist_master_iv"`
	Mnemonic          string    `boil:"mnemonic" json:"mnemonic" toml:"mnemonic" yaml:"mnemonic"`
	MnemonicIv        string    `boil:"mnemonic_iv" json:"mnemonic_iv" toml:"mnemonic_iv" yaml:"mnemonic_iv"`
	Wordlist          string    `boil:"wordlist" json:"wordlist" toml:"wordlist" yaml:"wordlist"`
	WordlistIv        string    `boil:"wordlist_iv" json:"wordlist_iv" toml:"wordlist_iv" yaml:"wordlist_iv"`
	PublicKey0        string    `boil:"public_key_0" json:"public_key_0" toml:"public_key_0" yaml:"public_key_0"`
	CreatedAt         time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt         time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy         string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`

	R *userSecurityR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userSecurityL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserSecurityColumns = struct {
	ID                string
	UserID            string
	KDFSalt           string
	MnemonicMasterKey string
	MnemonicMasterIv  string
	WordlistMasterKey string
	WordlistMasterIv  string
	Mnemonic          string
	MnemonicIv        string
	Wordlist          string
	WordlistIv        string
	PublicKey0        string
	CreatedAt         string
	UpdatedAt         string
	UpdatedBy         string
}{
	ID:                "id",
	UserID:            "user_id",
	KDFSalt:           "kdf_salt",
	MnemonicMasterKey: "mnemonic_master_key",
	MnemonicMasterIv:  "mnemonic_master_iv",
	WordlistMasterKey: "wordlist_master_key",
	WordlistMasterIv:  "wordlist_master_iv",
	Mnemonic:          "mnemonic",
	MnemonicIv:        "mnemonic_iv",
	Wordlist:          "wordlist",
	WordlistIv:        "wordlist_iv",
	PublicKey0:        "public_key_0",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	UpdatedBy:         "updated_by",
}

// UserSecurityRels is where relationship names are stored.
var UserSecurityRels = struct {
	User string
}{
	User: "User",
}

// userSecurityR is where relationships are stored.
type userSecurityR struct {
	User *UserProfile
}

// NewStruct creates a new relationship struct
func (*userSecurityR) NewStruct() *userSecurityR {
	return &userSecurityR{}
}

// userSecurityL is where Load methods for each relationship are stored.
type userSecurityL struct{}

var (
	userSecurityColumns               = []string{"id", "user_id", "kdf_salt", "mnemonic_master_key", "mnemonic_master_iv", "wordlist_master_key", "wordlist_master_iv", "mnemonic", "mnemonic_iv", "wordlist", "wordlist_iv", "public_key_0", "created_at", "updated_at", "updated_by"}
	userSecurityColumnsWithoutDefault = []string{"user_id", "kdf_salt", "mnemonic_master_key", "mnemonic_master_iv", "wordlist_master_key", "wordlist_master_iv", "mnemonic", "mnemonic_iv", "wordlist", "wordlist_iv", "public_key_0", "updated_by"}
	userSecurityColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	userSecurityPrimaryKeyColumns     = []string{"id"}
)

type (
	// UserSecuritySlice is an alias for a slice of pointers to UserSecurity.
	// This should generally be used opposed to []UserSecurity.
	UserSecuritySlice []*UserSecurity
	// UserSecurityHook is the signature for custom UserSecurity hook methods
	UserSecurityHook func(boil.Executor, *UserSecurity) error

	userSecurityQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userSecurityType                 = reflect.TypeOf(&UserSecurity{})
	userSecurityMapping              = queries.MakeStructMapping(userSecurityType)
	userSecurityPrimaryKeyMapping, _ = queries.BindMapping(userSecurityType, userSecurityMapping, userSecurityPrimaryKeyColumns)
	userSecurityInsertCacheMut       sync.RWMutex
	userSecurityInsertCache          = make(map[string]insertCache)
	userSecurityUpdateCacheMut       sync.RWMutex
	userSecurityUpdateCache          = make(map[string]updateCache)
	userSecurityUpsertCacheMut       sync.RWMutex
	userSecurityUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var userSecurityBeforeInsertHooks []UserSecurityHook
var userSecurityBeforeUpdateHooks []UserSecurityHook
var userSecurityBeforeDeleteHooks []UserSecurityHook
var userSecurityBeforeUpsertHooks []UserSecurityHook

var userSecurityAfterInsertHooks []UserSecurityHook
var userSecurityAfterSelectHooks []UserSecurityHook
var userSecurityAfterUpdateHooks []UserSecurityHook
var userSecurityAfterDeleteHooks []UserSecurityHook
var userSecurityAfterUpsertHooks []UserSecurityHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserSecurity) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserSecurity) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserSecurity) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserSecurity) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserSecurity) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserSecurity) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserSecurity) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserSecurity) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserSecurity) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userSecurityAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserSecurityHook registers your hook function for all future operations.
func AddUserSecurityHook(hookPoint boil.HookPoint, userSecurityHook UserSecurityHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userSecurityBeforeInsertHooks = append(userSecurityBeforeInsertHooks, userSecurityHook)
	case boil.BeforeUpdateHook:
		userSecurityBeforeUpdateHooks = append(userSecurityBeforeUpdateHooks, userSecurityHook)
	case boil.BeforeDeleteHook:
		userSecurityBeforeDeleteHooks = append(userSecurityBeforeDeleteHooks, userSecurityHook)
	case boil.BeforeUpsertHook:
		userSecurityBeforeUpsertHooks = append(userSecurityBeforeUpsertHooks, userSecurityHook)
	case boil.AfterInsertHook:
		userSecurityAfterInsertHooks = append(userSecurityAfterInsertHooks, userSecurityHook)
	case boil.AfterSelectHook:
		userSecurityAfterSelectHooks = append(userSecurityAfterSelectHooks, userSecurityHook)
	case boil.AfterUpdateHook:
		userSecurityAfterUpdateHooks = append(userSecurityAfterUpdateHooks, userSecurityHook)
	case boil.AfterDeleteHook:
		userSecurityAfterDeleteHooks = append(userSecurityAfterDeleteHooks, userSecurityHook)
	case boil.AfterUpsertHook:
		userSecurityAfterUpsertHooks = append(userSecurityAfterUpsertHooks, userSecurityHook)
	}
}

// OneG returns a single userSecurity record from the query using the global executor.
func (q userSecurityQuery) OneG() (*UserSecurity, error) {
	return q.One(boil.GetDB())
}

// One returns a single userSecurity record from the query.
func (q userSecurityQuery) One(exec boil.Executor) (*UserSecurity, error) {
	o := &UserSecurity{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_security")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all UserSecurity records from the query using the global executor.
func (q userSecurityQuery) AllG() (UserSecuritySlice, error) {
	return q.All(boil.GetDB())
}

// All returns all UserSecurity records from the query.
func (q userSecurityQuery) All(exec boil.Executor) (UserSecuritySlice, error) {
	var o []*UserSecurity

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserSecurity slice")
	}

	if len(userSecurityAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all UserSecurity records in the query, and panics on error.
func (q userSecurityQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all UserSecurity records in the query.
func (q userSecurityQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_security rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q userSecurityQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q userSecurityQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_security exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserSecurity) User(mods ...qm.QueryMod) userProfileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := UserProfiles(queryMods...)
	queries.SetFrom(query.Query, "\"user_profile\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userSecurityL) LoadUser(e boil.Executor, singular bool, maybeUserSecurity interface{}, mods queries.Applicator) error {
	var slice []*UserSecurity
	var object *UserSecurity

	if singular {
		object = maybeUserSecurity.(*UserSecurity)
	} else {
		slice = *maybeUserSecurity.(*[]*UserSecurity)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userSecurityR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userSecurityR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	query := NewQuery(qm.From(`user_profile`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserProfile")
	}

	var resultSlice []*UserProfile
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserProfile")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_profile")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_profile")
	}

	if len(userSecurityAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userProfileR{}
		}
		foreign.R.UserUserSecurity = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userProfileR{}
				}
				foreign.R.UserUserSecurity = local
				break
			}
		}
	}

	return nil
}

// SetUserG of the userSecurity to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserUserSecurity.
// Uses the global database handle.
func (o *UserSecurity) SetUserG(insert bool, related *UserProfile) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUser of the userSecurity to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserUserSecurity.
func (o *UserSecurity) SetUser(exec boil.Executor, insert bool, related *UserProfile) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_security\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userSecurityPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userSecurityR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userProfileR{
			UserUserSecurity: o,
		}
	} else {
		related.R.UserUserSecurity = o
	}

	return nil
}

// UserSecurities retrieves all the records using an executor.
func UserSecurities(mods ...qm.QueryMod) userSecurityQuery {
	mods = append(mods, qm.From("\"user_security\""))
	return userSecurityQuery{NewQuery(mods...)}
}

// FindUserSecurityG retrieves a single record by ID.
func FindUserSecurityG(iD int, selectCols ...string) (*UserSecurity, error) {
	return FindUserSecurity(boil.GetDB(), iD, selectCols...)
}

// FindUserSecurity retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserSecurity(exec boil.Executor, iD int, selectCols ...string) (*UserSecurity, error) {
	userSecurityObj := &UserSecurity{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_security\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, userSecurityObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_security")
	}

	return userSecurityObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *UserSecurity) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserSecurity) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_security provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userSecurityColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userSecurityInsertCacheMut.RLock()
	cache, cached := userSecurityInsertCache[key]
	userSecurityInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userSecurityColumns,
			userSecurityColumnsWithDefault,
			userSecurityColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userSecurityType, userSecurityMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userSecurityType, userSecurityMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_security\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_security\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_security")
	}

	if !cached {
		userSecurityInsertCacheMut.Lock()
		userSecurityInsertCache[key] = cache
		userSecurityInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single UserSecurity record using the global executor.
// See Update for more documentation.
func (o *UserSecurity) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the UserSecurity.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserSecurity) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userSecurityUpdateCacheMut.RLock()
	cache, cached := userSecurityUpdateCache[key]
	userSecurityUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userSecurityColumns,
			userSecurityPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_security, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_security\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userSecurityPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userSecurityType, userSecurityMapping, append(wl, userSecurityPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_security row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_security")
	}

	if !cached {
		userSecurityUpdateCacheMut.Lock()
		userSecurityUpdateCache[key] = cache
		userSecurityUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q userSecurityQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userSecurityQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_security")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_security")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserSecuritySlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserSecuritySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSecurityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_security\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userSecurityPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userSecurity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userSecurity")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *UserSecurity) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserSecurity) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_security provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userSecurityColumnsWithDefault, o)

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

	userSecurityUpsertCacheMut.RLock()
	cache, cached := userSecurityUpsertCache[key]
	userSecurityUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userSecurityColumns,
			userSecurityColumnsWithDefault,
			userSecurityColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userSecurityColumns,
			userSecurityPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert user_security, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userSecurityPrimaryKeyColumns))
			copy(conflict, userSecurityPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_security\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userSecurityType, userSecurityMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userSecurityType, userSecurityMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_security")
	}

	if !cached {
		userSecurityUpsertCacheMut.Lock()
		userSecurityUpsertCache[key] = cache
		userSecurityUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single UserSecurity record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *UserSecurity) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single UserSecurity record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserSecurity) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserSecurity provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userSecurityPrimaryKeyMapping)
	sql := "DELETE FROM \"user_security\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_security")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_security")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userSecurityQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userSecurityQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_security")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_security")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserSecuritySlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserSecuritySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserSecurity slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(userSecurityBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSecurityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_security\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userSecurityPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userSecurity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_security")
	}

	if len(userSecurityAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *UserSecurity) ReloadG() error {
	if o == nil {
		return errors.New("models: no UserSecurity provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserSecurity) Reload(exec boil.Executor) error {
	ret, err := FindUserSecurity(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSecuritySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty UserSecuritySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSecuritySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserSecuritySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSecurityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_security\".* FROM \"user_security\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userSecurityPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserSecuritySlice")
	}

	*o = slice

	return nil
}

// UserSecurityExistsG checks if the UserSecurity row exists.
func UserSecurityExistsG(iD int) (bool, error) {
	return UserSecurityExists(boil.GetDB(), iD)
}

// UserSecurityExists checks if the UserSecurity row exists.
func UserSecurityExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_security\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_security exists")
	}

	return exists, nil
}
