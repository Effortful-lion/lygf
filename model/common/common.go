package common

import (
    "encoding/json"
    "errors"
)

type UserType int

// 用户类型
const (
    USER UserType = iota // 0
    SHOP                 // 1
)

// 反序列化：将json数据反序列化为UserType类型
func (u *UserType) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }

    switch s {
    case "USER":
        *u = USER
    case "SHOP":
        *u = SHOP
    default:
        return errors.New("invalid user type")
    }
    return nil
}