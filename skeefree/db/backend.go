package db

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/startupheroes/kubebot/skeefree/core"
)

type Backend struct {
	db *sqlx.DB
}

type DBConfig struct {
	User string
	Passwd string
	Net string
	Addr string
	DBName string
	ParseTime bool
	InterpolateParams bool
	Timeout int
}

func NewBackend() *Backend {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Net = "tcp"
	cfg.Addr = "localhost"
	cfg.DBName = "kubebot"
	cfg.ParseTime = true
	cfg.InterpolateParams = true
	cfg.Timeout = 1 * time.Second
	db, _ := sqlx.Open("mysql", cfg.FormatDSN())
	return &Backend{
		db: db,
	}
}

func (backend *Backend) SubmitPR(pr *core.PullRequest) (err error) {
	sqlResult, err := backend.db.NamedExec(`
	    update pull_requests set
				title=:title,
				author=:author,
				priority=:priority,
				status=:status,
				is_open=:is_open,
				added_timestamp=now()
			where
				pull_request_number=:pull_request_number
    `, pr)
	if err != nil {
		return err
	}
	affected, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if affected > 0 {
		// good, row updated
		return nil
	}
		// No rows affected? try and isnert the row
		_, err = backend.db.NamedExec(`
	    insert into pull_requests (
				pull_request_number, title, author, priority, status, is_open,
				added_timestamp
			) values (
				:pull_request_number, :title, :author, :priority, :status, :is_open,
				now()
			)
    `, pr)
	return err
}

