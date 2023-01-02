# GO Web-Service with JWT

---

Atom Api CRUD Service with User Signup, Login and Roles Authentication

---

Start service with: 

```go run .```

Package used:

* HTTP server: GIN -> https://github.com/gin-gonic/gin
* Development compilation and run: CompileDaemon -> https://github.com/githubnemo/CompileDaemon
* Storing configuration in the environment: GODOTENV -> https://github.com/joho/godotenv
* Generate uuid string -> https://github.com/satori/go.uuid
* JWT manipulating -> https://github.com/golang-jwt/jwt
* Hashing password: bcryp -> https://pkg.go.dev/golang.org/x/crypto


Start CompileDaemon with command: 

```$GOPATH/bin/CompileDaemon -command="./go-web-jwt"```

Find GOPATH with: ```go env GOPATH```

In my case: 

```/home/milan/go/bin/CompileDaemon -command="./go-web-jwt"```