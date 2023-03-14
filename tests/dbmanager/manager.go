package dbmanager

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/galexrt/arpanet/query"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.uber.org/zap"
)

var TestDBManager *DBManager

type DBManager struct {
	db       *sql.DB
	pool     *dockertest.Pool
	resource *dockertest.Resource

	stopped bool
}

func init() {
	TestDBManager = &DBManager{}
}

func (m *DBManager) Setup() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	var err error
	m.pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = m.pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	m.resource, err = m.pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "docker.io/library/mariadb",
			Tag:        "10.11.2-jammy",
			Env: []string{
				"MYSQL_ROOT_PASSWORD=secret",
				"MYSQL_USER=arpanet",
				"MYSQL_PASSWORD=changeme",
				"MYSQL_DATABASE=arpanettest",
			},
			Cmd: []string{
				"mariadbd",
				"--innodb-ft-min-token-size=2",
				"--innodb-ft-max-token-size=50",
			},
		},
		func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := m.pool.Retry(func() error {
		var err error
		m.db, err = sql.Open("mysql", m.getDSN())
		if err != nil {
			return err
		}
		return m.db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	m.prepareDBForFirstUse()

	m.LoadTestData()
}

func (m *DBManager) DB() *sql.DB {
	if m.db == nil {
		log.Fatal("Test DB connection has not been established! You are accessing DB() method too early.")
	}

	return m.db
}

func (m *DBManager) getDSN() string {
	return fmt.Sprintf("arpanet:changeme@(localhost:%s)/arpanettest?collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", m.resource.GetPort("3306/tcp"))
}

func (m *DBManager) prepareDBForFirstUse() {
	// Load and apply premigrate.sql file
	m.loadSQLFile(filepath.Join("../../../", "tests", "testdata", "premigrate.sql"))

	// Use DB migrations to handle the rest
	if err := query.MigrateDB(zap.NewNop(), m.getDSN()); err != nil {
		log.Fatalf("Failed to migrate test database: %s", err)
	}
}

func (m *DBManager) getMultiStatementDB() *sql.DB {
	// Open db connection with multiStatements param so we can apply sql files
	initDB, err := sql.Open("mysql", m.getDSN()+"&multiStatements=true")
	if err != nil {
		log.Fatalf("Failed to open test database connection for premigrate.sql apply: %s", err)
	}
	return initDB
}

func (m *DBManager) loadSQLFile(file string) {
	initDB := m.getMultiStatementDB()

	c, ioErr := os.ReadFile(file)
	if ioErr != nil {
		log.Fatalf("Failed to read premigrate.sql for tests: %s", ioErr)
	}
	sqlBase := string(c)
	if _, err := initDB.Exec(sqlBase); err != nil {
		log.Fatalf("Failed to apply premigrate.sql for tests: %s", err)
	}
}

func (m *DBManager) LoadTestData() {
	m.loadSQLFile(filepath.Join("../../../", "tests", "testdata", "base.sql"))
}

func (m *DBManager) Stop() {
	if m.stopped {
		return
	}
	m.stopped = true

	// You can't defer this because os.Exit doesn't care for defer
	if err := m.pool.Purge(m.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

// Reset truncates all `arpanet_*` tables and loads the base test data
func (m *DBManager) Reset() {
	initDB := m.getMultiStatementDB()

	rows, err := initDB.Query("SHOW TABLES LIKE 'arpanet_%';")
	if err != nil {
		log.Fatalf("Failed to list arpanet tables in test database: %s", err)
	}

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Failed to scan table name to string: %s", err)
		}

		// Placeholders aren't supported for table names, see
		// https://github.com/go-sql-driver/mysql/issues/848#issuecomment-414910152`
		if _, err := initDB.Exec("SET FOREIGN_KEY_CHECKS = 0; TRUNCATE TABLE `" + tableName + "`; SET FOREIGN_KEY_CHECKS = 1;"); err != nil {
			log.Printf("Failed to truncate %s table: %s", tableName, err)
		}
	}

	m.LoadTestData()
}
