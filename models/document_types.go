// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// DocumentType is an object representing the database table.
type DocumentType struct {
	ID           int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	DocumentType string `boil:"document_type" json:"document_type" toml:"document_type" yaml:"document_type"`

	R *documentTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L documentTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DocumentTypeColumns = struct {
	ID           string
	DocumentType string
}{
	ID:           "id",
	DocumentType: "document_type",
}

var DocumentTypeTableColumns = struct {
	ID           string
	DocumentType string
}{
	ID:           "document_types.id",
	DocumentType: "document_types.document_type",
}

// Generated where

var DocumentTypeWhere = struct {
	ID           whereHelperint64
	DocumentType whereHelperstring
}{
	ID:           whereHelperint64{field: "\"document_types\".\"id\""},
	DocumentType: whereHelperstring{field: "\"document_types\".\"document_type\""},
}

// DocumentTypeRels is where relationship names are stored.
var DocumentTypeRels = struct {
	Documents string
}{
	Documents: "Documents",
}

// documentTypeR is where relationships are stored.
type documentTypeR struct {
	Documents DocumentSlice `boil:"Documents" json:"Documents" toml:"Documents" yaml:"Documents"`
}

// NewStruct creates a new relationship struct
func (*documentTypeR) NewStruct() *documentTypeR {
	return &documentTypeR{}
}

// documentTypeL is where Load methods for each relationship are stored.
type documentTypeL struct{}

var (
	documentTypeAllColumns            = []string{"id", "document_type"}
	documentTypeColumnsWithoutDefault = []string{"document_type"}
	documentTypeColumnsWithDefault    = []string{"id"}
	documentTypePrimaryKeyColumns     = []string{"id"}
	documentTypeGeneratedColumns      = []string{"id"}
)

type (
	// DocumentTypeSlice is an alias for a slice of pointers to DocumentType.
	// This should almost always be used instead of []DocumentType.
	DocumentTypeSlice []*DocumentType
	// DocumentTypeHook is the signature for custom DocumentType hook methods
	DocumentTypeHook func(context.Context, boil.ContextExecutor, *DocumentType) error

	documentTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	documentTypeType                 = reflect.TypeOf(&DocumentType{})
	documentTypeMapping              = queries.MakeStructMapping(documentTypeType)
	documentTypePrimaryKeyMapping, _ = queries.BindMapping(documentTypeType, documentTypeMapping, documentTypePrimaryKeyColumns)
	documentTypeInsertCacheMut       sync.RWMutex
	documentTypeInsertCache          = make(map[string]insertCache)
	documentTypeUpdateCacheMut       sync.RWMutex
	documentTypeUpdateCache          = make(map[string]updateCache)
	documentTypeUpsertCacheMut       sync.RWMutex
	documentTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var documentTypeAfterSelectHooks []DocumentTypeHook

var documentTypeBeforeInsertHooks []DocumentTypeHook
var documentTypeAfterInsertHooks []DocumentTypeHook

var documentTypeBeforeUpdateHooks []DocumentTypeHook
var documentTypeAfterUpdateHooks []DocumentTypeHook

var documentTypeBeforeDeleteHooks []DocumentTypeHook
var documentTypeAfterDeleteHooks []DocumentTypeHook

var documentTypeBeforeUpsertHooks []DocumentTypeHook
var documentTypeAfterUpsertHooks []DocumentTypeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DocumentType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DocumentType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DocumentType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DocumentType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DocumentType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DocumentType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DocumentType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DocumentType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DocumentType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDocumentTypeHook registers your hook function for all future operations.
func AddDocumentTypeHook(hookPoint boil.HookPoint, documentTypeHook DocumentTypeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		documentTypeAfterSelectHooks = append(documentTypeAfterSelectHooks, documentTypeHook)
	case boil.BeforeInsertHook:
		documentTypeBeforeInsertHooks = append(documentTypeBeforeInsertHooks, documentTypeHook)
	case boil.AfterInsertHook:
		documentTypeAfterInsertHooks = append(documentTypeAfterInsertHooks, documentTypeHook)
	case boil.BeforeUpdateHook:
		documentTypeBeforeUpdateHooks = append(documentTypeBeforeUpdateHooks, documentTypeHook)
	case boil.AfterUpdateHook:
		documentTypeAfterUpdateHooks = append(documentTypeAfterUpdateHooks, documentTypeHook)
	case boil.BeforeDeleteHook:
		documentTypeBeforeDeleteHooks = append(documentTypeBeforeDeleteHooks, documentTypeHook)
	case boil.AfterDeleteHook:
		documentTypeAfterDeleteHooks = append(documentTypeAfterDeleteHooks, documentTypeHook)
	case boil.BeforeUpsertHook:
		documentTypeBeforeUpsertHooks = append(documentTypeBeforeUpsertHooks, documentTypeHook)
	case boil.AfterUpsertHook:
		documentTypeAfterUpsertHooks = append(documentTypeAfterUpsertHooks, documentTypeHook)
	}
}

// One returns a single documentType record from the query.
func (q documentTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DocumentType, error) {
	o := &DocumentType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for document_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DocumentType records from the query.
func (q documentTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (DocumentTypeSlice, error) {
	var o []*DocumentType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DocumentType slice")
	}

	if len(documentTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DocumentType records in the query.
func (q documentTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count document_types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q documentTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if document_types exists")
	}

	return count > 0, nil
}

// Documents retrieves all the document's Documents with an executor.
func (o *DocumentType) Documents(mods ...qm.QueryMod) documentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"documents\".\"document_type_id\"=?", o.ID),
	)

	return Documents(queryMods...)
}

