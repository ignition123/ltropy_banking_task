# Eltropy Banking Task
Task Given By Eltropy For 1st Round Of Interview

Config for database mongodb + Redis
Mongodb":"mongodb://localhost:27017/?replicaSet=replicaSet",
"Redis":"localhost:6379",
"MongodbDatabase":"coynce"
HttpPort : 3500

This Project is not exactly the requirement you have asked but quite similar. Developing a crypto exchange for someone and it has all such features even more. Kindly look into the code.
It is working in my personal Golang Framework called Haste its fast and has many features like http2 streaming, RPC and middlewares etc. Same can be created using Golang Gin a lightweight open source framework with out much changes

Code are in the src folder

  -> cns : Contains all code retated to framework
  -> controllers: Contains all routes/apis
  -> lib: Contains misc methods like handling panics and writing logs etc.
  -> messages: contains all messages related to http response
  -> middlewares: contains session validation with redis
  -> models: contains all mongodb queries
  -> objects: contains all structure for request, response and parsing of data
  -> RequestValidation: contains all validation rules checking the incomming http request
  -> Web: Contains all routes list
  -> app.go is the init file

It contains all apis related to admin, super admin, permissions, updating profile, signin, redistration, ressetting password, ApproveKYC, blocking userId, Created New BrokerId, Creating new broker id for broadcasting and interactive order placement it is developed on the basis of crypto exchange.

All apis are json based request and response. 
Content-Type: application/json

All apis where session is required extra header must be used.
Authorization : "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1ZGVlcEBwb3VuemUuY29tIiwiZXhwIjoxNjEwMDE2NDI1fQ.9MazLsLcEACntdZIM31H3SMYI13SYhowzP_Ut3aMclI"

It contains all admin roles. Please find the request and response sample with postman collection.

It will have realtime pricefeed push to brokers and order placement in future cant share the entire code for now. The streaming is developed over Brahmaputra servers which is again an open source log streammer like kafka. It has a outgoing bandwidth of 1.5Gbps with 16core CPU with multiple topics and replication. Currently suppored for golang only.
Brahmaputra URL: 
Server / Client: https://github.com/ignition123/brahmaputra
Haste Framework: https://github.com/pounze/go_haste

After connecting to redis and mongodb click the go_build.exe for window platform. It is horizontal scalable using any Load balancer like HA Proxy or Nginx.
