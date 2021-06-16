# bookrepositorygo
 **Running service command on local** 

`go run /cmd/main.go`

**Command to get all books**

`curl http://localhost:9090 && echo`

**Command to add a new book**

`curl http://localhost:9090 -XPOST -d '{"id": 3, "name": "Animal Farm", "writer": "George Orwell", "type": "Classics", "description": "A farm is taken over by its overworked, mistreated animals."}
`

**Command to update existing book with id**

`curl -v localhost:9090/2 -XPUT -d '{"description": "Discovered in the attic in which she spent the last years of her life. Reminder of the horrors of war and an eloquent testament to the human spirit."}'`