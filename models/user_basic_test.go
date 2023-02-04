package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserBasic_TableName(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gin_chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&UserBasic{})
	// Create
	user := &UserBasic{}
	user.Name = "申专"
	user.LoginTime = time.Now()
	user.HeartbeatTime = time.Now()
	user.LoginOutTime = time.Now()
	tx := db.Create(user)
	if tx.Error != nil {
		panic("insert row fail")
	}

	// Read
	fmt.Println(db.First(user)) // 根据整型主键查找
	//db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// Update - 将 product 的 price 更新为 200
	db.Model(user).Update("PassWord", "1234")
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - 删除 product
	//db.Delete(&user)
}
