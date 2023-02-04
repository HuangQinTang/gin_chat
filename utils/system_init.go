package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() error {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	fmt.Println("config app inited ...")
	return nil
}
func InitMySQL() error {
	var err error
	if DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{}); err != nil {
		return err
	}
	fmt.Println(" MySQL inited ...")
	return nil
}
