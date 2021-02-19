package authenticated

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/shareholders/payments"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers/incomings"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers/outgoings"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithPK(pk signature.PrivateKey) Builder
	Now() (Application, error)
}

// Application represents an authenticated shareholder application
type Application interface {
	Proposition() Proposition
	Payment(amount uint, note string) error
	Transfer(amount uint, seed string, to []hash.Hash, note string) error
	View(amount uint, seed string, to []hash.Hash) (transfers.Section, error)
	Receive(view transfers.Section, pk signature.PrivateKey, note string) error
	Transaction(filter DateFilter) Transaction
}

// Proposition represents an authenticated proposition application
type Proposition interface {
	New(content propositions.Proposition, sigs []signature.RingSignature) error
	Approve(propositionHash hash.Hash) error
	Cancel(propositionHash hash.Hash) error
	Disapprove(propositionHash hash.Hash) error
}

// Transaction represents an authenticated transaction application
type Transaction interface {
	Payment() Payments
	Incoming() IncomingTransaction
	Outgoing() OutgoingTransaction
}

// Payments represents a payments application
type Payments interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) ([]payments.Payment, error)
}

// IncomingTransaction represents an incoming transaction application
type IncomingTransaction interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) ([]incomings.Incoming, error)
}

// OutgoingTransaction represents an outgoing transaction application
type OutgoingTransaction interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) ([]outgoings.Outgoing, error)
}

// DateFilterBuilder represents a date filter builder
type DateFilterBuilder interface {
	Create() DateFilterBuilder
	From(from time.Time) DateFilterBuilder
	To(to time.Time) DateFilterBuilder
	Now() (DateFilter, error)
}

// DateFilter represents a date filter
type DateFilter interface {
	HasFrom() bool
	From() *time.Time
	HasTo() bool
	To() *time.Time
}
