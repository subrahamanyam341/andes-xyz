package block

import "github.com/subrahamanyam341/andes-xyz/data"

type emptyHeaderV2Creator struct{}

// NewEmptyHeaderV2Creator is able to create empty header v2 instances
func NewEmptyHeaderV2Creator() *emptyHeaderV2Creator {
	return &emptyHeaderV2Creator{}
}

// CreateNewHeader creates a new empty header v2
func (creator *emptyHeaderV2Creator) CreateNewHeader() data.HeaderHandler {
	return &HeaderV2{
		Header: &Header{},
	}
}

// IsInterfaceNil returns true if there is no value under the interface
func (creator *emptyHeaderV2Creator) IsInterfaceNil() bool {
	return creator == nil
}
