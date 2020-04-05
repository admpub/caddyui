/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Package upload 监听上传文件事件
package upload

import (
	"fmt"
	"os"
	"context"

	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	"github.com/admpub/events"
	"github.com/admpub/nging/application/dbschema"
	"github.com/admpub/nging/application/library/common"
	"github.com/admpub/nging/application/library/config"
	"github.com/admpub/nging/application/registry/upload"
	"github.com/admpub/nging/application/registry/upload/convert"
)

func init() {
	// 当用户文件被删除
	config.Emitter.On(`user-file-deleted`, events.Callback(func(e events.Event) error {
		data := e.Context.Get(`data`).(*dbschema.NgingFile)
		ownerID := e.Context.Uint64(`ownerID`)
		userM := &dbschema.NgingUser{}
		err := userM.Get(nil, db.Cond{`id`: ownerID})
		if err != nil {
			if err == db.ErrNoMoreRows {
				return nil
			}
			return err
		}
		err = userM.SetFields(nil, map[string]interface{}{
			`file_size`: db.Raw(`file_size-` + fmt.Sprintf(`%d`, data.Size)),
			`file_num`:  db.Raw(`file_num-1`),
		}, db.And(
			db.Cond{`id`: ownerID},
			db.Cond{`file_size`: db.Gte(data.Size)},
			db.Cond{`file_num`: db.Gt(0)},
		))
		if err != nil {
			fileM := &dbschema.NgingFile{}
			recv := echo.H{}
			err = fileM.NewParam().SetMW(func(r db.Result) db.Result {
				return r.Select(db.Raw(`SUM(size) AS c`), db.Raw(`COUNT(1) AS n`))
			}).SetRecv(&recv).SetArgs(db.And(
				db.Cond{`owner_type`: `user`},
				db.Cond{`owner_id`: ownerID},
			)).One()
			if err != nil {
				return err
			}
			totalNum := recv.Uint64(`n`)
			totalSize := recv.Uint64(`c`)
			err = userM.SetFields(nil, map[string]interface{}{
				`file_size`: totalSize,
				`file_num`:  totalNum,
			}, db.Cond{`id`: ownerID})
		}
		return err
	}))
	// 当文件被删除
	config.Emitter.On(`file-deleted`, events.Callback(func(e events.Event) error {
		ctx := e.Context.Get(`ctx`).(echo.Context)
		files := e.Context.Get(`files`).([]string)
		data := e.Context.Get(`data`).(*dbschema.NgingFile)
		newStore := upload.StorerGet(data.StorerName)
		if newStore == nil {
			return ctx.E(`存储引擎“%s”未被登记`, data.StorerName)
		}
		key := `storerID`
		ctx.SetStdContext(context.WithValue(ctx, key, data.StorerId))
		storer := newStore(ctx, ``)
		defer storer.Close()
		var errs common.Errors
		otherFormatExtensions := convert.Extensions()
		for _, file := range files {
			if err := storer.Delete(file); err != nil && !os.IsNotExist(err) {
				errs = append(errs, err)
			}
			for _, extension := range otherFormatExtensions {
				if err := storer.Delete(file + extension); err != nil && !os.IsNotExist(err) {
					errs = append(errs, err)
				}
			}
		}
		if len(errs) == 0 {
			return nil
		}
		return errs
	}))
	// 当文件被移动
	config.Emitter.On(`file-moved`, events.Callback(func(e events.Event) error {
		return nil
	}))
}
