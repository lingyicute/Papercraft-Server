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
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/service/biz/user/user"
)

// UserUpdateEmojiStatus
// user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;
func (c *UserCore) UserUpdateEmojiStatus(in *user.TLUserUpdateEmojiStatus) (*mtproto.Bool, error) {
	rB := c.svcCtx.Dao.UpdateUserEmojiStatus(
		c.ctx,
		in.UserId,
		in.EmojiStatusDocumentId,
		in.EmojiStatusUntil)

	return mtproto.ToBool(rB), nil
}
