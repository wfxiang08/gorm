package demo

import (
	"testing"
	"database/sql"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/wfxiang08/cyutils/utils/rolling_log"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserSong struct {
	UserId               int64         // 7
	SongId               int64         // 8
	AppName              string        // 2
	CreatedOn            int64         // 1
	IsPermanent          bool          // 3
	HighScore            sql.NullInt64 // 4
	HighStarCount        sql.NullInt64 // 5
	TokensRedeemed       int           // 6
	HighScoreRecordingId sql.NullInt64 // 9
}

// go test github.com/jinzhu/gorm/demo -v -run "TestUserSong$"
func TestUserSong(t *testing.T) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?interpolateParams=true&autocommit=true&charset=utf8mb4,utf8,latin1", "root", "", "localhost", 3306, "node1")
	db, err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.ErrorErrorf(err, "Open database failed")
		return
	}
	db.DB().SetConnMaxLifetime(time.Hour * 4)
	db.DB().SetMaxOpenConns(2) // 设置最大的连接数（防止异常情况干死数据库)
	db.DB().SetMaxIdleConns(2)

	var songs []*UserSong
	log.Printf("Before query")
	db.Table("user_song").Limit(2).Find(&songs)

	for _, song := range songs {
		log.Printf("Song: %v %v %v", song.UserId, song.AppName, song.SongId)
	}

}
