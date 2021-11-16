package types

import (
	"fmt"
	"strings"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeSetSuperfluidAssets    = "SetSuperfluidAssets"
	ProposalTypeAddSuperfluidAssets    = "AddSuperfluidAssets"
	ProposalTypeRemoveSuperfluidAssets = "RemoveSuperfluidAssets"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeSetSuperfluidAssets)
	govtypes.RegisterProposalTypeCodec(&SetSuperfluidAssetsProposal{}, "osmosis/SetSuperfluidAssetsProposal")
	govtypes.RegisterProposalType(ProposalTypeAddSuperfluidAssets)
	govtypes.RegisterProposalTypeCodec(&AddSuperfluidAssetsProposal{}, "osmosis/AddSuperfluidAssetsProposal")
	govtypes.RegisterProposalType(ProposalTypeRemoveSuperfluidAssets)
	govtypes.RegisterProposalTypeCodec(&RemoveSuperfluidAssetsProposal{}, "osmosis/RemoveSuperfluidAssetsProposal")
}

var _ govtypes.Content = &SetSuperfluidAssetsProposal{}
var _ govtypes.Content = &AddSuperfluidAssetsProposal{}
var _ govtypes.Content = &RemoveSuperfluidAssetsProposal{}

func NewSetSuperfluidAssetsProposal(title, description string, assets []SuperfluidAsset) govtypes.Content {
	return &SetSuperfluidAssetsProposal{
		Title:       title,
		Description: description,
		Assets:      assets,
	}
}

func (p *SetSuperfluidAssetsProposal) GetTitle() string { return p.Title }

func (p *SetSuperfluidAssetsProposal) GetDescription() string { return p.Description }

func (p *SetSuperfluidAssetsProposal) ProposalRoute() string { return RouterKey }

func (p *SetSuperfluidAssetsProposal) ProposalType() string {
	return ProposalTypeSetSuperfluidAssets
}

func (p *SetSuperfluidAssetsProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p SetSuperfluidAssetsProposal) String() string {
	return fmt.Sprintf(`Set Superfluid Assets Proposal:
	Title:       %s
	Description: %s
	Assets:     %+v
  `, p.Title, p.Description, p.Assets)
}

func NewAddSuperfluidAssetsProposal(title, description string, assets []SuperfluidAsset) govtypes.Content {
	return &AddSuperfluidAssetsProposal{
		Title:            title,
		Description:      description,
		SuperfluidAssets: assets,
	}
}

func (p *AddSuperfluidAssetsProposal) GetTitle() string { return p.Title }

func (p *AddSuperfluidAssetsProposal) GetDescription() string { return p.Description }

func (p *AddSuperfluidAssetsProposal) ProposalRoute() string { return RouterKey }

func (p *AddSuperfluidAssetsProposal) ProposalType() string {
	return ProposalTypeAddSuperfluidAssets
}

func (p *AddSuperfluidAssetsProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p AddSuperfluidAssetsProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Add Superfluid Assets Proposal:
  Title:       %s
  Description: %s
  SuperfluidAssets:     %+v
`, p.Title, p.Description, p.SuperfluidAssets))
	return b.String()
}

func NewRemoveSuperfluidAssetsProposal(title, description string, denoms []string) govtypes.Content {
	return &RemoveSuperfluidAssetsProposal{
		Title:                 title,
		Description:           description,
		SuperfluidAssetDenoms: denoms,
	}
}

func (p *RemoveSuperfluidAssetsProposal) GetTitle() string { return p.Title }

func (p *RemoveSuperfluidAssetsProposal) GetDescription() string { return p.Description }

func (p *RemoveSuperfluidAssetsProposal) ProposalRoute() string { return RouterKey }

func (p *RemoveSuperfluidAssetsProposal) ProposalType() string {
	return ProposalTypeRemoveSuperfluidAssets
}

func (p *RemoveSuperfluidAssetsProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p RemoveSuperfluidAssetsProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Remove Superfluid Assets Proposal:
  Title:       %s
  Description: %s
  SuperfluidAssetDenoms:     %+v
`, p.Title, p.Description, p.SuperfluidAssetDenoms))
	return b.String()
}
