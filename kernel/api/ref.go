// SiYuan - Build Your Eternal Digital Garden
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package api

import (
	"net/http"

	"github.com/88250/gulu"
	"github.com/gin-gonic/gin"
	"github.com/siyuan-note/siyuan/kernel/model"
	"github.com/siyuan-note/siyuan/kernel/util"
)

func refreshBacklink(c *gin.Context) {
	ret := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, ret)

	arg, ok := util.JsonArg(c, ret)
	if !ok {
		return
	}

	id := arg["id"].(string)
	model.RefreshBacklink(id)
}

func getBacklink(c *gin.Context) {
	ret := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, ret)

	arg, ok := util.JsonArg(c, ret)
	if !ok {
		return
	}

	if nil == arg["id"] {
		return
	}

	id := arg["id"].(string)
	keyword := arg["k"].(string)
	mentionKeyword := arg["mk"].(string)
	beforeLen := arg["beforeLen"].(float64)
	boxID, backlinks, backmentions, linkRefsCount, mentionsCount := model.BuildTreeBacklink(id, keyword, mentionKeyword, int(beforeLen))
	ret.Data = map[string]interface{}{
		"backlinks":     backlinks,
		"linkRefsCount": linkRefsCount,
		"backmentions":  backmentions,
		"mentionsCount": mentionsCount,
		"k":             keyword,
		"mk":            mentionKeyword,
		"box":           boxID,
	}
	util.RandomSleep(200, 500)
}

func createBacklink(c *gin.Context) {
	ret := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, ret)

	arg, ok := util.JsonArg(c, ret)
	if !ok {
		return
	}

	defID := arg["defID"].(string)
	refID := arg["refID"].(string)
	refText := arg["refText"].(string)
	isDynamic := arg["isDynamic"].(bool)
	refRootID, err := model.CreateBacklink(defID, refID, refText, isDynamic)
	if nil != err {
		ret.Code = -1
		ret.Msg = err.Error()
		return
	}
	ret.Data = map[string]interface{}{
		"defID":     defID,
		"refID":     refID,
		"refRootID": refRootID,
		"refText":   refText,
	}
}
