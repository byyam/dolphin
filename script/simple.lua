-- wrk -t12 -c100 -d30s --script=simple.lua --latency http://127.0.0.1:12000/greeting

wrk.method = "POST"
wrk.body   = "{\"question\":\"j\"}"
wrk.headers["Content-Type"] = "application/json"

request = function()
   uid = math.random(1, 10000000)

   body = string.format("{\"question\":\"%d\"}", uid)
   return wrk.format('POST', '/greeting', nil, body)
end
