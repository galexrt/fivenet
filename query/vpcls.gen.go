// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/galexrt/arpanet/model"
)

func newVpcL(db *gorm.DB, opts ...gen.DOOption) vpcL {
	_vpcL := vpcL{}

	_vpcL.vpcLDo.UseDB(db, opts...)
	_vpcL.vpcLDo.UseModel(&model.VpcL{})

	tableName := _vpcL.vpcLDo.TableName()
	_vpcL.ALL = field.NewAsterisk(tableName)
	_vpcL.PlayerID = field.NewString(tableName, "playerId")
	_vpcL.Coordsx = field.NewString(tableName, "coordsx")
	_vpcL.Coordsy = field.NewString(tableName, "coordsy")
	_vpcL.NET = field.NewString(tableName, "NET")

	_vpcL.fillFieldMap()

	return _vpcL
}

type vpcL struct {
	vpcLDo

	ALL      field.Asterisk
	PlayerID field.String
	Coordsx  field.String
	Coordsy  field.String
	NET      field.String

	fieldMap map[string]field.Expr
}

func (v vpcL) Table(newTableName string) *vpcL {
	v.vpcLDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vpcL) As(alias string) *vpcL {
	v.vpcLDo.DO = *(v.vpcLDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vpcL) updateTableName(table string) *vpcL {
	v.ALL = field.NewAsterisk(table)
	v.PlayerID = field.NewString(table, "playerId")
	v.Coordsx = field.NewString(table, "coordsx")
	v.Coordsy = field.NewString(table, "coordsy")
	v.NET = field.NewString(table, "NET")

	v.fillFieldMap()

	return v
}

func (v *vpcL) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vpcL) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 4)
	v.fieldMap["playerId"] = v.PlayerID
	v.fieldMap["coordsx"] = v.Coordsx
	v.fieldMap["coordsy"] = v.Coordsy
	v.fieldMap["NET"] = v.NET
}

func (v vpcL) clone(db *gorm.DB) vpcL {
	v.vpcLDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vpcL) replaceDB(db *gorm.DB) vpcL {
	v.vpcLDo.ReplaceDB(db)
	return v
}

type vpcLDo struct{ gen.DO }

type IVpcLDo interface {
	gen.SubQuery
	Debug() IVpcLDo
	WithContext(ctx context.Context) IVpcLDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVpcLDo
	WriteDB() IVpcLDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVpcLDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVpcLDo
	Not(conds ...gen.Condition) IVpcLDo
	Or(conds ...gen.Condition) IVpcLDo
	Select(conds ...field.Expr) IVpcLDo
	Where(conds ...gen.Condition) IVpcLDo
	Order(conds ...field.Expr) IVpcLDo
	Distinct(cols ...field.Expr) IVpcLDo
	Omit(cols ...field.Expr) IVpcLDo
	Join(table schema.Tabler, on ...field.Expr) IVpcLDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVpcLDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVpcLDo
	Group(cols ...field.Expr) IVpcLDo
	Having(conds ...gen.Condition) IVpcLDo
	Limit(limit int) IVpcLDo
	Offset(offset int) IVpcLDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVpcLDo
	Unscoped() IVpcLDo
	Create(values ...*model.VpcL) error
	CreateInBatches(values []*model.VpcL, batchSize int) error
	Save(values ...*model.VpcL) error
	First() (*model.VpcL, error)
	Take() (*model.VpcL, error)
	Last() (*model.VpcL, error)
	Find() ([]*model.VpcL, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VpcL, err error)
	FindInBatches(result *[]*model.VpcL, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VpcL) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVpcLDo
	Assign(attrs ...field.AssignExpr) IVpcLDo
	Joins(fields ...field.RelationField) IVpcLDo
	Preload(fields ...field.RelationField) IVpcLDo
	FirstOrInit() (*model.VpcL, error)
	FirstOrCreate() (*model.VpcL, error)
	FindByPage(offset int, limit int) (result []*model.VpcL, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVpcLDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v vpcLDo) Debug() IVpcLDo {
	return v.withDO(v.DO.Debug())
}

func (v vpcLDo) WithContext(ctx context.Context) IVpcLDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vpcLDo) ReadDB() IVpcLDo {
	return v.Clauses(dbresolver.Read)
}

func (v vpcLDo) WriteDB() IVpcLDo {
	return v.Clauses(dbresolver.Write)
}

func (v vpcLDo) Session(config *gorm.Session) IVpcLDo {
	return v.withDO(v.DO.Session(config))
}

func (v vpcLDo) Clauses(conds ...clause.Expression) IVpcLDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vpcLDo) Returning(value interface{}, columns ...string) IVpcLDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vpcLDo) Not(conds ...gen.Condition) IVpcLDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vpcLDo) Or(conds ...gen.Condition) IVpcLDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vpcLDo) Select(conds ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vpcLDo) Where(conds ...gen.Condition) IVpcLDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vpcLDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IVpcLDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v vpcLDo) Order(conds ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vpcLDo) Distinct(cols ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vpcLDo) Omit(cols ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vpcLDo) Join(table schema.Tabler, on ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vpcLDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vpcLDo) RightJoin(table schema.Tabler, on ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vpcLDo) Group(cols ...field.Expr) IVpcLDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vpcLDo) Having(conds ...gen.Condition) IVpcLDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vpcLDo) Limit(limit int) IVpcLDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vpcLDo) Offset(offset int) IVpcLDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vpcLDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVpcLDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vpcLDo) Unscoped() IVpcLDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vpcLDo) Create(values ...*model.VpcL) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vpcLDo) CreateInBatches(values []*model.VpcL, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vpcLDo) Save(values ...*model.VpcL) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vpcLDo) First() (*model.VpcL, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VpcL), nil
	}
}

func (v vpcLDo) Take() (*model.VpcL, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VpcL), nil
	}
}

func (v vpcLDo) Last() (*model.VpcL, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VpcL), nil
	}
}

func (v vpcLDo) Find() ([]*model.VpcL, error) {
	result, err := v.DO.Find()
	return result.([]*model.VpcL), err
}

func (v vpcLDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VpcL, err error) {
	buf := make([]*model.VpcL, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vpcLDo) FindInBatches(result *[]*model.VpcL, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vpcLDo) Attrs(attrs ...field.AssignExpr) IVpcLDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vpcLDo) Assign(attrs ...field.AssignExpr) IVpcLDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vpcLDo) Joins(fields ...field.RelationField) IVpcLDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vpcLDo) Preload(fields ...field.RelationField) IVpcLDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vpcLDo) FirstOrInit() (*model.VpcL, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VpcL), nil
	}
}

func (v vpcLDo) FirstOrCreate() (*model.VpcL, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VpcL), nil
	}
}

func (v vpcLDo) FindByPage(offset int, limit int) (result []*model.VpcL, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v vpcLDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vpcLDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vpcLDo) Delete(models ...*model.VpcL) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vpcLDo) withDO(do gen.Dao) *vpcLDo {
	v.DO = *do.(*gen.DO)
	return v
}
