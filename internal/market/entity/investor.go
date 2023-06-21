package entity

type Investor struct {
	ID            string
	Name          string
	Assetposition []*InvestorAssetPosition
}

// add new investor
func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		Assetposition: []*InvestorAssetPosition{},
	}
}

// adds an asset position to the investor's portfolio
func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.Assetposition = append(i.Assetposition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.Assetposition = append(i.Assetposition, NewInvestorAssetPosition(assetID, qtdShares))

	} else {
		assetPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.Assetposition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}

	}
	return nil

}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares}
}
