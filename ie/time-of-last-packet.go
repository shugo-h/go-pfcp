// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"time"
)

// NewTimeOfLastPacket creates a new TimeOfLastPacket IE.
func NewTimeOfLastPacket(ts time.Time) *IE {
	u64sec := uint64(ts.Sub(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))) / 1000000000
	return newUint32ValIE(TimeOfLastPacket, uint32(u64sec))
}

// TimeOfLastPacket returns TimeOfLastPacket in time.Time if the type of IE matches.
func (i *IE) TimeOfLastPacket() (time.Time, error) {
	if len(i.Payload) < 4 {
		return time.Time{}, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case TimeOfLastPacket:
		return time.Unix(int64(binary.BigEndian.Uint32(i.Payload[0:4])-2208988800), 0), nil
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return time.Time{}, err
		}
		for _, x := range ies {
			if x.Type == TimeOfLastPacket {
				return x.TimeOfLastPacket()
			}
		}
		return time.Time{}, ErrIENotFound
	default:
		return time.Time{}, &InvalidTypeError{Type: i.Type}
	}
}
