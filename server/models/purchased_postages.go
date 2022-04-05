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

// PurchasedPostage is an object representing the database table.
type PurchasedPostage struct {
	ID               uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	DocumentID       uint      `boil:"document_id" json:"document_id" toml:"document_id" yaml:"document_id"`
	SenderPlace      string    `boil:"sender_place" json:"sender_place" toml:"sender_place" yaml:"sender_place"`
	DestinationPlace string    `boil:"destination_place" json:"destination_place" toml:"destination_place" yaml:"destination_place"`
	PostagePrice     int       `boil:"postage_price" json:"postage_price" toml:"postage_price" yaml:"postage_price"`
	Quantity         int       `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	Size             string    `boil:"size" json:"size" toml:"size" yaml:"size"`
	Unit             string    `boil:"unit" json:"unit" toml:"unit" yaml:"unit"`
	DeletedAt        null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	CreatedAt        null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt        null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *purchasedPostageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L purchasedPostageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PurchasedPostageColumns = struct {
	ID               string
	DocumentID       string
	SenderPlace      string
	DestinationPlace string
	PostagePrice     string
	Quantity         string
	Size             string
	Unit             string
	DeletedAt        string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	DocumentID:       "document_id",
	SenderPlace:      "sender_place",
	DestinationPlace: "destination_place",
	PostagePrice:     "postage_price",
	Quantity:         "quantity",
	Size:             "size",
	Unit:             "unit",
	DeletedAt:        "deleted_at",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

var PurchasedPostageTableColumns = struct {
	ID               string
	DocumentID       string
	SenderPlace      string
	DestinationPlace string
	PostagePrice     string
	Quantity         string
	Size             string
	Unit             string
	DeletedAt        string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "purchased_postages.id",
	DocumentID:       "purchased_postages.document_id",
	SenderPlace:      "purchased_postages.sender_place",
	DestinationPlace: "purchased_postages.destination_place",
	PostagePrice:     "purchased_postages.postage_price",
	Quantity:         "purchased_postages.quantity",
	Size:             "purchased_postages.size",
	Unit:             "purchased_postages.unit",
	DeletedAt:        "purchased_postages.deleted_at",
	CreatedAt:        "purchased_postages.created_at",
	UpdatedAt:        "purchased_postages.updated_at",
}

// Generated where

var PurchasedPostageWhere = struct {
	ID               whereHelperuint
	DocumentID       whereHelperuint
	SenderPlace      whereHelperstring
	DestinationPlace whereHelperstring
	PostagePrice     whereHelperint
	Quantity         whereHelperint
	Size             whereHelperstring
	Unit             whereHelperstring
	DeletedAt        whereHelpernull_Time
	CreatedAt        whereHelpernull_Time
	UpdatedAt        whereHelpernull_Time
}{
	ID:               whereHelperuint{field: "`purchased_postages`.`id`"},
	DocumentID:       whereHelperuint{field: "`purchased_postages`.`document_id`"},
	SenderPlace:      whereHelperstring{field: "`purchased_postages`.`sender_place`"},
	DestinationPlace: whereHelperstring{field: "`purchased_postages`.`destination_place`"},
	PostagePrice:     whereHelperint{field: "`purchased_postages`.`postage_price`"},
	Quantity:         whereHelperint{field: "`purchased_postages`.`quantity`"},
	Size:             whereHelperstring{field: "`purchased_postages`.`size`"},
	Unit:             whereHelperstring{field: "`purchased_postages`.`unit`"},
	DeletedAt:        whereHelpernull_Time{field: "`purchased_postages`.`deleted_at`"},
	CreatedAt:        whereHelpernull_Time{field: "`purchased_postages`.`created_at`"},
	UpdatedAt:        whereHelpernull_Time{field: "`purchased_postages`.`updated_at`"},
}

// PurchasedPostageRels is where relationship names are stored.
var PurchasedPostageRels = struct {
	Document string
}{
	Document: "Document",
}

// purchasedPostageR is where relationships are stored.
type purchasedPostageR struct {
	Document *Document `boil:"Document" json:"Document" toml:"Document" yaml:"Document"`
}

// NewStruct creates a new relationship struct
func (*purchasedPostageR) NewStruct() *purchasedPostageR {
	return &purchasedPostageR{}
}

// purchasedPostageL is where Load methods for each relationship are stored.
type purchasedPostageL struct{}

var (
	purchasedPostageAllColumns            = []string{"id", "document_id", "sender_place", "destination_place", "postage_price", "quantity", "size", "unit", "deleted_at", "created_at", "updated_at"}
	purchasedPostageColumnsWithoutDefault = []string{"document_id", "sender_place", "destination_place", "postage_price", "quantity", "size", "unit", "deleted_at", "created_at", "updated_at"}
	purchasedPostageColumnsWithDefault    = []string{"id"}
	purchasedPostagePrimaryKeyColumns     = []string{"id"}
)

type (
	// PurchasedPostageSlice is an alias for a slice of pointers to PurchasedPostage.
	// This should almost always be used instead of []PurchasedPostage.
	PurchasedPostageSlice []*PurchasedPostage
	// PurchasedPostageHook is the signature for custom PurchasedPostage hook methods
	PurchasedPostageHook func(context.Context, boil.ContextExecutor, *PurchasedPostage) error

	purchasedPostageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	purchasedPostageType                 = reflect.TypeOf(&PurchasedPostage{})
	purchasedPostageMapping              = queries.MakeStructMapping(purchasedPostageType)
	purchasedPostagePrimaryKeyMapping, _ = queries.BindMapping(purchasedPostageType, purchasedPostageMapping, purchasedPostagePrimaryKeyColumns)
	purchasedPostageInsertCacheMut       sync.RWMutex
	purchasedPostageInsertCache          = make(map[string]insertCache)
	purchasedPostageUpdateCacheMut       sync.RWMutex
	purchasedPostageUpdateCache          = make(map[string]updateCache)
	purchasedPostageUpsertCacheMut       sync.RWMutex
	purchasedPostageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var purchasedPostageBeforeInsertHooks []PurchasedPostageHook
var purchasedPostageBeforeUpdateHooks []PurchasedPostageHook
var purchasedPostageBeforeDeleteHooks []PurchasedPostageHook
var purchasedPostageBeforeUpsertHooks []PurchasedPostageHook

var purchasedPostageAfterInsertHooks []PurchasedPostageHook
var purchasedPostageAfterSelectHooks []PurchasedPostageHook
var purchasedPostageAfterUpdateHooks []PurchasedPostageHook
var purchasedPostageAfterDeleteHooks []PurchasedPostageHook
var purchasedPostageAfterUpsertHooks []PurchasedPostageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PurchasedPostage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PurchasedPostage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PurchasedPostage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PurchasedPostage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PurchasedPostage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PurchasedPostage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PurchasedPostage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PurchasedPostage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PurchasedPostage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchasedPostageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPurchasedPostageHook registers your hook function for all future operations.
func AddPurchasedPostageHook(hookPoint boil.HookPoint, purchasedPostageHook PurchasedPostageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		purchasedPostageBeforeInsertHooks = append(purchasedPostageBeforeInsertHooks, purchasedPostageHook)
	case boil.BeforeUpdateHook:
		purchasedPostageBeforeUpdateHooks = append(purchasedPostageBeforeUpdateHooks, purchasedPostageHook)
	case boil.BeforeDeleteHook:
		purchasedPostageBeforeDeleteHooks = append(purchasedPostageBeforeDeleteHooks, purchasedPostageHook)
	case boil.BeforeUpsertHook:
		purchasedPostageBeforeUpsertHooks = append(purchasedPostageBeforeUpsertHooks, purchasedPostageHook)
	case boil.AfterInsertHook:
		purchasedPostageAfterInsertHooks = append(purchasedPostageAfterInsertHooks, purchasedPostageHook)
	case boil.AfterSelectHook:
		purchasedPostageAfterSelectHooks = append(purchasedPostageAfterSelectHooks, purchasedPostageHook)
	case boil.AfterUpdateHook:
		purchasedPostageAfterUpdateHooks = append(purchasedPostageAfterUpdateHooks, purchasedPostageHook)
	case boil.AfterDeleteHook:
		purchasedPostageAfterDeleteHooks = append(purchasedPostageAfterDeleteHooks, purchasedPostageHook)
	case boil.AfterUpsertHook:
		purchasedPostageAfterUpsertHooks = append(purchasedPostageAfterUpsertHooks, purchasedPostageHook)
	}
}

// One returns a single purchasedPostage record from the query.
func (q purchasedPostageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PurchasedPostage, error) {
	o := &PurchasedPostage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for purchased_postages")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PurchasedPostage records from the query.
func (q purchasedPostageQuery) All(ctx context.Context, exec boil.ContextExecutor) (PurchasedPostageSlice, error) {
	var o []*PurchasedPostage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PurchasedPostage slice")
	}

	if len(purchasedPostageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PurchasedPostage records in the query.
func (q purchasedPostageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count purchased_postages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q purchasedPostageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if purchased_postages exists")
	}

	return count > 0, nil
}

// Document pointed to by the foreign key.
func (o *PurchasedPostage) Document(mods ...qm.QueryMod) documentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.DocumentID),
	}

	queryMods = append(queryMods, mods...)

	query := Documents(queryMods...)
	queries.SetFrom(query.Query, "`documents`")

	return query
}

// LoadDocument allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (purchasedPostageL) LoadDocument(ctx context.Context, e boil.ContextExecutor, singular bool, maybePurchasedPostage interface{}, mods queries.Applicator) error {
	var slice []*PurchasedPostage
	var object *PurchasedPostage

	if singular {
		object = maybePurchasedPostage.(*PurchasedPostage)
	} else {
		slice = *maybePurchasedPostage.(*[]*PurchasedPostage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &purchasedPostageR{}
		}
		args = append(args, object.DocumentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &purchasedPostageR{}
			}

			for _, a := range args {
				if a == obj.DocumentID {
					continue Outer
				}
			}

			args = append(args, obj.DocumentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`documents`),
		qm.WhereIn(`documents.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Document")
	}

	var resultSlice []*Document
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Document")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for documents")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for documents")
	}

	if len(purchasedPostageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Document = foreign
		if foreign.R == nil {
			foreign.R = &documentR{}
		}
		foreign.R.PurchasedPostages = append(foreign.R.PurchasedPostages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DocumentID == foreign.ID {
				local.R.Document = foreign
				if foreign.R == nil {
					foreign.R = &documentR{}
				}
				foreign.R.PurchasedPostages = append(foreign.R.PurchasedPostages, local)
				break
			}
		}
	}

	return nil
}

