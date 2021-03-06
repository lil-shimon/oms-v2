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

// BusinessPartnerCompany is an object representing the database table.
type BusinessPartnerCompany struct {
	ID              uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name            string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Furigana        null.String `boil:"furigana" json:"furigana,omitempty" toml:"furigana" yaml:"furigana,omitempty"`
	Department      null.String `boil:"department" json:"department,omitempty" toml:"department" yaml:"department,omitempty"`
	PersonInCharge  null.String `boil:"person_in_charge" json:"person_in_charge,omitempty" toml:"person_in_charge" yaml:"person_in_charge,omitempty"`
	MailAddress     null.String `boil:"mail_address" json:"mail_address,omitempty" toml:"mail_address" yaml:"mail_address,omitempty"`
	PostCode        null.String `boil:"post_code" json:"post_code,omitempty" toml:"post_code" yaml:"post_code,omitempty"`
	Address         null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	TelephoneNumber null.String `boil:"telephone_number" json:"telephone_number,omitempty" toml:"telephone_number" yaml:"telephone_number,omitempty"`
	HonorificTitle  string      `boil:"honorific_title" json:"honorific_title" toml:"honorific_title" yaml:"honorific_title"`
	DeletedAt       null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *businessPartnerCompanyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L businessPartnerCompanyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BusinessPartnerCompanyColumns = struct {
	ID              string
	Name            string
	Furigana        string
	Department      string
	PersonInCharge  string
	MailAddress     string
	PostCode        string
	Address         string
	TelephoneNumber string
	HonorificTitle  string
	DeletedAt       string
}{
	ID:              "id",
	Name:            "name",
	Furigana:        "furigana",
	Department:      "department",
	PersonInCharge:  "person_in_charge",
	MailAddress:     "mail_address",
	PostCode:        "post_code",
	Address:         "address",
	TelephoneNumber: "telephone_number",
	HonorificTitle:  "honorific_title",
	DeletedAt:       "deleted_at",
}

var BusinessPartnerCompanyTableColumns = struct {
	ID              string
	Name            string
	Furigana        string
	Department      string
	PersonInCharge  string
	MailAddress     string
	PostCode        string
	Address         string
	TelephoneNumber string
	HonorificTitle  string
	DeletedAt       string
}{
	ID:              "business_partner_companies.id",
	Name:            "business_partner_companies.name",
	Furigana:        "business_partner_companies.furigana",
	Department:      "business_partner_companies.department",
	PersonInCharge:  "business_partner_companies.person_in_charge",
	MailAddress:     "business_partner_companies.mail_address",
	PostCode:        "business_partner_companies.post_code",
	Address:         "business_partner_companies.address",
	TelephoneNumber: "business_partner_companies.telephone_number",
	HonorificTitle:  "business_partner_companies.honorific_title",
	DeletedAt:       "business_partner_companies.deleted_at",
}

// Generated where

var BusinessPartnerCompanyWhere = struct {
	ID              whereHelperuint
	Name            whereHelperstring
	Furigana        whereHelpernull_String
	Department      whereHelpernull_String
	PersonInCharge  whereHelpernull_String
	MailAddress     whereHelpernull_String
	PostCode        whereHelpernull_String
	Address         whereHelpernull_String
	TelephoneNumber whereHelpernull_String
	HonorificTitle  whereHelperstring
	DeletedAt       whereHelpernull_Time
}{
	ID:              whereHelperuint{field: "`business_partner_companies`.`id`"},
	Name:            whereHelperstring{field: "`business_partner_companies`.`name`"},
	Furigana:        whereHelpernull_String{field: "`business_partner_companies`.`furigana`"},
	Department:      whereHelpernull_String{field: "`business_partner_companies`.`department`"},
	PersonInCharge:  whereHelpernull_String{field: "`business_partner_companies`.`person_in_charge`"},
	MailAddress:     whereHelpernull_String{field: "`business_partner_companies`.`mail_address`"},
	PostCode:        whereHelpernull_String{field: "`business_partner_companies`.`post_code`"},
	Address:         whereHelpernull_String{field: "`business_partner_companies`.`address`"},
	TelephoneNumber: whereHelpernull_String{field: "`business_partner_companies`.`telephone_number`"},
	HonorificTitle:  whereHelperstring{field: "`business_partner_companies`.`honorific_title`"},
	DeletedAt:       whereHelpernull_Time{field: "`business_partner_companies`.`deleted_at`"},
}

// BusinessPartnerCompanyRels is where relationship names are stored.
var BusinessPartnerCompanyRels = struct {
}{}

// businessPartnerCompanyR is where relationships are stored.
type businessPartnerCompanyR struct {
}

// NewStruct creates a new relationship struct
func (*businessPartnerCompanyR) NewStruct() *businessPartnerCompanyR {
	return &businessPartnerCompanyR{}
}

// businessPartnerCompanyL is where Load methods for each relationship are stored.
type businessPartnerCompanyL struct{}

var (
	businessPartnerCompanyAllColumns            = []string{"id", "name", "furigana", "department", "person_in_charge", "mail_address", "post_code", "address", "telephone_number", "honorific_title", "deleted_at"}
	businessPartnerCompanyColumnsWithoutDefault = []string{"name", "furigana", "department", "person_in_charge", "mail_address", "post_code", "address", "telephone_number", "honorific_title", "deleted_at"}
	businessPartnerCompanyColumnsWithDefault    = []string{"id"}
	businessPartnerCompanyPrimaryKeyColumns     = []string{"id"}
)

type (
	// BusinessPartnerCompanySlice is an alias for a slice of pointers to BusinessPartnerCompany.
	// This should almost always be used instead of []BusinessPartnerCompany.
	BusinessPartnerCompanySlice []*BusinessPartnerCompany
	// BusinessPartnerCompanyHook is the signature for custom BusinessPartnerCompany hook methods
	BusinessPartnerCompanyHook func(context.Context, boil.ContextExecutor, *BusinessPartnerCompany) error

	businessPartnerCompanyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	businessPartnerCompanyType                 = reflect.TypeOf(&BusinessPartnerCompany{})
	businessPartnerCompanyMapping              = queries.MakeStructMapping(businessPartnerCompanyType)
	businessPartnerCompanyPrimaryKeyMapping, _ = queries.BindMapping(businessPartnerCompanyType, businessPartnerCompanyMapping, businessPartnerCompanyPrimaryKeyColumns)
	businessPartnerCompanyInsertCacheMut       sync.RWMutex
	businessPartnerCompanyInsertCache          = make(map[string]insertCache)
	businessPartnerCompanyUpdateCacheMut       sync.RWMutex
	businessPartnerCompanyUpdateCache          = make(map[string]updateCache)
	businessPartnerCompanyUpsertCacheMut       sync.RWMutex
	businessPartnerCompanyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var businessPartnerCompanyBeforeInsertHooks []BusinessPartnerCompanyHook
var businessPartnerCompanyBeforeUpdateHooks []BusinessPartnerCompanyHook
var businessPartnerCompanyBeforeDeleteHooks []BusinessPartnerCompanyHook
var businessPartnerCompanyBeforeUpsertHooks []BusinessPartnerCompanyHook

var businessPartnerCompanyAfterInsertHooks []BusinessPartnerCompanyHook
var businessPartnerCompanyAfterSelectHooks []BusinessPartnerCompanyHook
var businessPartnerCompanyAfterUpdateHooks []BusinessPartnerCompanyHook
var businessPartnerCompanyAfterDeleteHooks []BusinessPartnerCompanyHook
var businessPartnerCompanyAfterUpsertHooks []BusinessPartnerCompanyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BusinessPartnerCompany) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BusinessPartnerCompany) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BusinessPartnerCompany) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BusinessPartnerCompany) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BusinessPartnerCompany) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BusinessPartnerCompany) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BusinessPartnerCompany) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BusinessPartnerCompany) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BusinessPartnerCompany) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessPartnerCompanyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBusinessPartnerCompanyHook registers your hook function for all future operations.
func AddBusinessPartnerCompanyHook(hookPoint boil.HookPoint, businessPartnerCompanyHook BusinessPartnerCompanyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		businessPartnerCompanyBeforeInsertHooks = append(businessPartnerCompanyBeforeInsertHooks, businessPartnerCompanyHook)
	case boil.BeforeUpdateHook:
		businessPartnerCompanyBeforeUpdateHooks = append(businessPartnerCompanyBeforeUpdateHooks, businessPartnerCompanyHook)
	case boil.BeforeDeleteHook:
		businessPartnerCompanyBeforeDeleteHooks = append(businessPartnerCompanyBeforeDeleteHooks, businessPartnerCompanyHook)
	case boil.BeforeUpsertHook:
		businessPartnerCompanyBeforeUpsertHooks = append(businessPartnerCompanyBeforeUpsertHooks, businessPartnerCompanyHook)
	case boil.AfterInsertHook:
		businessPartnerCompanyAfterInsertHooks = append(businessPartnerCompanyAfterInsertHooks, businessPartnerCompanyHook)
	case boil.AfterSelectHook:
		businessPartnerCompanyAfterSelectHooks = append(businessPartnerCompanyAfterSelectHooks, businessPartnerCompanyHook)
	case boil.AfterUpdateHook:
		businessPartnerCompanyAfterUpdateHooks = append(businessPartnerCompanyAfterUpdateHooks, businessPartnerCompanyHook)
	case boil.AfterDeleteHook:
		businessPartnerCompanyAfterDeleteHooks = append(businessPartnerCompanyAfterDeleteHooks, businessPartnerCompanyHook)
	case boil.AfterUpsertHook:
		businessPartnerCompanyAfterUpsertHooks = append(businessPartnerCompanyAfterUpsertHooks, businessPartnerCompanyHook)
	}
}

