### How to reproduce

Confirm that the cookie is set without Traefik by running the react app in the web-app folder and the go server. You will see that the cookie can be set.

You will need to adjust the following lines:

In the React App:
```javascript
// app.js
fetch("http://api.localhost").then((res=> console.log(res)));
// To:
fetch("http://localhost:8089").then((res=> console.log(res)));
```

On the Go server:

```go
//main.go
panic(http.ListenAndServe("0.0.0.0:8089", r))
```

To run, from the "traefik-min-repro" folder, run ```docker-compose up```, and then ```docker-compose build```.

Go to api.localhost or app.localhost and you will see that no cookies are set.