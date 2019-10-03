// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
	
	"time"
)

type Slice_UserRole []*UserRole

func (s Slice_UserRole) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_UserRole) RangeRaw(fn func(m *UserRole) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// UserRole 用户角色
type UserRole struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*UserRole
	namer   func(string) string
	connID  int
	context echo.Context
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name       	string  	`db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created    	uint    	`db:"created" bson:"created" comment:"添加时间" json:"created" xml:"created"`
	Updated    	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled   	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	ParentId   	uint    	`db:"parent_id" bson:"parent_id" comment:"父级ID" json:"parent_id" xml:"parent_id"`
	PermCmd    	string  	`db:"perm_cmd" bson:"perm_cmd" comment:"指令集权限(多个用“,”分隔开)" json:"perm_cmd" xml:"perm_cmd"`
	PermAction 	string  	`db:"perm_action" bson:"perm_action" comment:"行为权限(多个用“,”分隔开)" json:"perm_action" xml:"perm_action"`
}

func (this *UserRole) Trans() *factory.Transaction {
	return this.trans
}

func (this *UserRole) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *UserRole) SetContext(ctx echo.Context) factory.Model {
	this.context = ctx
	return this
}

func (this *UserRole) Context() echo.Context {
	return this.context
}

func (this *UserRole) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *UserRole) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *UserRole) Objects() []*UserRole {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *UserRole) NewObjects() factory.Ranger {
	return &Slice_UserRole{}
}

func (this *UserRole) InitObjects() *[]*UserRole {
	this.objects = []*UserRole{}
	return &this.objects
}

func (this *UserRole) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *UserRole) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *UserRole) Short_() string {
	return "user_role"
}

func (this *UserRole) Struct_() string {
	return "UserRole"
}

func (this *UserRole) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *UserRole) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *UserRole) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *UserRole) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *UserRole) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *UserRole) GroupBy(keyField string, inputRows ...[]*UserRole) map[string][]*UserRole {
	var rows []*UserRole
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*UserRole{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*UserRole{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *UserRole) KeyBy(keyField string, inputRows ...[]*UserRole) map[string]*UserRole {
	var rows []*UserRole
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*UserRole{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *UserRole) AsKV(keyField string, valueField string, inputRows ...[]*UserRole) map[string]interface{} {
	var rows []*UserRole
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *UserRole) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *UserRole) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	err = DBI.EventFire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.EventFire("created", this, nil)
	}
	return
}

func (this *UserRole) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if err = DBI.EventFire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.EventFire("updated", this, mw, args...)
}

func (this *UserRole) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *UserRole) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *UserRole) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	m := *this
	m.FromMap(kvset)
	if err = DBI.EventFire("updating", &m, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.EventFire("updated", &m, mw, args...)
}

func (this *UserRole) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
		return DBI.EventFire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
		return DBI.EventFire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.EventFire("updated", this, mw, args...)
		} else {
			err = DBI.EventFire("created", this, nil)
		}
	} 
	return 
}

func (this *UserRole) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.EventFire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.EventFire("deleted", this, mw, args...)
}

func (this *UserRole) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *UserRole) Reset() *UserRole {
	this.Id = 0
	this.Name = ``
	this.Description = ``
	this.Created = 0
	this.Updated = 0
	this.Disabled = ``
	this.ParentId = 0
	this.PermCmd = ``
	this.PermAction = ``
	return this
}

func (this *UserRole) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["Description"] = this.Description
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["Disabled"] = this.Disabled
	r["ParentId"] = this.ParentId
	r["PermCmd"] = this.PermCmd
	r["PermAction"] = this.PermAction
	return r
}

func (this *UserRole) FromMap(rows map[string]interface{}) {
	for key, value := range rows {
		switch key {
			case "id": this.Id = param.AsUint(value)
			case "name": this.Name = param.AsString(value)
			case "description": this.Description = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
			case "updated": this.Updated = param.AsUint(value)
			case "disabled": this.Disabled = param.AsString(value)
			case "parent_id": this.ParentId = param.AsUint(value)
			case "perm_cmd": this.PermCmd = param.AsString(value)
			case "perm_action": this.PermAction = param.AsString(value)
		}
	}
}

func (this *UserRole) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "Description": this.Description = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "ParentId": this.ParentId = param.AsUint(vv)
				case "PermCmd": this.PermCmd = param.AsString(vv)
				case "PermAction": this.PermAction = param.AsString(vv)
			}
	}
}

func (this *UserRole) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["name"] = this.Name
	r["description"] = this.Description
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["disabled"] = this.Disabled
	r["parent_id"] = this.ParentId
	r["perm_cmd"] = this.PermCmd
	r["perm_action"] = this.PermAction
	return r
}

func (this *UserRole) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *UserRole) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

