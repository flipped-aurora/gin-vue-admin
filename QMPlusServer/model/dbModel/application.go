package dbModel

import (
	"github.com/jinzhu/gorm"
	"time"
)

//申请model 工作流实例

type Application struct {
	gorm.Model
	WorkFlowID           string    // 所属工作流ID
	WorkFlowStepInfoID   string    // 当前节点ID
	ApplicationName      string    // 申请人姓名
	ApplicationCause     string    // 请假原因
	ApplicationStartData time.Time // 请假开始日期
	ApplicationEndData   time.Time // 请假开始日期
}
