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

// Maintenance is an object representing the database table.
type Maintenance struct {
	ID           uint         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Date         null.Time    `boil:"date" json:"date,omitempty" toml:"date" yaml:"date,omitempty"`
	Title        null.String  `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Address      null.String  `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Note         null.String  `boil:"note" json:"note,omitempty" toml:"note" yaml:"note,omitempty"`
	WorkingHours null.Float64 `boil:"working_hours" json:"working_hours,omitempty" toml:"working_hours" yaml:"working_hours,omitempty"`
	DeletedAt    null.Time    `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	CreatedAt    null.Time    `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt    null.Time    `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Count        int          `boil:"count" json:"count" toml:"count" yaml:"count"`

	R *maintenanceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L maintenanceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MaintenanceColumns = struct {
	ID           string
	Date         string
	Title        string
	Address      string
	Note         string
	WorkingHours string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
	Count        string
}{
	ID:           "id",
	Date:         "date",
	Title:        "title",
	Address:      "address",
	Note:         "note",
	WorkingHours: "working_hours",
	DeletedAt:    "deleted_at",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Count:        "count",
}

var MaintenanceTableColumns = struct {
	ID           string
	Date         string
	Title        string
	Address      string
	Note         string
	WorkingHours string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
	Count        string
}{
	ID:           "maintenances.id",
	Date:         "maintenances.date",
	Title:        "maintenances.title",
	Address:      "maintenances.address",
	Note:         "maintenances.note",
	WorkingHours: "maintenances.working_hours",
	DeletedAt:    "maintenances.deleted_at",
	CreatedAt:    "maintenances.created_at",
	UpdatedAt:    "maintenances.updated_at",
	Count:        "maintenances.count",
}

// Generated where

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MaintenanceWhere = struct {
	ID           whereHelperuint
	Date         whereHelpernull_Time
	Title        whereHelpernull_String
	Address      whereHelpernull_String
	Note         whereHelpernull_String
	WorkingHours whereHelpernull_Float64
	DeletedAt    whereHelpernull_Time
	CreatedAt    whereHelpernull_Time
	UpdatedAt    whereHelpernull_Time
	Count        whereHelperint
}{
	ID:           whereHelperuint{field: "`maintenances`.`id`"},
	Date:         whereHelpernull_Time{field: "`maintenances`.`date`"},
	Title:        whereHelpernull_String{field: "`maintenances`.`title`"},
	Address:      whereHelpernull_String{field: "`maintenances`.`address`"},
	Note:         whereHelpernull_String{field: "`maintenances`.`note`"},
	WorkingHours: whereHelpernull_Float64{field: "`maintenances`.`working_hours`"},
	DeletedAt:    whereHelpernull_Time{field: "`maintenances`.`deleted_at`"},
	CreatedAt:    whereHelpernull_Time{field: "`maintenances`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`maintenances`.`updated_at`"},
	Count:        whereHelperint{field: "`maintenances`.`count`"},
}

// MaintenanceRels is where relationship names are stored.
var MaintenanceRels = struct {
}{}

// maintenanceR is where relationships are stored.
type maintenanceR struct {
}

// NewStruct creates a new relationship struct
func (*maintenanceR) NewStruct() *maintenanceR {
	return &maintenanceR{}
}

// maintenanceL is where Load methods for each relationship are stored.
type maintenanceL struct{}

var (
	maintenanceAllColumns            = []string{"id", "date", "title", "address", "note", "working_hours", "deleted_at", "created_at", "updated_at", "count"}
	maintenanceColumnsWithoutDefault = []string{"date", "title", "address", "note", "working_hours", "deleted_at", "created_at", "updated_at", "count"}
	maintenanceColumnsWithDefault    = []string{"id"}
	maintenancePrimaryKeyColumns     = []string{"id"}
)

type (
	// MaintenanceSlice is an alias for a slice of pointers to Maintenance.
	// This should almost always be used instead of []Maintenance.
	MaintenanceSlice []*Maintenance
	// MaintenanceHook is the signature for custom Maintenance hook methods
	MaintenanceHook func(context.Context, boil.ContextExecutor, *Maintenance) error

	maintenanceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	maintenanceType                 = reflect.TypeOf(&Maintenance{})
	maintenanceMapping              = queries.MakeStructMapping(maintenanceType)
	maintenancePrimaryKeyMapping, _ = queries.BindMapping(maintenanceType, maintenanceMapping, maintenancePrimaryKeyColumns)
	maintenanceInsertCacheMut       sync.RWMutex
	maintenanceInsertCache          = make(map[string]insertCache)
	maintenanceUpdateCacheMut       sync.RWMutex
	maintenanceUpdateCache          = make(map[string]updateCache)
	maintenanceUpsertCacheMut       sync.RWMutex
	maintenanceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var maintenanceBeforeInsertHooks []MaintenanceHook
var maintenanceBeforeUpdateHooks []MaintenanceHook
var maintenanceBeforeDeleteHooks []MaintenanceHook
var maintenanceBeforeUpsertHooks []MaintenanceHook

var maintenanceAfterInsertHooks []MaintenanceHook
var maintenanceAfterSelectHooks []MaintenanceHook
var maintenanceAfterUpdateHooks []MaintenanceHook
var maintenanceAfterDeleteHooks []MaintenanceHook
var maintenanceAfterUpsertHooks []MaintenanceHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Maintenance) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Maintenance) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Maintenance) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Maintenance) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Maintenance) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Maintenance) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Maintenance) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Maintenance) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Maintenance) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range maintenanceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMaintenanceHook registers your hook function for all future operations.
func AddMaintenanceHook(hookPoint boil.HookPoint, maintenanceHook MaintenanceHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		maintenanceBeforeInsertHooks = append(maintenanceBeforeInsertHooks, maintenanceHook)
	case boil.BeforeUpdateHook:
		maintenanceBeforeUpdateHooks = append(maintenanceBeforeUpdateHooks, maintenanceHook)
	case boil.BeforeDeleteHook:
		maintenanceBeforeDeleteHooks = append(maintenanceBeforeDeleteHooks, maintenanceHook)
	case boil.BeforeUpsertHook:
		maintenanceBeforeUpsertHooks = append(maintenanceBeforeUpsertHooks, maintenanceHook)
	case boil.AfterInsertHook:
		maintenanceAfterInsertHooks = append(maintenanceAfterInsertHooks, maintenanceHook)
	case boil.AfterSelectHook:
		maintenanceAfterSelectHooks = append(maintenanceAfterSelectHooks, maintenanceHook)
	case boil.AfterUpdateHook:
		maintenanceAfterUpdateHooks = append(maintenanceAfterUpdateHooks, maintenanceHook)
	case boil.AfterDeleteHook:
		maintenanceAfterDeleteHooks = append(maintenanceAfterDeleteHooks, maintenanceHook)
	case boil.AfterUpsertHook:
		maintenanceAfterUpsertHooks = append(maintenanceAfterUpsertHooks, maintenanceHook)
	}
}

// One returns a single maintenance record from the query.
func (q maintenanceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Maintenance, error) {
	o := &Maintenance{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for maintenances")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Maintenance records from the query.
func (q maintenanceQuery) All(ctx context.Context, exec boil.ContextExecutor) (MaintenanceSlice, error) {
	var o []*Maintenance

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Maintenance slice")
	}

	if len(maintenanceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Maintenance records in the query.
func (q maintenanceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count maintenances rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q maintenanceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if maintenances exists")
	}

	return count > 0, nil
}

// Maintenances retrieves all the records using an executor.
func Maintenances(mods ...qm.QueryMod) maintenanceQuery {
	mods = append(mods, qm.From("`maintenances`"))
	return maintenanceQuery{NewQuery(mods...)}
}

// FindMaintenance retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMaintenance(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Maintenance, error) {
	maintenanceObj := &Maintenance{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `maintenances` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, maintenanceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from maintenances")
	}

	if err = maintenanceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return maintenanceObj, err
	}

	return maintenanceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Maintenance) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no maintenances provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(maintenanceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	maintenanceInsertCacheMut.RLock()
	cache, cached := maintenanceInsertCache[key]
	maintenanceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			maintenanceAllColumns,
			maintenanceColumnsWithDefault,
			maintenanceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(maintenanceType, maintenanceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(maintenanceType, maintenanceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `maintenances` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `maintenances` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `maintenances` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, maintenancePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into maintenances")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == maintenanceMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for maintenances")
	}

CacheNoHooks:
	if !cached {
		maintenanceInsertCacheMut.Lock()
		maintenanceInsertCache[key] = cache
		maintenanceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Maintenance.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Maintenance) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	maintenanceUpdateCacheMut.RLock()
	cache, cached := maintenanceUpdateCache[key]
	maintenanceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			maintenanceAllColumns,
			maintenancePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update maintenances, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `maintenances` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, maintenancePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(maintenanceType, maintenanceMapping, append(wl, maintenancePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update maintenances row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for maintenances")
	}

	if !cached {
		maintenanceUpdateCacheMut.Lock()
		maintenanceUpdateCache[key] = cache
		maintenanceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q maintenanceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for maintenances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for maintenances")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MaintenanceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), maintenancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `maintenances` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, maintenancePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in maintenance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all maintenance")
	}
	return rowsAff, nil
}

var mySQLMaintenanceUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Maintenance) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no maintenances provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(maintenanceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMaintenanceUniqueColumns, o)

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

	maintenanceUpsertCacheMut.RLock()
	cache, cached := maintenanceUpsertCache[key]
	maintenanceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			maintenanceAllColumns,
			maintenanceColumnsWithDefault,
			maintenanceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			maintenanceAllColumns,
			maintenancePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert maintenances, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`maintenances`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `maintenances` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(maintenanceType, maintenanceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(maintenanceType, maintenanceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for maintenances")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == maintenanceMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(maintenanceType, maintenanceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for maintenances")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for maintenances")
	}

CacheNoHooks:
	if !cached {
		maintenanceUpsertCacheMut.Lock()
		maintenanceUpsertCache[key] = cache
		maintenanceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Maintenance record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Maintenance) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Maintenance provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), maintenancePrimaryKeyMapping)
	sql := "DELETE FROM `maintenances` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from maintenances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for maintenances")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q maintenanceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no maintenanceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from maintenances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for maintenances")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MaintenanceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(maintenanceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), maintenancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `maintenances` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, maintenancePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from maintenance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for maintenances")
	}

	if len(maintenanceAfterDeleteHooks) != 0 {
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
func (o *Maintenance) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMaintenance(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MaintenanceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MaintenanceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), maintenancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `maintenances`.* FROM `maintenances` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, maintenancePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MaintenanceSlice")
	}

	*o = slice

	return nil
}

// MaintenanceExists checks if the Maintenance row exists.
func MaintenanceExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `maintenances` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if maintenances exists")
	}

	return exists, nil
}
