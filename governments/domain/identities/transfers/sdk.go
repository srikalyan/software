package transfers

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents an outgoing builder
type Builder interface {
	Create() Builder
	WithTransfer(transfer views.Transfer) Builder
	WithNote(note string) Builder
	Now() (Transfer, error)
}

// Transfer represents a transfer
type Transfer interface {
	Hash() hash.Hash
	Transfer() views.Transfer
	Note() string
}

// Service represents a transfer service
type Service interface {
	Insert(ins Transfer) error
}
