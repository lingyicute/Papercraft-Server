/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package tos_helper

import (
	"github.com/lingyicute/papercraft-server/app/bff/tos/internal/config"
	"github.com/lingyicute/papercraft-server/app/bff/tos/internal/server/grpc/service"
	"github.com/lingyicute/papercraft-server/app/bff/tos/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
