package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCall = "call"

var _ sdk.Msg = &MsgCall{}

func NewMsgCall(creator string, data string) *MsgCall {
	return &MsgCall{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgCall) Route() string {
	return RouterKey
}

func (msg *MsgCall) Type() string {
	return TypeMsgCall
}

func (msg *MsgCall) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCall) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCall) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
