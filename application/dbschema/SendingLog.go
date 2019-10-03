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

type Slice_SendingLog []*SendingLog

func (s Slice_SendingLog) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_SendingLog) RangeRaw(fn func(m *SendingLog) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// SendingLog 邮件短信等发送日志
type SendingLog struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*SendingLog
	namer   func(string) string
	connID  int
	context echo.Context
	
	Id              	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Created         	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	SentAt          	uint    	`db:"sent_at" bson:"sent_at" comment:"发送时间" json:"sent_at" xml:"sent_at"`
	SourceId        	uint64  	`db:"source_id" bson:"source_id" comment:"来源ID" json:"source_id" xml:"source_id"`
	SourceType      	string  	`db:"source_type" bson:"source_type" comment:"来源类型" json:"source_type" xml:"source_type"`
	Disabled        	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Method          	string  	`db:"method" bson:"method" comment:"发送方式(mobile-手机;email-邮箱)" json:"method" xml:"method"`
	To              	string  	`db:"to" bson:"to" comment:"发送目标" json:"to" xml:"to"`
	Provider        	string  	`db:"provider" bson:"provider" comment:"发送平台" json:"provider" xml:"provider"`
	Result          	string  	`db:"result" bson:"result" comment:"发送结果描述" json:"result" xml:"result"`
	Status          	string  	`db:"status" bson:"status" comment:"发送状态(none-无需发送)" json:"status" xml:"status"`
	Retries         	uint    	`db:"retries" bson:"retries" comment:"重试次数" json:"retries" xml:"retries"`
	Content         	string  	`db:"content" bson:"content" comment:"发送消息内容" json:"content" xml:"content"`
	Params          	string  	`db:"params" bson:"params" comment:"发送消息参数(JSON)" json:"params" xml:"params"`
	AppointmentTime 	uint    	`db:"appointment_time" bson:"appointment_time" comment:"预约发送时间" json:"appointment_time" xml:"appointment_time"`
}

func (this *SendingLog) Trans() *factory.Transaction {
	return this.trans
}

func (this *SendingLog) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *SendingLog) SetContext(ctx echo.Context) factory.Model {
	this.context = ctx
	return this
}

func (this *SendingLog) Context() echo.Context {
	return this.context
}

func (this *SendingLog) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *SendingLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *SendingLog) Objects() []*SendingLog {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *SendingLog) NewObjects() factory.Ranger {
	return &Slice_SendingLog{}
}

func (this *SendingLog) InitObjects() *[]*SendingLog {
	this.objects = []*SendingLog{}
	return &this.objects
}

func (this *SendingLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *SendingLog) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *SendingLog) Short_() string {
	return "sending_log"
}

func (this *SendingLog) Struct_() string {
	return "SendingLog"
}

func (this *SendingLog) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *SendingLog) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *SendingLog) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *SendingLog) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *SendingLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *SendingLog) GroupBy(keyField string, inputRows ...[]*SendingLog) map[string][]*SendingLog {
	var rows []*SendingLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*SendingLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*SendingLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *SendingLog) KeyBy(keyField string, inputRows ...[]*SendingLog) map[string]*SendingLog {
	var rows []*SendingLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*SendingLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *SendingLog) AsKV(keyField string, valueField string, inputRows ...[]*SendingLog) map[string]interface{} {
	var rows []*SendingLog
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

func (this *SendingLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *SendingLog) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.Status) == 0 { this.Status = "waiting" }
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

func (this *SendingLog) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.Status) == 0 { this.Status = "waiting" }
	if err = DBI.EventFire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.EventFire("updated", this, mw, args...)
}

func (this *SendingLog) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *SendingLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *SendingLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	if val, ok := kvset["source_type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["source_type"] = "user" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["method"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["method"] = "mobile" } }
	if val, ok := kvset["status"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["status"] = "waiting" } }
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

func (this *SendingLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { 
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.Status) == 0 { this.Status = "waiting" }
		return DBI.EventFire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.Status) == 0 { this.Status = "waiting" }
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

func (this *SendingLog) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.EventFire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.EventFire("deleted", this, mw, args...)
}

func (this *SendingLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *SendingLog) Reset() *SendingLog {
	this.Id = 0
	this.Created = 0
	this.SentAt = 0
	this.SourceId = 0
	this.SourceType = ``
	this.Disabled = ``
	this.Method = ``
	this.To = ``
	this.Provider = ``
	this.Result = ``
	this.Status = ``
	this.Retries = 0
	this.Content = ``
	this.Params = ``
	this.AppointmentTime = 0
	return this
}

func (this *SendingLog) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Created"] = this.Created
	r["SentAt"] = this.SentAt
	r["SourceId"] = this.SourceId
	r["SourceType"] = this.SourceType
	r["Disabled"] = this.Disabled
	r["Method"] = this.Method
	r["To"] = this.To
	r["Provider"] = this.Provider
	r["Result"] = this.Result
	r["Status"] = this.Status
	r["Retries"] = this.Retries
	r["Content"] = this.Content
	r["Params"] = this.Params
	r["AppointmentTime"] = this.AppointmentTime
	return r
}

func (this *SendingLog) FromMap(rows map[string]interface{}) {
	for key, value := range rows {
		switch key {
			case "id": this.Id = param.AsUint64(value)
			case "created": this.Created = param.AsUint(value)
			case "sent_at": this.SentAt = param.AsUint(value)
			case "source_id": this.SourceId = param.AsUint64(value)
			case "source_type": this.SourceType = param.AsString(value)
			case "disabled": this.Disabled = param.AsString(value)
			case "method": this.Method = param.AsString(value)
			case "to": this.To = param.AsString(value)
			case "provider": this.Provider = param.AsString(value)
			case "result": this.Result = param.AsString(value)
			case "status": this.Status = param.AsString(value)
			case "retries": this.Retries = param.AsUint(value)
			case "content": this.Content = param.AsString(value)
			case "params": this.Params = param.AsString(value)
			case "appointment_time": this.AppointmentTime = param.AsUint(value)
		}
	}
}

func (this *SendingLog) Set(key interface{}, value ...interface{}) {
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
				case "Created": this.Created = param.AsUint(vv)
				case "SentAt": this.SentAt = param.AsUint(vv)
				case "SourceId": this.SourceId = param.AsUint64(vv)
				case "SourceType": this.SourceType = param.AsString(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "Method": this.Method = param.AsString(vv)
				case "To": this.To = param.AsString(vv)
				case "Provider": this.Provider = param.AsString(vv)
				case "Result": this.Result = param.AsString(vv)
				case "Status": this.Status = param.AsString(vv)
				case "Retries": this.Retries = param.AsUint(vv)
				case "Content": this.Content = param.AsString(vv)
				case "Params": this.Params = param.AsString(vv)
				case "AppointmentTime": this.AppointmentTime = param.AsUint(vv)
			}
	}
}

func (this *SendingLog) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["created"] = this.Created
	r["sent_at"] = this.SentAt
	r["source_id"] = this.SourceId
	r["source_type"] = this.SourceType
	r["disabled"] = this.Disabled
	r["method"] = this.Method
	r["to"] = this.To
	r["provider"] = this.Provider
	r["result"] = this.Result
	r["status"] = this.Status
	r["retries"] = this.Retries
	r["content"] = this.Content
	r["params"] = this.Params
	r["appointment_time"] = this.AppointmentTime
	return r
}

func (this *SendingLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *SendingLog) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

