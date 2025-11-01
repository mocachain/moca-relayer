package relayer

import (
	"github.com/mocachain/moca-relayer/assembler"
	"github.com/mocachain/moca-relayer/executor"
	"github.com/mocachain/moca-relayer/listener"
	"github.com/mocachain/moca-relayer/vote"
)

type MocaRelayer struct {
	Listener      *listener.MocaListener
	MocaExecutor  *executor.MocaExecutor
	bscExecutor   *executor.BSCExecutor
	voteProcessor *vote.MocaVoteProcessor
	mocaAssembler *assembler.MocaAssembler
}

func NewMocaRelayer(listener *listener.MocaListener, mocaExecutor *executor.MocaExecutor, bscExecutor *executor.BSCExecutor, voteProcessor *vote.MocaVoteProcessor, mocaAssembler *assembler.MocaAssembler,
) *MocaRelayer {
	return &MocaRelayer{
		Listener:      listener,
		MocaExecutor:  mocaExecutor,
		bscExecutor:   bscExecutor,
		voteProcessor: voteProcessor,
		mocaAssembler: mocaAssembler,
	}
}

func (r *MocaRelayer) Start() {
	go r.MonitorEventsLoop()
	go r.SignAndBroadcastLoop()
	go r.CollectVotesLoop()
	go r.AssembleTransactionsLoop()
	go r.UpdateCachedLatestValidatorsLoop()
	go r.PurgeLoop()
}

// MonitorEventsLoop will monitor cross chain events for every block and persist into DB
func (r *MocaRelayer) MonitorEventsLoop() {
	r.Listener.StartLoop()
}

func (r *MocaRelayer) SignAndBroadcastLoop() {
	r.voteProcessor.SignAndBroadcastLoop()
}

func (r *MocaRelayer) CollectVotesLoop() {
	r.voteProcessor.CollectVotesLoop()
}

func (r *MocaRelayer) AssembleTransactionsLoop() {
	r.mocaAssembler.AssembleTransactionsLoop()
}

func (r *MocaRelayer) UpdateCachedLatestValidatorsLoop() {
	r.MocaExecutor.UpdateCachedLatestValidatorsLoop() // cache validators queried from moca, update it every 1 minute
}

func (r *MocaRelayer) PurgeLoop() {
	r.Listener.PurgeLoop()
}
