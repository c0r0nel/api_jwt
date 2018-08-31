
## API_JWT

## Purpose

This is a small API written in golang that implements JSON Web Token through several external libraries.

- https://github.com/go-chi/chi
- https://github.com/go-chi/jwtauth
- https://github.com/mattn/go-sqlite3


## Usage

Create the sqlite3 database with some testing data in it.
```
$ sqlite3 users.db < users.sql
```

Before running the server first you need to get all the dependencies needed, this might take some minutes. Be patience.
```
$ go get ./...
```

Finally you can run the server doing:

```
$ go run api_jwt.go
```

For testing purpose, you can try to access the API root first:

```
$ curl -v http://localhost:3001/
```
Accessing the restricted endpoint /admin needs that you get authenticated first:
```
$ curl -v http://testuser:supersecret@localhost:3001/auth
```
This will give you a JSON Web Token that will last 3 minutes.
Finally, to use the token do:
```
$ curl -H"Authorization: BEARER YOUeXtrEmlY.LongJSONWEbToken" -v http://localhost:3001/admin
```

