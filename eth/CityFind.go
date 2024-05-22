package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"service/eth/citynft"
)

// 查询链上

/*
*
链id是：421614
合约地址：0xc7E3E80ADf70263e3B46B94687828B54cB88ac27
链的url：https://api-sepolia.arbiscan.io/api
*/
func contractQuerySample() {
	client, err := ethclient.DialContext(context.Background(), "https://sepolia.infura.io/v3/659b3553bb42497c8c38724899fa489b")
	if err != nil {
		log.Fatal(err)
	}

	//privateKey, err := crypto.HexToECDSA("634ae752918922040eba9aefb4ad95521712e1ef0a114ec12d01dd8bd604e2cf")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//
	//chainID := big.NewInt(11155111) // Sepolia 测试网络的 chainID
	//_, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	//if err != nil {
	//	log.Fatalf("Failed to create authorized transactor: %v", err)
	//}
	//
	//// 合约地址
	address := common.HexToAddress("0x80D6Fc288aE2F93b259a3d63e3d8E3C610d83802")
	//// 创建合约实例
	contract, err := citynft.NewCitynft(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Smart Contract: %v", err)
	}

	fromAddress := common.HexToAddress("0x459115d03992914fDc832A14a03da78a1e4d87B8")

	city, err := contract.GetUserMintedCity(&bind.CallOpts{
		From:    fromAddress,
		Context: context.Background(),
	}, "China")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("city: ", city)

}
