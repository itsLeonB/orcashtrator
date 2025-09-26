package provider

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/meq"
	"github.com/itsLeonB/orcashtrator/internal/message"
)

type Queues struct {
	ExpenseBillUploaded meq.TaskQueue[message.ExpenseBillUploaded]
}

func ProvideQueues(logger ezutil.Logger, db meq.DB) *Queues {
	return &Queues{
		ExpenseBillUploaded: meq.NewTaskQueue[message.ExpenseBillUploaded](logger, db),
	}
}
