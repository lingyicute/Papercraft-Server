// Copyright 2022 Yomi
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
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// AccountGetAccountTTL
// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (c *AccountCore) AccountGetAccountTTL(in *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	days, err := c.svcCtx.Dao.UserClient.UserGetAccountDaysTTL(c.ctx, &userpb.TLUserGetAccountDaysTTL{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("account.getAccountTTL - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLAccountDaysTTL(&mtproto.AccountDaysTTL{
		Days: days.Days,
	}).To_AccountDaysTTL(), nil
}
