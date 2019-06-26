# crud-redis-go


API
--------

```
POST localhost:3000/students?department=CS&code=101&section=1
```

```
GET localhost:3000/students?department=CS&code=101&section=1
```

```
DELETE localhost:3000/students/example?department=CS&code=101&section=1
```

Request Body Structure for Post Request

```json
  {"id":"example"}
```


Testing
--------
POST and DELETE requests can be tested by using curl commands

- Example Post Request

```bash
$ curl -i -X POST -H 'Content-Type: application/json' -d '{"id":"41301749"}' "http://localhost:3000/students?department=CS&code=101&section=1"
```

- Example Delete Request

```bash
$ curl -i -X DELETE -H 'Content-Type: application/json' "http://localhost:3000/students/41301749?department=CS&code=101&section=1"
```
