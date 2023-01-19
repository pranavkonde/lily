// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package messages

import (
	"fmt"
	"io"
	"math"
	"sort"

	processor "github.com/filecoin-project/lily/pkg/extract/processor"
	types "github.com/filecoin-project/lotus/chain/types"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *FullBlockIPLDContainer) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.BlockHeader (types.BlockHeader) (struct)
	if len("block_header") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"block_header\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("block_header"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("block_header")); err != nil {
		return err
	}

	if err := t.BlockHeader.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.SecpMessages (cid.Cid) (struct)
	if len("secp_messages") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"secp_messages\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("secp_messages"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("secp_messages")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.SecpMessages); err != nil {
		return xerrors.Errorf("failed to write cid field t.SecpMessages: %w", err)
	}

	// t.BlsMessages (cid.Cid) (struct)
	if len("bls_messages") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"bls_messages\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("bls_messages"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("bls_messages")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.BlsMessages); err != nil {
		return xerrors.Errorf("failed to write cid field t.BlsMessages: %w", err)
	}

	return nil
}

func (t *FullBlockIPLDContainer) UnmarshalCBOR(r io.Reader) (err error) {
	*t = FullBlockIPLDContainer{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("FullBlockIPLDContainer: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.BlockHeader (types.BlockHeader) (struct)
		case "block_header":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.BlockHeader = new(types.BlockHeader)
					if err := t.BlockHeader.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.BlockHeader pointer: %w", err)
					}
				}

			}
			// t.SecpMessages (cid.Cid) (struct)
		case "secp_messages":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.SecpMessages: %w", err)
				}

				t.SecpMessages = c

			}
			// t.BlsMessages (cid.Cid) (struct)
		case "bls_messages":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.BlsMessages: %w", err)
				}

				t.BlsMessages = c

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ChainMessageIPLDContainer) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.Message (types.Message) (struct)
	if len("message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"message\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("message")); err != nil {
		return err
	}

	if err := t.Message.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Receipt (processor.ChainMessageReceipt) (struct)
	if len("receipt") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"receipt\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("receipt"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("receipt")); err != nil {
		return err
	}

	if err := t.Receipt.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.VmMessagesAmt (cid.Cid) (struct)
	if len("vm_messages") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"vm_messages\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("vm_messages"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("vm_messages")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.VmMessagesAmt); err != nil {
		return xerrors.Errorf("failed to write cid field t.VmMessagesAmt: %w", err)
	}

	return nil
}

func (t *ChainMessageIPLDContainer) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ChainMessageIPLDContainer{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ChainMessageIPLDContainer: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Message (types.Message) (struct)
		case "message":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Message = new(types.Message)
					if err := t.Message.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Message pointer: %w", err)
					}
				}

			}
			// t.Receipt (processor.ChainMessageReceipt) (struct)
		case "receipt":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Receipt = new(processor.ChainMessageReceipt)
					if err := t.Receipt.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Receipt pointer: %w", err)
					}
				}

			}
			// t.VmMessagesAmt (cid.Cid) (struct)
		case "vm_messages":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.VmMessagesAmt: %w", err)
				}

				t.VmMessagesAmt = c

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SignedChainMessageIPLDContainer) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.Message (types.SignedMessage) (struct)
	if len("message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"message\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("message")); err != nil {
		return err
	}

	if err := t.Message.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Receipt (processor.ChainMessageReceipt) (struct)
	if len("receipt") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"receipt\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("receipt"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("receipt")); err != nil {
		return err
	}

	if err := t.Receipt.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.VmMessagesAmt (cid.Cid) (struct)
	if len("vm_messages") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"vm_messages\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("vm_messages"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("vm_messages")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.VmMessagesAmt); err != nil {
		return xerrors.Errorf("failed to write cid field t.VmMessagesAmt: %w", err)
	}

	return nil
}

func (t *SignedChainMessageIPLDContainer) UnmarshalCBOR(r io.Reader) (err error) {
	*t = SignedChainMessageIPLDContainer{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SignedChainMessageIPLDContainer: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Message (types.SignedMessage) (struct)
		case "message":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Message = new(types.SignedMessage)
					if err := t.Message.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Message pointer: %w", err)
					}
				}

			}
			// t.Receipt (processor.ChainMessageReceipt) (struct)
		case "receipt":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Receipt = new(processor.ChainMessageReceipt)
					if err := t.Receipt.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Receipt pointer: %w", err)
					}
				}

			}
			// t.VmMessagesAmt (cid.Cid) (struct)
		case "vm_messages":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.VmMessagesAmt: %w", err)
				}

				t.VmMessagesAmt = c

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ImplicitMessageIPLDContainer) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.Message (types.Message) (struct)
	if len("message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"message\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("message")); err != nil {
		return err
	}

	if err := t.Message.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Receipt (processor.ImplicitMessageReceipt) (struct)
	if len("receipt") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"receipt\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("receipt"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("receipt")); err != nil {
		return err
	}

	if err := t.Receipt.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.VmMessagesAmt (cid.Cid) (struct)
	if len("vm_messages") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"vm_messages\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("vm_messages"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("vm_messages")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.VmMessagesAmt); err != nil {
		return xerrors.Errorf("failed to write cid field t.VmMessagesAmt: %w", err)
	}

	return nil
}

func (t *ImplicitMessageIPLDContainer) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ImplicitMessageIPLDContainer{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ImplicitMessageIPLDContainer: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Message (types.Message) (struct)
		case "message":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Message = new(types.Message)
					if err := t.Message.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Message pointer: %w", err)
					}
				}

			}
			// t.Receipt (processor.ImplicitMessageReceipt) (struct)
		case "receipt":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Receipt = new(processor.ImplicitMessageReceipt)
					if err := t.Receipt.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Receipt pointer: %w", err)
					}
				}

			}
			// t.VmMessagesAmt (cid.Cid) (struct)
		case "vm_messages":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.VmMessagesAmt: %w", err)
				}

				t.VmMessagesAmt = c

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
