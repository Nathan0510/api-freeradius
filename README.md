# FreeRADIUS API (Golang)

API développée en golang qui permet d'administrer un freeradius (utilisateurs/nas).
---

## Fonctionnalités

- Gestion des utilisateurs RADIUS
- Gestion des nas RADIUS
- Documentation Swagger / OpenAPI

---

## Technologies utilisées

- Go
- Gin Gonic
- FreeRADIUS
- PostgreSQL
- Swagger

---

## Configuration

Edit cmd/api/.env and add your variable

# Launch the API

git clone https://github.com/naruto0510/api-freeradius.git
go mod tidy
go run cmd/api/main.go

The API will be available with localhost:8080

# Documentation Swagger

Available at localhost:8080/swagger/index.html

# Example API with curl

For NAS :

<pre>
Post Nas :
curl -X POST http://192.168.1.240:8080/api/nas -H "Content-Type: application/json" -d '{"nasname": "1.1.1.1","shortname": "LNS1","secret": "naruto"}'
{"message":"Nas 1.1.1.1 added successfully"}


Get all Nas :
curl http://192.168.1.240:8080/api/nas
[{"ID":14,"Nasname":"1.1.1.1","Shortname":"LNS1","Secret":"naruto","Type":"","Ports":0,"Server":"","Community":"","Description":""},{"ID":15,"Nasname":"2.2.2.2","Shortname":"LNS2","Secret":"naruto","Type":"","Ports":0,"Server":"","Community":"","Description":""}]


Get Nas :
curl http://192.168.1.240:8080/api/nas/1.1.1.1
{"ID":14,"Nasname":"1.1.1.1","Shortname":"LNS1","Secret":"naruto","Type":"","Ports":0,"Server":"","Community":"","Description":""}


Patch Nas :
curl -X PATCH http://192.168.1.240:8080/api/nas/1.1.1.1 -H "Content-Type: application/json" -d '{"secret": "sasuke"}'
{"message":"Nas 1.1.1.1 updated successfully"}


Delete Nas :
curl -X DELETE http://192.168.1.240:8080/api/nas/1.1.1.1
{"message":"Nas 1.1.1.1 deleted successfully"}
</pre>

For Users :
<pre>
Post user :
curl -X POST http://192.168.1.240:8080/api/users -H "Content-Type: application/json" -d '{"Username":"naruto@naruto.ninja","Attribute": "Cleartext-Password","Value":"beaugoss","Options":[{"Attribute": "Framed-IP-Address","Value": "100.127.0.1"},{"Attribute": "Profile-Mikrotik","Value": "Profile-Internet"}]}'
{"message":"User naruto@naruto.ninja added successfully"}


Get all user :
[{"ID":1,"Username":"minato@naruto.ninja","Attribute":"Cleartext-Password","Value":"surcote","Options":[]},{"ID":2,"Username":"naruto@naruto.ninja","Attribute":"Cleartext-Password","Value":"beaugoss","Options":[{"Attribute":"Framed-IP-Address","Value":"100.127.0.1"},{"Attribute":"Mikrotik-Group","Value":"Profile-Internet"}]}]


Get user :
curl http://192.168.1.240:8080/api/users/naruto@naruto.ninja
{"ID":27,"Username":"naruto@naruto.ninja","Attribute":"Cleartext-Password","Value":"sasuke2","Options":[{"Attribute":"Framed-IP-Address","Value":"100.64.0.10"},{"Attribute":"Framed-IP-Address","Value":"100.127.0.1"},{"Attribute":"Profile-Mikrotik","Value":"Profile-Internet"},{"Attribute":"Profile-Mikrotik","Value":"Profile-Internet"}]}


Patch user :
curl -X PATCH http://192.168.1.240:8080/api/users/naruto@naruto.ninja -H "Content-Type: application/json" -d '{"Options":[{"Attribute": "Framed-IP-Address","Value": "100.127.0.10"}]}'                                     {"message":"User naruto@naruto.ninja updated successfully"}


Delete option user :
curl -X DELETE http://192.168.1.240:8080/api/users/option/naruto@naruto.ninja -H "Content-Type: application/json" -d '{"Attribute": "Framed-IP-Address","Value": "100.127.0.10"}'
{"message":"Option Framed-IP-Address value 100.127.0.10 user naruto@naruto.ninja deleted successfully"}


Delete user :
curl -X DELETE http://192.168.1.240:8080/api/users/naruto@naruto.ninja
{"message":"User naruto@naruto.ninja deleted successfully"}
</pre>

Enjoy !