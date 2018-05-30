### NOUGHTS AND CROSSES

This is a naive implementation of noughts and crosses for 2 players
on the console. It's a primitive first pass using simple sums and loops.

Porting this online to a online application supporting multiple concurrent games could involve:

* A websocket driven front end talking to a back end driven by a Go service making use of web sockets (most probably using [gorilla websocket](https://github.com/gorilla/websocket) and potentially using 
a real time database such as [rethink](https://www.rethinkdb.com/) for
persistence.
* A cloud solution backed by AWS Lambda (for game logic) and DynamoDB (for storage). This would be a cost-effective solution, especially on
the Lambda side where the first million requests per month are free.

Deployment into the cloud could either make use of Docker and Docker Compose to encapsulate the required services for deployment as a multi-container application via ECS, or potentially the Go program could just be compiled and shipped direct to AWS. 

