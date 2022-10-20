package conf

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB = &gorm.DB{}

func InitDatabase() {
	var dbConfig = Conf().Database
	username := dbConfig.User
	password := dbConfig.Password
	host := dbConfig.Host
	port := dbConfig.Port
	Dbname := dbConfig.DBName

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true, QueryFields: true, CreateBatchSize: 200})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	Db = _db
}
