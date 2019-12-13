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
		{0x67, 0x4d, 0x00, 0x2a, 0x95, 0xa8, 0x1e, 0x00, 0x89, 0xf9, 0x66, 0xe0, 0x20, 0x20, 0x28, 0x00, 0x00, 0x03, 0x00, 0x08, 0x00, 0x00, 0x03, 0x01, 0x94, 0x20},
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
