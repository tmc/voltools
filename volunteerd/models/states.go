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

// State is an object representing the database table.
type State struct {
	Gid      int          `boil:"gid" json:"gid" toml:"gid" yaml:"gid"`
	Region   null.String  `boil:"region" json:"region,omitempty" toml:"region" yaml:"region,omitempty"`
	Division null.String  `boil:"division" json:"division,omitempty" toml:"division" yaml:"division,omitempty"`
	Statefp  string       `boil:"statefp" json:"statefp" toml:"statefp" yaml:"statefp"`
	Statens  null.String  `boil:"statens" json:"statens,omitempty" toml:"statens" yaml:"statens,omitempty"`
	Geoid    null.String  `boil:"geoid" json:"geoid,omitempty" toml:"geoid" yaml:"geoid,omitempty"`
	Stusps   null.String  `boil:"stusps" json:"stusps,omitempty" toml:"stusps" yaml:"stusps,omitempty"`
	Name     null.String  `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Lsad     null.String  `boil:"lsad" json:"lsad,omitempty" toml:"lsad" yaml:"lsad,omitempty"`
	MTFCC    null.String  `boil:"mtfcc" json:"mtfcc,omitempty" toml:"mtfcc" yaml:"mtfcc,omitempty"`
	Funcstat null.String  `boil:"funcstat" json:"funcstat,omitempty" toml:"funcstat" yaml:"funcstat,omitempty"`
	Aland    null.Float64 `boil:"aland" json:"aland,omitempty" toml:"aland" yaml:"aland,omitempty"`
	Awater   null.Float64 `boil:"awater" json:"awater,omitempty" toml:"awater" yaml:"awater,omitempty"`
	Intptlat null.String  `boil:"intptlat" json:"intptlat,omitempty" toml:"intptlat" yaml:"intptlat,omitempty"`
	Intptlon null.String  `boil:"intptlon" json:"intptlon,omitempty" toml:"intptlon" yaml:"intptlon,omitempty"`

	R *stateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StateColumns = struct {
	Gid      string
	Region   string
	Division string
	Statefp  string
	Statens  string
	Geoid    string
	Stusps   string
	Name     string
	Lsad     string
	MTFCC    string
	Funcstat string
	Aland    string
	Awater   string
	Intptlat string
	Intptlon string
}{
	Gid:      "gid",
	Region:   "region",
	Division: "division",
	Statefp:  "statefp",
	Statens:  "statens",
	Geoid:    "geoid",
	Stusps:   "stusps",
	Name:     "name",
	Lsad:     "lsad",
	MTFCC:    "mtfcc",
	Funcstat: "funcstat",
	Aland:    "aland",
	Awater:   "awater",
	Intptlat: "intptlat",
	Intptlon: "intptlon",
}

// Generated where

var StateWhere = struct {
	Gid      whereHelperint
	Region   whereHelpernull_String
	Division whereHelpernull_String
	Statefp  whereHelperstring
	Statens  whereHelpernull_String
	Geoid    whereHelpernull_String
	Stusps   whereHelpernull_String
	Name     whereHelpernull_String
	Lsad     whereHelpernull_String
	MTFCC    whereHelpernull_String
	Funcstat whereHelpernull_String
	Aland    whereHelpernull_Float64
	Awater   whereHelpernull_Float64
	Intptlat whereHelpernull_String
	Intptlon whereHelpernull_String
}{
	Gid:      whereHelperint{field: "\"states\".\"gid\""},
	Region:   whereHelpernull_String{field: "\"states\".\"region\""},
	Division: whereHelpernull_String{field: "\"states\".\"division\""},
	Statefp:  whereHelperstring{field: "\"states\".\"statefp\""},
	Statens:  whereHelpernull_String{field: "\"states\".\"statens\""},
	Geoid:    whereHelpernull_String{field: "\"states\".\"geoid\""},
	Stusps:   whereHelpernull_String{field: "\"states\".\"stusps\""},
	Name:     whereHelpernull_String{field: "\"states\".\"name\""},
	Lsad:     whereHelpernull_String{field: "\"states\".\"lsad\""},
	MTFCC:    whereHelpernull_String{field: "\"states\".\"mtfcc\""},
	Funcstat: whereHelpernull_String{field: "\"states\".\"funcstat\""},
	Aland:    whereHelpernull_Float64{field: "\"states\".\"aland\""},
	Awater:   whereHelpernull_Float64{field: "\"states\".\"awater\""},
	Intptlat: whereHelpernull_String{field: "\"states\".\"intptlat\""},
	Intptlon: whereHelpernull_String{field: "\"states\".\"intptlon\""},
}

// StateRels is where relationship names are stored.
var StateRels = struct {
	StatefpCongressionalDistricts string
}{
	StatefpCongressionalDistricts: "StatefpCongressionalDistricts",
}

// stateR is where relationships are stored.
type stateR struct {
	StatefpCongressionalDistricts CongressionalDistrictSlice `boil:"StatefpCongressionalDistricts" json:"StatefpCongressionalDistricts" toml:"StatefpCongressionalDistricts" yaml:"StatefpCongressionalDistricts"`
}

// NewStruct creates a new relationship struct
func (*stateR) NewStruct() *stateR {
	return &stateR{}
}

// stateL is where Load methods for each relationship are stored.
type stateL struct{}

var (
	stateAllColumns            = []string{"gid", "region", "division", "statefp", "statens", "geoid", "stusps", "name", "lsad", "mtfcc", "funcstat", "aland", "awater", "intptlat", "intptlon"}
	stateColumnsWithoutDefault = []string{"region", "division", "statefp", "statens", "geoid", "stusps", "name", "lsad", "mtfcc", "funcstat", "aland", "awater", "intptlat", "intptlon"}
	stateColumnsWithDefault    = []string{"gid"}
	statePrimaryKeyColumns     = []string{"statefp"}
)

type (
	// StateSlice is an alias for a slice of pointers to State.
	// This should generally be used opposed to []State.
	StateSlice []*State
	// StateHook is the signature for custom State hook methods
	StateHook func(context.Context, boil.ContextExecutor, *State) error

	stateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stateType                 = reflect.TypeOf(&State{})
	stateMapping              = queries.MakeStructMapping(stateType)
	statePrimaryKeyMapping, _ = queries.BindMapping(stateType, stateMapping, statePrimaryKeyColumns)
	stateInsertCacheMut       sync.RWMutex
	stateInsertCache          = make(map[string]insertCache)
	stateUpdateCacheMut       sync.RWMutex
	stateUpdateCache          = make(map[string]updateCache)
	stateUpsertCacheMut       sync.RWMutex
	stateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var stateBeforeInsertHooks []StateHook
var stateBeforeUpdateHooks []StateHook
var stateBeforeDeleteHooks []StateHook
var stateBeforeUpsertHooks []StateHook

var stateAfterInsertHooks []StateHook
var stateAfterSelectHooks []StateHook
var stateAfterUpdateHooks []StateHook
var stateAfterDeleteHooks []StateHook
var stateAfterUpsertHooks []StateHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *State) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *State) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *State) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *State) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *State) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *State) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *State) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *State) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *State) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStateHook registers your hook function for all future operations.
func AddStateHook(hookPoint boil.HookPoint, stateHook StateHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stateBeforeInsertHooks = append(stateBeforeInsertHooks, stateHook)
	case boil.BeforeUpdateHook:
		stateBeforeUpdateHooks = append(stateBeforeUpdateHooks, stateHook)
	case boil.BeforeDeleteHook:
		stateBeforeDeleteHooks = append(stateBeforeDeleteHooks, stateHook)
	case boil.BeforeUpsertHook:
		stateBeforeUpsertHooks = append(stateBeforeUpsertHooks, stateHook)
	case boil.AfterInsertHook:
		stateAfterInsertHooks = append(stateAfterInsertHooks, stateHook)
	case boil.AfterSelectHook:
		stateAfterSelectHooks = append(stateAfterSelectHooks, stateHook)
	case boil.AfterUpdateHook:
		stateAfterUpdateHooks = append(stateAfterUpdateHooks, stateHook)
	case boil.AfterDeleteHook:
		stateAfterDeleteHooks = append(stateAfterDeleteHooks, stateHook)
	case boil.AfterUpsertHook:
		stateAfterUpsertHooks = append(stateAfterUpsertHooks, stateHook)
	}
}

// One returns a single state record from the query.
func (q stateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*State, error) {
	o := &State{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for states")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all State records from the query.
func (q stateQuery) All(ctx context.Context, exec boil.ContextExecutor) (StateSlice, error) {
	var o []*State

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to State slice")
	}

	if len(stateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all State records in the query.
func (q stateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count states rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if states exists")
	}

	return count > 0, nil
}

// StatefpCongressionalDistricts retrieves all the congressional_district's CongressionalDistricts with an executor via statefp column.
func (o *State) StatefpCongressionalDistricts(mods ...qm.QueryMod) congressionalDistrictQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"congressional_districts\".\"statefp\"=?", o.Statefp),
	)

	query := CongressionalDistricts(queryMods...)
	queries.SetFrom(query.Query, "\"congressional_districts\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"congressional_districts\".*"})
	}

	return query
}

// LoadStatefpCongressionalDistricts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (stateL) LoadStatefpCongressionalDistricts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeState interface{}, mods queries.Applicator) error {
	var slice []*State
	var object *State

	if singular {
		object = maybeState.(*State)
	} else {
		slice = *maybeState.(*[]*State)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stateR{}
		}
		args = append(args, object.Statefp)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stateR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.Statefp) {
					continue Outer
				}
			}

			args = append(args, obj.Statefp)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`congressional_districts`),
		qm.WhereIn(`congressional_districts.statefp in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load congressional_districts")
	}

	var resultSlice []*CongressionalDistrict
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice congressional_districts")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on congressional_districts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for congressional_districts")
	}

	if len(congressionalDistrictAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.StatefpCongressionalDistricts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &congressionalDistrictR{}
			}
			foreign.R.Statefp = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.Statefp, foreign.Statefp) {
				local.R.StatefpCongressionalDistricts = append(local.R.StatefpCongressionalDistricts, foreign)
				if foreign.R == nil {
					foreign.R = &congressionalDistrictR{}
				}
				foreign.R.Statefp = local
				break
			}
		}
	}

	return nil
}

