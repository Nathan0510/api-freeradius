package models

type Radreply struct {
    ID        int    `gorm:"primaryKey;autoIncrement"`
    Username  string
    Attribute string
    Op        string
    Value     string
}

func (Radreply) TableName() string {
    return "radreply"
}
