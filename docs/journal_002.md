# REPL Server Development Journal 001
__June 24, 2015 02:31 AM / Los Angeles, California__

## Boilerplate Stuff
Okay, so so far I've set up the basics of a go web server project:

1. router.go declares a struct called replserver.Router, which implements the http.Handler interface. The router is responsible for properly handling requests to the API. I'd like to implement a simple HTTP REST interface, as well as a Websocket interface.
2. The main() function starts an HTTP server using Router as the handler
3. datastore.go handles requests for data and is implemented using redis. So far it just connects to redis on init()



## What's next?
1. Figure out data structures
2. Figure out the API interface
3. Figure out how to structure the code so that both REST and WebSocket interfaces work seamlessly
4. Figure out authentication / crypto
