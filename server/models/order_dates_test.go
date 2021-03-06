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

func testOrderDates(t *testing.T) {
	t.Parallel()

	query := OrderDates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrderDatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
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

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrderDatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrderDates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrderDatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderDateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrderDatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrderDateExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OrderDate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrderDateExists to return true, but got false.")
	}
}

func testOrderDatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	orderDateFound, err := FindOrderDate(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if orderDateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrderDatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrderDates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrderDatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrderDates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrderDatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	orderDateOne := &OrderDate{}
	orderDateTwo := &OrderDate{}
	if err = randomize.Struct(seed, orderDateOne, orderDateDBTypes, false, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}
	if err = randomize.Struct(seed, orderDateTwo, orderDateDBTypes, false, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderDateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderDateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrderDates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrderDatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	orderDateOne := &OrderDate{}
	orderDateTwo := &OrderDate{}
	if err = randomize.Struct(seed, orderDateOne, orderDateDBTypes, false, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}
	if err = randomize.Struct(seed, orderDateTwo, orderDateDBTypes, false, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderDateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderDateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func orderDateBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func orderDateAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrderDate) error {
	*o = OrderDate{}
	return nil
}

func testOrderDatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrderDate{}
	o := &OrderDate{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, orderDateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrderDate object: %s", err)
	}

	AddOrderDateHook(boil.BeforeInsertHook, orderDateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	orderDateBeforeInsertHooks = []OrderDateHook{}

	AddOrderDateHook(boil.AfterInsertHook, orderDateAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	orderDateAfterInsertHooks = []OrderDateHook{}

	AddOrderDateHook(boil.AfterSelectHook, orderDateAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	orderDateAfterSelectHooks = []OrderDateHook{}

	AddOrderDateHook(boil.BeforeUpdateHook, orderDateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	orderDateBeforeUpdateHooks = []OrderDateHook{}

	AddOrderDateHook(boil.AfterUpdateHook, orderDateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	orderDateAfterUpdateHooks = []OrderDateHook{}

	AddOrderDateHook(boil.BeforeDeleteHook, orderDateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	orderDateBeforeDeleteHooks = []OrderDateHook{}

	AddOrderDateHook(boil.AfterDeleteHook, orderDateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	orderDateAfterDeleteHooks = []OrderDateHook{}

	AddOrderDateHook(boil.BeforeUpsertHook, orderDateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	orderDateBeforeUpsertHooks = []OrderDateHook{}

	AddOrderDateHook(boil.AfterUpsertHook, orderDateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	orderDateAfterUpsertHooks = []OrderDateHook{}
}

func testOrderDatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrderDatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(orderDateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrderDatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
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

func testOrderDatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderDateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrderDatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrderDates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	orderDateDBTypes = map[string]string{`ID`: `bigint`, `Name`: `varchar`}
	_                = bytes.MinRead
)

func testOrderDatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(orderDatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(orderDateAllColumns) == len(orderDatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrderDatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(orderDateAllColumns) == len(orderDatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrderDate{}
	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderDateDBTypes, true, orderDatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(orderDateAllColumns, orderDatePrimaryKeyColumns) {
		fields = orderDateAllColumns
	} else {
		fields = strmangle.SetComplement(
			orderDateAllColumns,
			orderDatePrimaryKeyColumns,
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

	slice := OrderDateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrderDatesUpsert(t *testing.T) {
	t.Parallel()

	if len(orderDateAllColumns) == len(orderDatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOrderDateUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrderDate{}
	if err = randomize.Struct(seed, &o, orderDateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrderDate: %s", err)
	}

	count, err := OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, orderDateDBTypes, false, orderDatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrderDate struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrderDate: %s", err)
	}

	count, err = OrderDates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
