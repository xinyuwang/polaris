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

package defaultauth

import (
	"context"

	"github.com/polarismesh/polaris-server/auth"
	api "github.com/polarismesh/polaris-server/common/api/v1"
	"github.com/polarismesh/polaris-server/common/utils"
)

// UserServer 用户数据管理 server
type userServerAuth struct {
	authMgn *defaultAuthManager
	target  auth.UserServer
}

// newUserServerWithAuth 构建一个带鉴权检查的 UserServer, 无论鉴权逻辑是否开启，这里的操作都必须经过用户角色检查
func newUserServerWithAuth(authMgn *defaultAuthManager, target auth.UserServer) auth.UserServer {
	return &userServerAuth{
		authMgn: authMgn,
		target:  target,
	}
}

// CreateUser
func (svr *userServerAuth) CreateUsers(ctx context.Context, req []*api.User) *api.BatchWriteResponse {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewBatchWriteResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		resp := api.NewBatchWriteResponse(api.ExecuteSuccess)
		resp.Collect(errResp)
		return resp
	}

	return svr.target.CreateUsers(ctx, req)
}

// UpdateUser
func (svr *userServerAuth) UpdateUser(ctx context.Context, user *api.User) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, false)
	if errResp != nil {
		errResp.User = user
		return errResp
	}

	return svr.target.UpdateUser(ctx, user)
}

// DeleteUser
func (svr *userServerAuth) DeleteUser(ctx context.Context, user *api.User) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		errResp.User = user
		return errResp
	}

	return svr.target.DeleteUser(ctx, user)
}

// ListUsers
func (svr *userServerAuth) ListUsers(ctx context.Context, query map[string]string) *api.BatchQueryResponse {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewBatchQueryResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		return api.NewBatchQueryResponseWithMsg(errResp.GetCode().Value, errResp.Info.Value)
	}

	return svr.target.ListUsers(ctx, query)
}

// GetUserToken
func (svr *userServerAuth) GetUserToken(ctx context.Context, filter map[string]string) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, false)
	if errResp != nil {
		return errResp
	}

	return svr.target.GetUserToken(ctx, filter)
}

// ChangeUserTokenStatus
func (svr *userServerAuth) ChangeUserTokenStatus(ctx context.Context, user *api.User) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		errResp.User = user
		return errResp
	}

	return svr.target.ChangeUserTokenStatus(ctx, user)
}

// RefreshUserToken
func (svr *userServerAuth) RefreshUserToken(ctx context.Context, user *api.User) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, false)
	if errResp != nil {
		errResp.User = user
		return errResp
	}

	return svr.target.RefreshUserToken(ctx, user)
}

// CreateUserGroup
func (svr *userServerAuth) CreateUserGroup(ctx context.Context, group *api.UserGroup) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		errResp.UserGroup = group
		return errResp
	}

	return svr.target.CreateUserGroup(ctx, group)
}

// UpdateUserGroup
func (svr *userServerAuth) UpdateUserGroup(ctx context.Context, group *api.ModifyUserGroup) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		errResp.ModifyUserGroup = group
		return errResp
	}

	return svr.target.UpdateUserGroup(ctx, group)
}

// DeleteUserGroup
func (svr *userServerAuth) DeleteUserGroup(ctx context.Context, group *api.UserGroup) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		errResp.UserGroup = group
		return errResp
	}

	return svr.target.DeleteUserGroup(ctx, group)
}

// ListGroups 查看用户组
func (svr *userServerAuth) ListGroups(ctx context.Context, query map[string]string) *api.BatchQueryResponse {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewBatchQueryResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, false)
	if errResp != nil {
		return api.NewBatchQueryResponseWithMsg(errResp.GetCode().Value, errResp.Info.Value)
	}

	return svr.target.ListGroups(ctx, query)
}

// ListUserByGroup
func (svr *userServerAuth) ListUserByGroup(ctx context.Context, query map[string]string) *api.BatchQueryResponse {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewBatchQueryResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, false)
	if errResp != nil {
		return api.NewBatchQueryResponseWithMsg(errResp.GetCode().Value, errResp.Info.Value)
	}

	return svr.target.ListUserByGroup(ctx, query)
}

// ListUserLinkGroups
func (svr *userServerAuth) ListUserLinkGroups(ctx context.Context, query map[string]string) *api.BatchQueryResponse {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewBatchQueryResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, false)
	if errResp != nil {
		return api.NewBatchQueryResponseWithMsg(errResp.GetCode().Value, errResp.Info.Value)
	}

	return svr.target.ListUserLinkGroups(ctx, query)
}

// GetUserGroupToken
func (svr *userServerAuth) GetUserGroupToken(ctx context.Context, filter map[string]string) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, false)
	if errResp != nil {
		return errResp
	}

	return svr.target.GetUserGroupToken(ctx, filter)
}

// ChangeUserGroupTokenStatus
func (svr *userServerAuth) ChangeUserGroupTokenStatus(ctx context.Context, group *api.UserGroup) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		errResp.UserGroup = group
		return errResp
	}

	return svr.target.ChangeUserGroupTokenStatus(ctx, group)
}

// RefreshUserGroupToken
func (svr *userServerAuth) RefreshUserGroupToken(ctx context.Context, group *api.UserGroup) *api.Response {
	authToken := utils.ParseAuthToken(ctx)
	if authToken == "" {
		return api.NewResponse(api.NotAllowedAccess)
	}

	ctx, errResp := verifyAuth(ctx, svr.authMgn, authToken, true)
	if errResp != nil {
		errResp.UserGroup = group
		return errResp
	}

	return svr.target.RefreshUserGroupToken(ctx, group)
}
