package provider

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/meq"
	"github.com/itsLeonB/orcashtrator/internal/message"
	"github.com/rotisserie/eris"
)

type Queues struct {
	ExpenseBillUploaded meq.TaskQueue[message.ExpenseBillUploaded]
}

func ProvideQueues(logger ezutil.Logger, db meq.DB) (*Queues, error) {
	if logger == nil {
		return nil, eris.New("logger cannot be nil")
	}
	if db == nil {
		return nil, eris.New("db cannot be nil")
	}
	return &Queues{
		ExpenseBillUploaded: meq.NewTaskQueue[message.ExpenseBillUploaded](logger, db),
	}, nil
}
