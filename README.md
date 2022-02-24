# Sample Booking App in Golang

- Built in go 1.16

This project use the following packages:

Routing:
- [pat router](https://github.com/bmizerany/pat)
- [chi router](https://github.com/go-chi/chi) (switched to this one to see how easy it is to replace module
 in go and this one have some interesting middlewares)
Middlewares:
- [nosurf](https://github.com/justinas/nosurf) CSRF protection middleware for Go.
State Management with Sessions:
- [scs](https://github.com/alexedwards/scs) HTTP Session Management for Go


Note: For more options on the routing side of things
      looks this interesting benchmarking
https://github.com/julienschmidt/go-http-routing-benchmark
