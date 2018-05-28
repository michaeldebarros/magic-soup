### Golang Routing

The idea this repo is to understand routing in the simplest way possible.  When comming from the node.js world there is Express to meke web servers, which is very flexible in terms of routing and specially easy in terms of building a milddlewre pipeline.

In Golang, as well as in Node.js, there is a native package responsible for working with web connections.  Golang's standard library, despite beeing very complete, could be more efficient in terms of routing.  That's where Julien Schmidt's [Httprouter](https://github.com/julienschmidt/httprouter) comes in place "github.com/julienschmidt/httprouter".  

In the example in this repo I am doing all the routing through the httprouter, even staitc files.