### Magic Soup

This repo started [here](https://github.com/michaeldebarros/go-routing), in a simple attempt to implement routing in Go.  I used Julien Schmidt's [Httprouter](https://github.com/julienschmidt/httprouter) for routing. 

After that, I decided to build a simple CRUD app and in the process transfer some Node.js skills into Go. So I added Mongo DB, validation, sessions, middleware, etc. 

The ORM to connect do MongoDB id [mgo](https://godoc.org/labix.org/v2/mgo). 

There is a login wall that is implemented via middleware. Also, session management is done via a map (Go's native implementation of maps) of valid sassions to user id. This lets we get the user Id in every request.


There is a documenting branch also, which contains the documentation that works with go doc command.

I did a little walk through the app and recorded. The videos are as follows:

1) [Intro](https://www.youtube.com/watch?v=ZkjOEUYEYBI)

[![THUMBNAIL](https://img.youtube.com/vi/ZkjOEUYEYBI/mqdefault.jpg)](https://www.youtube.com/watch?v=ZkjOEUYEYBI)

2) [Init Function](https://www.youtube.com/watch?v=pf2VQChWiC0&t=79s)

[![THUMBNAIL](https://img.youtube.com/vi/pf2VQChWiC0/mqdefault.jpg)](https://www.youtube.com/watch?v=pf2VQChWiC0)

3) [Routing](https://youtu.be/Egzt5Bnl414)

[![THUMBNAIL](https://img.youtube.com/vi/Egzt5Bnl414/mqdefault.jpg)](https://youtu.be/Egzt5Bnl414)

4) [Mongo DB](https://youtu.be/Hd7XAf7D2JA)

[![THUMBNAIL](https://img.youtube.com/vi/Hd7XAf7D2JA/mqdefault.jpg)](https://youtu.be/Hd7XAf7D2JA)

5) [Packages](https://youtu.be/EaaEijKiu54)

[![THUMBNAIL](https://img.youtube.com/vi/EaaEijKiu54/mqdefault.jpg)](https://youtu.be/EaaEijKiu54)

6) [Templates](https://youtu.be/s1dyiXoSC-E)

[![THUMBNAIL](https://img.youtube.com/vi/s1dyiXoSC-E/mqdefault.jpg)](https://youtu.be/s1dyiXoSC-E)



