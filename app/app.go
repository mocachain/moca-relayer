package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mocachain/moca-relayer/assembler"
	"github.com/mocachain/moca-relayer/config"
	"github.com/mocachain/moca-relayer/db/dao"
	"github.com/mocachain/moca-relayer/db/model"
	"github.com/mocachain/moca-relayer/executor"
	"github.com/mocachain/moca-relayer/listener"
	"github.com/mocachain/moca-relayer/metric"
	"github.com/mocachain/moca-relayer/relayer"
	"github.com/mocachain/moca-relayer/vote"
)

type App struct {
	BSCRelayer    *relayer.BSCRelayer
	GnfdRelayer   *relayer.MocaRelayer
	metricService *metric.MetricService
}

func NewApp(cfg *config.Config) *App {
	username := viper.GetString(config.FlagConfigDbUsername)
	if username == "" {
		username = os.Getenv(config.ConfigDBUserName)
		if username == "" {
			username = cfg.DBConfig.Username
		}
	}
	password := viper.GetString(config.FlagConfigDbPass)
	if password == "" {
		password = os.Getenv(config.ConfigDBPass)
		if password == "" {
			password = getDBPass(&cfg.DBConfig)
		}
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	var db *gorm.DB
	var err error
	var dialector gorm.Dialector

	if cfg.DBConfig.Dialect == config.DBDialectMysql {
		url := cfg.DBConfig.Url
		dbPath := fmt.Sprintf("%s:%s@%s", username, password, url)
		dialector = mysql.Open(dbPath)
	} else {
		panic(fmt.Sprintf("unexpected DB dialect %s", cfg.DBConfig.Dialect))
	}
	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	dbConfig, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbConfig.SetMaxIdleConns(cfg.DBConfig.MaxIdleConns)
	dbConfig.SetMaxOpenConns(cfg.DBConfig.MaxOpenConns)

	model.InitBSCTables(db)
	model.InitMocaTables(db)
	model.InitVoteTables(db)

	metricService := metric.NewMetricService(cfg)

	mocaDao := dao.NewMocaDao(db)
	bscDao := dao.NewBSCDao(db)
	voteDao := dao.NewVoteDao(db)
	daoManager := dao.NewDaoManager(mocaDao, bscDao, voteDao)

	mocaExecutor := executor.NewMocaExecutor(cfg)
	bscExecutor := executor.NewBSCExecutor(cfg, metricService)

	mocaExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetMocaExecutor(mocaExecutor)

	// vote signer
	signer := vote.NewVoteSigner(mocaExecutor.BlsPrivateKey)

	// voteProcessors
	mocaVoteProcessor := vote.NewMocaVoteProcessor(cfg, daoManager, signer, mocaExecutor)
	bscVoteProcessor := vote.NewBSCVoteProcessor(cfg, daoManager, signer, bscExecutor)

	// listeners
	mocaListener := listener.NewMocaListener(cfg, mocaExecutor, bscExecutor, daoManager, metricService)
	bscListener := listener.NewBSCListener(cfg, bscExecutor, mocaExecutor, daoManager, metricService)

	// assemblers
	mocaAssembler := assembler.NewMocaAssembler(cfg, mocaExecutor, daoManager, bscExecutor, metricService)
	bscAssembler := assembler.NewBSCAssembler(cfg, bscExecutor, daoManager, mocaExecutor, metricService)

	// relayers
	gnfdRelayer := relayer.NewMocaRelayer(mocaListener, mocaExecutor, bscExecutor, mocaVoteProcessor, mocaAssembler)
	bscRelayer := relayer.NewBSCRelayer(bscListener, mocaExecutor, bscExecutor, bscVoteProcessor, bscAssembler)

	return &App{
		BSCRelayer:    bscRelayer,
		GnfdRelayer:   gnfdRelayer,
		metricService: metricService,
	}
}

func (a *App) Start() {
	a.GnfdRelayer.Start()
	a.BSCRelayer.Start()
	a.metricService.Start()
}

func getDBPass(cfg *config.DBConfig) string {
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSSecretName, cfg.AWSRegion)
		if err != nil {
			panic(err)
		}
		type DBPass struct {
			DbPass string `json:"db_pass"`
		}
		var dbPassword DBPass
		err = json.Unmarshal([]byte(result), &dbPassword)
		if err != nil {
			panic(err)
		}
		return dbPassword.DbPass
	}
	return cfg.Password
}
