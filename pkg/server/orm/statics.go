package orm

type UserAndHaohuoCount struct {
	TotalUsers   uint64
	TotalHaohuos uint64
}

const countUserAndHaohuoSQL = "SELECT u.total_users, h.total_haohuos from (select count(1) as total_users from user) as u, (select count(1) as total_haohuos from haohuo) as h"

func CountUserAndHaohuo() (*UserAndHaohuoCount, error) {
	var ret UserAndHaohuoCount
	if e := database.Raw(countUserAndHaohuoSQL).Scan(&ret).Error; e != nil {
		return nil, e
	}
	return &ret, nil
}
