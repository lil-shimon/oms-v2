// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Order is an object representing the database table.
type Order struct {
	ID                      uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	StartDate               null.String `boil:"start_date" json:"start_date,omitempty" toml:"start_date" yaml:"start_date,omitempty"`
	EndDate                 null.String `boil:"end_date" json:"end_date,omitempty" toml:"end_date" yaml:"end_date,omitempty"`
	ExpectedStartDate       null.String `boil:"expected_start_date" json:"expected_start_date,omitempty" toml:"expected_start_date" yaml:"expected_start_date,omitempty"`
	ExpectedEndDate         null.String `boil:"expected_end_date" json:"expected_end_date,omitempty" toml:"expected_end_date" yaml:"expected_end_date,omitempty"`
	InvoiceNote             null.String `boil:"invoice_note" json:"invoice_note,omitempty" toml:"invoice_note" yaml:"invoice_note,omitempty"`
	SaleNote                null.String `boil:"sale_note" json:"sale_note,omitempty" toml:"sale_note" yaml:"sale_note,omitempty"`
	MaintainanceNote        null.String `boil:"maintainance_note" json:"maintainance_note,omitempty" toml:"maintainance_note" yaml:"maintainance_note,omitempty"`
	SimNumber               null.Int    `boil:"sim_number" json:"sim_number,omitempty" toml:"sim_number" yaml:"sim_number,omitempty"`
	SignnageID              null.Int    `boil:"signnage_id" json:"signnage_id,omitempty" toml:"signnage_id" yaml:"signnage_id,omitempty"`
	ConditionID             uint        `boil:"condition_id" json:"condition_id" toml:"condition_id" yaml:"condition_id"`
	UserID                  uint        `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	DeletedAt               null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	CreatedAt               null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt               null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	CompanyName             null.String `boil:"company_name" json:"company_name,omitempty" toml:"company_name" yaml:"company_name,omitempty"`
	SiteName                null.String `boil:"site_name" json:"site_name,omitempty" toml:"site_name" yaml:"site_name,omitempty"`
	AdditionalInvoiceID     uint        `boil:"additional_invoice_id" json:"additional_invoice_id" toml:"additional_invoice_id" yaml:"additional_invoice_id"`
	ConditionName           null.String `boil:"condition_name" json:"condition_name,omitempty" toml:"condition_name" yaml:"condition_name,omitempty"`
	AdditionalInvoice       null.String `boil:"additional_invoice" json:"additional_invoice,omitempty" toml:"additional_invoice" yaml:"additional_invoice,omitempty"`
	Payment                 null.String `boil:"payment" json:"payment,omitempty" toml:"payment" yaml:"payment,omitempty"`
	Comment                 null.String `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	PhoneNumber             null.String `boil:"phone_number" json:"phone_number,omitempty" toml:"phone_number" yaml:"phone_number,omitempty"`
	SiteRepresentative      null.String `boil:"site_representative" json:"site_representative,omitempty" toml:"site_representative" yaml:"site_representative,omitempty"`
	SiteRepresentativePhone null.String `boil:"site_representative_phone" json:"site_representative_phone,omitempty" toml:"site_representative_phone" yaml:"site_representative_phone,omitempty"`
	SiteMail                null.String `boil:"site_mail" json:"site_mail,omitempty" toml:"site_mail" yaml:"site_mail,omitempty"`
	SiteAddress             null.String `boil:"site_address" json:"site_address,omitempty" toml:"site_address" yaml:"site_address,omitempty"`
	TotalPrice              null.Int    `boil:"total_price" json:"total_price,omitempty" toml:"total_price" yaml:"total_price,omitempty"`
	SiteOfficeAddress       null.String `boil:"site_office_address" json:"site_office_address,omitempty" toml:"site_office_address" yaml:"site_office_address,omitempty"`
	DocumentID              null.Uint   `boil:"document_id" json:"document_id,omitempty" toml:"document_id" yaml:"document_id,omitempty"`

	R *orderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderColumns = struct {
	ID                      string
	StartDate               string
	EndDate                 string
	ExpectedStartDate       string
	ExpectedEndDate         string
	InvoiceNote             string
	SaleNote                string
	MaintainanceNote        string
	SimNumber               string
	SignnageID              string
	ConditionID             string
	UserID                  string
	DeletedAt               string
	CreatedAt               string
	UpdatedAt               string
	CompanyName             string
	SiteName                string
	AdditionalInvoiceID     string
	ConditionName           string
	AdditionalInvoice       string
	Payment                 string
	Comment                 string
	PhoneNumber             string
	SiteRepresentative      string
	SiteRepresentativePhone string
	SiteMail                string
	SiteAddress             string
	TotalPrice              string
	SiteOfficeAddress       string
	DocumentID              string
}{
	ID:                      "id",
	StartDate:               "start_date",
	EndDate:                 "end_date",
	ExpectedStartDate:       "expected_start_date",
	ExpectedEndDate:         "expected_end_date",
	InvoiceNote:             "invoice_note",
	SaleNote:                "sale_note",
	MaintainanceNote:        "maintainance_note",
	SimNumber:               "sim_number",
	SignnageID:              "signnage_id",
	ConditionID:             "condition_id",
	UserID:                  "user_id",
	DeletedAt:               "deleted_at",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	CompanyName:             "company_name",
	SiteName:                "site_name",
	AdditionalInvoiceID:     "additional_invoice_id",
	ConditionName:           "condition_name",
	AdditionalInvoice:       "additional_invoice",
	Payment:                 "payment",
	Comment:                 "comment",
	PhoneNumber:             "phone_number",
	SiteRepresentative:      "site_representative",
	SiteRepresentativePhone: "site_representative_phone",
	SiteMail:                "site_mail",
	SiteAddress:             "site_address",
	TotalPrice:              "total_price",
	SiteOfficeAddress:       "site_office_address",
	DocumentID:              "document_id",
}

var OrderTableColumns = struct {
	ID                      string
	StartDate               string
	EndDate                 string
	ExpectedStartDate       string
	ExpectedEndDate         string
	InvoiceNote             string
	SaleNote                string
	MaintainanceNote        string
	SimNumber               string
	SignnageID              string
	ConditionID             string
	UserID                  string
	DeletedAt               string
	CreatedAt               string
	UpdatedAt               string
	CompanyName             string
	SiteName                string
	AdditionalInvoiceID     string
	ConditionName           string
	AdditionalInvoice       string
	Payment                 string
	Comment                 string
	PhoneNumber             string
	SiteRepresentative      string
	SiteRepresentativePhone string
	SiteMail                string
	SiteAddress             string
	TotalPrice              string
	SiteOfficeAddress       string
	DocumentID              string
}{
	ID:                      "orders.id",
	StartDate:               "orders.start_date",
	EndDate:                 "orders.end_date",
	ExpectedStartDate:       "orders.expected_start_date",
	ExpectedEndDate:         "orders.expected_end_date",
	InvoiceNote:             "orders.invoice_note",
	SaleNote:                "orders.sale_note",
	MaintainanceNote:        "orders.maintainance_note",
	SimNumber:               "orders.sim_number",
	SignnageID:              "orders.signnage_id",
	ConditionID:             "orders.condition_id",
	UserID:                  "orders.user_id",
	DeletedAt:               "orders.deleted_at",
	CreatedAt:               "orders.created_at",
	UpdatedAt:               "orders.updated_at",
	CompanyName:             "orders.company_name",
	SiteName:                "orders.site_name",
	AdditionalInvoiceID:     "orders.additional_invoice_id",
	ConditionName:           "orders.condition_name",
	AdditionalInvoice:       "orders.additional_invoice",
	Payment:                 "orders.payment",
	Comment:                 "orders.comment",
	PhoneNumber:             "orders.phone_number",
	SiteRepresentative:      "orders.site_representative",
	SiteRepresentativePhone: "orders.site_representative_phone",
	SiteMail:                "orders.site_mail",
	SiteAddress:             "orders.site_address",
	TotalPrice:              "orders.total_price",
	SiteOfficeAddress:       "orders.site_office_address",
	DocumentID:              "orders.document_id",
}

// Generated where

var OrderWhere = struct {
	ID                      whereHelperuint
	StartDate               whereHelpernull_String
	EndDate                 whereHelpernull_String
	ExpectedStartDate       whereHelpernull_String
	ExpectedEndDate         whereHelpernull_String
	InvoiceNote             whereHelpernull_String
	SaleNote                whereHelpernull_String
	MaintainanceNote        whereHelpernull_String
	SimNumber               whereHelpernull_Int
	SignnageID              whereHelpernull_Int
	ConditionID             whereHelperuint
	UserID                  whereHelperuint
	DeletedAt               whereHelpernull_Time
	CreatedAt               whereHelpernull_Time
	UpdatedAt               whereHelpernull_Time
	CompanyName             whereHelpernull_String
	SiteName                whereHelpernull_String
	AdditionalInvoiceID     whereHelperuint
	ConditionName           whereHelpernull_String
	AdditionalInvoice       whereHelpernull_String
	Payment                 whereHelpernull_String
	Comment                 whereHelpernull_String
	PhoneNumber             whereHelpernull_String
	SiteRepresentative      whereHelpernull_String
	SiteRepresentativePhone whereHelpernull_String
	SiteMail                whereHelpernull_String
	SiteAddress             whereHelpernull_String
	TotalPrice              whereHelpernull_Int
	SiteOfficeAddress       whereHelpernull_String
	DocumentID              whereHelpernull_Uint
}{
	ID:                      whereHelperuint{field: "`orders`.`id`"},
	StartDate:               whereHelpernull_String{field: "`orders`.`start_date`"},
	EndDate:                 whereHelpernull_String{field: "`orders`.`end_date`"},
	ExpectedStartDate:       whereHelpernull_String{field: "`orders`.`expected_start_date`"},
	ExpectedEndDate:         whereHelpernull_String{field: "`orders`.`expected_end_date`"},
	InvoiceNote:             whereHelpernull_String{field: "`orders`.`invoice_note`"},
	SaleNote:                whereHelpernull_String{field: "`orders`.`sale_note`"},
	MaintainanceNote:        whereHelpernull_String{field: "`orders`.`maintainance_note`"},
	SimNumber:               whereHelpernull_Int{field: "`orders`.`sim_number`"},
	SignnageID:              whereHelpernull_Int{field: "`orders`.`signnage_id`"},
	ConditionID:             whereHelperuint{field: "`orders`.`condition_id`"},
	UserID:                  whereHelperuint{field: "`orders`.`user_id`"},
	DeletedAt:               whereHelpernull_Time{field: "`orders`.`deleted_at`"},
	CreatedAt:               whereHelpernull_Time{field: "`orders`.`created_at`"},
	UpdatedAt:               whereHelpernull_Time{field: "`orders`.`updated_at`"},
	CompanyName:             whereHelpernull_String{field: "`orders`.`company_name`"},
	SiteName:                whereHelpernull_String{field: "`orders`.`site_name`"},
	AdditionalInvoiceID:     whereHelperuint{field: "`orders`.`additional_invoice_id`"},
	ConditionName:           whereHelpernull_String{field: "`orders`.`condition_name`"},
	AdditionalInvoice:       whereHelpernull_String{field: "`orders`.`additional_invoice`"},
	Payment:                 whereHelpernull_String{field: "`orders`.`payment`"},
	Comment:                 whereHelpernull_String{field: "`orders`.`comment`"},
	PhoneNumber:             whereHelpernull_String{field: "`orders`.`phone_number`"},
	SiteRepresentative:      whereHelpernull_String{field: "`orders`.`site_representative`"},
	SiteRepresentativePhone: whereHelpernull_String{field: "`orders`.`site_representative_phone`"},
	SiteMail:                whereHelpernull_String{field: "`orders`.`site_mail`"},
	SiteAddress:             whereHelpernull_String{field: "`orders`.`site_address`"},
	TotalPrice:              whereHelpernull_Int{field: "`orders`.`total_price`"},
	SiteOfficeAddress:       whereHelpernull_String{field: "`orders`.`site_office_address`"},
	DocumentID:              whereHelpernull_Uint{field: "`orders`.`document_id`"},
}

// OrderRels is where relationship names are stored.
var OrderRels = struct {
}{}

// orderR is where relationships are stored.
type orderR struct {
}

// NewStruct creates a new relationship struct
func (*orderR) NewStruct() *orderR {
	return &orderR{}
}

// orderL is where Load methods for each relationship are stored.
type orderL struct{}

var (
	orderAllColumns            = []string{"id", "start_date", "end_date", "expected_start_date", "expected_end_date", "invoice_note", "sale_note", "maintainance_note", "sim_number", "signnage_id", "condition_id", "user_id", "deleted_at", "created_at", "updated_at", "company_name", "site_name", "additional_invoice_id", "condition_name", "additional_invoice", "payment", "comment", "phone_number", "site_representative", "site_representative_phone", "site_mail", "site_address", "total_price", "site_office_address", "document_id"}
	orderColumnsWithoutDefault = []string{"start_date", "end_date", "expected_start_date", "expected_end_date", "invoice_note", "sale_note", "maintainance_note", "sim_number", "signnage_id", "condition_id", "user_id", "deleted_at", "created_at", "updated_at", "company_name", "site_name", "additional_invoice_id", "condition_name", "additional_invoice", "payment", "comment", "phone_number", "site_representative", "site_representative_phone", "site_mail", "site_address", "total_price", "site_office_address", "document_id"}
	orderColumnsWithDefault    = []string{"id"}
	orderPrimaryKeyColumns     = []string{"id"}
)

type (
	// OrderSlice is an alias for a slice of pointers to Order.
	// This should almost always be used instead of []Order.
	OrderSlice []*Order
	// OrderHook is the signature for custom Order hook methods
	OrderHook func(context.Context, boil.ContextExecutor, *Order) error

	orderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderType                 = reflect.TypeOf(&Order{})
	orderMapping              = queries.MakeStructMapping(orderType)
	orderPrimaryKeyMapping, _ = queries.BindMapping(orderType, orderMapping, orderPrimaryKeyColumns)
	orderInsertCacheMut       sync.RWMutex
	orderInsertCache          = make(map[string]insertCache)
	orderUpdateCacheMut       sync.RWMutex
	orderUpdateCache          = make(map[string]updateCache)
	orderUpsertCacheMut       sync.RWMutex
	orderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderBeforeInsertHooks []OrderHook
var orderBeforeUpdateHooks []OrderHook
var orderBeforeDeleteHooks []OrderHook
var orderBeforeUpsertHooks []OrderHook

var orderAfterInsertHooks []OrderHook
var orderAfterSelectHooks []OrderHook
var orderAfterUpdateHooks []OrderHook
var orderAfterDeleteHooks []OrderHook
var orderAfterUpsertHooks []OrderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Order) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Order) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Order) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Order) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Order) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Order) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Order) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Order) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Order) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderHook registers your hook function for all future operations.
func AddOrderHook(hookPoint boil.HookPoint, orderHook OrderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		orderBeforeInsertHooks = append(orderBeforeInsertHooks, orderHook)
	case boil.BeforeUpdateHook:
		orderBeforeUpdateHooks = append(orderBeforeUpdateHooks, orderHook)
	case boil.BeforeDeleteHook:
		orderBeforeDeleteHooks = append(orderBeforeDeleteHooks, orderHook)
	case boil.BeforeUpsertHook:
		orderBeforeUpsertHooks = append(orderBeforeUpsertHooks, orderHook)
	case boil.AfterInsertHook:
		orderAfterInsertHooks = append(orderAfterInsertHooks, orderHook)
	case boil.AfterSelectHook:
		orderAfterSelectHooks = append(orderAfterSelectHooks, orderHook)
	case boil.AfterUpdateHook:
		orderAfterUpdateHooks = append(orderAfterUpdateHooks, orderHook)
	case boil.AfterDeleteHook:
		orderAfterDeleteHooks = append(orderAfterDeleteHooks, orderHook)
	case boil.AfterUpsertHook:
		orderAfterUpsertHooks = append(orderAfterUpsertHooks, orderHook)
	}
}

