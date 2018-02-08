package transactions

type (
	// TransactionType represents a transaction type and specifies the associated action
	TransactionType byte

	// Transaction represents a lisk network transaction
	Transaction struct {
		Type                          TransactionType
		Amount                        uint64
		RecipientID                   string
		Timestamp                     uint32
		Asset                         Asset
		SenderPublicKey               []byte
		TransactionRequesterPublicKey []byte
		signature                     []byte
		secondSignature               []byte
	}

	// serializableTransaction is a transaction model that can be serialized to JSON
	serializableTransaction struct {
		Type                          TransactionType `json:"type"`
		ID                            string          `json:"id"`
		SenderID                      string          `json:"senderId"`
		Amount                        uint64          `json:"amount"`
		Fee                           int             `json:"fee"`
		RecipientID                   string          `json:"recipientId"`
		Timestamp                     uint32          `json:"timestamp"`
		Asset                         Asset           `json:"asset,omitempty"`
		SenderPublicKey               string          `json:"senderPublicKey"`
		TransactionRequesterPublicKey string          `json:"transactionRequesterPublicKey,omitempty"`
		Signature                     string          `json:"signature,omitempty"`
		SecondSignature               string          `json:"secondSignature,omitempty"`
	}

	// Asset is asset data that can be attached to a transaction
	Asset interface {
		serialize() ([]byte, error)
		IsValid() (bool, error)
	}

	// Dapp represents a Dapp on the Lisk blockchain
	Dapp struct {
		Name        string
		Link        string
		Type        string
		Category    string
		Description string
		Tags        []string
		Icon        string
	}
)

const (
	// TransactionTypeNormal is used to transfer Lisk
	TransactionTypeNormal TransactionType = 0

	// TransactionTypeSecondSecretRegistration is used to register a second secret
	TransactionTypeSecondSecretRegistration TransactionType = 1

	// TransactionTypeDelegateRegistration is used to register a delegate
	TransactionTypeDelegateRegistration TransactionType = 2

	// TransactionTypeVote is used to change vote for delegates
	TransactionTypeVote TransactionType = 3

	// TransactionTypeMultisignatureRegistration is used to register and modify a multisignature wallet
	TransactionTypeMultisignatureRegistration TransactionType = 4

	// TransactionTypeDappRegistration is used to register a DApp
	TransactionTypeDappRegistration TransactionType = 5

	// TransactionTypeTransferInSidechain is used to transfer lisk into a sidechain
	TransactionTypeTransferInSidechain TransactionType = 6

	// TransactionTypeTransferOutSidechain is used to transfer lisk out of a sidechain
	TransactionTypeTransferOutSidechain TransactionType = 7
)
