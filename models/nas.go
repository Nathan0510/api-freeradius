package models

type Nas struct {
    ID          int    `gorm:"primaryKey;autoIncrement"`
    Nasname     string
    Shortname   string
    Secret      string
    Type        string
    Ports       int
    Server      string
    Community   string
    Description string
}

func (Nas) TableName() string {
    return "nas"
}

type NewNas struct {
    Nasname     string  `json:"Nasname"`
    Shortname   string  `json:"shortname"`
    Secret      string  `json:"secret"`
}