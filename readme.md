# REST API GOLANG
This is an application built with golang, jwt, gorm, postgresql.

route list :
- / | Method "GET" | Homepage
- /login | Method "POST" | Halaman Login
Json data : {
	"Username" : "xxxxx",
	"Password" : "xxxx"
}
- /register | Method "POST" | Halaman Register
Json data : {
	"Username" : "xxxx",
	"Password" : "xxxx",
	"NamaLengkap" : "xxxx"
}
- /users | Method "GET" | Halaman get all user data
- /users/{id} | Method "GET" | Halaman get user with spesific id
- /users/{id} | Method "PUT" | Halaman update user
Json data : {
	"Username" : "xxxx",
	"Password" : "xxxx",
	"NamaLengkap" : "xxxx"
}
- /users/foto/{id} | Method "PUT" | Halaman update foto user
Json data : {
    "Foto" : "xxxx"
}
- /users/{id} | Method "DELETE" | Halaman delete user