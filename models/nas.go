package models

type Nas struct {
    ID          int    `gorm:"primaryKey;autoIncrement"`
    Nasname     string `json:"Username"`
    Shortname   string `json:"Shortname"`
    Secret      string `json:"Secret"`
    Type        string `json:"Type"`
    Ports       int    `json:"Ports"`
    Server      string `json:"Server"`
    Community   string `json:"Community"`
    Description string `json:"Description"`
}

func (Nas) TableName() string {
    return "nas"
}