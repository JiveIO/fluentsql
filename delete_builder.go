package fluentsql

import "strings"

// ===========================================================================================================
//										Delete Builder :: Structure
// ===========================================================================================================

// DeleteBuilder struct
/*
DELETE [LOW_PRIORITY] [QUICK] [IGNORE] FROM tbl_name [[AS] tbl_alias]
    [PARTITION (partition_name [, partition_name] ...)]
    [WHERE where_condition]
    [ORDER BY ...]
    [LIMIT row_count]
*/
type DeleteBuilder struct {
	deleteStatement  Delete
	whereStatement   Where
	orderByStatement OrderBy
	limitStatement   Limit
}

// DeleteInstance Delete Builder constructor
func DeleteInstance() *DeleteBuilder {
	return &DeleteBuilder{}
}

// ===========================================================================================================
//										Update Builder :: Operators
// ===========================================================================================================

func (db *DeleteBuilder) String() string {
	var queryParts []string

	queryParts = append(queryParts, db.deleteStatement.String())

	whereSql := db.whereStatement.String()
	if whereSql != "" {
		queryParts = append(queryParts, whereSql)
	}

	orderBySql := db.orderByStatement.String()
	if orderBySql != "" {
		queryParts = append(queryParts, orderBySql)
	}

	limitSql := db.limitStatement.String()
	if limitSql != "" {
		queryParts = append(queryParts, limitSql)
	}

	sql := strings.Join(queryParts, " ")

	return sql
}

// Delete builder
func (db *DeleteBuilder) Delete(table string, alias ...string) *DeleteBuilder {
	db.deleteStatement.Table = table

	if len(alias) > 0 {
		db.deleteStatement.Alias = alias[0]
	}

	return db
}

// Where builder
func (db *DeleteBuilder) Where(Field any, Opt WhereOpt, Value any) *DeleteBuilder {
	db.whereStatement.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return db
}

// WhereOr builder
func (db *DeleteBuilder) WhereOr(Field any, Opt WhereOpt, Value any) *DeleteBuilder {
	db.whereStatement.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: Or,
	})

	return db
}