// One returns a single businessPartnerCompany record from the query.
func (q businessPartnerCompanyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BusinessPartnerCompany, error) {
	o := &BusinessPartnerCompany{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for business_partner_companies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BusinessPartnerCompany records from the query.
func (q businessPartnerCompanyQuery) All(ctx context.Context, exec boil.ContextExecutor) (BusinessPartnerCompanySlice, error) {
	var o []*BusinessPartnerCompany

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BusinessPartnerCompany slice")
	}

	if len(businessPartnerCompanyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BusinessPartnerCompany records in the query.
func (q businessPartnerCompanyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count business_partner_companies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q businessPartnerCompanyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if business_partner_companies exists")
	}

	return count > 0, nil
}

// BusinessPartnerCompanies retrieves all the records using an executor.
func BusinessPartnerCompanies(mods ...qm.QueryMod) businessPartnerCompanyQuery {
	mods = append(mods, qm.From("`business_partner_companies`"))
	return businessPartnerCompanyQuery{NewQuery(mods...)}
}

// FindBusinessPartnerCompany retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBusinessPartnerCompany(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*BusinessPartnerCompany, error) {
	businessPartnerCompanyObj := &BusinessPartnerCompany{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `business_partner_companies` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, businessPartnerCompanyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from business_partner_companies")
	}

	if err = businessPartnerCompanyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return businessPartnerCompanyObj, err
	}

	return businessPartnerCompanyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BusinessPartnerCompany) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no business_partner_companies provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(businessPartnerCompanyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	businessPartnerCompanyInsertCacheMut.RLock()
	cache, cached := businessPartnerCompanyInsertCache[key]
	businessPartnerCompanyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			businessPartnerCompanyAllColumns,
			businessPartnerCompanyColumnsWithDefault,
			businessPartnerCompanyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(businessPartnerCompanyType, businessPartnerCompanyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(businessPartnerCompanyType, businessPartnerCompanyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `business_partner_companies` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `business_partner_companies` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `business_partner_companies` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, businessPartnerCompanyPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into business_partner_companies")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == businessPartnerCompanyMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for business_partner_companies")
	}

CacheNoHooks:
	if !cached {
		businessPartnerCompanyInsertCacheMut.Lock()
		businessPartnerCompanyInsertCache[key] = cache
		businessPartnerCompanyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BusinessPartnerCompany.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BusinessPartnerCompany) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	businessPartnerCompanyUpdateCacheMut.RLock()
	cache, cached := businessPartnerCompanyUpdateCache[key]
	businessPartnerCompanyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			businessPartnerCompanyAllColumns,
			businessPartnerCompanyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update business_partner_companies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `business_partner_companies` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, businessPartnerCompanyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(businessPartnerCompanyType, businessPartnerCompanyMapping, append(wl, businessPartnerCompanyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update business_partner_companies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for business_partner_companies")
	}

	if !cached {
		businessPartnerCompanyUpdateCacheMut.Lock()
		businessPartnerCompanyUpdateCache[key] = cache
		businessPartnerCompanyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q businessPartnerCompanyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for business_partner_companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for business_partner_companies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BusinessPartnerCompanySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), businessPartnerCompanyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `business_partner_companies` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, businessPartnerCompanyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in businessPartnerCompany slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all businessPartnerCompany")
	}
	return rowsAff, nil
}

var mySQLBusinessPartnerCompanyUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BusinessPartnerCompany) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no business_partner_companies provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(businessPartnerCompanyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBusinessPartnerCompanyUniqueColumns, o)

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

	businessPartnerCompanyUpsertCacheMut.RLock()
	cache, cached := businessPartnerCompanyUpsertCache[key]
	businessPartnerCompanyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			businessPartnerCompanyAllColumns,
			businessPartnerCompanyColumnsWithDefault,
			businessPartnerCompanyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			businessPartnerCompanyAllColumns,
			businessPartnerCompanyPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert business_partner_companies, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`business_partner_companies`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `business_partner_companies` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(businessPartnerCompanyType, businessPartnerCompanyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(businessPartnerCompanyType, businessPartnerCompanyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for business_partner_companies")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == businessPartnerCompanyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(businessPartnerCompanyType, businessPartnerCompanyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for business_partner_companies")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for business_partner_companies")
	}

CacheNoHooks:
	if !cached {
		businessPartnerCompanyUpsertCacheMut.Lock()
		businessPartnerCompanyUpsertCache[key] = cache
		businessPartnerCompanyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BusinessPartnerCompany record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BusinessPartnerCompany) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BusinessPartnerCompany provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), businessPartnerCompanyPrimaryKeyMapping)
	sql := "DELETE FROM `business_partner_companies` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from business_partner_companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for business_partner_companies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q businessPartnerCompanyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no businessPartnerCompanyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from business_partner_companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for business_partner_companies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BusinessPartnerCompanySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(businessPartnerCompanyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), businessPartnerCompanyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `business_partner_companies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, businessPartnerCompanyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from businessPartnerCompany slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for business_partner_companies")
	}

	if len(businessPartnerCompanyAfterDeleteHooks) != 0 {
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
func (o *BusinessPartnerCompany) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBusinessPartnerCompany(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BusinessPartnerCompanySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BusinessPartnerCompanySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), businessPartnerCompanyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `business_partner_companies`.* FROM `business_partner_companies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, businessPartnerCompanyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BusinessPartnerCompanySlice")
	}

	*o = slice

	return nil
}

// BusinessPartnerCompanyExists checks if the BusinessPartnerCompany row exists.
func BusinessPartnerCompanyExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `business_partner_companies` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if business_partner_companies exists")
	}

	return exists, nil
}
