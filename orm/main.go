package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// 设置数据库连接参数
	dsn := "root:031119@tcp(127.0.0.1:3306)/information?charset=utf8&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}

	// 执行数据库迁移（创建表）
	db.AutoMigrate(&User{})

	// 创建用户
	user := User{ID: 222, Name: "Doe", Age: 30}
	db.Create(&user)

	users := []*User{{ID: 333, Name: "John", Age: 25},
		{ID: 111, Name: "Lily", Age: 18},
		{ID: 666, Name: "Helen", Age: 20},
		{ID: 888, Name: "Jay", Age: 15}}
	db.Create(&users)
	// 查询用户
	var result1 User
	db.First(&result1, 666)
	fmt.Println(result1)

	var result2 User
	db.First(&result2, 333)
	fmt.Println(result2)

	// 更新用户
	db.Model(&result1).Update("Age", 40)

	// 删除用户
	db.Delete(&result2)

}
