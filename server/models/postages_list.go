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

// PostagesList is an object representing the database table.
type PostagesList struct {
	ID               uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	SenderPlace      string    `boil:"sender_place" json:"sender_place" toml:"sender_place" yaml:"sender_place"`
	DestinationPlace string    `boil:"destination_place" json:"destination_place" toml:"destination_place" yaml:"destination_place"`
	PostagePrice     int       `boil:"postage_price" json:"postage_price" toml:"postage_price" yaml:"postage_price"`
	Tax              float64   `boil:"tax" json:"tax" toml:"tax" yaml:"tax"`
	Size             string    `boil:"size" json:"size" toml:"size" yaml:"size"`
	Unit             string    `boil:"unit" json:"unit" toml:"unit" yaml:"unit"`
	DeletedAt        null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	CreatedAt        null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt        null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *postagesListR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L postagesListL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PostagesListColumns = struct {
	ID               string
	SenderPlace      string
	DestinationPlace string
	PostagePrice     string
	Tax              string
	Size             string
	Unit             string
	DeletedAt        string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	SenderPlace:      "sender_place",
	DestinationPlace: "destination_place",
	PostagePrice:     "postage_price",
	Tax:              "tax",
	Size:             "size",
	Unit:             "unit",
	DeletedAt:        "deleted_at",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

var PostagesListTableColumns = struct {
	ID               string
	SenderPlace      string
	DestinationPlace string
	PostagePrice     string
	Tax              string
	Size             string
	Unit             string
	DeletedAt        string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "postages_list.id",
	SenderPlace:      "postages_list.sender_place",
	DestinationPlace: "postages_list.destination_place",
	PostagePrice:     "postages_list.postage_price",
	Tax:              "postages_list.tax",
	Size:             "postages_list.size",
	Unit:             "postages_list.unit",
	DeletedAt:        "postages_list.deleted_at",
	CreatedAt:        "postages_list.created_at",
	UpdatedAt:        "postages_list.updated_at",
}

// Generated where

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var PostagesListWhere = struct {
	ID               whereHelperuint
	SenderPlace      whereHelperstring
	DestinationPlace whereHelperstring
	PostagePrice     whereHelperint
	Tax              whereHelperfloat64
	Size             whereHelperstring
	Unit             whereHelperstring
	DeletedAt        whereHelpernull_Time
	CreatedAt        whereHelpernull_Time
	UpdatedAt        whereHelpernull_Time
}{
	ID:               whereHelperuint{field: "`postages_list`.`id`"},
	SenderPlace:      whereHelperstring{field: "`postages_list`.`sender_place`"},
	DestinationPlace: whereHelperstring{field: "`postages_list`.`destination_place`"},
	PostagePrice:     whereHelperint{field: "`postages_list`.`postage_price`"},
	Tax:              whereHelperfloat64{field: "`postages_list`.`tax`"},
	Size:             whereHelperstring{field: "`postages_list`.`size`"},
	Unit:             whereHelperstring{field: "`postages_list`.`unit`"},
	DeletedAt:        whereHelpernull_Time{field: "`postages_list`.`deleted_at`"},
	CreatedAt:        whereHelpernull_Time{field: "`postages_list`.`created_at`"},
	UpdatedAt:        whereHelpernull_Time{field: "`postages_list`.`updated_at`"},
}

// PostagesListRels is where relationship names are stored.
var PostagesListRels = struct {
}{}

// postagesListR is where relationships are stored.
type postagesListR struct {
}

// NewStruct creates a new relationship struct
func (*postagesListR) NewStruct() *postagesListR {
	return &postagesListR{}
}

// postagesListL is where Load methods for each relationship are stored.
type postagesListL struct{}

var (
	postagesListAllColumns            = []string{"id", "sender_place", "destination_place", "postage_price", "tax", "size", "unit", "deleted_at", "created_at", "updated_at"}
	postagesListColumnsWithoutDefault = []string{"sender_place", "destination_place", "postage_price", "tax", "size", "unit", "deleted_at", "created_at", "updated_at"}
	postagesListColumnsWithDefault    = []string{"id"}
	postagesListPrimaryKeyColumns     = []string{"id"}
)

type (
	// PostagesListSlice is an alias for a slice of pointers to PostagesList.
	// This should almost always be used instead of []PostagesList.
	PostagesListSlice []*PostagesList
	// PostagesListHook is the signature for custom PostagesList hook methods
	PostagesListHook func(context.Context, boil.ContextExecutor, *PostagesList) error

	postagesListQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	postagesListType                 = reflect.TypeOf(&PostagesList{})
	postagesListMapping              = queries.MakeStructMapping(postagesListType)
	postagesListPrimaryKeyMapping, _ = queries.BindMapping(postagesListType, postagesListMapping, postagesListPrimaryKeyColumns)
	postagesListInsertCacheMut       sync.RWMutex
	postagesListInsertCache          = make(map[string]insertCache)
	postagesListUpdateCacheMut       sync.RWMutex
	postagesListUpdateCache          = make(map[string]updateCache)
	postagesListUpsertCacheMut       sync.RWMutex
	postagesListUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var postagesListBeforeInsertHooks []PostagesListHook
var postagesListBeforeUpdateHooks []PostagesListHook
var postagesListBeforeDeleteHooks []PostagesListHook
var postagesListBeforeUpsertHooks []PostagesListHook

var postagesListAfterInsertHooks []PostagesListHook
var postagesListAfterSelectHooks []PostagesListHook
var postagesListAfterUpdateHooks []PostagesListHook
var postagesListAfterDeleteHooks []PostagesListHook
var postagesListAfterUpsertHooks []PostagesListHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PostagesList) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PostagesList) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PostagesList) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PostagesList) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PostagesList) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PostagesList) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PostagesList) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PostagesList) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PostagesList) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postagesListAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPostagesListHook registers your hook function for all future operations.
func AddPostagesListHook(hookPoint boil.HookPoint, postagesListHook PostagesListHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		postagesListBeforeInsertHooks = append(postagesListBeforeInsertHooks, postagesListHook)
	case boil.BeforeUpdateHook:
		postagesListBeforeUpdateHooks = append(postagesListBeforeUpdateHooks, postagesListHook)
	case boil.BeforeDeleteHook:
		postagesListBeforeDeleteHooks = append(postagesListBeforeDeleteHooks, postagesListHook)
	case boil.BeforeUpsertHook:
		postagesListBeforeUpsertHooks = append(postagesListBeforeUpsertHooks, postagesListHook)
	case boil.AfterInsertHook:
		postagesListAfterInsertHooks = append(postagesListAfterInsertHooks, postagesListHook)
	case boil.AfterSelectHook:
		postagesListAfterSelectHooks = append(postagesListAfterSelectHooks, postagesListHook)
	case boil.AfterUpdateHook:
		postagesListAfterUpdateHooks = append(postagesListAfterUpdateHooks, postagesListHook)
	case boil.AfterDeleteHook:
		postagesListAfterDeleteHooks = append(postagesListAfterDeleteHooks, postagesListHook)
	case boil.AfterUpsertHook:
		postagesListAfterUpsertHooks = append(postagesListAfterUpsertHooks, postagesListHook)
	}
}

