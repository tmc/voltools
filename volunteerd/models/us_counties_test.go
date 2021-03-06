// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUsCounties(t *testing.T) {
	t.Parallel()

	query := UsCounties()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUsCountiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsCountiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UsCounties().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsCountiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UsCountySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsCountiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UsCountyExists(ctx, tx, o.Gid)
	if err != nil {
		t.Errorf("Unable to check if UsCounty exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UsCountyExists to return true, but got false.")
	}
}

func testUsCountiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	usCountyFound, err := FindUsCounty(ctx, tx, o.Gid)
	if err != nil {
		t.Error(err)
	}

	if usCountyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUsCountiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UsCounties().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUsCountiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UsCounties().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUsCountiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	usCountyOne := &UsCounty{}
	usCountyTwo := &UsCounty{}
	if err = randomize.Struct(seed, usCountyOne, usCountyDBTypes, false, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}
	if err = randomize.Struct(seed, usCountyTwo, usCountyDBTypes, false, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = usCountyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = usCountyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UsCounties().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUsCountiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	usCountyOne := &UsCounty{}
	usCountyTwo := &UsCounty{}
	if err = randomize.Struct(seed, usCountyOne, usCountyDBTypes, false, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}
	if err = randomize.Struct(seed, usCountyTwo, usCountyDBTypes, false, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = usCountyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = usCountyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func usCountyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func usCountyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UsCounty) error {
	*o = UsCounty{}
	return nil
}

func testUsCountiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UsCounty{}
	o := &UsCounty{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, usCountyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UsCounty object: %s", err)
	}

	AddUsCountyHook(boil.BeforeInsertHook, usCountyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	usCountyBeforeInsertHooks = []UsCountyHook{}

	AddUsCountyHook(boil.AfterInsertHook, usCountyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	usCountyAfterInsertHooks = []UsCountyHook{}

	AddUsCountyHook(boil.AfterSelectHook, usCountyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	usCountyAfterSelectHooks = []UsCountyHook{}

	AddUsCountyHook(boil.BeforeUpdateHook, usCountyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	usCountyBeforeUpdateHooks = []UsCountyHook{}

	AddUsCountyHook(boil.AfterUpdateHook, usCountyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	usCountyAfterUpdateHooks = []UsCountyHook{}

	AddUsCountyHook(boil.BeforeDeleteHook, usCountyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	usCountyBeforeDeleteHooks = []UsCountyHook{}

	AddUsCountyHook(boil.AfterDeleteHook, usCountyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	usCountyAfterDeleteHooks = []UsCountyHook{}

	AddUsCountyHook(boil.BeforeUpsertHook, usCountyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	usCountyBeforeUpsertHooks = []UsCountyHook{}

	AddUsCountyHook(boil.AfterUpsertHook, usCountyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	usCountyAfterUpsertHooks = []UsCountyHook{}
}

func testUsCountiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsCountiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(usCountyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsCountiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUsCountiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UsCountySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUsCountiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UsCounties().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	usCountyDBTypes = map[string]string{`Gid`: `integer`, `Statefp`: `character varying`, `Countyfp`: `character varying`, `Countyns`: `character varying`, `Geoid`: `character varying`, `Name`: `character varying`, `Namelsad`: `character varying`, `Lsad`: `character varying`, `Classfp`: `character varying`, `MTFCC`: `character varying`, `Csafp`: `character varying`, `Cbsafp`: `character varying`, `Metdivfp`: `character varying`, `Funcstat`: `character varying`, `Aland`: `double precision`, `Awater`: `double precision`, `Intptlat`: `character varying`, `Intptlon`: `character varying`, `Geom`: `USER-DEFINED`}
	_               = bytes.MinRead
)

func testUsCountiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(usCountyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(usCountyAllColumns) == len(usCountyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUsCountiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(usCountyAllColumns) == len(usCountyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UsCounty{}
	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, usCountyDBTypes, true, usCountyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(usCountyAllColumns, usCountyPrimaryKeyColumns) {
		fields = usCountyAllColumns
	} else {
		fields = strmangle.SetComplement(
			usCountyAllColumns,
			usCountyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UsCountySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUsCountiesUpsert(t *testing.T) {
	t.Parallel()

	if len(usCountyAllColumns) == len(usCountyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UsCounty{}
	if err = randomize.Struct(seed, &o, usCountyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UsCounty: %s", err)
	}

	count, err := UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, usCountyDBTypes, false, usCountyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UsCounty struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UsCounty: %s", err)
	}

	count, err = UsCounties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
