// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

// Enum values for channel_status
const (
	ChannelStatusFree          = "free"
	ChannelStatusInUse         = "in_use"
	ChannelStatusMergeReserved = "merge_reserved"
	ChannelStatusMerged        = "merged"
)

// Enum values for exchange_currency_type
const (
	ExchangeCurrencyTypeCrypto = "crypto"
	ExchangeCurrencyTypeFiat   = "fiat"
)

// Enum values for payment_network
const (
	PaymentNetworkFiat     = "fiat"
	PaymentNetworkStellar  = "stellar"
	PaymentNetworkEthereum = "ethereum"
	PaymentNetworkBitcoin  = "bitcoin"
)

// Enum values for ico_status
const (
	IcoStatusPlanning  = "planning"
	IcoStatusReady     = "ready"
	IcoStatusActive    = "active"
	IcoStatusFinished  = "finished"
	IcoStatusCompleted = "completed"
	IcoStatusStopped   = "stopped"
)

// Enum values for ico_sales_model
const (
	IcoSalesModelFixed = "fixed"
)

// Enum values for ico_phase_status
const (
	IcoPhaseStatusPlanning  = "planning"
	IcoPhaseStatusReady     = "ready"
	IcoPhaseStatusActive    = "active"
	IcoPhaseStatusFinished  = "finished"
	IcoPhaseStatusCompleted = "completed"
	IcoPhaseStatusStopped   = "stopped"
)

// Enum values for transaction_status
const (
	TransactionStatusNew     = "new"
	TransactionStatusError   = "error"
	TransactionStatusRefund  = "refund"
	TransactionStatusCashout = "cashout"
)

// Enum values for message_type
const (
	MessageTypeIos     = "ios"
	MessageTypeAndroid = "android"
	MessageTypeMail    = "mail"
)

// Enum values for mail_content_type
const (
	MailContentTypeText = "text"
	MailContentTypeHTML = "html"
)

// Enum values for notification_status_code
const (
	NotificationStatusCodeNew     = "new"
	NotificationStatusCodeSuccess = "success"
	NotificationStatusCodeError   = "error"
)

// Enum values for memo_type
const (
	MemoTypeMEMO_TEXT   = "MEMO_TEXT"
	MemoTypeMEMO_ID     = "MEMO_ID"
	MemoTypeMEMO_HASH   = "MEMO_HASH"
	MemoTypeMEMO_RETURN = "MEMO_RETURN"
)

// Enum values for kyc_document_type
const (
	KycDocumentTypePassport         = "passport"
	KycDocumentTypeDriversLicense   = "drivers_license"
	KycDocumentTypeIDCard           = "id_card"
	KycDocumentTypeProofOfResidence = "proof_of_residence"
)

// Enum values for kyc_document_format
const (
	KycDocumentFormatPNG  = "png"
	KycDocumentFormatPDF  = "pdf"
	KycDocumentFormatJPG  = "jpg"
	KycDocumentFormatJpeg = "jpeg"
)

// Enum values for kyc_document_side
const (
	KycDocumentSideFront = "front"
	KycDocumentSideBack  = "back"
)

// Enum values for order_status
const (
	OrderStatusWaitingForPayment      = "waiting_for_payment"
	OrderStatusPaymentReceived        = "payment_received"
	OrderStatusWaitingUserTransaction = "waiting_user_transaction"
	OrderStatusPaymentError           = "payment_error"
	OrderStatusFinished               = "finished"
	OrderStatusError                  = "error"
	OrderStatusUnderPay               = "under_pay"
	OrderStatusOverPay                = "over_pay"
	OrderStatusNoCoinsLeft            = "no_coins_left"
	OrderStatusPhaseExpired           = "phase_expired"
)

// Enum values for payment_state
const (
	PaymentStateOpen  = "open"
	PaymentStateClose = "close"
)

// Enum values for kyc_status
const (
	KycStatusNotSupported     = "not_supported"
	KycStatusWaitingForData   = "waiting_for_data"
	KycStatusWaitingForReview = "waiting_for_review"
	KycStatusInReview         = "in_review"
	KycStatusPending          = "pending"
	KycStatusRejected         = "rejected"
	KycStatusApproved         = "approved"
)

// Enum values for device_type
const (
	DeviceTypeApple  = "apple"
	DeviceTypeGoogle = "google"
)

// Enum values for wallet_type
const (
	WalletTypeInternal = "internal"
	WalletTypeExternal = "external"
)
