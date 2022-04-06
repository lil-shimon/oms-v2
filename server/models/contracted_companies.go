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

// ContractedCompany is an object representing the database table.
type ContractedCompany struct {
	ID                      uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	DocumentID              uint        `boil:"document_id" json:"document_id" toml:"document_id" yaml:"document_id"`
	Name                    string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	PersonInCharge          null.String `boil:"person_in_charge" json:"person_in_charge,omitempty" toml:"person_in_charge" yaml:"person_in_charge,omitempty"`
	HonorificTitle          string      `boil:"honorific_title" json:"honorific_title" toml:"honorific_title" yaml:"honorific_title"`
	DeletedAt               null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	OrderID                 uint        `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	Tel                     string      `boil:"tel" json:"tel" toml:"tel" yaml:"tel"`
	SiteName                string      `boil:"site_name" json:"site_name" toml:"site_name" yaml:"site_name"`
	SiteRepresentative      string      `boil:"site_representative" json:"site_representative" toml:"site_representative" yaml:"site_representative"`
	SiteRepresentativePhone string      `boil:"site_representative_phone" json:"site_representative_phone" toml:"site_representative_phone" yaml:"site_representative_phone"`
	Mail                    string      `boil:"mail" json:"mail" toml:"mail" yaml:"mail"`
	Address                 string      `boil:"address" json:"address" toml:"address" yaml:"address"`

	R *contractedCompanyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L contractedCompanyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ContractedCompanyColumns = struct {
	ID                      string
	DocumentID              string
	Name                    string
	PersonInCharge          string
	HonorificTitle          string
	DeletedAt               string
	OrderID                 string
	Tel                     string
	SiteName                string
	SiteRepresentative      string
	SiteRepresentativePhone string
	Mail                    string
	Address                 string
}{
	ID:                      "id",
	DocumentID:              "document_id",
	Name:                    "name",
	PersonInCharge:          "person_in_charge",
	HonorificTitle:          "honorific_title",
	DeletedAt:               "deleted_at",
	OrderID:                 "order_id",
	Tel:                     "tel",
	SiteName:                "site_name",
	SiteRepresentative:      "site_representative",
	SiteRepresentativePhone: "site_representative_phone",
	Mail:                    "mail",
	Address:                 "address",
}

var ContractedCompanyTableColumns = struct {
	ID                      string
	DocumentID              string
	Name                    string
	PersonInCharge          string
	HonorificTitle          string
	DeletedAt               string
	OrderID                 string
	Tel                     string
	SiteName                string
	SiteRepresentative      string
	SiteRepresentativePhone string
	Mail                    string
	Address                 string
}{
	ID:                      "contracted_companies.id",
	DocumentID:              "contracted_companies.document_id",
	Name:                    "contracted_companies.name",
	PersonInCharge:          "contracted_companies.person_in_charge",
	HonorificTitle:          "contracted_companies.honorific_title",
	DeletedAt:               "contracted_companies.deleted_at",
	OrderID:                 "contracted_companies.order_id",
	Tel:                     "contracted_companies.tel",
	SiteName:                "contracted_companies.site_name",
	SiteRepresentative:      "contracted_companies.site_representative",
	SiteRepresentativePhone: "contracted_companies.site_representative_phone",
	Mail:                    "contracted_companies.mail",
	Address:                 "contracted_companies.address",
}

// Generated where

var ContractedCompanyWhere = struct {
	ID                      whereHelperuint
	DocumentID              whereHelperuint
	Name                    whereHelperstring
	PersonInCharge          whereHelpernull_String
	HonorificTitle          whereHelperstring
	DeletedAt               whereHelpernull_Time
	OrderID                 whereHelperuint
	Tel                     whereHelperstring
	SiteName                whereHelperstring
	SiteRepresentative      whereHelperstring
	SiteRepresentativePhone whereHelperstring
	Mail                    whereHelperstring
	Address                 whereHelperstring
}{
	ID:                      whereHelperuint{field: "`contracted_companies`.`id`"},
	DocumentID:              whereHelperuint{field: "`contracted_companies`.`document_id`"},
	Name:                    whereHelperstring{field: "`contracted_companies`.`name`"},
	PersonInCharge:          whereHelpernull_String{field: "`contracted_companies`.`person_in_charge`"},
	HonorificTitle:          whereHelperstring{field: "`contracted_companies`.`honorific_title`"},
	DeletedAt:               whereHelpernull_Time{field: "`contracted_companies`.`deleted_at`"},
	OrderID:                 whereHelperuint{field: "`contracted_companies`.`order_id`"},
	Tel:                     whereHelperstring{field: "`contracted_companies`.`tel`"},
	SiteName:                whereHelperstring{field: "`contracted_companies`.`site_name`"},
	SiteRepresentative:      whereHelperstring{field: "`contracted_companies`.`site_representative`"},
	SiteRepresentativePhone: whereHelperstring{field: "`contracted_companies`.`site_representative_phone`"},
	Mail:                    whereHelperstring{field: "`contracted_companies`.`mail`"},
	Address:                 whereHelperstring{field: "`contracted_companies`.`address`"},
}

// ContractedCompanyRels is where relationship names are stored.
var ContractedCompanyRels = struct {
	Document string
}{
	Document: "Document",
}

// contractedCompanyR is where relationships are stored.
type contractedCompanyR struct {
	Document *Document `boil:"Document" json:"Document" toml:"Document" yaml:"Document"`
}

// NewStruct creates a new relationship struct
func (*contractedCompanyR) NewStruct() *contractedCompanyR {
	return &contractedCompanyR{}
}

// contractedCompanyL is where Load methods for each relationship are stored.
type contractedCompanyL struct{}

var (
	contractedCompanyAllColumns            = []string{"id", "document_id", "name", "person_in_charge", "honorific_title", "deleted_at", "order_id", "tel", "site_name", "site_representative", "site_representative_phone", "mail", "address"}
	contractedCompanyColumnsWithoutDefault = []string{"document_id", "name", "person_in_charge", "honorific_title", "deleted_at", "order_id", "tel", "site_name", "site_representative", "site_representative_phone", "mail", "address"}
	contractedCompanyColumnsWithDefault    = []string{"id"}
	contractedCompanyPrimaryKeyColumns     = []string{"id"}
)

type (
	// ContractedCompanySlice is an alias for a slice of pointers to ContractedCompany.
	// This should almost always be used instead of []ContractedCompany.
	ContractedCompanySlice []*ContractedCompany
	// ContractedCompanyHook is the signature for custom ContractedCompany hook methods
	ContractedCompanyHook func(context.Context, boil.ContextExecutor, *ContractedCompany) error

	contractedCompanyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	contractedCompanyType                 = reflect.TypeOf(&ContractedCompany{})
	contractedCompanyMapping              = queries.MakeStructMapping(contractedCompanyType)
	contractedCompanyPrimaryKeyMapping, _ = queries.BindMapping(contractedCompanyType, contractedCompanyMapping, contractedCompanyPrimaryKeyColumns)
	contractedCompanyInsertCacheMut       sync.RWMutex
	contractedCompanyInsertCache          = make(map[string]insertCache)
	contractedCompanyUpdateCacheMut       sync.RWMutex
	contractedCompanyUpdateCache          = make(map[string]updateCache)
	contractedCompanyUpsertCacheMut       sync.RWMutex
	contractedCompanyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var contractedCompanyBeforeInsertHooks []ContractedCompanyHook
var contractedCompanyBeforeUpdateHooks []ContractedCompanyHook
var contractedCompanyBeforeDeleteHooks []ContractedCompanyHook
var contractedCompanyBeforeUpsertHooks []ContractedCompanyHook

var contractedCompanyAfterInsertHooks []ContractedCompanyHook
var contractedCompanyAfterSelectHooks []ContractedCompanyHook
var contractedCompanyAfterUpdateHooks []ContractedCompanyHook
var contractedCompanyAfterDeleteHooks []ContractedCompanyHook
var contractedCompanyAfterUpsertHooks []ContractedCompanyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ContractedCompany) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ContractedCompany) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ContractedCompany) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ContractedCompany) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ContractedCompany) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ContractedCompany) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ContractedCompany) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ContractedCompany) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ContractedCompany) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contractedCompanyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddContractedCompanyHook registers your hook function for all future operations.
func AddContractedCompanyHook(hookPoint boil.HookPoint, contractedCompanyHook ContractedCompanyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		contractedCompanyBeforeInsertHooks = append(contractedCompanyBeforeInsertHooks, contractedCompanyHook)
	case boil.BeforeUpdateHook:
		contractedCompanyBeforeUpdateHooks = append(contractedCompanyBeforeUpdateHooks, contractedCompanyHook)
	case boil.BeforeDeleteHook:
		contractedCompanyBeforeDeleteHooks = append(contractedCompanyBeforeDeleteHooks, contractedCompanyHook)
	case boil.BeforeUpsertHook:
		contractedCompanyBeforeUpsertHooks = append(contractedCompanyBeforeUpsertHooks, contractedCompanyHook)
	case boil.AfterInsertHook:
		contractedCompanyAfterInsertHooks = append(contractedCompanyAfterInsertHooks, contractedCompanyHook)
	case boil.AfterSelectHook:
		contractedCompanyAfterSelectHooks = append(contractedCompanyAfterSelectHooks, contractedCompanyHook)
	case boil.AfterUpdateHook:
		contractedCompanyAfterUpdateHooks = append(contractedCompanyAfterUpdateHooks, contractedCompanyHook)
	case boil.AfterDeleteHook:
		contractedCompanyAfterDeleteHooks = append(contractedCompanyAfterDeleteHooks, contractedCompanyHook)
	case boil.AfterUpsertHook:
		contractedCompanyAfterUpsertHooks = append(contractedCompanyAfterUpsertHooks, contractedCompanyHook)
	}
}

// One returns a single contractedCompany record from the query.
func (q contractedCompanyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ContractedCompany, error) {
	o := &ContractedCompany{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for contracted_companies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ContractedCompany records from the query.
func (q contractedCompanyQuery) All(ctx context.Context, exec boil.ContextExecutor) (ContractedCompanySlice, error) {
	var o []*ContractedCompany

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ContractedCompany slice")
	}

	if len(contractedCompanyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ContractedCompany records in the query.
func (q contractedCompanyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count contracted_companies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q contractedCompanyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if contracted_companies exists")
	}

	return count > 0, nil
}

// Document pointed to by the foreign key.
func (o *ContractedCompany) Document(mods ...qm.QueryMod) documentQuery {
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
func (contractedCompanyL) LoadDocument(ctx context.Context, e boil.ContextExecutor, singular bool, maybeContractedCompany interface{}, mods queries.Applicator) error {
	var slice []*ContractedCompany
	var object *ContractedCompany

	if singular {
		object = maybeContractedCompany.(*ContractedCompany)
	} else {
		slice = *maybeContractedCompany.(*[]*ContractedCompany)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &contractedCompanyR{}
		}
		args = append(args, object.DocumentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &contractedCompanyR{}
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

	if len(contractedCompanyAfterSelectHooks) != 0 {
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
		foreign.R.ContractedCompanies = append(foreign.R.ContractedCompanies, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DocumentID == foreign.ID {
				local.R.Document = foreign
				if foreign.R == nil {
					foreign.R = &documentR{}
				}
				foreign.R.ContractedCompanies = append(foreign.R.ContractedCompanies, local)
				break
			}
		}
	}

	return nil
}

// SetDocument of the contractedCompany to the related item.
// Sets o.R.Document to related.
// Adds o to related.R.ContractedCompanies.
func (o *ContractedCompany) SetDocument(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Document) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `contracted_companies` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"document_id"}),
		strmangle.WhereClause("`", "`", 0, contractedCompanyPrimaryKeyColumns),
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
		o.R = &contractedCompanyR{
			Document: related,
		}
	} else {
		o.R.Document = related
	}

	if related.R == nil {
		related.R = &documentR{
			ContractedCompanies: ContractedCompanySlice{o},
		}
	} else {
		related.R.ContractedCompanies = append(related.R.ContractedCompanies, o)
	}

	return nil
}

// ContractedCompanies retrieves all the records using an executor.
func ContractedCompanies(mods ...qm.QueryMod) contractedCompanyQuery {
	mods = append(mods, qm.From("`contracted_companies`"))
	return contractedCompanyQuery{NewQuery(mods...)}
}

// FindContractedCompany retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindContractedCompany(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*ContractedCompany, error) {
	contractedCompanyObj := &ContractedCompany{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `contracted_companies` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, contractedCompanyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from contracted_companies")
	}

	if err = contractedCompanyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return contractedCompanyObj, err
	}

	return contractedCompanyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ContractedCompany) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no contracted_companies provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contractedCompanyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	contractedCompanyInsertCacheMut.RLock()
	cache, cached := contractedCompanyInsertCache[key]
	contractedCompanyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			contractedCompanyAllColumns,
			contractedCompanyColumnsWithDefault,
			contractedCompanyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(contractedCompanyType, contractedCompanyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(contractedCompanyType, contractedCompanyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `contracted_companies` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `contracted_companies` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `contracted_companies` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, contractedCompanyPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into contracted_companies")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == contractedCompanyMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for contracted_companies")
	}

CacheNoHooks:
	if !cached {
		contractedCompanyInsertCacheMut.Lock()
		contractedCompanyInsertCache[key] = cache
		contractedCompanyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ContractedCompany.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ContractedCompany) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	contractedCompanyUpdateCacheMut.RLock()
	cache, cached := contractedCompanyUpdateCache[key]
	contractedCompanyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			contractedCompanyAllColumns,
			contractedCompanyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update contracted_companies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `contracted_companies` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, contractedCompanyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(contractedCompanyType, contractedCompanyMapping, append(wl, contractedCompanyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update contracted_companies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for contracted_companies")
	}

	if !cached {
		contractedCompanyUpdateCacheMut.Lock()
		contractedCompanyUpdateCache[key] = cache
		contractedCompanyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q contractedCompanyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for contracted_companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for contracted_companies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ContractedCompanySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contractedCompanyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `contracted_companies` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, contractedCompanyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in contractedCompany slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all contractedCompany")
	}
	return rowsAff, nil
}

var mySQLContractedCompanyUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ContractedCompany) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no contracted_companies provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contractedCompanyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLContractedCompanyUniqueColumns, o)

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

	contractedCompanyUpsertCacheMut.RLock()
	cache, cached := contractedCompanyUpsertCache[key]
	contractedCompanyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			contractedCompanyAllColumns,
			contractedCompanyColumnsWithDefault,
			contractedCompanyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			contractedCompanyAllColumns,
			contractedCompanyPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert contracted_companies, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`contracted_companies`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `contracted_companies` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(contractedCompanyType, contractedCompanyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(contractedCompanyType, contractedCompanyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for contracted_companies")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == contractedCompanyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(contractedCompanyType, contractedCompanyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for contracted_companies")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for contracted_companies")
	}

CacheNoHooks:
	if !cached {
		contractedCompanyUpsertCacheMut.Lock()
		contractedCompanyUpsertCache[key] = cache
		contractedCompanyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ContractedCompany record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ContractedCompany) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ContractedCompany provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), contractedCompanyPrimaryKeyMapping)
	sql := "DELETE FROM `contracted_companies` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from contracted_companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for contracted_companies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q contractedCompanyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no contractedCompanyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from contracted_companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for contracted_companies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ContractedCompanySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(contractedCompanyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contractedCompanyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `contracted_companies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, contractedCompanyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from contractedCompany slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for contracted_companies")
	}

	if len(contractedCompanyAfterDeleteHooks) != 0 {
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
func (o *ContractedCompany) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindContractedCompany(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContractedCompanySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ContractedCompanySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contractedCompanyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `contracted_companies`.* FROM `contracted_companies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, contractedCompanyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ContractedCompanySlice")
	}

	*o = slice

	return nil
}

// ContractedCompanyExists checks if the ContractedCompany row exists.
func ContractedCompanyExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `contracted_companies` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if contracted_companies exists")
	}

	return exists, nil
}