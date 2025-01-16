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

// 商家的属性表

type Category int

const (
	FRESH	Category = iota		// 生鲜
	CATERING				    // 餐饮
	SPOT					    // 景点
)

// json tag 自动调用
func (c *Category) UnmarshalJSON(data []byte) error {
    var s string
    // json解析为string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }

    // 根据string类型具体判断并赋值
    switch s {
        case "FRESH":
            *c = FRESH
        case "CATERING":
            *c = CATERING
        case "SPOT":
            *c = SPOT
        default:
            return errors.New("无效的类别")
    }

    return nil
}

// 将指定类型转换为字符串
func (c Category) String() string {
    switch c {
        case FRESH:
            return "生鲜"
        case CATERING:
            return "餐饮"
        case SPOT:
            return "景点"
    }
    return "未知类别"
}

// 一级栏目：热门、推荐、分类
type Cate1 int
const (
    HOT Cate1 = iota
    RECOMMEND
    CATEGORY
)

func (c *Cate1) UnmarshalJSON(data []byte) error {
    var s string
    // json解析为string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }

    // 根据string类型具体判断并赋值
    switch s {
        case "HOT":
            *c = HOT
        case "RECOMMEND":
            *c = RECOMMEND
        case "CATEGORY":
            *c = CATEGORY
        default: return nil
    }

    return nil
}

func (c Cate1) String() string {
    switch c {
        case HOT:
            return "热门"
        case RECOMMEND:
            return "推荐"
        case CATEGORY:
            return "分类"
        default: return "未知类别"
    }
}