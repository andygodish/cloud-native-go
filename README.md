# cloud-native-go
Book - Cloud Native Applications with Go
https://github.com/PacktPublishing/Cloud-Native-programming-with-Golang
---

What are examples of stateful applications. A cloud native application should not maintain any state. This is fairly obvious considering the need to build and tear down container running your application on a frequent basis. 

ex: in-memory, or on the file system
---

Applications need to be designed in a way to handle abrupt termination. 
---

Scenario: you need to retrieve an env variable.

os.Getenv("VAR_NAME")
--- 

Logs should be treated as event streams, simple solutions invole writing your log stream to process' standart output stream:

ex: fmt.Println(...)
---

Scenario: how would you craete a subrouter for URLs prefixed with `/events` using gorilla/mux

```
# Create a new router: 

r := mux.NewRouter()

# Create a new subrouter:

eventsrouter := r.PathPrefix("/events").Subrouter()
```
---

`handler` is a Go struct object that belongs to a Go struct type called `eventServiceHandler`.
---



