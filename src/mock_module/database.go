package mockModule

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	dbGorm, errGorm := gorm.Open(dialector, &gorm.Config{})
	if errGorm != nil {
		panic(errGorm)
	}

	return dbGorm, mock
}