// One returns a single postagesList record from the query.
func (q postagesListQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PostagesList, error) {
	o := &PostagesList{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for postages_list")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PostagesList records from the query.
func (q postagesListQuery) All(ctx context.Context, exec boil.ContextExecutor) (PostagesListSlice, error) {
	var o []*PostagesList

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PostagesList slice")
	}

	if len(postagesListAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PostagesList records in the query.
func (q postagesListQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count postages_list rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q postagesListQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if postages_list exists")
	}

	return count > 0, nil
}

// PostagesLists retrieves all the records using an executor.
func PostagesLists(mods ...qm.QueryMod) postagesListQuery {
	mods = append(mods, qm.From("`postages_list`"))
	return postagesListQuery{NewQuery(mods...)}
}

// FindPostagesList retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPostagesList(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*PostagesList, error) {
	postagesListObj := &PostagesList{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `postages_list` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, postagesListObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from postages_list")
	}

	if err = postagesListObj.doAfterSelectHooks(ctx, exec); err != nil {
		return postagesListObj, err
	}

	return postagesListObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PostagesList) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no postages_list provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(postagesListColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	postagesListInsertCacheMut.RLock()
	cache, cached := postagesListInsertCache[key]
	postagesListInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			postagesListAllColumns,
			postagesListColumnsWithDefault,
			postagesListColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(postagesListType, postagesListMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(postagesListType, postagesListMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `postages_list` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `postages_list` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `postages_list` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, postagesListPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into postages_list")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == postagesListMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for postages_list")
	}

CacheNoHooks:
	if !cached {
		postagesListInsertCacheMut.Lock()
		postagesListInsertCache[key] = cache
		postagesListInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PostagesList.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PostagesList) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	postagesListUpdateCacheMut.RLock()
	cache, cached := postagesListUpdateCache[key]
	postagesListUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			postagesListAllColumns,
			postagesListPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update postages_list, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `postages_list` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, postagesListPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(postagesListType, postagesListMapping, append(wl, postagesListPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update postages_list row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for postages_list")
	}

	if !cached {
		postagesListUpdateCacheMut.Lock()
		postagesListUpdateCache[key] = cache
		postagesListUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q postagesListQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for postages_list")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for postages_list")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PostagesListSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postagesListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `postages_list` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, postagesListPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in postagesList slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all postagesList")
	}
	return rowsAff, nil
}

var mySQLPostagesListUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PostagesList) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no postages_list provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(postagesListColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPostagesListUniqueColumns, o)

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

	postagesListUpsertCacheMut.RLock()
	cache, cached := postagesListUpsertCache[key]
	postagesListUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			postagesListAllColumns,
			postagesListColumnsWithDefault,
			postagesListColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			postagesListAllColumns,
			postagesListPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert postages_list, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`postages_list`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `postages_list` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(postagesListType, postagesListMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(postagesListType, postagesListMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for postages_list")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == postagesListMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(postagesListType, postagesListMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for postages_list")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for postages_list")
	}

CacheNoHooks:
	if !cached {
		postagesListUpsertCacheMut.Lock()
		postagesListUpsertCache[key] = cache
		postagesListUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PostagesList record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PostagesList) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PostagesList provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), postagesListPrimaryKeyMapping)
	sql := "DELETE FROM `postages_list` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from postages_list")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for postages_list")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q postagesListQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no postagesListQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from postages_list")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for postages_list")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PostagesListSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(postagesListBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postagesListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `postages_list` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, postagesListPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from postagesList slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for postages_list")
	}

	if len(postagesListAfterDeleteHooks) != 0 {
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
func (o *PostagesList) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPostagesList(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PostagesListSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PostagesListSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postagesListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `postages_list`.* FROM `postages_list` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, postagesListPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PostagesListSlice")
	}

	*o = slice

	return nil
}

// PostagesListExists checks if the PostagesList row exists.
func PostagesListExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `postages_list` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if postages_list exists")
	}

	return exists, nil
}
