package micro

import "github.com/metadiv-io/saas/constant"

var ErrMap = &errMap{}

type errMap map[string]errMsg

type errMsg map[string]string

func (m *errMap) Register(code string, locale string, message string) {
	if _, ok := (*m)[locale]; !ok {
		(*m)[locale] = make(errMsg)
	}
	(*m)[locale][code] = message
}

func (m *errMap) Get(code string, locale string) string {
	if _, ok := (*m)[locale]; !ok {
		return ""
	}
	if _, ok := (*m)[locale][code]; !ok {
		return ""
	}
	return (*m)[locale][code]
}

func init() {
	ErrMap.Register(constant.ERR_CODE_UNAUTHORIZED, constant.LOCALE_EN, "Unauthorized")
	ErrMap.Register(constant.ERR_CODE_UNAUTHORIZED, constant.LOCALE_ZHT, "未經授權 (Unauthorized))")
	ErrMap.Register(constant.ERR_CODE_UNAUTHORIZED, constant.LOCALE_ZHS, "未经授权 (Unauthorized)")

	ErrMap.Register(constant.ERR_CODE_FORBIDDEN, constant.LOCALE_EN, "Forbidden")
	ErrMap.Register(constant.ERR_CODE_FORBIDDEN, constant.LOCALE_ZHT, "禁止 (Forbidden)")
	ErrMap.Register(constant.ERR_CODE_FORBIDDEN, constant.LOCALE_ZHS, "禁止 (Forbidden)")

	ErrMap.Register(constant.ERR_CODE_NOT_ENOUGH_CREDIT, constant.LOCALE_EN, "Not enough credit")
	ErrMap.Register(constant.ERR_CODE_NOT_ENOUGH_CREDIT, constant.LOCALE_ZHT, "餘額不足")
	ErrMap.Register(constant.ERR_CODE_NOT_ENOUGH_CREDIT, constant.LOCALE_ZHS, "余额不足")

	ErrMap.Register(constant.ERR_CODE_WORKSPACE_NOT_FOUND, constant.LOCALE_EN, "Workspace not found")
	ErrMap.Register(constant.ERR_CODE_WORKSPACE_NOT_FOUND, constant.LOCALE_ZHT, "找不到工作區")
	ErrMap.Register(constant.ERR_CODE_WORKSPACE_NOT_FOUND, constant.LOCALE_ZHS, "找不到工作区")

	ErrMap.Register(constant.ERR_CODE_INTERNAL_SERVER_ERROR, constant.LOCALE_EN, "Internal server error")
	ErrMap.Register(constant.ERR_CODE_INTERNAL_SERVER_ERROR, constant.LOCALE_ZHT, "內部伺服器錯誤")
	ErrMap.Register(constant.ERR_CODE_INTERNAL_SERVER_ERROR, constant.LOCALE_ZHS, "内部服务器错误")
}
