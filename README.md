
api_jwt

## Purpose

This a small API written in golang that implements JSON Web Token through several external libraries.


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

For testing from a client you could first try to access the root of the api:

```
$ curl -v http://localhost:3001/
```
For accessing the restricted endpoint /admin you need to authenticate first:
```
$ curl -v http://testuser:supersecret@localhost:3001/auth
```
This will give you a Token that will last 3 minutes.
Finally to use the token you could do:
```
$ curl -H"Authorization: BEARER YOUeXtrEmlY.LongJSONWEbToken" -v http://testuser:supersecret@localhost:3001/admin
```

