# Freeradius API (Golang)

API developed in golang that allows you to administer a freeradius server (group, user, NAS and accounting management).

---

## Features

- Radius user management
- Radius NAS management
- Radius group managemen
- Swagger documentation

---

## Technologies used

- Go
- Gin Gonic
- FreeRADIUS
- PostgreSQL
- Swagger

---

## Configuration

Edit .env file (cmd/api/.env)
<pre>
DB_HOST=localhost
DB_USER=user
DB_PASSWORD=password
DB_NAME=radius
DB_PORT=5432
DB_SSLMODE=disable
</pre>

# Launch the API

<pre>
git clone https://github.com/Nathan0510/api-freeradius.git
cd api-freeradius/
go mod tidy
go run cmd/api/main.go
</pre>

The API will be available with localhost:8080/api

# Documentation Swagger

Available at localhost:8080/swagger/index.html
![alt text](image.png)

# Example API with curl

For NAS :
<pre>
Post Nas :
curl -X POST http://localhost:8080/api/nas -H "Content-Type: application/json" -d '{"nasname": "1.1.1.1","shortname": "LNS1","secret": "naruto"}'
{"message":"Nas 1.1.1.1 added successfully"}


Get all Nas :
curl http://localhost:8080/api/nas
[{"ID":14,"Nasname":"1.1.1.1","Shortname":"LNS1","Secret":"naruto","Type":"","Ports":0,"Server":"","Community":"","Description":""},{"ID":15,"Nasname":"2.2.2.2","Shortname":"LNS2","Secret":"naruto","Type":"","Ports":0,"Server":"","Community":"","Description":""}]


Get Nas :
curl http://localhost:8080/api/nas/1.1.1.1
{"ID":14,"Nasname":"1.1.1.1","Shortname":"LNS1","Secret":"naruto","Type":"","Ports":0,"Server":"","Community":"","Description":""}


Patch Nas :
curl -X PATCH http://localhost:8080/api/nas/1.1.1.1 -H "Content-Type: application/json" -d '{"secret": "sasuke"}'
{"message":"Nas 1.1.1.1 updated successfully"}


Delete Nas :
curl -X DELETE http://localhost:8080/api/nas/1.1.1.1
{"message":"Nas 1.1.1.1 deleted successfully"}
</pre>

For Users :
<pre>
Post user :
curl -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d '{"Username":"naruto@naruto.ninja","Attribute": "Cleartext-Password","Value":"beaugoss","Options":[{"Attribute": "Framed-IP-Address","Value": "100.127.0.1"},{"Attribute": "Profile-Mikrotik","Value": "Profile-Internet"}],"Groups":[{"Groupname":"Mikrotik-1G","Priority":1}]}'
{"message":"User naruto@naruto.ninja added successfully"}


Get all user :
curl http://localhost:8080/api/users
[{"ID":1,"Username":"minato@naruto.ninja","Attribute":"Cleartext-Password","Value":"surcote","Options":[]},{"ID":2,"Username":"naruto@naruto.ninja","Attribute":"Cleartext-Password","Value":"beaugoss","Options":[{"Attribute":"Framed-IP-Address","Value":"100.127.0.1"},{"Attribute":"Mikrotik-Group","Value":"Profile-Internet"}]}]


Get user :
curl http://localhost:8080/api/users/naruto@naruto.ninja
{"ID":27,"Username":"naruto@naruto.ninja","Attribute":"Cleartext-Password","Value":"sasuke2","Options":[{"Attribute":"Framed-IP-Address","Value":"100.64.0.10"},{"Attribute":"Framed-IP-Address","Value":"100.127.0.1"},{"Attribute":"Profile-Mikrotik","Value":"Profile-Internet"},{"Attribute":"Profile-Mikrotik","Value":"Profile-Internet"}]}


Patch user :
curl -X PATCH http://localhost:8080/api/users/naruto@naruto.ninja -H "Content-Type: application/json" -d '{"Options":[{"Attribute": "Framed-IP-Address","Value": "100.127.0.10"}]}'
{"message":"User naruto@naruto.ninja updated successfully"}


Delete option user :
curl -X DELETE http://localhost:8080/api/users/naruto@naruto.ninja/options/Framed-IP-Address?value=100.127.0.10
{"message":"Option Framed-IP-Address value 100.127.0.10 user naruto@naruto.ninja deleted successfully"}

Delete user :
curl -X DELETE http://localhost:8080/api/users/naruto@naruto.ninja
{"message":"User naruto@naruto.ninja deleted successfully"}
</pre>

For Groups :

<pre>
Post group :
curl -X POST http://192.168.1.241:8080/api/groups -H "Content-Type: application/json" -d '{"Groupname":"Mikrotik-500M","Attribute":"Mikrotik-Rate-Limit","Value":"500M/500M"}'
{"message":"GroupMikrotik-500M added successfully"}

Get all groups :
curl -X GET http://192.168.1.241:8080/api/groups
[{"ID":2,"Groupname":"Mikrotik-1G","Attribute":"Mikrotik-Rate-Limit","Value":"1G/1G"},{"ID":4,"Groupname":"Mikrotik-500M","Attribute":"Mikrotik-Rate-Limit","Value":"500M/500M"}]

Get group :
curl -X GET http://192.168.1.241:8080/api/groups/Mikrotik-500M
{"ID":4,"Groupname":"Mikrotik-500M","Attribute":"Mikrotik-Rate-Limit","Value":"500M/500M"}

Patch group :
curl -X PATCH http://192.168.1.241:8080/api/groups/Mikrotik-500M -H "Content-Type: application/json" -d '{"Value":"50M/50M"}'
{"message":"GroupMikrotik-500M updated successfully"}

Delete group :
curl -X DELETE http://192.168.1.241:8080/api/groups/Mikrotik-500M
{"message":"Group Mikrotik-500M deleted successfully"}
</pre>


For Accounting :
<pre>
Get accounting of a username :
curl -X GET curl -X GET http://localhost:8080/api/accounting/naruto@naruto.ninja
[{"Username":"naruto@naruto.ninja","Nasipaddress":"192.168.10.98","Acctstarttime":"2026-09-28T18:26:44Z","Acctstoptime":"","Acctinputoctets":"2148","Acctoutputoctets":"182","Callingstationid":"52:54:00:e6:56:01","Acctterminatecause":""},{"Username":"naruto@naruto.ninja","Nasipaddress":"192.168.10.98","Acctstarttime":"2026-09-27T18:58:41Z","Acctstoptime":"2026-09-27T19:06:34Z","Acctinputoctets":"554168","Acctoutputoctets":"540582","Callingstationid":"52:54:00:68:fb:01","Acctterminatecause":"NAS-Error"}]
</pre>
Enjoy !