// AddStatefpCongressionalDistricts adds the given related objects to the existing relationships
// of the state, optionally inserting them as new records.
// Appends related to o.R.StatefpCongressionalDistricts.
// Sets related.R.Statefp appropriately.
func (o *State) AddStatefpCongressionalDistricts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CongressionalDistrict) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.Statefp, o.Statefp)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"congressional_districts\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"statefp"}),
				strmangle.WhereClause("\"", "\"", 2, congressionalDistrictPrimaryKeyColumns),
			)
			values := []interface{}{o.Statefp, rel.Geoid}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.Statefp, o.Statefp)
		}
	}

	if o.R == nil {
		o.R = &stateR{
			StatefpCongressionalDistricts: related,
		}
	} else {
		o.R.StatefpCongressionalDistricts = append(o.R.StatefpCongressionalDistricts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &congressionalDistrictR{
				Statefp: o,
			}
		} else {
			rel.R.Statefp = o
		}
	}
	return nil
}

// SetStatefpCongressionalDistricts removes all previously related items of the
// state replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Statefp's StatefpCongressionalDistricts accordingly.
// Replaces o.R.StatefpCongressionalDistricts with related.
// Sets related.R.Statefp's StatefpCongressionalDistricts accordingly.
func (o *State) SetStatefpCongressionalDistricts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CongressionalDistrict) error {
	query := "update \"congressional_districts\" set \"statefp\" = null where \"statefp\" = $1"
	values := []interface{}{o.Statefp}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.StatefpCongressionalDistricts {
			queries.SetScanner(&rel.Statefp, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Statefp = nil
		}

		o.R.StatefpCongressionalDistricts = nil
	}
	return o.AddStatefpCongressionalDistricts(ctx, exec, insert, related...)
}