// One returns a single order record from the query.
func (q orderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Order, error) {
	o := &Order{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for orders")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Order records from the query.
func (q orderQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderSlice, error) {
	var o []*Order

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Order slice")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Order records in the query.
func (q orderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count orders rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if orders exists")
	}

	return count > 0, nil
}

// Orders retrieves all the records using an executor.
func Orders(mods ...qm.QueryMod) orderQuery {
	mods = append(mods, qm.From("`orders`"))
	return orderQuery{NewQuery(mods...)}
}

// FindOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrder(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Order, error) {
	orderObj := &Order{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `orders` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from orders")
	}

	if err = orderObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderObj, err
	}

	return orderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Order) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no orders provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderInsertCacheMut.RLock()
	cache, cached := orderInsertCache[key]
	orderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderAllColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderType, orderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `orders` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `orders` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `orders` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, orderPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into orders")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == orderMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for orders")
	}

CacheNoHooks:
	if !cached {
		orderInsertCacheMut.Lock()
		orderInsertCache[key] = cache
		orderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Order.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Order) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderUpdateCacheMut.RLock()
	cache, cached := orderUpdateCache[key]
	orderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update orders, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `orders` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, orderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, append(wl, orderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update orders row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for orders")
	}

	if !cached {
		orderUpdateCacheMut.Lock()
		orderUpdateCache[key] = cache
		orderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q orderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for orders")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `orders` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all order")
	}
	return rowsAff, nil
}

var mySQLOrderUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Order) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no orders provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOrderUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	orderUpsertCacheMut.RLock()
	cache, cached := orderUpsertCache[key]
	orderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orderAllColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert orders, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`orders`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `orders` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderType, orderMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for orders")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == orderMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(orderType, orderMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for orders")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for orders")
	}

CacheNoHooks:
	if !cached {
		orderUpsertCacheMut.Lock()
		orderUpsertCache[key] = cache
		orderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Order record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Order) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Order provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderPrimaryKeyMapping)
	sql := "DELETE FROM `orders` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for orders")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no orderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orders")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `orders` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orders")
	}

	if len(orderAfterDeleteHooks) != 0 {
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
func (o *Order) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `orders`.* FROM `orders` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrderSlice")
	}

	*o = slice

	return nil
}

// OrderExists checks if the Order row exists.
func OrderExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `orders` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if orders exists")
	}

	return exists, nil
}