#Api docs

Documentation for http api methods.

##Account list

Available on `/account`.
Returns list of all accounts in database.

```
$ curl localhost:8080/account
{"accounts":[{"id":1,"name":"Alice","balance":0,"currency":"USD"},{"id":2,"name":"Bob","balance":10,"currency":"USD"},{"id":3,"name":"Boris","balance":1000,"currency":"RUR"}]}
```

##Send

Available on `/send`.
Makes money transfer between given accounts on given amount.
On success returns newly created payment.

Returns non-empty error if:
- one of accounts not exists
- balance of source account less than given amount
- account's currencies is different
- given amount less or equal to zero

Success example:
```
$ curl -d '{"from":2,"to":1,"amount":1}' -H "Content-Type: application/json" -X POST localhost:8080/payment
{"payment":{"id":1,"from_id":2,"to_id":1,"created":"2019-10-07T11:33:17.162036Z","amount":1}}
```
Error example:
```
$ curl -d '{"from":3,"to":1,"amount":1}' -H "Content-Type: application/json" -X POST localhost:8080/payment
{"error":"only transactions between accounts with the same currencies are allowed"}
```

##Payment list

Available on `/payment`.
Returns list of all payments in database.
```
$ curl localhost:8080/payment
{"payments":[{"id":1,"from_id":2,"to_id":1,"created":"2019-10-07T11:33:17.162036Z","amount":1}]}
```

##Account by id

Available on `/account/{id}`.
Returns single account from database.

If account not found, returns error.

Success example:
```
$ curl localhost:8080/account/1
{"account":{"id":1,"name":"Alice","balance":1,"currency":"USD"}}
```
Error example:
```
$ curl localhost:8080/account/77
{"error":"sql: no rows in result set"}

```
