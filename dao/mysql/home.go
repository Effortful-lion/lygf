package mysql

import (
	"lygf/backend/model/entity"
)

func GetShops(cateID1, cateID2 string) *[]entity.Shop {
	data := &[]entity.Shop{}
	// 根据一级分类ID判断：0：热门（hot==1）；1：推荐(score前10)；2：分类
	condition := make(map[string]interface{})
	switch cateID1 {
		case "0":
			condition["hot"] = true
		case "1":
			condition["score"] = GetShopScoreTop10()
		default:break
	}
	data = GetShopsByCategory(condition,cateID2)
	return data
}

// 根据商铺的类型分类
func GetShopsByCategory(condition map[string]interface{},cateID2 string) *[]entity.Shop {
	shops := &[]entity.Shop{}
	if condition["hot"] != nil {
		db.Where("hot = ? and category = ?", true, cateID2).Find(shops)
	}else if condition["score"] != nil {
		db.Where("score >= ? and category = ?", condition["score"], cateID2).Find(shops)
	}else{
		db.Where("category = ?", cateID2).Find(shops)
	}
	return shops
}

// 获取商铺评分前10: 返回分数第10位的分数
func GetShopScoreTop10() float64 {
	var score float64
	result := db.Model(&entity.Shop{}).Order("score desc").Limit(1).Offset(4).Pluck("score", &score)
	if result.Error != nil { 
		return 0
	}
	return score
}

