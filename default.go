package ebby

import (
	"ebby/control"
)

// Init : 初始化
func defaultInitialFunc(ctrl *control.Control) {
	ctrl.SData().Tool.Logger.Info("[control] Started")
}

// BeforeExit : 关闭前行为（保存数据等）
func defaultBeforeExitFunc(c *control.Control) {
	ctrl.SData().Tool.Logger.Info("[control] Ternimated")
}

// emptyFunc : 空函数
func emptyFunc(ctrl *control.Control) {}

// emptyData : 空数据
var emptyData = new(struct{})
