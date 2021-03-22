package paymentchannelmanager

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/specs-actors/actors/builtin/paych"
	"github.com/ipfs/go-cid"
)

type PCHDir int

const (
	PCHUndef PCHDir = iota
	PCHInbound
	PCHOutbound
)

type PaychStatus struct {
	// PCHDir indicates whether this channel is an inbound channel or an
	// outbound channel
	Direction PCHDir
	// ControlAddr
	ControlAddr address.Address
	// Channel is the address of the channel
	Channel address.Address
	// From is the from address of the channel (channel creator)
	From address.Address
	// To is the to address of the channel
	To address.Address
	// ConfirmedAmt is the amount of funds that have been confirmed on-chain
	// for the channel
	ConfirmedAmt big.Int
	// PendingAmt is the amount of funds that are pending confirmation on-chain
	PendingAmt big.Int
	// PendingChainMessage
	PendingChainMessage *cid.Cid
	// QueuedAmt is the amount that is queued up behind a pending request
	QueuedAmt big.Int
	// VoucherRedeemedAmt is the amount that is redeemed by vouchers on-chain
	// and in the local datastore
	VoucherReedeemedAmt big.Int
	// ReservedAmt is the portion of the confirmed funds previously reserved by
	// consumers of this payment channel
	ReservedAmt big.Int
}

// VoucherCreateResult is the response to calling PaychVoucherCreate
type VoucherCreateResult struct {
	// Voucher that was created, or nil if there was an error or if there
	// were insufficient funds in the channel
	Voucher *paych.SignedVoucher
	// Shortfall is the additional amount that would be needed in the channel
	// in order to be able to create the voucher
	Shortfall big.Int
}

type PaymentChannelManagerAPI interface {
	// PaychCreate creates a new payment channel and deposits funds in the
	// channel
	PaychCreate(
		ctx context.Context,
		from, to address.Address,
		amt big.Int,
	) (cid.Cid, error)
	// PaychAddFunds deposits additional payment into a payment channel
	// It returns the cid of the chain message submitted for the additional
	// funds
	PaychAddFunds(
		ctx context.Context,
		ch address.Address,
		amt big.Int,
	) (cid.Cid, error)
	// PaychReserveFunds locks up the specified amount of funds in the given
	// channel for
	// a future transaction. If there are not enough available funds for
	// the channel, submits a message on chain to top up available funds.
	// Returns the cid of the message that was submitted on chain, or cid.Undef
	// if
	// the required funds were already available.
	PaychReserveFunds(
		ctx context.Context,
		ch address.Address,
		amt big.Int,
	) (cid.Cid, error)
	// PaychReleaseFunds frees channel funds so another transaction can use them
	PaychReleaseFunds(
		ctx context.Context,
		addr address.Address,
		amt big.Int,
	) error
	// PaychStatus looks up all available information about a payment channel
	PaychStatus(ctx context.Context, ch address.Address) (*PaychStatus, error)
	// PaychLookup searches for a channel matching the given from and to
	// addresses that has available lanes
	PaychLookup(
		ctx context.Context,
		from, to address.Address,
	) (address.Address, error)
	// PaychList lists the channels this manager is tracking
	PaychList(context.Context) ([]address.Address, error)
	// PaychSettle begins the settlement period for a payment channel
	PaychSettle(context.Context, address.Address) (cid.Cid, error)
	// PaychCollect attempts to collect payment on a payment channel at the end
	// of a settlement period
	PaychCollect(context.Context, address.Address) (cid.Cid, error)
	// PaychAllocateLane identifies the next free lane for a payment channel
	PaychAllocateLane(ctx context.Context, ch address.Address) (uint64, error)
	// PaychVoucherCheckValid verifies if a voucher appears to meet requirements
	// for submission
	PaychVoucherCheckValid(
		ctx context.Context,
		ch address.Address,
		sv *paych.SignedVoucher,
	) error
	// PaychVoucherCheckSpendable verifies if a voucher appears to meet
	// requirements for submission and also verifies
	// the given secret against the pre-image
	PaychVoucherCheckSpendable(
		ctx context.Context,
		ch address.Address,
		sv *paych.SignedVoucher,
		secret []byte,
		proof []byte,
	) (bool, error)
	// PaychVoucherCreate creates a new signed voucher on the given payment
	// channel
	// with the given lane and amount.  The value passed in is exactly the value
	// that will be used to create the voucher, so if previous vouchers exist,
	// the
	// actual additional value of this voucher will only be the difference
	// between
	// the two.
	// If there are insufficient funds in the channel to create the voucher,
	// returns a nil voucher and the shortfall.
	PaychVoucherCreate(
		ctx context.Context,
		pch address.Address,
		amt big.Int,
		lane uint64,
	) (*VoucherCreateResult, error)
	// PaychVoucherAdd records a new inbound voucher but does not save it to the
	// chain. In addition to the channel, voucher, and proof for the voucher
	// pre-image
	// if present, payment voucher add takes a minDelta, or minimum increase in
	// lane value required to save the voucher.
	// If successful, the function returns the delta in amound redeemed on the
	// given lane in the payment channel
	PaychVoucherAdd(
		ctx context.Context,
		ch address.Address,
		sv *paych.SignedVoucher,
		proof []byte,
		minDelta big.Int,
	) (big.Int, error)
	// PaymentVoucherList returns all vouchers for a given payment channel
	PaychVoucherList(
		context.Context,
		address.Address,
	) ([]*paych.SignedVoucher, error)
	// PaymentVoucherSubmit submits a given voucher to the blockchain
	PaychVoucherSubmit(
		ctx context.Context,
		ch address.Address,
		sv *paych.SignedVoucher,
		secret []byte,
		proof []byte,
	) (cid.Cid, error)
}
