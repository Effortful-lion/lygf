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
        *u = USER       // 0
    case "SHOP":
        *u = SHOP       // 1
    default:
        return errors.New("无效的用户类型")     // 如果前端采用单选框，这里可以返回nil
    }
    return nil
}