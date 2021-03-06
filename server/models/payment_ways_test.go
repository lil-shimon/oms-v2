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

func testPaymentWays(t *testing.T) {
	t.Parallel()

	query := PaymentWays()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPaymentWaysDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
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

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPaymentWaysQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PaymentWays().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPaymentWaysSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PaymentWaySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPaymentWaysExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PaymentWayExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PaymentWay exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PaymentWayExists to return true, but got false.")
	}
}

func testPaymentWaysFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	paymentWayFound, err := FindPaymentWay(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if paymentWayFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPaymentWaysBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PaymentWays().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPaymentWaysOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PaymentWays().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPaymentWaysAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	paymentWayOne := &PaymentWay{}
	paymentWayTwo := &PaymentWay{}
	if err = randomize.Struct(seed, paymentWayOne, paymentWayDBTypes, false, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}
	if err = randomize.Struct(seed, paymentWayTwo, paymentWayDBTypes, false, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = paymentWayOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = paymentWayTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PaymentWays().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPaymentWaysCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	paymentWayOne := &PaymentWay{}
	paymentWayTwo := &PaymentWay{}
	if err = randomize.Struct(seed, paymentWayOne, paymentWayDBTypes, false, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}
	if err = randomize.Struct(seed, paymentWayTwo, paymentWayDBTypes, false, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = paymentWayOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = paymentWayTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func paymentWayBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func paymentWayAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PaymentWay) error {
	*o = PaymentWay{}
	return nil
}

func testPaymentWaysHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PaymentWay{}
	o := &PaymentWay{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, paymentWayDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PaymentWay object: %s", err)
	}

	AddPaymentWayHook(boil.BeforeInsertHook, paymentWayBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	paymentWayBeforeInsertHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.AfterInsertHook, paymentWayAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	paymentWayAfterInsertHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.AfterSelectHook, paymentWayAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	paymentWayAfterSelectHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.BeforeUpdateHook, paymentWayBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	paymentWayBeforeUpdateHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.AfterUpdateHook, paymentWayAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	paymentWayAfterUpdateHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.BeforeDeleteHook, paymentWayBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	paymentWayBeforeDeleteHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.AfterDeleteHook, paymentWayAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	paymentWayAfterDeleteHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.BeforeUpsertHook, paymentWayBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	paymentWayBeforeUpsertHooks = []PaymentWayHook{}

	AddPaymentWayHook(boil.AfterUpsertHook, paymentWayAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	paymentWayAfterUpsertHooks = []PaymentWayHook{}
}

func testPaymentWaysInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPaymentWaysInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(paymentWayColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPaymentWayToManyCompanies(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PaymentWay
	var b, c Company

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, companyDBTypes, false, companyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, companyDBTypes, false, companyColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.PaymentWayID, a.ID)
	queries.Assign(&c.PaymentWayID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Companies().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.PaymentWayID, b.PaymentWayID) {
			bFound = true
		}
		if queries.Equal(v.PaymentWayID, c.PaymentWayID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PaymentWaySlice{&a}
	if err = a.L.LoadCompanies(ctx, tx, false, (*[]*PaymentWay)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Companies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Companies = nil
	if err = a.L.LoadCompanies(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Companies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPaymentWayToManyAddOpCompanies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PaymentWay
	var b, c, d, e Company

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, paymentWayDBTypes, false, strmangle.SetComplement(paymentWayPrimaryKeyColumns, paymentWayColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Company{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Company{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCompanies(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.PaymentWayID) {
			t.Error("foreign key was wrong value", a.ID, first.PaymentWayID)
		}
		if !queries.Equal(a.ID, second.PaymentWayID) {
			t.Error("foreign key was wrong value", a.ID, second.PaymentWayID)
		}

		if first.R.PaymentWay != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.PaymentWay != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Companies[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Companies[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Companies().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPaymentWayToManySetOpCompanies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PaymentWay
	var b, c, d, e Company

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, paymentWayDBTypes, false, strmangle.SetComplement(paymentWayPrimaryKeyColumns, paymentWayColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Company{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetCompanies(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Companies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCompanies(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Companies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PaymentWayID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PaymentWayID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.PaymentWayID) {
		t.Error("foreign key was wrong value", a.ID, d.PaymentWayID)
	}
	if !queries.Equal(a.ID, e.PaymentWayID) {
		t.Error("foreign key was wrong value", a.ID, e.PaymentWayID)
	}

	if b.R.PaymentWay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.PaymentWay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.PaymentWay != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.PaymentWay != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Companies[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Companies[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPaymentWayToManyRemoveOpCompanies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PaymentWay
	var b, c, d, e Company

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, paymentWayDBTypes, false, strmangle.SetComplement(paymentWayPrimaryKeyColumns, paymentWayColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Company{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddCompanies(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Companies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCompanies(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Companies().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PaymentWayID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PaymentWayID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.PaymentWay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.PaymentWay != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.PaymentWay != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.PaymentWay != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Companies) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Companies[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Companies[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPaymentWaysReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
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

func testPaymentWaysReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PaymentWaySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPaymentWaysSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PaymentWays().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	paymentWayDBTypes = map[string]string{`ID`: `bigint`, `Name`: `varchar`}
	_                 = bytes.MinRead
)

func testPaymentWaysUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(paymentWayPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(paymentWayAllColumns) == len(paymentWayPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPaymentWaysSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(paymentWayAllColumns) == len(paymentWayPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PaymentWay{}
	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, paymentWayDBTypes, true, paymentWayPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(paymentWayAllColumns, paymentWayPrimaryKeyColumns) {
		fields = paymentWayAllColumns
	} else {
		fields = strmangle.SetComplement(
			paymentWayAllColumns,
			paymentWayPrimaryKeyColumns,
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

	slice := PaymentWaySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPaymentWaysUpsert(t *testing.T) {
	t.Parallel()

	if len(paymentWayAllColumns) == len(paymentWayPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLPaymentWayUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PaymentWay{}
	if err = randomize.Struct(seed, &o, paymentWayDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PaymentWay: %s", err)
	}

	count, err := PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, paymentWayDBTypes, false, paymentWayPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PaymentWay struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PaymentWay: %s", err)
	}

	count, err = PaymentWays().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
