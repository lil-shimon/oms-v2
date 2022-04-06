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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// FailedJob is an object representing the database table.
type FailedJob struct {
	ID         uint64    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Connection string    `boil:"connection" json:"connection" toml:"connection" yaml:"connection"`
	Queue      string    `boil:"queue" json:"queue" toml:"queue" yaml:"queue"`
	Payload    string    `boil:"payload" json:"payload" toml:"payload" yaml:"payload"`
	Exception  string    `boil:"exception" json:"exception" toml:"exception" yaml:"exception"`
	FailedAt   time.Time `boil:"failed_at" json:"failed_at" toml:"failed_at" yaml:"failed_at"`

	R *failedJobR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L failedJobL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FailedJobColumns = struct {
	ID         string
	Connection string
	Queue      string
	Payload    string
	Exception  string
	FailedAt   string
}{
	ID:         "id",
	Connection: "connection",
	Queue:      "queue",
	Payload:    "payload",
	Exception:  "exception",
	FailedAt:   "failed_at",
}

var FailedJobTableColumns = struct {
	ID         string
	Connection string
	Queue      string
	Payload    string
	Exception  string
	FailedAt   string
}{
	ID:         "failed_jobs.id",
	Connection: "failed_jobs.connection",
	Queue:      "failed_jobs.queue",
	Payload:    "failed_jobs.payload",
	Exception:  "failed_jobs.exception",
	FailedAt:   "failed_jobs.failed_at",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var FailedJobWhere = struct {
	ID         whereHelperuint64
	Connection whereHelperstring
	Queue      whereHelperstring
	Payload    whereHelperstring
	Exception  whereHelperstring
	FailedAt   whereHelpertime_Time
}{
	ID:         whereHelperuint64{field: "`failed_jobs`.`id`"},
	Connection: whereHelperstring{field: "`failed_jobs`.`connection`"},
	Queue:      whereHelperstring{field: "`failed_jobs`.`queue`"},
	Payload:    whereHelperstring{field: "`failed_jobs`.`payload`"},
	Exception:  whereHelperstring{field: "`failed_jobs`.`exception`"},
	FailedAt:   whereHelpertime_Time{field: "`failed_jobs`.`failed_at`"},
}

// FailedJobRels is where relationship names are stored.
var FailedJobRels = struct {
}{}

// failedJobR is where relationships are stored.
type failedJobR struct {
}

// NewStruct creates a new relationship struct
func (*failedJobR) NewStruct() *failedJobR {
	return &failedJobR{}
}

// failedJobL is where Load methods for each relationship are stored.
type failedJobL struct{}

var (
	failedJobAllColumns            = []string{"id", "connection", "queue", "payload", "exception", "failed_at"}
	failedJobColumnsWithoutDefault = []string{"connection", "queue", "payload", "exception"}
	failedJobColumnsWithDefault    = []string{"id", "failed_at"}
	failedJobPrimaryKeyColumns     = []string{"id"}
)

type (
	// FailedJobSlice is an alias for a slice of pointers to FailedJob.
	// This should almost always be used instead of []FailedJob.
	FailedJobSlice []*FailedJob
	// FailedJobHook is the signature for custom FailedJob hook methods
	FailedJobHook func(context.Context, boil.ContextExecutor, *FailedJob) error

	failedJobQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	failedJobType                 = reflect.TypeOf(&FailedJob{})
	failedJobMapping              = queries.MakeStructMapping(failedJobType)
	failedJobPrimaryKeyMapping, _ = queries.BindMapping(failedJobType, failedJobMapping, failedJobPrimaryKeyColumns)
	failedJobInsertCacheMut       sync.RWMutex
	failedJobInsertCache          = make(map[string]insertCache)
	failedJobUpdateCacheMut       sync.RWMutex
	failedJobUpdateCache          = make(map[string]updateCache)
	failedJobUpsertCacheMut       sync.RWMutex
	failedJobUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var failedJobBeforeInsertHooks []FailedJobHook
var failedJobBeforeUpdateHooks []FailedJobHook
var failedJobBeforeDeleteHooks []FailedJobHook
var failedJobBeforeUpsertHooks []FailedJobHook

var failedJobAfterInsertHooks []FailedJobHook
var failedJobAfterSelectHooks []FailedJobHook
var failedJobAfterUpdateHooks []FailedJobHook
var failedJobAfterDeleteHooks []FailedJobHook
var failedJobAfterUpsertHooks []FailedJobHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FailedJob) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FailedJob) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FailedJob) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FailedJob) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FailedJob) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FailedJob) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FailedJob) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FailedJob) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FailedJob) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range failedJobAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFailedJobHook registers your hook function for all future operations.
func AddFailedJobHook(hookPoint boil.HookPoint, failedJobHook FailedJobHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		failedJobBeforeInsertHooks = append(failedJobBeforeInsertHooks, failedJobHook)
	case boil.BeforeUpdateHook:
		failedJobBeforeUpdateHooks = append(failedJobBeforeUpdateHooks, failedJobHook)
	case boil.BeforeDeleteHook:
		failedJobBeforeDeleteHooks = append(failedJobBeforeDeleteHooks, failedJobHook)
	case boil.BeforeUpsertHook:
		failedJobBeforeUpsertHooks = append(failedJobBeforeUpsertHooks, failedJobHook)
	case boil.AfterInsertHook:
		failedJobAfterInsertHooks = append(failedJobAfterInsertHooks, failedJobHook)
	case boil.AfterSelectHook:
		failedJobAfterSelectHooks = append(failedJobAfterSelectHooks, failedJobHook)
	case boil.AfterUpdateHook:
		failedJobAfterUpdateHooks = append(failedJobAfterUpdateHooks, failedJobHook)
	case boil.AfterDeleteHook:
		failedJobAfterDeleteHooks = append(failedJobAfterDeleteHooks, failedJobHook)
	case boil.AfterUpsertHook:
		failedJobAfterUpsertHooks = append(failedJobAfterUpsertHooks, failedJobHook)
	}
}

