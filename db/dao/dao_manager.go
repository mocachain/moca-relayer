package dao

type DaoManager struct {
	MocaDao *MocaDao
	VoteDao *VoteDao
	BSCDao  *BSCDao
}

func NewDaoManager(mocaDao *MocaDao, bscDao *BSCDao, voteDao *VoteDao) *DaoManager {
	return &DaoManager{
		MocaDao: mocaDao,
		VoteDao: voteDao,
		BSCDao:  bscDao,
	}
}
