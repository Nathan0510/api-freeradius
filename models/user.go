package models

type Option struct {
    Attribute string `json:"attribute"`
    Value     string `json:"value"`
}

type User struct {
    Username string  `json:"username"`
    Password string  `json:"password"`
    Options  []Option `json:"Options"`
}