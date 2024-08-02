package internal

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/event"
	opt "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var Mongo = new(mongo)

type mongo struct{}

func (m *mongo) GetClientOptions() []options.ClientOptions {
	cmdMonitor := &event.CommandMonitor{
		Started: func(ctx context.Context, event *event.CommandStartedEvent) {
			zap.L().Info(fmt.Sprintf("[MongoDB][RequestID:%d][database:%s] %s\n", event.RequestID, event.DatabaseName, event.Command), zap.String("business", "mongo"))
		},
		Succeeded: func(ctx context.Context, event *event.CommandSucceededEvent) {
			zap.L().Info(fmt.Sprintf("[MongoDB][RequestID:%d] [%s] %s\n", event.RequestID, event.Duration.String(), event.Reply), zap.String("business", "mongo"))
		},
		Failed: func(ctx context.Context, event *event.CommandFailedEvent) {
			zap.L().Error(fmt.Sprintf("[MongoDB][RequestID:%d] [%s] %s\n", event.RequestID, event.Duration.String(), event.Failure), zap.String("business", "mongo"))
		},
	}
	return []options.ClientOptions{{ClientOptions: &opt.ClientOptions{Monitor: cmdMonitor}}}
}
