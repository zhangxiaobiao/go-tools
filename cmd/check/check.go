package check

import (
	"bb/cmd/global"
	"fmt"

	"github.com/golang-module/carbon"
)

type Check struct {
	Id    int    `json:"id"`
	Date  string `json:"date"`
	Start string `json:"start"`
	End   string `json:"end"`
}

var check *Check

func GetCheck() *Check {
	return check
}

/* 保存起始时间 */
func (c *Check) SaveTime(start string, end string) int {
	startDate := carbon.Parse(start).ToDateString()
	startTime := carbon.Parse(start).ToTimeString()
	endTime := carbon.Parse(end).ToTimeString()
	//查询是否有 该日期记录，如果有，则更新
	result := global.Db.Where("date = ?", startDate).First(&check)
	if int(result.RowsAffected) > 0 {
		check.Start = startTime
		check.End = endTime
		result = global.Db.Save(&check)
	} else {
		/* 插入一条新记录 */
		check := Check{Date: startDate, Start: startTime, End: endTime}
		result = global.Db.Create(&check)
	}
	return int(result.RowsAffected)
}

/* 获取当月所有数据 */
func (c *Check) AllChecks(date string) []Check {
	var checks []Check
	startDate := carbon.Parse(date).Format("Y-m")
	result := global.Db.Model(&Check{}).Where("date Like ?", startDate+"%").Find(&checks)
	fmt.Println(result.Error)
	return checks
}

/* 删除某一天数据 */
func (c *Check) DelChecks(date string) int {
	result := global.Db.Where("date = ?", date).Delete(Check{})
	fmt.Println(result.Error)
	return int(result.RowsAffected)
}