// RemoveStatefpCongressionalDistricts relationships from objects passed in.
// Removes related items from R.StatefpCongressionalDistricts (uses pointer comparison, removal does not keep order)
// Sets related.R.Statefp.
func (o *State) RemoveStatefpCongressionalDistricts(ctx context.Context, exec boil.ContextExecutor, related ...*CongressionalDistrict) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.Statefp, nil)
		if rel.R != nil {
			rel.R.Statefp = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("statefp")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.StatefpCongressionalDistricts {
			if rel != ri {
				continue
			}

			ln := len(o.R.StatefpCongressionalDistricts)
			if ln > 1 && i < ln-1 {
				o.R.StatefpCongressionalDistricts[i] = o.R.StatefpCongressionalDistricts[ln-1]
			}
			o.R.StatefpCongressionalDistricts = o.R.StatefpCongressionalDistricts[:ln-1]
			break
		}
	}

	return nil
}

// States retrieves all the records using an executor.
func States(mods ...qm.QueryMod) stateQuery {
	mods = append(mods, qm.From("\"states\""))
	return stateQuery{NewQuery(mods...)}
}

// FindState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindState(ctx context.Context, exec boil.ContextExecutor, statefp string, selectCols ...string) (*State, error) {
	stateObj := &State{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"states\" where \"statefp\"=$1", sel,
	)

	q := queries.Raw(query, statefp)

	err := q.Bind(ctx, exec, stateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from states")
	}

	return stateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *State) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no states provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stateInsertCacheMut.RLock()
	cache, cached := stateInsertCache[key]
	stateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stateAllColumns,
			stateColumnsWithDefault,
			stateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stateType, stateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stateType, stateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"states\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"states\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into states")
	}

	if !cached {
		stateInsertCacheMut.Lock()
		stateInsertCache[key] = cache
		stateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the State.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *State) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	stateUpdateCacheMut.RLock()
	cache, cached := stateUpdateCache[key]
	stateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stateAllColumns,
			statePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update states, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"states\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, statePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stateType, stateMapping, append(wl, statePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update states row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for states")
	}

	if !cached {
		stateUpdateCacheMut.Lock()
		stateUpdateCache[key] = cache
		stateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q stateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for states")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, statePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in state slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all state")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *State) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no states provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stateColumnsWithDefault, o)

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

	stateUpsertCacheMut.RLock()
	cache, cached := stateUpsertCache[key]
	stateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stateAllColumns,
			stateColumnsWithDefault,
			stateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stateAllColumns,
			statePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert states, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(statePrimaryKeyColumns))
			copy(conflict, statePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"states\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stateType, stateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stateType, stateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert states")
	}

	if !cached {
		stateUpsertCacheMut.Lock()
		stateUpsertCache[key] = cache
		stateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single State record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *State) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no State provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), statePrimaryKeyMapping)
	sql := "DELETE FROM \"states\" WHERE \"statefp\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for states")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for states")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(stateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"states\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, statePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from state slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for states")
	}

	if len(stateAfterDeleteHooks) != 0 {
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
func (o *State) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindState(ctx, exec, o.Statefp)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"states\".* FROM \"states\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, statePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StateSlice")
	}

	*o = slice

	return nil
}

// StateExists checks if the State row exists.
func StateExists(ctx context.Context, exec boil.ContextExecutor, statefp string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"states\" where \"statefp\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, statefp)
	}
	row := exec.QueryRowContext(ctx, sql, statefp)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if states exists")
	}

	return exists, nil
}