package entity

type Check struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//日期
	Date string `gorm:"column:date;default:'';not null;comment:'日期';INDEX:idx_date"`
	//开始时间
	Start string `gorm:"column:start;default:'';not null;comment:'开始时间'"`
	//结束时间
	End string `gorm:"column:end;default:'';not null;comment:'结束时间'"`
}

func (Check) TableName() string {
	return "checks"
}
