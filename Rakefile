#perfomance test
require 'rake'
require 'net/http'
require 'uri'
require 'json'
require 'benchmark'

task :perfomance, [:port] do |task, args|
  Benchmark.bm 10 do |r|
    r.report do
      10000.times do
        uri = URI.parse("http://localhost:#{args.port}")
        json = Net::HTTP.get(uri)

        puts json
      end
    end
  end
end
