/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package core

import (
	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/messenger/sync/sync"
)

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (c *SyncCore) SyncPushUpdates(in *sync.TLSyncPushUpdates) (*mtproto.Void, error) {
	var (
		userId  = in.GetUserId()
		updates = in.GetUpdates()
	)

	notification, err := c.processUpdates(syncTypeUser, userId, false, updates)
	if err != nil {
		c.Logger.Errorf("sync.updatesNotMe - error: %v", err)
		return nil, err
	}

	c.pushUpdatesToSession(syncTypeUser, userId, 0, nil, nil, nil, updates, notification)

	return mtproto.EmptyVoid, nil
}
