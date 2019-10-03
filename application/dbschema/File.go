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

type Slice_File []*File

func (s Slice_File) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_File) RangeRaw(fn func(m *File) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// File 文件表
type File struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*File
	namer   func(string) string
	connID  int
	context echo.Context
	
	Id         	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"文件ID" json:"id" xml:"id"`
	OwnerType  	string  	`db:"owner_type" bson:"owner_type" comment:"用户类型" json:"owner_type" xml:"owner_type"`
	OwnerId    	uint64  	`db:"owner_id" bson:"owner_id" comment:"用户ID" json:"owner_id" xml:"owner_id"`
	Name       	string  	`db:"name" bson:"name" comment:"原始文件名" json:"name" xml:"name"`
	SaveName   	string  	`db:"save_name" bson:"save_name" comment:"保存名称" json:"save_name" xml:"save_name"`
	SavePath   	string  	`db:"save_path" bson:"save_path" comment:"文件保存路径" json:"save_path" xml:"save_path"`
	ViewUrl    	string  	`db:"view_url" bson:"view_url" comment:"查看链接" json:"view_url" xml:"view_url"`
	Ext        	string  	`db:"ext" bson:"ext" comment:"文件后缀" json:"ext" xml:"ext"`
	Mime       	string  	`db:"mime" bson:"mime" comment:"文件mime类型" json:"mime" xml:"mime"`
	Type       	string  	`db:"type" bson:"type" comment:"文件类型" json:"type" xml:"type"`
	Size       	uint64  	`db:"size" bson:"size" comment:"文件大小" json:"size" xml:"size"`
	Width      	uint    	`db:"width" bson:"width" comment:"宽度(像素)" json:"width" xml:"width"`
	Height     	uint    	`db:"height" bson:"height" comment:"高度(像素)" json:"height" xml:"height"`
	Dpi        	uint    	`db:"dpi" bson:"dpi" comment:"分辨率" json:"dpi" xml:"dpi"`
	Md5        	string  	`db:"md5" bson:"md5" comment:"文件md5" json:"md5" xml:"md5"`
	StorerName 	string  	`db:"storer_name" bson:"storer_name" comment:"文件保存位置" json:"storer_name" xml:"storer_name"`
	StorerId   	string  	`db:"storer_id" bson:"storer_id" comment:"位置ID" json:"storer_id" xml:"storer_id"`
	Created    	uint    	`db:"created" bson:"created" comment:"上传时间" json:"created" xml:"created"`
	Updated    	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Project    	string  	`db:"project" bson:"project" comment:"项目名称" json:"project" xml:"project"`
	TableId    	uint64  	`db:"table_id" bson:"table_id" comment:"关联表数据id" json:"table_id" xml:"table_id"`
	TableName  	string  	`db:"table_name" bson:"table_name" comment:"关联表名称" json:"table_name" xml:"table_name"`
	FieldName  	string  	`db:"field_name" bson:"field_name" comment:"关联表字段名" json:"field_name" xml:"field_name"`
	Sort       	int64   	`db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Status     	int     	`db:"status" bson:"status" comment:"状态(1-已审核/0-未审核)" json:"status" xml:"status"`
	CategoryId 	uint    	`db:"category_id" bson:"category_id" comment:"分类ID" json:"category_id" xml:"category_id"`
	UsedTimes  	uint    	`db:"used_times" bson:"used_times" comment:"被使用的次数" json:"used_times" xml:"used_times"`
}

func (this *File) Trans() *factory.Transaction {
	return this.trans
}

func (this *File) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *File) SetContext(ctx echo.Context) factory.Model {
	this.context = ctx
	return this
}

func (this *File) Context() echo.Context {
	return this.context
}

func (this *File) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *File) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *File) Objects() []*File {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *File) NewObjects() factory.Ranger {
	return &Slice_File{}
}

func (this *File) InitObjects() *[]*File {
	this.objects = []*File{}
	return &this.objects
}

func (this *File) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *File) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *File) Short_() string {
	return "file"
}

func (this *File) Struct_() string {
	return "File"
}

func (this *File) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *File) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *File) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *File) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *File) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *File) GroupBy(keyField string, inputRows ...[]*File) map[string][]*File {
	var rows []*File
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*File{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*File{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *File) KeyBy(keyField string, inputRows ...[]*File) map[string]*File {
	var rows []*File
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*File{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *File) AsKV(keyField string, valueField string, inputRows ...[]*File) map[string]interface{} {
	var rows []*File
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

func (this *File) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *File) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Type) == 0 { this.Type = "image" }
	err = DBI.EventFire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	if err == nil {
		err = DBI.EventFire("created", this, nil)
	}
	return
}

func (this *File) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	this.Updated = uint(time.Now().Unix())
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Type) == 0 { this.Type = "image" }
	if err = DBI.EventFire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.EventFire("updated", this, mw, args...)
}

func (this *File) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *File) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *File) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	if val, ok := kvset["owner_type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["owner_type"] = "user" } }
	if val, ok := kvset["type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["type"] = "image" } }
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

func (this *File) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { this.Updated = uint(time.Now().Unix())
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Type) == 0 { this.Type = "image" }
		return DBI.EventFire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Type) == 0 { this.Type = "image" }
		return DBI.EventFire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
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

func (this *File) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.EventFire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.EventFire("deleted", this, mw, args...)
}

