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

// DocumentsType is an object representing the database table.
type DocumentsType struct {
	ID   uint   `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *documentsTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L documentsTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DocumentsTypeColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var DocumentsTypeTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "documents_types.id",
	Name: "documents_types.name",
}

// Generated where

var DocumentsTypeWhere = struct {
	ID   whereHelperuint
	Name whereHelperstring
}{
	ID:   whereHelperuint{field: "`documents_types`.`id`"},
	Name: whereHelperstring{field: "`documents_types`.`name`"},
}

// DocumentsTypeRels is where relationship names are stored.
var DocumentsTypeRels = struct {
}{}

// documentsTypeR is where relationships are stored.
type documentsTypeR struct {
}

// NewStruct creates a new relationship struct
func (*documentsTypeR) NewStruct() *documentsTypeR {
	return &documentsTypeR{}
}

// documentsTypeL is where Load methods for each relationship are stored.
type documentsTypeL struct{}

var (
	documentsTypeAllColumns            = []string{"id", "name"}
	documentsTypeColumnsWithoutDefault = []string{"name"}
	documentsTypeColumnsWithDefault    = []string{"id"}
	documentsTypePrimaryKeyColumns     = []string{"id"}
)

type (
	// DocumentsTypeSlice is an alias for a slice of pointers to DocumentsType.
	// This should almost always be used instead of []DocumentsType.
	DocumentsTypeSlice []*DocumentsType
	// DocumentsTypeHook is the signature for custom DocumentsType hook methods
	DocumentsTypeHook func(context.Context, boil.ContextExecutor, *DocumentsType) error

	documentsTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	documentsTypeType                 = reflect.TypeOf(&DocumentsType{})
	documentsTypeMapping              = queries.MakeStructMapping(documentsTypeType)
	documentsTypePrimaryKeyMapping, _ = queries.BindMapping(documentsTypeType, documentsTypeMapping, documentsTypePrimaryKeyColumns)
	documentsTypeInsertCacheMut       sync.RWMutex
	documentsTypeInsertCache          = make(map[string]insertCache)
	documentsTypeUpdateCacheMut       sync.RWMutex
	documentsTypeUpdateCache          = make(map[string]updateCache)
	documentsTypeUpsertCacheMut       sync.RWMutex
	documentsTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var documentsTypeBeforeInsertHooks []DocumentsTypeHook
var documentsTypeBeforeUpdateHooks []DocumentsTypeHook
var documentsTypeBeforeDeleteHooks []DocumentsTypeHook
var documentsTypeBeforeUpsertHooks []DocumentsTypeHook

var documentsTypeAfterInsertHooks []DocumentsTypeHook
var documentsTypeAfterSelectHooks []DocumentsTypeHook
var documentsTypeAfterUpdateHooks []DocumentsTypeHook
var documentsTypeAfterDeleteHooks []DocumentsTypeHook
var documentsTypeAfterUpsertHooks []DocumentsTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DocumentsType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DocumentsType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DocumentsType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DocumentsType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DocumentsType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DocumentsType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DocumentsType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DocumentsType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DocumentsType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentsTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDocumentsTypeHook registers your hook function for all future operations.
func AddDocumentsTypeHook(hookPoint boil.HookPoint, documentsTypeHook DocumentsTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		documentsTypeBeforeInsertHooks = append(documentsTypeBeforeInsertHooks, documentsTypeHook)
	case boil.BeforeUpdateHook:
		documentsTypeBeforeUpdateHooks = append(documentsTypeBeforeUpdateHooks, documentsTypeHook)
	case boil.BeforeDeleteHook:
		documentsTypeBeforeDeleteHooks = append(documentsTypeBeforeDeleteHooks, documentsTypeHook)
	case boil.BeforeUpsertHook:
		documentsTypeBeforeUpsertHooks = append(documentsTypeBeforeUpsertHooks, documentsTypeHook)
	case boil.AfterInsertHook:
		documentsTypeAfterInsertHooks = append(documentsTypeAfterInsertHooks, documentsTypeHook)
	case boil.AfterSelectHook:
		documentsTypeAfterSelectHooks = append(documentsTypeAfterSelectHooks, documentsTypeHook)
	case boil.AfterUpdateHook:
		documentsTypeAfterUpdateHooks = append(documentsTypeAfterUpdateHooks, documentsTypeHook)
	case boil.AfterDeleteHook:
		documentsTypeAfterDeleteHooks = append(documentsTypeAfterDeleteHooks, documentsTypeHook)
	case boil.AfterUpsertHook:
		documentsTypeAfterUpsertHooks = append(documentsTypeAfterUpsertHooks, documentsTypeHook)
	}
}

// One returns a single documentsType record from the query.
func (q documentsTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DocumentsType, error) {
	o := &DocumentsType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for documents_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DocumentsType records from the query.
func (q documentsTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (DocumentsTypeSlice, error) {
	var o []*DocumentsType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DocumentsType slice")
	}

	if len(documentsTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DocumentsType records in the query.
func (q documentsTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count documents_types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q documentsTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if documents_types exists")
	}

	return count > 0, nil
}

// DocumentsTypes retrieves all the records using an executor.
func DocumentsTypes(mods ...qm.QueryMod) documentsTypeQuery {
	mods = append(mods, qm.From("`documents_types`"))
	return documentsTypeQuery{NewQuery(mods...)}
}

// FindDocumentsType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDocumentsType(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DocumentsType, error) {
	documentsTypeObj := &DocumentsType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `documents_types` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, documentsTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from documents_types")
	}

	if err = documentsTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return documentsTypeObj, err
	}

	return documentsTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DocumentsType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no documents_types provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentsTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	documentsTypeInsertCacheMut.RLock()
	cache, cached := documentsTypeInsertCache[key]
	documentsTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			documentsTypeAllColumns,
			documentsTypeColumnsWithDefault,
			documentsTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(documentsTypeType, documentsTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(documentsTypeType, documentsTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `documents_types` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `documents_types` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `documents_types` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, documentsTypePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into documents_types")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == documentsTypeMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for documents_types")
	}

CacheNoHooks:
	if !cached {
		documentsTypeInsertCacheMut.Lock()
		documentsTypeInsertCache[key] = cache
		documentsTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DocumentsType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DocumentsType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	documentsTypeUpdateCacheMut.RLock()
	cache, cached := documentsTypeUpdateCache[key]
	documentsTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			documentsTypeAllColumns,
			documentsTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update documents_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `documents_types` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, documentsTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(documentsTypeType, documentsTypeMapping, append(wl, documentsTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update documents_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for documents_types")
	}

	if !cached {
		documentsTypeUpdateCacheMut.Lock()
		documentsTypeUpdateCache[key] = cache
		documentsTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q documentsTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for documents_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for documents_types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DocumentsTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentsTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `documents_types` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentsTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in documentsType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all documentsType")
	}
	return rowsAff, nil
}

var mySQLDocumentsTypeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DocumentsType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no documents_types provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentsTypeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDocumentsTypeUniqueColumns, o)

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

	documentsTypeUpsertCacheMut.RLock()
	cache, cached := documentsTypeUpsertCache[key]
	documentsTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			documentsTypeAllColumns,
			documentsTypeColumnsWithDefault,
			documentsTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			documentsTypeAllColumns,
			documentsTypePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert documents_types, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`documents_types`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `documents_types` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(documentsTypeType, documentsTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(documentsTypeType, documentsTypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for documents_types")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == documentsTypeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(documentsTypeType, documentsTypeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for documents_types")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for documents_types")
	}

CacheNoHooks:
	if !cached {
		documentsTypeUpsertCacheMut.Lock()
		documentsTypeUpsertCache[key] = cache
		documentsTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DocumentsType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DocumentsType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DocumentsType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), documentsTypePrimaryKeyMapping)
	sql := "DELETE FROM `documents_types` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from documents_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for documents_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q documentsTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no documentsTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from documents_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for documents_types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DocumentsTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(documentsTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentsTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `documents_types` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentsTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from documentsType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for documents_types")
	}

	if len(documentsTypeAfterDeleteHooks) != 0 {
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
func (o *DocumentsType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDocumentsType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DocumentsTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DocumentsTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentsTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `documents_types`.* FROM `documents_types` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentsTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DocumentsTypeSlice")
	}

	*o = slice

	return nil
}

// DocumentsTypeExists checks if the DocumentsType row exists.
func DocumentsTypeExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `documents_types` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if documents_types exists")
	}

	return exists, nil
}
