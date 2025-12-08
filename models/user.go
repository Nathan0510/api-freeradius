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

type Option struct {
    Attribute string `json:"attribute"`
    Value     string `json:"value"`
}

type User struct {
    ID       int     `gorm:"primaryKey"`
    Username string  `json:"username"`
    Password string  `json:"password"`
    Options  []Option `json:"options" gorm:"type:jsonb"`
}

func (User) TableName() string {
    return "radcheck"
}
