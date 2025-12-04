package models

type Radcheck struct {
    ID        int    `gorm:"primaryKey;autoIncrement"`
    Username  string
    Attribute string
    Op        string
    Value     string
}

func (Radcheck) TableName() string {
    return "radcheck"
}
