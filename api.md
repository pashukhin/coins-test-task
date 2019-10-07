```
$ curl localhost:8080/account
{"accounts":[{"id":1,"name":"Alice","balance":0,"currency":"USD"},{"id":2,"name":"Bob","balance":10,"currency":"USD"},{"id":3,"name":"Boris","balance":1000,"currency":"RUR"},{"id":4,"name":"Eve","balance":1000000,"currency":"USD"},{"id":5,"name":"Natasha","balance":10000,"currency":"RUR"},{"id":6,"name":"Vladimir","balance":1000000,"currency":"RUR"}]}
$ curl localhost:8080/payment
{"payments":null}
$ curl -d '{"from":2,"to":1,"amount":1}' -H "Content-Type: application/json" -X POST localhost:8080/payment
{"payment":{"id":1,"from_id":2,"to_id":1,"created":"2019-10-07T11:33:17.162036Z","amount":1}}
$ curl localhost:8080/payment
{"payments":[{"id":1,"from_id":2,"to_id":1,"created":"2019-10-07T11:33:17.162036Z","amount":1}]}
$ curl localhost:8080/account/1
{"account":{"id":1,"name":"Alice","balance":1,"currency":"USD"}}
$ curl -d '{"from":3,"to":1,"amount":1}' -H "Content-Type: application/json" -X POST localhost:8080/payment
{"error":"only transactions between accounts with the same currencies are allowed"}
```
