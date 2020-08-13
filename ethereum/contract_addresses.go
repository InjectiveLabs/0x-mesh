package ethereum

import (
	"fmt"

	"github.com/0xProject/0x-mesh/constants"
	"github.com/ethereum/go-ethereum/common"
)

// ContractAddresses maps a contract's name to it's Ethereum address
type ContractAddresses struct {
	ERC20Proxy          common.Address `json:"erc20Proxy"`
	ERC721Proxy         common.Address `json:"erc721Proxy"`
	ERC1155Proxy        common.Address `json:"erc1155Proxy"`
	Exchange            common.Address `json:"exchange"`
	Coordinator         common.Address `json:"coordinator"`
	CoordinatorRegistry common.Address `json:"coordinatorRegistry"`
	DevUtils            common.Address `json:"devUtils"`
	WETH9               common.Address `json:"weth9"`
	ZRXToken            common.Address `json:"zrxToken"`
	ChaiBridge          common.Address `json:"chaiBridge"`
	ChaiToken           common.Address `json:"chaiToken"`
	MaximumGasPrice     common.Address `json:"maximumGasPrice"`
}

// GanacheAddresses The addresses that the 0x contracts were deployed to on the Ganache snapshot (chainID = 1337).
var GanacheAddresses = ganacheAddresses()

// NewContractAddressesForChainID The default contract addresses for the standard chainIDs.
func NewContractAddressesForChainID(chainID int) (ContractAddresses, error) {
	switch chainID {
	case 1:
		return ContractAddresses{
			ERC20Proxy:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ERC721Proxy:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Exchange:            common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ERC1155Proxy:        common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Coordinator:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
			CoordinatorRegistry: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			DevUtils:            common.HexToAddress("0x0000000000000000000000000000000000000000"),
			WETH9:               common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ZRXToken:            common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ChaiBridge:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ChaiToken:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
			MaximumGasPrice:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
		}, nil
	case 15001:
		return ContractAddresses{
			ERC20Proxy:          common.HexToAddress("0x979934a6e6500f5ab9a0625def49148017eefa03"),
			ERC721Proxy:         common.HexToAddress("0x8bef4a095f0188dfaf448a8e0b646e972cfc8a65"),
			Exchange:            common.HexToAddress("0xa5ab15972d62a5cd212d5ca791137bc6770373de"),
			ERC1155Proxy:        common.HexToAddress("0x8ace39ed9d17987782e283abe2973c568bf59a92"),
			Coordinator:         common.HexToAddress("0x2e64217707ffa5e17627160a9feeff83fe8e2da3"),
			CoordinatorRegistry: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			DevUtils:            common.HexToAddress("0x988e8d73acd4f9ab84cf4b0d6c264a6f0f275807"),
			WETH9:               common.HexToAddress("0xe11c000a97cd12843dbb6da2d61b4342e1f24772"),
			ZRXToken:            common.HexToAddress("0x7dc9fb4c428675e5ad08697d191847fe71e96700"),
			ChaiBridge:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ChaiToken:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
			MaximumGasPrice:     common.HexToAddress("0x2c668051f237caa8aba4277143ac5f663bdbfeca"),
		}, nil
	case 1337:
		return ganacheAddresses(), nil
	default:
		return ContractAddresses{}, fmt.Errorf("Cannot create contract addresses for non-standard chainID")
	}
}

func ValidateContractAddressesForChainID(chainID int, addresses ContractAddresses) error {
	if chainID == 1 {
		return fmt.Errorf("cannot add contract addresses for chainID 1: addresses for mainnet are hard-coded and cannot be changed")
	}
	if addresses.Exchange == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: Exchange address is required", chainID)
	}
	if addresses.DevUtils == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: DevUtils address is required", chainID)
	}
	if addresses.ERC20Proxy == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: ERC20Proxy address is required", chainID)
	}
	if addresses.ERC721Proxy == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: ERC721Proxy address is required", chainID)
	}
	if addresses.ERC1155Proxy == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: ERC1155Proxy address is required", chainID)
	}
	// TODO(albrow): Uncomment this if we re-add coordinator support.
	// if addresses.CoordinatorRegistry == constants.NullAddress {
	// 	return fmt.Errorf("cannot add contract addresses for chain ID %d: CoordinatorRegistry address is required", chainID)
	// }
	return nil
}

// ganacheAddresses Returns the addresses of the deployed contracts on the Ganache snapshot. This
// function allows these addresses to only be defined in one place.
func ganacheAddresses() ContractAddresses {
	return ContractAddresses{
		ERC20Proxy:          common.HexToAddress("0x1dc4c1cefef38a777b15aa20260a54e584b16c48"),
		ERC721Proxy:         common.HexToAddress("0x1d7022f5b17d2f8b695918fb48fa1089c9f85401"),
		ERC1155Proxy:        common.HexToAddress("0x6a4a62e5a7ed13c361b176a5f62c2ee620ac0df8"),
		Exchange:            common.HexToAddress("0x48bacb9266a570d521063ef5dd96e61686dbe788"),
		Coordinator:         common.HexToAddress("0x4d3d5c850dd5bd9d6f4adda3dd039a3c8054ca29"),
		CoordinatorRegistry: common.HexToAddress("0xaa86dda78e9434aca114b6676fc742a18d15a1cc"),
		DevUtils:            common.HexToAddress("0xb23672f74749bf7916ba6827c64111a4d6de7f11"),
		WETH9:               common.HexToAddress("0x0b1ba0af832d7c05fd64161e0db78e85978e8082"),
		ZRXToken:            common.HexToAddress("0x871dd7c2b4b25e1aa18728e9d5f2af4c4e431f5c"),
		ChaiBridge:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
		ChaiToken:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
		MaximumGasPrice:     common.HexToAddress("0x2c530e4ecc573f11bd72cf5fdf580d134d25f15f"),
	}
}
