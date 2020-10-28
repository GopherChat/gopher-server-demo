package sqlstore

import (
	"github.com/GopherChat/gopher-server-demo/model"
	"github.com/GopherChat/gopher-server-demo/module/klog"
	"github.com/GopherChat/gopher-server-demo/store"
	"xorm.io/xorm"
)

var (
	tables []interface{}
)

func init() {
	tables = append(tables,
		new(model.User),
	)
}

type collection struct {
	store.UserStore
}

type SqlStore struct {
	engine *xorm.Engine
	logger *klog.Logger
	coll   collection
}

func NewSqlStore(l *klog.Logger) *SqlStore {
	sqlStore := &SqlStore{
		logger: l,
	}

	// sqlStore.initConnection()

	sqlStore.coll.UserStore = newSqlUserStore(sqlStore)

	return sqlStore
}

func (ss *SqlStore) User() store.UserStore {
	return ss.coll.UserStore
}

// func (ss *SqlStore) initConnection() {
// 	klog.Info("Beginning SQL engine initialization.")

// 	reties := ss.cfg.ConnectRetries
// 	backOff := ss.cfg.ConnectBackoff * time.Second

// 	for i := 0; i < reties; i++ {
// 		klog.Info(fmt.Sprintf("SQL engine initialization attempt #%d/%d...", i+1, reties))

// 		e, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@%s", ss.cfg.User, ss.cfg.Password, ss.cfg.DSN))
// 		if err != nil {
// 			klog.Error(fmt.Sprintf("SQL engine initialization attempt #%d/%d failed", i+1, reties), err)
// 			if i == reties-1 {
// 				klog.Error("SQL engine connection retries ended", err)
// 				os.Exit(1)
// 			}
// 		} else {
// 			if err := e.Sync2(tables...); err != nil {
// 				klog.Error("SQL engine auto migration failed", err)
// 				os.Exit(1)
// 			}
// 			ss.engine = e
// 			return
// 		}

// 		klog.Info(fmt.Sprintf("Backing off for %d seconds", backOff))
// 		time.Sleep(backOff)
// 	}
// }
