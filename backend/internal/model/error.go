// backend/internal/model/error.go
package model

import "errors"

var (
    ErrInvalidCredentials = errors.New("無効な認証情報です")
    ErrUserNotFound      = errors.New("ユーザーが見つかりません")
    ErrUserAlreadyExists = errors.New("ユーザーは既に存在します")
    ErrInvalidToken      = errors.New("無効なトークンです")
)
