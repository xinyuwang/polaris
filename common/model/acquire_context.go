/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package model

import (
	"context"

	api "github.com/polarismesh/polaris-server/common/api/v1"
)

// AcquireContext 每次鉴权请求上下文信息
type AcquireContext struct {

	// RequestContext 请求上下文
	requestContext context.Context

	// Token 本次请求的访问凭据
	token string

	// Module 来自那个业务层（服务注册与服务治理、配置模块）
	module BzModule

	// Operation 本次操作涉及的动作
	operation ResourceOperation

	// Resources 本次
	accessResources map[api.ResourceType][]ResourceEntry

	// Attachment 携带信息，用于操作完权限检查和资源操作的后置处理逻辑，解决信息需要二次查询问题
	attachment map[string]interface{}
}

type acquireContextOption func(authCtx *AcquireContext)

func NewAcquireContext(options ...acquireContextOption) *AcquireContext {
	authCtx := &AcquireContext{
		attachment:      make(map[string]interface{}),
		accessResources: make(map[api.ResourceType][]ResourceEntry),
	}

	for index := range options {
		opt := options[index]
		opt(authCtx)
	}

	return authCtx
}

func WithRequestContext(ctx context.Context) acquireContextOption {
	return func(authCtx *AcquireContext) {
		authCtx.requestContext = ctx
	}
}

func WithToken(token string) acquireContextOption {
	return func(authCtx *AcquireContext) {
		authCtx.token = token
	}
}

func WithModule(module BzModule) acquireContextOption {
	return func(authCtx *AcquireContext) {
		authCtx.module = module
	}
}

func WithOperation(operation ResourceOperation) acquireContextOption {
	return func(authCtx *AcquireContext) {
		authCtx.operation = operation
	}
}

func WithAccessResources(accessResources map[api.ResourceType][]ResourceEntry) acquireContextOption {
	return func(authCtx *AcquireContext) {
		authCtx.accessResources = accessResources
	}
}

func WithAttachment(attachment map[string]interface{}) acquireContextOption {
	return func(authCtx *AcquireContext) {
		for k, v := range attachment {
			authCtx.attachment[k] = v
		}
	}
}

func (authCtx *AcquireContext) GetRequestContext() context.Context {
	return authCtx.requestContext
}

func (authCtx *AcquireContext) SetRequestContext(requestContext context.Context) {
	authCtx.requestContext = requestContext
}

func (authCtx *AcquireContext) GetToken() string {
	return authCtx.token
}

func (authCtx *AcquireContext) GetModule() BzModule {
	return authCtx.module
}

func (authCtx *AcquireContext) GetOperation() ResourceOperation {
	return authCtx.operation
}

func (authCtx *AcquireContext) GetAccessResources() map[api.ResourceType][]ResourceEntry {
	return authCtx.accessResources
}

func (authCtx *AcquireContext) GetAttachment() map[string]interface{} {
	return authCtx.attachment
}
