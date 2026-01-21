package models

type Radcheck struct {
    ID        int    `gorm:"primaryKey;autoIncrement"`
    Username  string `json:"Username"`
    Attribute string `json:"Attribute"`
    Op        string `json:"-"`
    Value     string `json:"Value"`
    Options []Radreply `json:"Options" gorm:"foreignKey:Username;references:Username"`
}

func (Radcheck) TableName() string {
    return "radcheck"
}

type Radreply struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `json:"-"`
    Attribute string `json:"Attribute"`
    Op        string `json:"-"`
    Value     string `json:"Value"`
}

func (Radreply) TableName() string {
    return "radreply"
}