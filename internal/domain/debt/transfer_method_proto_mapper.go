package debt

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/drex-protos/gen/go/transaction/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/rotisserie/eris"
)

func fromTransferMethodProto(tm *transaction.TransferMethodResponse) (TransferMethod, error) {
	if tm == nil {
		return TransferMethod{}, eris.New("transfer method is nil")
	}

	id, err := ezutil.Parse[uuid.UUID](tm.GetId())
	if err != nil {
		return TransferMethod{}, err
	}

	return TransferMethod{
		ID:        id,
		Name:      tm.GetName(),
		Display:   tm.GetDisplay(),
		CreatedAt: ezutil.FromProtoTime(tm.GetCreatedAt()),
		UpdatedAt: ezutil.FromProtoTime(tm.GetUpdatedAt()),
		DeletedAt: ezutil.FromProtoTime(tm.GetDeletedAt()),
	}, nil
}
