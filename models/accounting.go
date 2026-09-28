package models

type Radacct struct {
    Username 			 	string `json:"Username"`
    Nasipaddress 		 	string `json:"Nasipaddress"`
    Acctstarttime        	string `json:"Acctstarttime"`
    Acctstoptime			string `json:"Acctstoptime"`
    Acctinputoctets     	string `json:"Acctinputoctets"`
	Acctoutputoctets     	string `json:"Acctoutputoctets"`
	Callingstationid     	string `json:"Callingstationid"`
	Acctterminatecause     	string `json:"Acctterminatecause"`
}

func (Radacct) TableName() string {
    return "radacct"
}