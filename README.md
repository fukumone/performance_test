# performance_test

### Go
- beego => bee run beego_sample
- gin => go run server.go
- goji => go run server.go
- martini => go run server.go
- revel => revel run github.com/myaccount/revel_sample

### Ruby
- Sinatra => ruby server.rb
- Ruby On Rails => rails s

### python
- flask => python server.py


### test command

```
go run http_request -port="port number"
# example => go run http_request.go -port=3000

rake perfomance["port number"]
# example => rake perfomance["3000"]
```
