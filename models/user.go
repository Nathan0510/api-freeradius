package models

type Radcheck struct {
    ID        int    `gorm:"primaryKey;autoIncrement"`
    Username  string
    Attribute string
    Op        string `json:"-"`
    Value     string
    Options []Radreply `json:"Options" gorm:"foreignKey:Username;references:Username"`
}

func (Radcheck) TableName() string {
    return "radcheck"
}

type Radreply struct {
    ID        uint   `gorm:"primaryKey" json:"-"` 
    Username  string `json:"-"`
    Attribute string `json:"Attribute"`
    Op        string `json:"-"`
    Value     string `json:"Value"`
}

func (Radreply) TableName() string {
    return "radreply"
}