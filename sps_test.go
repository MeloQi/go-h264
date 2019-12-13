package go_h264

import (
	"testing"
)

func TestNewSPS(t *testing.T) {
	d := [][]byte{
		{0x64, 0x00, 0x2A, 0xAD, 0x84, 0x01, 0x0C, 0x20, 0x08, 0x61, 0x00,
			0x43, 0x08, 0x02, 0x18, 0x40, 0x10, 0xC2, 0x00, 0x84, 0x2B, 0x50, 0x3C, 0x01, 0x13, 0xF2, 0xCD,
			0xC0, 0x40, 0x40, 0x50, 0x00, 0x00, 0x03, 0x00, 0x10, 0x00, 0x00, 0x03, 0x03, 0x28, 0x40},
		{0x67, 0x64, 0x00, 0x2A, 0xAD, 0x84, 0x01, 0x0C, 0x20, 0x08, 0x61, 0x00,
			0x43, 0x08, 0x02, 0x18, 0x40, 0x10, 0xC2, 0x00, 0x84, 0x2B, 0x50, 0x3C, 0x01, 0x13, 0xF2, 0xCD,
			0xC0, 0x40, 0x40, 0x50, 0x00, 0x00, 0x03, 0x00, 0x10, 0x00, 0x00, 0x03, 0x03, 0x28, 0x40},
		{0x00, 0x00, 0x01, 0x67, 0x64, 0x00, 0x2A, 0xAD, 0x84, 0x01, 0x0C, 0x20, 0x08, 0x61, 0x00,
			0x43, 0x08, 0x02, 0x18, 0x40, 0x10, 0xC2, 0x00, 0x84, 0x2B, 0x50, 0x3C, 0x01, 0x13, 0xF2, 0xCD,
			0xC0, 0x40, 0x40, 0x50, 0x00, 0x00, 0x03, 0x00, 0x10, 0x00, 0x00, 0x03, 0x03, 0x28, 0x40},
		{0x00, 0x00, 0x00, 0x01, 0x67, 0x64, 0x00, 0x2A, 0xAD, 0x84, 0x01, 0x0C, 0x20, 0x08, 0x61, 0x00,
			0x43, 0x08, 0x02, 0x18, 0x40, 0x10, 0xC2, 0x00, 0x84, 0x2B, 0x50, 0x3C, 0x01, 0x13, 0xF2, 0xCD,
			0xC0, 0x40, 0x40, 0x50, 0x00, 0x00, 0x03, 0x00, 0x10, 0x00, 0x00, 0x03, 0x03, 0x28, 0x40},
	}
	for _, v := range d {
		if w, h, err := GetWidthHeight(v); err != nil {
			t.Error(err)
			t.Failed()
			return
		} else {
			t.Log("width:", w, "  height:", h)
		}
	}
}
