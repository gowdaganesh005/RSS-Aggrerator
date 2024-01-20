This is the docs that literraly explain the topics mentioned in detail to the comments mentioned in the actual code use the reference points "1(1)" to the actual code

[1(1)] [PORT] - In the context of network programming and web servers, a "port" refers to a logical endpoint for communication. It is a numeric identifier that helps direct network traffic to the appropriate application or service running on a computer.

A port is identified by a numerical value called a "port number." Port numbers range from 0 to 65535. 
When a server application starts, it "binds" to a specific port, indicating that it is ready to receive incoming network traffic on that port. Clients (such as web browsers) use the port number to address their requests to the appropriate service on the server.

One computer can host multiple services, and each service can be associated with a different port number. This allows for multiplexing, where a single machine handles various types of network communication simultaneously.

The combination of a protocol and a port number defines a unique communication endpoint. For instance, HTTP typically uses port 80, and the combination "HTTP on port 80" signifies a specific type of communication.

but when we search in search engine we never mention port we just write google.com so does it work?
>>>When you enter a domain like "google.com" into a web browser, you're using the default port for HTTP, which is port 80. Most web browsers automatically assume port 80 if you don't explicitly specify a port in the URL. For example, "http://google.com" is equivalent to "http://google.com:80".

In the provided Go program, specifying the port is essential because it determines the communication endpoint for the HTTP server. The port number is crucial for the server to listen for incoming requests. If you skip specifying the port, the program won't know on which port to listen, and as a result, it won't be able to receive and handle incoming HTTP requests.


[1(2)] [Router] In the context of web development, a "route" refers to a mechanism that associates a URL pattern with a specific function or handler. Routes define how an application should respond to different URLs or URIs (Uniform Resource Identifiers). The primary purpose of routes is to map incoming requests to the appropriate code that should handle those requests.

URL Mapping:

    Routes establish a mapping between specific URLs and the code that should be executed when a request is made to that URL. Each route typically corresponds to a specific resource or functionality in your web application.

[1(3)] Handler Functions: [2(1)]

    When a request matches a defined route, a corresponding function, often called a "handler function," is executed. This function is responsible for processing the request, generating a response, and possibly interacting with data or other parts of the application.

HTTP Methods:

    Routes can be associated with specific HTTP methods (GET, POST, PUT, DELETE, etc.). This means that different code can be executed based on the type of request being made. For example, a route might handle GET requests for retrieving data and POST requests for creating new data.


[1(4)]     &http.Server:

        & is the address-of operator in Go. It is used to obtain the memory address of a variable. In this case, it's used to create a pointer to an instance of the http.Server struct.

        http.Server is a struct provided by the Go standard library to configure and run an HTTP server. It has various fields to set up the server, such as Handler, Addr, ReadTimeout, WriteTimeout, etc.

    Handler: router:
        Handler is a field of the http.Server struct, and it specifies the handler to use for processing incoming requests. In this case, the router variable (an instance of chi.Router) is set as the handler. This means that the chi router will handle incoming HTTP requests and direct them to the appropriate routes.

    Addr: ":" + portString:

        Addr is another field of the http.Server struct, and it specifies the network address the server should listen on. It's typically in the form host:port.

        ":" + portString constructs the full network address by concatenating a colon (:) and the portString retrieved from the environment variables. For example, if the port is "8080," the address becomes ":8080," indicating that the server should listen on port 8080.



[Database]

[2(1)] [Goose] -  Goose is a popular tool for DB migrations in Go. It simplifies the process of creating and managing database schema changes by providing a structured way to define migration scripts.
In Goose, DB migrations are a process of updating the database schema by creating new tables, modifying existing ones, or dropping them. Migrations are necessary because they help ensure that the database schema is consistent with the application's codebase, and they provide a way to track changes made to the database structure over time.

[2(1)] [2(2)]  The `up` migration creates new tables or modifies existing ones in the database, while the `down` migration reverses those changes, restoring the database to a previous state. Keeping both `up` and `down` migrations is necessary because it allows for easy rollback of changes in case something goes wrong during deployment or if the application's requirements change.