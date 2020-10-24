// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UsCounty is an object representing the database table.
type UsCounty struct {
	Gid      int          `boil:"gid" json:"gid" toml:"gid" yaml:"gid"`
	Statefp  null.String  `boil:"statefp" json:"statefp,omitempty" toml:"statefp" yaml:"statefp,omitempty"`
	Countyfp null.String  `boil:"countyfp" json:"countyfp,omitempty" toml:"countyfp" yaml:"countyfp,omitempty"`
	Countyns null.String  `boil:"countyns" json:"countyns,omitempty" toml:"countyns" yaml:"countyns,omitempty"`
	Geoid    null.String  `boil:"geoid" json:"geoid,omitempty" toml:"geoid" yaml:"geoid,omitempty"`
	Name     null.String  `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Namelsad null.String  `boil:"namelsad" json:"namelsad,omitempty" toml:"namelsad" yaml:"namelsad,omitempty"`
	Lsad     null.String  `boil:"lsad" json:"lsad,omitempty" toml:"lsad" yaml:"lsad,omitempty"`
	Classfp  null.String  `boil:"classfp" json:"classfp,omitempty" toml:"classfp" yaml:"classfp,omitempty"`
	MTFCC    null.String  `boil:"mtfcc" json:"mtfcc,omitempty" toml:"mtfcc" yaml:"mtfcc,omitempty"`
	Csafp    null.String  `boil:"csafp" json:"csafp,omitempty" toml:"csafp" yaml:"csafp,omitempty"`
	Cbsafp   null.String  `boil:"cbsafp" json:"cbsafp,omitempty" toml:"cbsafp" yaml:"cbsafp,omitempty"`
	Metdivfp null.String  `boil:"metdivfp" json:"metdivfp,omitempty" toml:"metdivfp" yaml:"metdivfp,omitempty"`
	Funcstat null.String  `boil:"funcstat" json:"funcstat,omitempty" toml:"funcstat" yaml:"funcstat,omitempty"`
	Aland    null.Float64 `boil:"aland" json:"aland,omitempty" toml:"aland" yaml:"aland,omitempty"`
	Awater   null.Float64 `boil:"awater" json:"awater,omitempty" toml:"awater" yaml:"awater,omitempty"`
	Intptlat null.String  `boil:"intptlat" json:"intptlat,omitempty" toml:"intptlat" yaml:"intptlat,omitempty"`
	Intptlon null.String  `boil:"intptlon" json:"intptlon,omitempty" toml:"intptlon" yaml:"intptlon,omitempty"`
	Geom     string       `boil:"geom" json:"geom,omitempty" toml:"geom" yaml:"geom,omitempty"`

	R *usCountyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L usCountyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UsCountyColumns = struct {
	Gid      string
	Statefp  string
	Countyfp string
	Countyns string
	Geoid    string
	Name     string
	Namelsad string
	Lsad     string
	Classfp  string
	MTFCC    string
	Csafp    string
	Cbsafp   string
	Metdivfp string
	Funcstat string
	Aland    string
	Awater   string
	Intptlat string
	Intptlon string
	Geom     string
}{
	Gid:      "gid",
	Statefp:  "statefp",
	Countyfp: "countyfp",
	Countyns: "countyns",
	Geoid:    "geoid",
	Name:     "name",
	Namelsad: "namelsad",
	Lsad:     "lsad",
	Classfp:  "classfp",
	MTFCC:    "mtfcc",
	Csafp:    "csafp",
	Cbsafp:   "cbsafp",
	Metdivfp: "metdivfp",
	Funcstat: "funcstat",
	Aland:    "aland",
	Awater:   "awater",
	Intptlat: "intptlat",
	Intptlon: "intptlon",
	Geom:     "geom",
}

// Generated where

var UsCountyWhere = struct {
	Gid      whereHelperint
	Statefp  whereHelpernull_String
	Countyfp whereHelpernull_String
	Countyns whereHelpernull_String
	Geoid    whereHelpernull_String
	Name     whereHelpernull_String
	Namelsad whereHelpernull_String
	Lsad     whereHelpernull_String
	Classfp  whereHelpernull_String
	MTFCC    whereHelpernull_String
	Csafp    whereHelpernull_String
	Cbsafp   whereHelpernull_String
	Metdivfp whereHelpernull_String
	Funcstat whereHelpernull_String
	Aland    whereHelpernull_Float64
	Awater   whereHelpernull_Float64
	Intptlat whereHelpernull_String
	Intptlon whereHelpernull_String
	Geom     whereHelperstring
}{
	Gid:      whereHelperint{field: "\"us_counties\".\"gid\""},
	Statefp:  whereHelpernull_String{field: "\"us_counties\".\"statefp\""},
	Countyfp: whereHelpernull_String{field: "\"us_counties\".\"countyfp\""},
	Countyns: whereHelpernull_String{field: "\"us_counties\".\"countyns\""},
	Geoid:    whereHelpernull_String{field: "\"us_counties\".\"geoid\""},
	Name:     whereHelpernull_String{field: "\"us_counties\".\"name\""},
	Namelsad: whereHelpernull_String{field: "\"us_counties\".\"namelsad\""},
	Lsad:     whereHelpernull_String{field: "\"us_counties\".\"lsad\""},
	Classfp:  whereHelpernull_String{field: "\"us_counties\".\"classfp\""},
	MTFCC:    whereHelpernull_String{field: "\"us_counties\".\"mtfcc\""},
	Csafp:    whereHelpernull_String{field: "\"us_counties\".\"csafp\""},
	Cbsafp:   whereHelpernull_String{field: "\"us_counties\".\"cbsafp\""},
	Metdivfp: whereHelpernull_String{field: "\"us_counties\".\"metdivfp\""},
	Funcstat: whereHelpernull_String{field: "\"us_counties\".\"funcstat\""},
	Aland:    whereHelpernull_Float64{field: "\"us_counties\".\"aland\""},
	Awater:   whereHelpernull_Float64{field: "\"us_counties\".\"awater\""},
	Intptlat: whereHelpernull_String{field: "\"us_counties\".\"intptlat\""},
	Intptlon: whereHelpernull_String{field: "\"us_counties\".\"intptlon\""},
	Geom:     whereHelperstring{field: "\"us_counties\".\"geom\""},
}

// UsCountyRels is where relationship names are stored.
var UsCountyRels = struct {
}{}

// usCountyR is where relationships are stored.
type usCountyR struct {
}

// NewStruct creates a new relationship struct
func (*usCountyR) NewStruct() *usCountyR {
	return &usCountyR{}
}

// usCountyL is where Load methods for each relationship are stored.
type usCountyL struct{}

var (
	usCountyAllColumns            = []string{"gid", "statefp", "countyfp", "countyns", "geoid", "name", "namelsad", "lsad", "classfp", "mtfcc", "csafp", "cbsafp", "metdivfp", "funcstat", "aland", "awater", "intptlat", "intptlon", "geom"}
	usCountyColumnsWithoutDefault = []string{"statefp", "countyfp", "countyns", "geoid", "name", "namelsad", "lsad", "classfp", "mtfcc", "csafp", "cbsafp", "metdivfp", "funcstat", "aland", "awater", "intptlat", "intptlon", "geom"}
	usCountyColumnsWithDefault    = []string{"gid"}
	usCountyPrimaryKeyColumns     = []string{"gid"}
)

type (
	// UsCountySlice is an alias for a slice of pointers to UsCounty.
	// This should generally be used opposed to []UsCounty.
	UsCountySlice []*UsCounty
	// UsCountyHook is the signature for custom UsCounty hook methods
	UsCountyHook func(context.Context, boil.ContextExecutor, *UsCounty) error

	usCountyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	usCountyType                 = reflect.TypeOf(&UsCounty{})
	usCountyMapping              = queries.MakeStructMapping(usCountyType)
	usCountyPrimaryKeyMapping, _ = queries.BindMapping(usCountyType, usCountyMapping, usCountyPrimaryKeyColumns)
	usCountyInsertCacheMut       sync.RWMutex
	usCountyInsertCache          = make(map[string]insertCache)
	usCountyUpdateCacheMut       sync.RWMutex
	usCountyUpdateCache          = make(map[string]updateCache)
	usCountyUpsertCacheMut       sync.RWMutex
	usCountyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var usCountyBeforeInsertHooks []UsCountyHook
var usCountyBeforeUpdateHooks []UsCountyHook
var usCountyBeforeDeleteHooks []UsCountyHook
var usCountyBeforeUpsertHooks []UsCountyHook

var usCountyAfterInsertHooks []UsCountyHook
var usCountyAfterSelectHooks []UsCountyHook
var usCountyAfterUpdateHooks []UsCountyHook
var usCountyAfterDeleteHooks []UsCountyHook
var usCountyAfterUpsertHooks []UsCountyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UsCounty) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UsCounty) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UsCounty) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UsCounty) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UsCounty) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UsCounty) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UsCounty) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UsCounty) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UsCounty) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usCountyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUsCountyHook registers your hook function for all future operations.
func AddUsCountyHook(hookPoint boil.HookPoint, usCountyHook UsCountyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		usCountyBeforeInsertHooks = append(usCountyBeforeInsertHooks, usCountyHook)
	case boil.BeforeUpdateHook:
		usCountyBeforeUpdateHooks = append(usCountyBeforeUpdateHooks, usCountyHook)
	case boil.BeforeDeleteHook:
		usCountyBeforeDeleteHooks = append(usCountyBeforeDeleteHooks, usCountyHook)
	case boil.BeforeUpsertHook:
		usCountyBeforeUpsertHooks = append(usCountyBeforeUpsertHooks, usCountyHook)
	case boil.AfterInsertHook:
		usCountyAfterInsertHooks = append(usCountyAfterInsertHooks, usCountyHook)
	case boil.AfterSelectHook:
		usCountyAfterSelectHooks = append(usCountyAfterSelectHooks, usCountyHook)
	case boil.AfterUpdateHook:
		usCountyAfterUpdateHooks = append(usCountyAfterUpdateHooks, usCountyHook)
	case boil.AfterDeleteHook:
		usCountyAfterDeleteHooks = append(usCountyAfterDeleteHooks, usCountyHook)
	case boil.AfterUpsertHook:
		usCountyAfterUpsertHooks = append(usCountyAfterUpsertHooks, usCountyHook)
	}
}

// One returns a single usCounty record from the query.
func (q usCountyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UsCounty, error) {
	o := &UsCounty{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for us_counties")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UsCounty records from the query.
func (q usCountyQuery) All(ctx context.Context, exec boil.ContextExecutor) (UsCountySlice, error) {
	var o []*UsCounty

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UsCounty slice")
	}

	if len(usCountyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UsCounty records in the query.
func (q usCountyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count us_counties rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q usCountyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if us_counties exists")
	}

	return count > 0, nil
}

// UsCounties retrieves all the records using an executor.
func UsCounties(mods ...qm.QueryMod) usCountyQuery {
	mods = append(mods, qm.From("\"us_counties\""))
	return usCountyQuery{NewQuery(mods...)}
}

// FindUsCounty retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUsCounty(ctx context.Context, exec boil.ContextExecutor, gid int, selectCols ...string) (*UsCounty, error) {
	usCountyObj := &UsCounty{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"us_counties\" where \"gid\"=$1", sel,
	)

	q := queries.Raw(query, gid)

	err := q.Bind(ctx, exec, usCountyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from us_counties")
	}

	return usCountyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UsCounty) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no us_counties provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(usCountyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	usCountyInsertCacheMut.RLock()
	cache, cached := usCountyInsertCache[key]
	usCountyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			usCountyAllColumns,
			usCountyColumnsWithDefault,
			usCountyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(usCountyType, usCountyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(usCountyType, usCountyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"us_counties\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"us_counties\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into us_counties")
	}

	if !cached {
		usCountyInsertCacheMut.Lock()
		usCountyInsertCache[key] = cache
		usCountyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UsCounty.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UsCounty) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	usCountyUpdateCacheMut.RLock()
	cache, cached := usCountyUpdateCache[key]
	usCountyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			usCountyAllColumns,
			usCountyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update us_counties, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"us_counties\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, usCountyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(usCountyType, usCountyMapping, append(wl, usCountyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update us_counties row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for us_counties")
	}

	if !cached {
		usCountyUpdateCacheMut.Lock()
		usCountyUpdateCache[key] = cache
		usCountyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q usCountyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for us_counties")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for us_counties")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UsCountySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usCountyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"us_counties\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, usCountyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in usCounty slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all usCounty")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UsCounty) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no us_counties provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(usCountyColumnsWithDefault, o)

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

	usCountyUpsertCacheMut.RLock()
	cache, cached := usCountyUpsertCache[key]
	usCountyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			usCountyAllColumns,
			usCountyColumnsWithDefault,
			usCountyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			usCountyAllColumns,
			usCountyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert us_counties, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(usCountyPrimaryKeyColumns))
			copy(conflict, usCountyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"us_counties\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(usCountyType, usCountyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(usCountyType, usCountyMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert us_counties")
	}

	if !cached {
		usCountyUpsertCacheMut.Lock()
		usCountyUpsertCache[key] = cache
		usCountyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UsCounty record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UsCounty) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UsCounty provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), usCountyPrimaryKeyMapping)
	sql := "DELETE FROM \"us_counties\" WHERE \"gid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from us_counties")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for us_counties")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q usCountyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no usCountyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from us_counties")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for us_counties")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UsCountySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(usCountyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usCountyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"us_counties\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, usCountyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from usCounty slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for us_counties")
	}

	if len(usCountyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UsCounty) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUsCounty(ctx, exec, o.Gid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UsCountySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UsCountySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usCountyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"us_counties\".* FROM \"us_counties\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, usCountyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UsCountySlice")
	}

	*o = slice

	return nil
}

// UsCountyExists checks if the UsCounty row exists.
func UsCountyExists(ctx context.Context, exec boil.ContextExecutor, gid int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"us_counties\" where \"gid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gid)
	}
	row := exec.QueryRowContext(ctx, sql, gid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if us_counties exists")
	}

	return exists, nil
}