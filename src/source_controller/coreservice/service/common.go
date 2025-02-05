/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"configcenter/src/common/http/rest"
	"configcenter/src/common/metadata"
)

func (s *coreService) GetDistinctField(ctx *rest.Contexts) {
	option := new(metadata.DistinctFieldOption)
	if err := ctx.DecodeInto(option); err != nil {
		ctx.RespAutoError(err)
		return
	}

	rawErr := option.Validate()
	if rawErr.ErrCode != 0 {
		ctx.RespAutoError(rawErr.ToCCError(ctx.Kit.CCError))
		return
	}

	ret, err := s.core.CommonOperation().GetDistinctField(ctx.Kit, option)

	if err != nil {
		ctx.RespAutoError(err)
		return
	}

	ctx.RespEntity(ret)
}

// GetDistinctCount 根据条件获取指定表中满足条件数据的数量
func (s *coreService) GetDistinctCount(ctx *rest.Contexts) {
	option := new(metadata.DistinctFieldOption)
	if err := ctx.DecodeInto(option); err != nil {
		ctx.RespAutoError(err)
		return
	}

	rawErr := option.Validate()
	if rawErr.ErrCode != 0 {
		ctx.RespAutoError(rawErr.ToCCError(ctx.Kit.CCError))
		return
	}

	count, err := s.core.CommonOperation().GetDistinctCount(ctx.Kit, option)

	if err != nil {
		ctx.RespAutoError(err)
		return
	}

	ctx.RespEntity(count)
}
