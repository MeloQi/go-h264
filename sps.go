package go-h264

import (
	"fmt"
	"github.com/MeloQi/go-h264/h264"
)

func GetWidthHeight(sps []byte) (w, h int, err error) {
	if sps == nil || len(sps) < 6 {
		return 0, 0, fmt.Errorf("sps format err")
	}
	spsOffset := 0
	if sps[0]&0x1f == 7 {
		spsOffset = 1
	} else if sps[0] == 0 && sps[1] == 0 && sps[2] == 1 {
		spsOffset = 4
	} else if sps[0] == 0 && sps[1] == 0 && sps[2] == 0 && sps[3] == 1 {
		spsOffset = 5
	}
	s := h264.SequenceParameterSet{}
	if err := s.UnmarshalBinary(sps[spsOffset:]); err != nil {
		return 0, 0, err
	}
	width := (s.PicWidthInMbsMinus1 + 1) * 16
	frameMbsOnlyFlag := 0
	if s.FrameMbsOnlyFlag {
		frameMbsOnlyFlag = 1
	}
	height := uint64((2 - frameMbsOnlyFlag) * ((int(s.PicHeightInMapUnitsMinus1) + 1) * 16))
	if s.FrameCroppingFlag {
		crop_unit_x := 0
		crop_unit_y := 0
		if (s.ChromaFormatIDC == 0) {
			crop_unit_x = 1
			crop_unit_y = 2 - frameMbsOnlyFlag
		} else if s.ChromaFormatIDC == 1 {
			crop_unit_x = 2
			crop_unit_y = 2 * (2 - frameMbsOnlyFlag)
		} else if s.ChromaFormatIDC == 2 {
			crop_unit_x = 2
			crop_unit_y = 2 - frameMbsOnlyFlag
		} else {
			crop_unit_x = 1
			crop_unit_y = 2 - frameMbsOnlyFlag
		}
		width -= uint64(crop_unit_x) * (s.FrameCropLeftOffset + s.FrameCropRightOffset)
		height -= uint64(crop_unit_y) * (s.FrameCropTopOffset + s.FrameCropBottomOffset)
	}
	return int(width), int(height), nil
}

