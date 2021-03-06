package requests

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Request represents a connection request
type Request interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.Signature
}

// Content represents a content request
type Content interface {
	Hash() hash.Hash
	Requestor() hash.Hash
	PublicKey() public.Key
	Server() servers.Server
	CreatedOn() time.Time
}