// LoadDocuments allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (documentTypeL) LoadDocuments(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDocumentType interface{}, mods queries.Applicator) error {
	var slice []*DocumentType
	var object *DocumentType

	if singular {
		object = maybeDocumentType.(*DocumentType)
	} else {
		slice = *maybeDocumentType.(*[]*DocumentType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &documentTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &documentTypeR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`documents`),
		qm.WhereIn(`documents.document_type_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load documents")
	}

	var resultSlice []*Document
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice documents")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on documents")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for documents")
	}

	if len(documentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Documents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &documentR{}
			}
			foreign.R.DocumentType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.DocumentTypeID {
				local.R.Documents = append(local.R.Documents, foreign)
				if foreign.R == nil {
					foreign.R = &documentR{}
				}
				foreign.R.DocumentType = local
				break
			}
		}
	}

	return nil
}

// AddDocuments adds the given related objects to the existing relationships
// of the document_type, optionally inserting them as new records.
// Appends related to o.R.Documents.
// Sets related.R.DocumentType appropriately.
func (o *DocumentType) AddDocuments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Document) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DocumentTypeID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"documents\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"document_type_id"}),
				strmangle.WhereClause("\"", "\"", 0, documentPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DocumentTypeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &documentTypeR{
			Documents: related,
		}
	} else {
		o.R.Documents = append(o.R.Documents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &documentR{
				DocumentType: o,
			}
		} else {
			rel.R.DocumentType = o
		}
	}
	return nil
}

// DocumentTypes retrieves all the records using an executor.
func DocumentTypes(mods ...qm.QueryMod) documentTypeQuery {
	mods = append(mods, qm.From("\"document_types\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"document_types\".*"})
	}

	return documentTypeQuery{q}
}

// FindDocumentType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDocumentType(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*DocumentType, error) {
	documentTypeObj := &DocumentType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"document_types\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, documentTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from document_types")
	}

	if err = documentTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return documentTypeObj, err
	}

	return documentTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DocumentType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no document_types provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	documentTypeInsertCacheMut.RLock()
	cache, cached := documentTypeInsertCache[key]
	documentTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			documentTypeAllColumns,
			documentTypeColumnsWithDefault,
			documentTypeColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, documentTypeGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(documentTypeType, documentTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(documentTypeType, documentTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"document_types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"document_types\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into document_types")
	}

	if !cached {
		documentTypeInsertCacheMut.Lock()
		documentTypeInsertCache[key] = cache
		documentTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DocumentType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DocumentType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	documentTypeUpdateCacheMut.RLock()
	cache, cached := documentTypeUpdateCache[key]
	documentTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			documentTypeAllColumns,
			documentTypePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, documentTypeGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update document_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"document_types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, documentTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(documentTypeType, documentTypeMapping, append(wl, documentTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update document_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for document_types")
	}

	if !cached {
		documentTypeUpdateCacheMut.Lock()
		documentTypeUpdateCache[key] = cache
		documentTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q documentTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for document_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for document_types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DocumentTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"document_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in documentType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all documentType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DocumentType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no document_types provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentTypeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	documentTypeUpsertCacheMut.RLock()
	cache, cached := documentTypeUpsertCache[key]
	documentTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			documentTypeAllColumns,
			documentTypeColumnsWithDefault,
			documentTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			documentTypeAllColumns,
			documentTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert document_types, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(documentTypePrimaryKeyColumns))
			copy(conflict, documentTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"document_types\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(documentTypeType, documentTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(documentTypeType, documentTypeMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert document_types")
	}

	if !cached {
		documentTypeUpsertCacheMut.Lock()
		documentTypeUpsertCache[key] = cache
		documentTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DocumentType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DocumentType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DocumentType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), documentTypePrimaryKeyMapping)
	sql := "DELETE FROM \"document_types\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from document_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for document_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q documentTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no documentTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from document_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for document_types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DocumentTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(documentTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"document_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from documentType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for document_types")
	}

	if len(documentTypeAfterDeleteHooks) != 0 {
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
func (o *DocumentType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDocumentType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DocumentTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DocumentTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"document_types\".* FROM \"document_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, documentTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DocumentTypeSlice")
	}

	*o = slice

	return nil
}

// DocumentTypeExists checks if the DocumentType row exists.
func DocumentTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"document_types\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if document_types exists")
	}

	return exists, nil
}
