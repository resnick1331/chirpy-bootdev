# chirpy-bootdev

For HTTP Servers in Go course at Boot.dev

01-Assigment

The Go standard library makes it easy to build a simple server. Your task is to build and run a server that binds to localhost:8080 and always responds with a 404 Not Found response.

Steps
Create a new http.ServeMux
Create a new http.Server struct.
Use the new "ServeMux" as the server's handler
Set the .Addr field to ":8080"
Use the server's ListenAndServe method to start the server
Build and run your server (e.g. go build -o out && ./out)
Open http://localhost:8080 in your browser. You should see a 404 error because we haven't connected any handler logic yet. Don't worry, that's what is expected for the tests to pass for now.
