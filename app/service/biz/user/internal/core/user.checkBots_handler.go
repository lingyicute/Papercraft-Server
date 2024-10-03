// Copyright 2022 Papercraft Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package core

import (
	"github.com/lingyicute/papercraft-server/app/service/biz/user/user"
)

// UserCheckBots
// user.checkBots id:Vector<long> = Vector<long>;
func (c *UserCore) UserCheckBots(in *user.TLUserCheckBots) (*user.Vector_Long, error) {
	var (
		rVals []int64
	)

	if len(in.Id) > 0 {
		rVals, _ = c.svcCtx.Dao.UsersDAO.SelectBots(c.ctx, in.Id)
		if rVals == nil {
			rVals = []int64{}
		}
	}

	return &user.Vector_Long{
		Datas: rVals,
	}, nil
}