// One returns a single failedJob record from the query.
func (q failedJobQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FailedJob, error) {
	o := &FailedJob{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for failed_jobs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FailedJob records from the query.
func (q failedJobQuery) All(ctx context.Context, exec boil.ContextExecutor) (FailedJobSlice, error) {
	var o []*FailedJob

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FailedJob slice")
	}

	if len(failedJobAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FailedJob records in the query.
func (q failedJobQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count failed_jobs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q failedJobQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if failed_jobs exists")
	}

	return count > 0, nil
}

// FailedJobs retrieves all the records using an executor.
func FailedJobs(mods ...qm.QueryMod) failedJobQuery {
	mods = append(mods, qm.From("`failed_jobs`"))
	return failedJobQuery{NewQuery(mods...)}
}

// FindFailedJob retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFailedJob(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*FailedJob, error) {
	failedJobObj := &FailedJob{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `failed_jobs` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, failedJobObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from failed_jobs")
	}

	if err = failedJobObj.doAfterSelectHooks(ctx, exec); err != nil {
		return failedJobObj, err
	}

	return failedJobObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FailedJob) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no failed_jobs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(failedJobColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	failedJobInsertCacheMut.RLock()
	cache, cached := failedJobInsertCache[key]
	failedJobInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			failedJobAllColumns,
			failedJobColumnsWithDefault,
			failedJobColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(failedJobType, failedJobMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(failedJobType, failedJobMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `failed_jobs` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `failed_jobs` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `failed_jobs` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, failedJobPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into failed_jobs")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == failedJobMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for failed_jobs")
	}

CacheNoHooks:
	if !cached {
		failedJobInsertCacheMut.Lock()
		failedJobInsertCache[key] = cache
		failedJobInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the FailedJob.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FailedJob) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	failedJobUpdateCacheMut.RLock()
	cache, cached := failedJobUpdateCache[key]
	failedJobUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			failedJobAllColumns,
			failedJobPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update failed_jobs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `failed_jobs` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, failedJobPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(failedJobType, failedJobMapping, append(wl, failedJobPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update failed_jobs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for failed_jobs")
	}

	if !cached {
		failedJobUpdateCacheMut.Lock()
		failedJobUpdateCache[key] = cache
		failedJobUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q failedJobQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for failed_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for failed_jobs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FailedJobSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), failedJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `failed_jobs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, failedJobPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in failedJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all failedJob")
	}
	return rowsAff, nil
}

var mySQLFailedJobUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *FailedJob) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no failed_jobs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(failedJobColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFailedJobUniqueColumns, o)

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

	failedJobUpsertCacheMut.RLock()
	cache, cached := failedJobUpsertCache[key]
	failedJobUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			failedJobAllColumns,
			failedJobColumnsWithDefault,
			failedJobColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			failedJobAllColumns,
			failedJobPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert failed_jobs, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`failed_jobs`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `failed_jobs` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(failedJobType, failedJobMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(failedJobType, failedJobMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for failed_jobs")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == failedJobMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(failedJobType, failedJobMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for failed_jobs")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for failed_jobs")
	}

CacheNoHooks:
	if !cached {
		failedJobUpsertCacheMut.Lock()
		failedJobUpsertCache[key] = cache
		failedJobUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single FailedJob record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FailedJob) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FailedJob provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), failedJobPrimaryKeyMapping)
	sql := "DELETE FROM `failed_jobs` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from failed_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for failed_jobs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q failedJobQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no failedJobQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from failed_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for failed_jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FailedJobSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(failedJobBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), failedJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `failed_jobs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, failedJobPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from failedJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for failed_jobs")
	}

	if len(failedJobAfterDeleteHooks) != 0 {
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
func (o *FailedJob) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFailedJob(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FailedJobSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FailedJobSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), failedJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `failed_jobs`.* FROM `failed_jobs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, failedJobPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FailedJobSlice")
	}

	*o = slice

	return nil
}

// FailedJobExists checks if the FailedJob row exists.
func FailedJobExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `failed_jobs` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if failed_jobs exists")
	}

	return exists, nil
}