// SetDocument of the purchasedPostage to the related item.
// Sets o.R.Document to related.
// Adds o to related.R.PurchasedPostages.
func (o *PurchasedPostage) SetDocument(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Document) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `purchased_postages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"document_id"}),
		strmangle.WhereClause("`", "`", 0, purchasedPostagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DocumentID = related.ID
	if o.R == nil {
		o.R = &purchasedPostageR{
			Document: related,
		}
	} else {
		o.R.Document = related
	}

	if related.R == nil {
		related.R = &documentR{
			PurchasedPostages: PurchasedPostageSlice{o},
		}
	} else {
		related.R.PurchasedPostages = append(related.R.PurchasedPostages, o)
	}

	return nil
}

// PurchasedPostages retrieves all the records using an executor.
func PurchasedPostages(mods ...qm.QueryMod) purchasedPostageQuery {
	mods = append(mods, qm.From("`purchased_postages`"))
	return purchasedPostageQuery{NewQuery(mods...)}
}

// FindPurchasedPostage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPurchasedPostage(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*PurchasedPostage, error) {
	purchasedPostageObj := &PurchasedPostage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `purchased_postages` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, purchasedPostageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from purchased_postages")
	}

	if err = purchasedPostageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return purchasedPostageObj, err
	}

	return purchasedPostageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PurchasedPostage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no purchased_postages provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(purchasedPostageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	purchasedPostageInsertCacheMut.RLock()
	cache, cached := purchasedPostageInsertCache[key]
	purchasedPostageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			purchasedPostageAllColumns,
			purchasedPostageColumnsWithDefault,
			purchasedPostageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(purchasedPostageType, purchasedPostageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(purchasedPostageType, purchasedPostageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `purchased_postages` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `purchased_postages` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `purchased_postages` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, purchasedPostagePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into purchased_postages")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == purchasedPostageMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for purchased_postages")
	}

CacheNoHooks:
	if !cached {
		purchasedPostageInsertCacheMut.Lock()
		purchasedPostageInsertCache[key] = cache
		purchasedPostageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PurchasedPostage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PurchasedPostage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	purchasedPostageUpdateCacheMut.RLock()
	cache, cached := purchasedPostageUpdateCache[key]
	purchasedPostageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			purchasedPostageAllColumns,
			purchasedPostagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update purchased_postages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `purchased_postages` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, purchasedPostagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(purchasedPostageType, purchasedPostageMapping, append(wl, purchasedPostagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update purchased_postages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for purchased_postages")
	}

	if !cached {
		purchasedPostageUpdateCacheMut.Lock()
		purchasedPostageUpdateCache[key] = cache
		purchasedPostageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q purchasedPostageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for purchased_postages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for purchased_postages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PurchasedPostageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), purchasedPostagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `purchased_postages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, purchasedPostagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in purchasedPostage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all purchasedPostage")
	}
	return rowsAff, nil
}

var mySQLPurchasedPostageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PurchasedPostage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no purchased_postages provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(purchasedPostageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPurchasedPostageUniqueColumns, o)

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

	purchasedPostageUpsertCacheMut.RLock()
	cache, cached := purchasedPostageUpsertCache[key]
	purchasedPostageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			purchasedPostageAllColumns,
			purchasedPostageColumnsWithDefault,
			purchasedPostageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			purchasedPostageAllColumns,
			purchasedPostagePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert purchased_postages, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`purchased_postages`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `purchased_postages` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(purchasedPostageType, purchasedPostageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(purchasedPostageType, purchasedPostageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for purchased_postages")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == purchasedPostageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(purchasedPostageType, purchasedPostageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for purchased_postages")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for purchased_postages")
	}

CacheNoHooks:
	if !cached {
		purchasedPostageUpsertCacheMut.Lock()
		purchasedPostageUpsertCache[key] = cache
		purchasedPostageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PurchasedPostage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PurchasedPostage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PurchasedPostage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), purchasedPostagePrimaryKeyMapping)
	sql := "DELETE FROM `purchased_postages` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from purchased_postages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for purchased_postages")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q purchasedPostageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no purchasedPostageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from purchased_postages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for purchased_postages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PurchasedPostageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(purchasedPostageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), purchasedPostagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `purchased_postages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, purchasedPostagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from purchasedPostage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for purchased_postages")
	}

	if len(purchasedPostageAfterDeleteHooks) != 0 {
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
func (o *PurchasedPostage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPurchasedPostage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PurchasedPostageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PurchasedPostageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), purchasedPostagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `purchased_postages`.* FROM `purchased_postages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, purchasedPostagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PurchasedPostageSlice")
	}

	*o = slice

	return nil
}

// PurchasedPostageExists checks if the PurchasedPostage row exists.
func PurchasedPostageExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `purchased_postages` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if purchased_postages exists")
	}

	return exists, nil
}
