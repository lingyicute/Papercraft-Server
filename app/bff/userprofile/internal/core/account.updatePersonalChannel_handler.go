// Copyright 2024 Papercraft Authors
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
)

// AccountUpdatePersonalChannel
// account.updatePersonalChannel#d94305e0 channel:InputChannel = Bool;
func (c *UserProfileCore) AccountUpdatePersonalChannel(in *mtproto.TLAccountUpdatePersonalChannel) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("account.updatePersonalChannel blocked, License key from https://papercraft.net required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
