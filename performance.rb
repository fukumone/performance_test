#perfomance test

require 'net/http'
require 'uri'
require 'json'
require 'benchmark'

Benchmark.bm 10 do |r|
  r.report do
    100.times do

      uri = URI.parse('http://localhost:3000')
      json = Net::HTTP.get(uri)

      puts json
    end
  end
end
