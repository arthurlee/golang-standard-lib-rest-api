# How to Create a RESTful API With Only The Golang Standard Library
This is the code for the post [How to Create a RESTful API With Only The Golang Standard Library](https://ryanmccue.ca/how-to-create-restful-api-golang-standard-library/).
You can get the code of author's from [github](https://github.com/rymccue/golang-standard-lib-rest-api).


# main function duties
1. Create a database connection
2. Create a caching connection for authentication
3. Create a mux
4. Load in our controllers
5. Create our routes with the mux and controllers
6. Start the server

# Install go packages
- go get -u github.com/go-redis/redis
- go get github.com/lib/pq

# Run Postgress in Mac
- [Postgress.app](http://postgresapp.com/documentation/)
