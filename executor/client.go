package executor

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	sdkclient "github.com/mocachain/moca-go-sdk/client"
	"github.com/mocachain/moca-go-sdk/types"
	"github.com/mocachain/moca-relayer/contract/universalVerifier"
	"github.com/mocachain/moca-relayer/contract/mocacrosschainupgradeable"
	"github.com/mocachain/moca-relayer/logging"
)

type MocaClient struct {
	sdkclient.IClient
	ethClient               *ethclient.Client
	mocaSBTCrossChainClient *mocacrosschainupgradeable.IMOCACrossChainUpgradeable
	mocaVCCrossChainClient  *universalVerifier.UniversalVerifier
	Height                  int64
}

type GnfdCompositeClients struct {
	clients []*MocaClient
}

// func getEthRPCAddress(rpcAddr string) string {
// 	u, err := url.Parse(rpcAddr)
// 	if err != nil {
// 		return rpcAddr
// 	}
// 	host, portStr, _ := net.SplitHostPort(u.Host)
// 	if portStr == "" {
// 		return fmt.Sprintf("%s://%s:8545%s", u.Scheme, host, u.Path)
// 	}
// 	port, err := strconv.Atoi(portStr)
// 	if err != nil {
// 		return rpcAddr
// 	}
// 	if port != 8545 {
// 		return fmt.Sprintf("%s://%s:8545%s", u.Scheme, host, u.Path)
// 	}
// 	return rpcAddr
// }

func NewGnfdCompositClients(rpcAddrs []string, evmAddrs []string, chainId string, privateKey string, account *types.Account, useWebsocket bool, srcMocaSBTContractAddr, srcMocaVCContractAddr string) GnfdCompositeClients {
	clients := make([]*MocaClient, 0)
	for i := 0; i < len(rpcAddrs); i++ {
		sdkClient, err := sdkclient.New(chainId, rpcAddrs[i], evmAddrs[i], privateKey, sdkclient.Option{DefaultAccount: account, UseWebSocketConn: useWebsocket})
		if err != nil {
			logging.Logger.Errorf("rpc node %s is not available", rpcAddrs[i])
			continue
		}

		ethClient, err := ethclient.Dial((evmAddrs[i]))
		if err != nil {
			panic("new eth client error")
		}
		mocaSBTCrossChainClient, err := mocacrosschainupgradeable.NewIMOCACrossChainUpgradeable(
			common.HexToAddress(srcMocaSBTContractAddr),
			ethClient)
		if err != nil {
			panic("new mocaCrossChain client error")
		}
		mocaVCCrossChainClient, err := universalVerifier.NewUniversalVerifier(
			common.HexToAddress(srcMocaVCContractAddr),
			ethClient)
		if err != nil {
			panic("new mocaVCCrossChain client error")
		}
		clients = append(clients, &MocaClient{
			IClient:                 sdkClient,
			ethClient:               ethClient,
			mocaSBTCrossChainClient: mocaSBTCrossChainClient,
			mocaVCCrossChainClient:  mocaVCCrossChainClient,
		})
		if len(clients) == 0 {
			panic("no Moca client available")
		}
	}
	return GnfdCompositeClients{
		clients: clients,
	}
}

func (gc *GnfdCompositeClients) GetClient() *MocaClient {
	wg := new(sync.WaitGroup)
	wg.Add(len(gc.clients))
	clientCh := make(chan *MocaClient)
	waitCh := make(chan struct{})
	go func() {
		for _, c := range gc.clients {
			go getClientBlockHeight(clientCh, wg, c)
		}
		wg.Wait()
		close(waitCh)
	}()
	var maxHeight int64
	maxHeightClient := gc.clients[0]
	for {
		select {
		case c := <-clientCh:
			if c.Height > maxHeight {
				maxHeight = c.Height
				maxHeightClient = c
			}
		case <-waitCh:
			return maxHeightClient
		}
	}
}

func getClientBlockHeight(clientChan chan *MocaClient, wg *sync.WaitGroup, client *MocaClient) {
	defer wg.Done()
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.GetStatus(ctxWithTimeout)
	if err != nil {
		return
	}
	latestHeight := status.SyncInfo.LatestBlockHeight
	client.Height = latestHeight
	clientChan <- client
}