func (this *File) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *File) Reset() *File {
	this.Id = 0
	this.OwnerType = ``
	this.OwnerId = 0
	this.Name = ``
	this.SaveName = ``
	this.SavePath = ``
	this.ViewUrl = ``
	this.Ext = ``
	this.Mime = ``
	this.Type = ``
	this.Size = 0
	this.Width = 0
	this.Height = 0
	this.Dpi = 0
	this.Md5 = ``
	this.StorerName = ``
	this.StorerId = ``
	this.Created = 0
	this.Updated = 0
	this.Project = ``
	this.TableId = 0
	this.TableName = ``
	this.FieldName = ``
	this.Sort = 0
	this.Status = 0
	this.CategoryId = 0
	this.UsedTimes = 0
	return this
}

func (this *File) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["OwnerType"] = this.OwnerType
	r["OwnerId"] = this.OwnerId
	r["Name"] = this.Name
	r["SaveName"] = this.SaveName
	r["SavePath"] = this.SavePath
	r["ViewUrl"] = this.ViewUrl
	r["Ext"] = this.Ext
	r["Mime"] = this.Mime
	r["Type"] = this.Type
	r["Size"] = this.Size
	r["Width"] = this.Width
	r["Height"] = this.Height
	r["Dpi"] = this.Dpi
	r["Md5"] = this.Md5
	r["StorerName"] = this.StorerName
	r["StorerId"] = this.StorerId
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["Project"] = this.Project
	r["TableId"] = this.TableId
	r["TableName"] = this.TableName
	r["FieldName"] = this.FieldName
	r["Sort"] = this.Sort
	r["Status"] = this.Status
	r["CategoryId"] = this.CategoryId
	r["UsedTimes"] = this.UsedTimes
	return r
}

func (this *File) FromMap(rows map[string]interface{}) {
	for key, value := range rows {
		switch key {
			case "id": this.Id = param.AsUint64(value)
			case "owner_type": this.OwnerType = param.AsString(value)
			case "owner_id": this.OwnerId = param.AsUint64(value)
			case "name": this.Name = param.AsString(value)
			case "save_name": this.SaveName = param.AsString(value)
			case "save_path": this.SavePath = param.AsString(value)
			case "view_url": this.ViewUrl = param.AsString(value)
			case "ext": this.Ext = param.AsString(value)
			case "mime": this.Mime = param.AsString(value)
			case "type": this.Type = param.AsString(value)
			case "size": this.Size = param.AsUint64(value)
			case "width": this.Width = param.AsUint(value)
			case "height": this.Height = param.AsUint(value)
			case "dpi": this.Dpi = param.AsUint(value)
			case "md5": this.Md5 = param.AsString(value)
			case "storer_name": this.StorerName = param.AsString(value)
			case "storer_id": this.StorerId = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
			case "updated": this.Updated = param.AsUint(value)
			case "project": this.Project = param.AsString(value)
			case "table_id": this.TableId = param.AsUint64(value)
			case "table_name": this.TableName = param.AsString(value)
			case "field_name": this.FieldName = param.AsString(value)
			case "sort": this.Sort = param.AsInt64(value)
			case "status": this.Status = param.AsInt(value)
			case "category_id": this.CategoryId = param.AsUint(value)
			case "used_times": this.UsedTimes = param.AsUint(value)
		}
	}
}

func (this *File) Set(key interface{}, value ...interface{}) {
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
				case "Id": this.Id = param.AsUint64(vv)
				case "OwnerType": this.OwnerType = param.AsString(vv)
				case "OwnerId": this.OwnerId = param.AsUint64(vv)
				case "Name": this.Name = param.AsString(vv)
				case "SaveName": this.SaveName = param.AsString(vv)
				case "SavePath": this.SavePath = param.AsString(vv)
				case "ViewUrl": this.ViewUrl = param.AsString(vv)
				case "Ext": this.Ext = param.AsString(vv)
				case "Mime": this.Mime = param.AsString(vv)
				case "Type": this.Type = param.AsString(vv)
				case "Size": this.Size = param.AsUint64(vv)
				case "Width": this.Width = param.AsUint(vv)
				case "Height": this.Height = param.AsUint(vv)
				case "Dpi": this.Dpi = param.AsUint(vv)
				case "Md5": this.Md5 = param.AsString(vv)
				case "StorerName": this.StorerName = param.AsString(vv)
				case "StorerId": this.StorerId = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "Project": this.Project = param.AsString(vv)
				case "TableId": this.TableId = param.AsUint64(vv)
				case "TableName": this.TableName = param.AsString(vv)
				case "FieldName": this.FieldName = param.AsString(vv)
				case "Sort": this.Sort = param.AsInt64(vv)
				case "Status": this.Status = param.AsInt(vv)
				case "CategoryId": this.CategoryId = param.AsUint(vv)
				case "UsedTimes": this.UsedTimes = param.AsUint(vv)
			}
	}
}

func (this *File) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["owner_type"] = this.OwnerType
	r["owner_id"] = this.OwnerId
	r["name"] = this.Name
	r["save_name"] = this.SaveName
	r["save_path"] = this.SavePath
	r["view_url"] = this.ViewUrl
	r["ext"] = this.Ext
	r["mime"] = this.Mime
	r["type"] = this.Type
	r["size"] = this.Size
	r["width"] = this.Width
	r["height"] = this.Height
	r["dpi"] = this.Dpi
	r["md5"] = this.Md5
	r["storer_name"] = this.StorerName
	r["storer_id"] = this.StorerId
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["project"] = this.Project
	r["table_id"] = this.TableId
	r["table_name"] = this.TableName
	r["field_name"] = this.FieldName
	r["sort"] = this.Sort
	r["status"] = this.Status
	r["category_id"] = this.CategoryId
	r["used_times"] = this.UsedTimes
	return r
}

func (this *File) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *File) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

