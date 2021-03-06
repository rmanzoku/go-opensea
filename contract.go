package opensea

import (
	"context"
	"encoding/json"
)

type Contract struct {
	// Collection                  Collection  `json:"collection"`
	Address                     Address     `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NFTVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       int64       `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int64       `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int64       `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int64       `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int64       `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int64       `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int64       `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

func (o Opensea) GetSingleContract(assetContractAddress string) (*Contract, error) {
	ctx := context.TODO()
	return o.GetSingleContractWithContext(ctx, assetContractAddress)
}

func (o Opensea) GetSingleContractWithContext(ctx context.Context, assetContractAddress string) (contract *Contract, err error) {
	path := "/api/v1/asset_contract/" + assetContractAddress
	b, err := o.getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	contract = new(Contract)
	err = json.Unmarshal(b, contract)
	return
}
