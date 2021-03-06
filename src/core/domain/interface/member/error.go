/**
 * Copyright 2014 @ S1N1 Team.
 * name :
 * author : jarryliu
 * date : 2014-02-05 20:12
 * description :
 * history :
 */
package member

import (
	"go2o/src/core/infrastructure/domain"
)

var (
	ErrSessionTimeout *domain.DomainError = domain.NewDomainError(
		"member_session_time_out", "会员会话超时")

	ErrInvalidSession *domain.DomainError = domain.NewDomainError(
		"member_invalid_session", "异常会话")

	ErrNoSuchDeliverAddress *domain.DomainError = domain.NewDomainError(
		"member_no_such_deliver_address", "配送地址错误")

	ErrNoSuchMember *domain.DomainError = domain.NewDomainError(
		"member_no_such_member", "会员不存在")

	ErrDeliverAddressLen *domain.DomainError = domain.NewDomainError(
		"err_deliver_address_len", "请填写详细的配送地址")

	ErrDeliverRealNameIsNull *domain.DomainError = domain.NewDomainError(
		"err_deliver_real_name_is_null", "收货人不能为空")

	ErrDeliverPhoneIsNull *domain.DomainError = domain.NewDomainError(
		"err_deliver_phone_is_null", "电话不能为空")

	ErrPwdCannotSame *domain.DomainError = domain.NewDomainError(
		"Err_Pwd_Can_not_Same", "新密码不能与旧密码相同")

	ErrPwdPldPwdNotRight *domain.DomainError = domain.NewDomainError(
		"Err_Pwd_Pld_Pwd_Not_Right", "原密码不正确")
)
