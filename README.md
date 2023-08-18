# Getting started 
- To build the binary file run ```make build```
- To just serve the server run ```make server```
- After running the server there should be an API docs server at e.g. http://localhost:8000/swagger/index.html

# Migration
When the server is started, migration is automatically applied, in case you want to manually tinker with its up and down please install this first [migration tool](https://github.com/golang-migrate).

You can run these commands depending on what you need:
- ```make migrate-up``` for incrementally running the sql up scripts in **migrations** folder
- ```make migrate-down``` for incrementally running the sql down scripts in **migrations** folder

# Testing the API(s)
Only possible after running DB migration

<h3>Admin Credentials</h3>
username = "don.juan"<br>
password = "don.juan"

<h3>User Credentials</h3>
username = "jon.doe"<br>
password = "jon.doe"