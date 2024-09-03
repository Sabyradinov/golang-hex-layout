package gorm

import (
	"github.com/Sabyradinov/golang-hex-layout/config"
	"github.com/Sabyradinov/golang-hex-layout/internal/adapter/storage/gorm/repo"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/repository"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/storage"
	"gorm.io/gorm"
)

// dbClient instance to db client
type dbClient struct {
	DB  *gorm.DB
	cfg *config.Configs
}

// InitDB initialize DB
func (dbm *dbClient) InitDB(cfg *config.Configs) (err error) {
	//logMode := logger.Error
	//if cfg.DB.LogMode {
	//	logMode = logger.Info
	//}

	dbm.cfg = cfg
	dbm.DB = &gorm.DB{}
	//dbm.DB, err = gorm.Open(postgres.Open(cfg.DB.ConnectionString), &gorm.Config{
	//	Logger:      logger.Default.LogMode(logMode),
	//	PrepareStmt: true,
	//})

	return err
}

// InitRepo initialize repositories
func (dbm *dbClient) InitRepo() *repository.Manager {
	return &repository.Manager{
		Product:  repo.NewProduct(dbm.DB),
		Shipment: repo.NewShipment(dbm.DB),
	}
}

// Close DB client close
func (dbm *dbClient) Close() (err error) {
	db, err := dbm.DB.DB()
	if err != nil {
		return
	}

	err = db.Close()
	return
}

// RegisterMetrics register prometheus metrics
func (dbm *dbClient) RegisterMetrics(dbname string) (err error) {

	return
}

// NewDBManager DB manager constructor
func NewDBManager(cfg *config.Configs) (storage.IDB, error) {
	dbm := dbClient{}
	err := dbm.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	err = dbm.RegisterMetrics(cfg.DB.Name)

	return &dbm, err
}
