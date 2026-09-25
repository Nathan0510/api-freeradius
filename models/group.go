package models

type Radgroupreply struct {
    ID        int    `gorm:"primaryKey;autoIncrement"`
    Groupname  string `json:"Groupname"`
    Attribute string `json:"Attribute"`
    Op        string `json:"-"`
    Value     string `json:"Value"`
}

func (Radgroupreply) TableName() string {
    return "radgroupreply"
}

type Radusergroup struct {
    ID        int    `gorm:"primaryKey;autoIncrement"`
    Username  string `json:"-"`
    Groupname string `json:"Groupname"`
    Priority int `json:"Priority"`
}

func (Radusergroup) TableName() string {
    return "radusergroup"
}