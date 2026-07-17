package system

import (
	"context"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/datascope"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

type DataAccessLogService struct{}

var DataAccessLogServiceApp = new(DataAccessLogService)

// 异步写入: 有界缓冲 + 后台批量落表, 满则丢弃(审计尽力而为, 绝不阻塞业务请求)
var (
	dataAccessLogCh         = make(chan system.SysDataAccessLog, 1024)
	dataAccessLogWriterOnce sync.Once // reload 会重复调注册入口, 写协程只起一次
)

// Enqueue 非阻塞入队(注册为 datascope 的审计 hook)
func (s *DataAccessLogService) Enqueue(evt datascope.AuditEvent) {
	rec := system.SysDataAccessLog{
		EventType:   evt.EventType,
		TargetTable: evt.TargetTable,
		Operation:   evt.Operation,
		UserID:      evt.UserID,
		AuthorityID: evt.AuthorityID,
		Scope:       evt.Scope,
		RequestID:   evt.RequestID,
		Method:      evt.Method,
		Path:        evt.Path,
		Detail:      evt.Detail,
	}
	select {
	case dataAccessLogCh <- rec:
	default: // 缓冲已满, 丢弃并告警
		logger.Bg().Mod("datascope").Warn("数据权限审计缓冲已满, 事件被丢弃: " + evt.EventType + " " + evt.TargetTable)
	}
}

// StartWriter 启动后台批量写入协程(每 2 秒或攒满 100 条落一次表), 幂等
func (s *DataAccessLogService) StartWriter() {
	dataAccessLogWriterOnce.Do(s.startWriter)
}

func (s *DataAccessLogService) startWriter() {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		batch := make([]system.SysDataAccessLog, 0, 128)
		flush := func() {
			if len(batch) == 0 {
				return
			}
			ctx := datascope.WithSystem(context.Background())
			if err := global.GVA_DB.WithContext(ctx).Create(&batch).Error; err != nil {
				logger.Bg().Mod("datascope").Err(err).Warn("数据权限审计批量写入失败")
			}
			batch = batch[:0]
		}
		for {
			select {
			case rec := <-dataAccessLogCh:
				batch = append(batch, rec)
				if len(batch) >= 100 {
					flush()
				}
			case <-ticker.C:
				flush()
			}
		}
	}()
}

// GetDataAccessLogList 分页查询审计日志
func (s *DataAccessLogService) GetDataAccessLogList(ctx context.Context, info systemReq.SysDataAccessLogSearch) (list []system.SysDataAccessLog, total int64, err error) {
	limit, offset := info.LimitOffset()
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysDataAccessLog{})
	if info.EventType != "" {
		db = db.Where("event_type = ?", info.EventType)
	}
	if info.TargetTable != "" {
		db = db.Where("target_table LIKE ?", "%"+info.TargetTable+"%")
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

// DeleteDataAccessLogByIds 批量删除审计日志
func (s *DataAccessLogService) DeleteDataAccessLogByIds(ctx context.Context, ids []int) error {
	return global.GVA_DB.WithContext(ctx).Delete(&[]system.SysDataAccessLog{}, "id in ?", ids).Error
}
