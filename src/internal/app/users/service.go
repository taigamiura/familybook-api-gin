package users

import (
  "gorm.io/gorm"
)

func GetAllUsersService(db *gorm.DB) ([]User, error) {
  // ユーザー情報を格納するためのスライスを定義します。
  var users []User
  // データベースから全ユーザー情報を取得します。
  tx := db.Table("users").Find(&users)
  // エラーが発生した場合は、エラーを返します。
  if tx.Error != nil {
    return nil, tx.Error
  }
  // ユーザー情報を返します。
  return users, nil
}
