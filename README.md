# Go Notebook

Go demo with gorilla/mux, sqlite and dependency injection.

Typescript/React frontend.

Deployed to Digitalocean with Docker.

# Curl

```
curl -X POST http://localhost:9090/notes -d '{"contents":"Read War and Peace."}'

curl -X PUT http://localhost:9090/notes -d '{"contents":"Read War and Peace again."}'

curl -X DELETE http://localhost:9090/notes -d '{"id":1}'
```
