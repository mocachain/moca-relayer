package relayer

import (
	"github.com/mocachain/moca-relayer/assembler"
	"github.com/mocachain/moca-relayer/executor"
	"github.com/mocachain/moca-relayer/listener"
	"github.com/mocachain/moca-relayer/vote"
)

type BSCRelayer struct {
	Listener      *listener.BSCListener
	MocaExecutor  *executor.MocaExecutor
	bscExecutor   *executor.BSCExecutor
	voteProcessor *vote.BSCVoteProcessor
	assembler     *assembler.BSCAssembler
}

func NewBSCRelayer(listener *listener.BSCListener, mocaExecutor *executor.MocaExecutor,
	bscExecutor *executor.BSCExecutor, voteProcessor *vote.BSCVoteProcessor,
	bscAssembler *assembler.BSCAssembler,
) *BSCRelayer {
	return &BSCRelayer{
		Listener:      listener,
		MocaExecutor:  mocaExecutor,
		bscExecutor:   bscExecutor,
		voteProcessor: voteProcessor,
		assembler:     bscAssembler,
	}
}

func (r *BSCRelayer) Start() {
	go r.MonitorEventsLoop()
	go r.SignAndBroadcastVoteLoop()
	go r.CollectVotesLoop()
	go r.AssemblePackagesLoop()
	go r.UpdateCachedLatestValidatorsLoop()
	go r.UpdateClientLoop()
	go r.ClaimRewardLoop()
	go r.PurgeLoop()
}

// MonitorEventsLoop will monitor cross chain events for every block and persist into DB
func (r *BSCRelayer) MonitorEventsLoop() {
	r.Listener.StartLoop()
}

func (r *BSCRelayer) SignAndBroadcastVoteLoop() {
	r.voteProcessor.SignAndBroadcastVoteLoop()
}

func (r *BSCRelayer) CollectVotesLoop() {
	r.voteProcessor.CollectVotesLoop()
}

func (r *BSCRelayer) AssemblePackagesLoop() {
	r.assembler.AssemblePackagesAndClaimLoop()
}

func (r *BSCRelayer) UpdateCachedLatestValidatorsLoop() {
	r.bscExecutor.UpdateCachedLatestValidatorsLoop() // cache validators queried from moca, update it every 1 minute
}

func (r *BSCRelayer) UpdateClientLoop() {
	r.bscExecutor.UpdateClientLoop()
}

func (r *BSCRelayer) ClaimRewardLoop() {
	r.bscExecutor.ClaimRewardLoop()
}

func (r *BSCRelayer) PurgeLoop() {
	r.Listener.PurgeLoop()
}
