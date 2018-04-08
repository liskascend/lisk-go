package transactions

import "time"

const (
	fixedPoint = 100000000

	feeSend           = 0.1 * fixedPoint
	feeData           = 0.1 * fixedPoint
	feeTransferIn     = 0.1 * fixedPoint
	feeTransferOut    = 0.1 * fixedPoint
	feeSignature      = 5 * fixedPoint
	feeDelegate       = 25 * fixedPoint
	feeVote           = 1 * fixedPoint
	feeMultisignature = 5 * fixedPoint
	feeDapp           = 25 * fixedPoint

	//byteSizeTimestamp                  = 4
	byteSizeRecipientID = 8
	//byteSizeAmount                     = 8
	byteSizeSignatureTransaction       = 64
	byteSizeSecondSignatureTransaction = 64
	byteSizeData                       = 64
)

var (
	epochTime   = time.Date(2016, 5, 24, 17, 0, 0, 0, time.UTC)
	epochTimeMs = epochTime.UTC().UnixNano() / int64(time.Millisecond)
)
