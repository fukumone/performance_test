#perfomance test

require 'net/http'
require 'uri'

url = URI.parse('127.0.0.0.1:3000')
res = Net::HTTP.start(url.host, url.port) {|http|
  http.get('/index.html')
}

puts res.body
