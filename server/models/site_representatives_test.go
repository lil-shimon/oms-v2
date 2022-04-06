// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testSiteRepresentatives(t *testing.T) {
	t.Parallel()

	query := SiteRepresentatives()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSiteRepresentativesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
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

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSiteRepresentativesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SiteRepresentatives().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSiteRepresentativesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SiteRepresentativeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSiteRepresentativesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SiteRepresentativeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if SiteRepresentative exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SiteRepresentativeExists to return true, but got false.")
	}
}

func testSiteRepresentativesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	siteRepresentativeFound, err := FindSiteRepresentative(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if siteRepresentativeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSiteRepresentativesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SiteRepresentatives().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSiteRepresentativesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SiteRepresentatives().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSiteRepresentativesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	siteRepresentativeOne := &SiteRepresentative{}
	siteRepresentativeTwo := &SiteRepresentative{}
	if err = randomize.Struct(seed, siteRepresentativeOne, siteRepresentativeDBTypes, false, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}
	if err = randomize.Struct(seed, siteRepresentativeTwo, siteRepresentativeDBTypes, false, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = siteRepresentativeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = siteRepresentativeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SiteRepresentatives().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSiteRepresentativesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	siteRepresentativeOne := &SiteRepresentative{}
	siteRepresentativeTwo := &SiteRepresentative{}
	if err = randomize.Struct(seed, siteRepresentativeOne, siteRepresentativeDBTypes, false, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}
	if err = randomize.Struct(seed, siteRepresentativeTwo, siteRepresentativeDBTypes, false, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = siteRepresentativeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = siteRepresentativeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func siteRepresentativeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func siteRepresentativeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SiteRepresentative) error {
	*o = SiteRepresentative{}
	return nil
}

func testSiteRepresentativesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SiteRepresentative{}
	o := &SiteRepresentative{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative object: %s", err)
	}

	AddSiteRepresentativeHook(boil.BeforeInsertHook, siteRepresentativeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeBeforeInsertHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.AfterInsertHook, siteRepresentativeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeAfterInsertHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.AfterSelectHook, siteRepresentativeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeAfterSelectHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.BeforeUpdateHook, siteRepresentativeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeBeforeUpdateHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.AfterUpdateHook, siteRepresentativeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeAfterUpdateHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.BeforeDeleteHook, siteRepresentativeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeBeforeDeleteHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.AfterDeleteHook, siteRepresentativeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeAfterDeleteHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.BeforeUpsertHook, siteRepresentativeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeBeforeUpsertHooks = []SiteRepresentativeHook{}

	AddSiteRepresentativeHook(boil.AfterUpsertHook, siteRepresentativeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	siteRepresentativeAfterUpsertHooks = []SiteRepresentativeHook{}
}

func testSiteRepresentativesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSiteRepresentativesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(siteRepresentativeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSiteRepresentativesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
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

func testSiteRepresentativesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SiteRepresentativeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSiteRepresentativesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SiteRepresentatives().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	siteRepresentativeDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `Mail`: `varchar`, `Phone`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_                         = bytes.MinRead
)

func testSiteRepresentativesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(siteRepresentativePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(siteRepresentativeAllColumns) == len(siteRepresentativePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSiteRepresentativesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(siteRepresentativeAllColumns) == len(siteRepresentativePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SiteRepresentative{}
	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, siteRepresentativeDBTypes, true, siteRepresentativePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(siteRepresentativeAllColumns, siteRepresentativePrimaryKeyColumns) {
		fields = siteRepresentativeAllColumns
	} else {
		fields = strmangle.SetComplement(
			siteRepresentativeAllColumns,
			siteRepresentativePrimaryKeyColumns,
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

	slice := SiteRepresentativeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSiteRepresentativesUpsert(t *testing.T) {
	t.Parallel()

	if len(siteRepresentativeAllColumns) == len(siteRepresentativePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLSiteRepresentativeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SiteRepresentative{}
	if err = randomize.Struct(seed, &o, siteRepresentativeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SiteRepresentative: %s", err)
	}

	count, err := SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, siteRepresentativeDBTypes, false, siteRepresentativePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SiteRepresentative struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SiteRepresentative: %s", err)
	}

	count, err = SiteRepresentatives().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}