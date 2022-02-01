package unittest

import (
	"testing"

	"github.com/andiahmads/go-api/config"
	"gorm.io/gorm"
)

func TestGetBookWithDbRaw(t *testing.T) {
	var db *gorm.DB = config.SetupDatabaseConnection()
	_, err := db.DB()
	if err != nil {
		t.Log(err)
	}

}
