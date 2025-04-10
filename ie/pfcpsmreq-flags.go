// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPFCPSMReqFlags creates a new PFCPSMReqFlags IE.
func NewPFCPSMReqFlags(flag uint8) *IE {
	return newUint8ValIE(PFCPSMReqFlags, flag)
}

// PFCPSMReqFlags returns PFCPSMReqFlags in uint8 if the type of IE matches.
func (i *IE) PFCPSMReqFlags() (uint8, error) {
	switch i.Type {
	case PFCPSMReqFlags:
		return i.Payload[0], nil
	case ForwardingParameters:
		ies, err := i.ForwardingParameters()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PFCPSMReqFlags {
				return x.PFCPSMReqFlags()
			}
		}
		return 0, ErrIENotFound
	case UpdateForwardingParameters:
		ies, err := i.UpdateForwardingParameters()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PFCPSMReqFlags {
				return x.PFCPSMReqFlags()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// HasDROBU reports whether an IE has DROBU bit.
func (i *IE) HasDROBU() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case PFCPSMReqFlags:
		v, err := i.PFCPSMReqFlags()
		if err != nil {
			return false
		}

		return has1stBit(v)
	case PFCPSRRspFlags:
		v, err := i.PFCPSRRspFlags()
		if err != nil {
			return false
		}

		return has1stBit(v)
	default:
		return false
	}
}

// HasSNDEM reports whether an IE has SNDEM bit.
func (i *IE) HasSNDEM() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case PFCPSMReqFlags:
		v, err := i.PFCPSMReqFlags()
		if err != nil {
			return false
		}

		return has2ndBit(v)
	default:
		return false
	}
}

// HasQAURR reports whether an IE has QAURR bit.
func (i *IE) HasQAURR() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case PFCPSMReqFlags:
		v, err := i.PFCPSMReqFlags()
		if err != nil {
			return false
		}

		return has3rdBit(v)
	default:
		return false
	}
}
