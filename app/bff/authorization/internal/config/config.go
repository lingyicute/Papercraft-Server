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

package config

import (
	kafka "github.com/papercraft/marmota/pkg/mq"
	"github.com/lingyicute/papercraft-server/pkg/code/conf"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	KV                        kv.KvConf
	Code                      *conf.SmsVerifyCodeConfig
	UserClient                zrpc.RpcClientConf
	AuthsessionClient         zrpc.RpcClientConf
	ChatClient                zrpc.RpcClientConf
	StatusClient              zrpc.RpcClientConf
	UsernameClient            zrpc.RpcClientConf
	MsgClient                 zrpc.RpcClientConf
	SyncClient                *kafka.KafkaProducerConf
	SignInServiceNotification []conf.MessageEntityConfig `json:",optional"`
	SignInMessage             []conf.MessageEntityConfig `json:",optional"`
}
