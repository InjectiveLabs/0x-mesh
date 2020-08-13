package p2p

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

const DHTProtocolID = protocol.ID("/injective-0x-mesh-dht/version/1")

// DefaultBootstrapList is a list of addresses to use by default for
// bootstrapping the DHT.
var DefaultBootstrapList = []string{
	"/dns4/mesh-bootstrap.injective.dev/tcp/40558/ipfs/16Uiu2HAm2VAoigVxsZy1i5anRc1DdkCpDAJRkqQZh6j3wsPcH2Pb",
	"/dns4/mesh-bootstrap.injective.dev/tcp/40559/ws/ipfs/16Uiu2HAm2VAoigVxsZy1i5anRc1DdkCpDAJRkqQZh6j3wsPcH2Pb",

	// These nodes are provided by the libp2p community on a best-effort basis.
	// We're using them as a backup for increased redundancy.
	"/ip4/34.201.54.78/tcp/4001/ipfs/12D3KooWHwJDdbx73qiBpSCJfg4RuYyzqnLUwfLBqzn77TSy7kRX",
	"/ip4/18.204.221.103/tcp/4001/ipfs/12D3KooWQS6Gsr2kLZvF7DVtoRFtj24aar5jvz88LvJePrawM3EM",
}

func BootstrapListToAddrInfos(bootstrapList []string) ([]peer.AddrInfo, error) {
	maddrs := make([]ma.Multiaddr, len(bootstrapList))
	for i, addrString := range bootstrapList {
		maddr, err := ma.NewMultiaddr(addrString)
		if err != nil {
			return nil, err
		}
		maddrs[i] = maddr
	}
	return peer.AddrInfosFromP2pAddrs(maddrs...)
}

func ConnectToBootstrapList(ctx context.Context, host host.Host, bootstrapList []string) error {
	log.WithField("bootstrapList", bootstrapList).Info("connecting to bootstrap peers")
	bootstrapAddrInfos, err := BootstrapListToAddrInfos(bootstrapList)
	if err != nil {
		return err
	}
	connectCtx, cancel := context.WithTimeout(ctx, defaultNetworkTimeout)
	defer cancel()
	wg := sync.WaitGroup{}
	for _, peerInfo := range bootstrapAddrInfos {
		if peerInfo.ID == host.ID() {
			// Don't connect to self.
			continue
		}
		wg.Add(1)
		go func(peerInfo peer.AddrInfo) {
			defer wg.Done()
			if err := host.Connect(connectCtx, peerInfo); err != nil {
				log.WithFields(map[string]interface{}{
					"error":    err.Error(),
					"peerInfo": peerInfo,
				}).Warn("failed to connect to bootstrap peer")
			}
		}(peerInfo)
	}
	wg.Wait()

	// It is recommended to wait for 2 seconds after connecting to all the
	// bootstrap peers to give time for the relevant notifees to trigger and the
	// DHT to fully initialize.
	// See: https://github.com/0xProject/0x-mesh/pull/69#discussion_r286849679
	time.Sleep(2 * time.Second)

	return nil
}
