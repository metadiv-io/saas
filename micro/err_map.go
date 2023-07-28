package micro

import "github.com/metadiv-io/ginger"

const (
	ERR_CODE_WORKSPACE_NOT_FOUND = "ERR-A2arjntTnywfPFRyJFbKh"
	ERR_CODE_NOT_ENOUGH_CREDIT   = "ERR-3CKQfiqirajbCYtjzNHpF"
)

func init() {
	ginger.ErrMap.Register(ERR_CODE_WORKSPACE_NOT_FOUND, ginger.LOCALE_EN, "Workspace not found")
	ginger.ErrMap.Register(ERR_CODE_WORKSPACE_NOT_FOUND, ginger.LOCALE_ZHT, "找不到工作區")
	ginger.ErrMap.Register(ERR_CODE_WORKSPACE_NOT_FOUND, ginger.LOCALE_ZHS, "找不到工作区")

	ginger.ErrMap.Register(ERR_CODE_NOT_ENOUGH_CREDIT, ginger.LOCALE_EN, "Not enough credit")
	ginger.ErrMap.Register(ERR_CODE_NOT_ENOUGH_CREDIT, ginger.LOCALE_ZHT, "餘額不足")
	ginger.ErrMap.Register(ERR_CODE_NOT_ENOUGH_CREDIT, ginger.LOCALE_ZHS, "余额不足")
}
