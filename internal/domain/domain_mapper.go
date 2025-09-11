package domain

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/audit/gen/go/audit/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/rotisserie/eris"
)

func FromAuditMetadataProto(amp *audit.Metadata) (AuditMetadata, error) {
	if amp == nil {
		return AuditMetadata{}, eris.New("audit metadata is nil")
	}

	id, err := ezutil.Parse[uuid.UUID](amp.GetId())
	if err != nil {
		return AuditMetadata{}, err
	}

	return AuditMetadata{
		ID:        id,
		CreatedAt: ezutil.FromProtoTime(amp.GetCreatedAt()),
		UpdatedAt: ezutil.FromProtoTime(amp.GetUpdatedAt()),
		DeletedAt: ezutil.FromProtoTime(amp.GetDeletedAt()),
	}, nil
}
