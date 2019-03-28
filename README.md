Blog
====

I would like to:

Build a real world "production" REST or gRPC API:
Scalable, must be able to run more than one instance.
Dockerized, runnable on minikube.
Unit tested, must be able to run "go test ./..." directly from clone.
Integration tested, recommend docker-compose.
OpenAPI/Swagger (or similar for gRPC) documented.
dep vendored, but using the standard library often, instead of piling on dependencies.
Authenticated and Authorized via apikeys and/or user accounts.
Built and tested via CI: Travis, CircleCi, etc. Recommend Makefile for task documentation.
Flag & ENV config, API keys, ports, dev mode, etc.
"why" comments, not "what" or "how" which should be clear through func/variable names and godoc comments.
Use of Context to limit request time.
Leveled JSON logging, logrus or similar.
Postgres/MySQL, sqlx or an ORM.
Redis/memcache for scalable caching.
Datadog, New Relic, AppDynamics, etc for monitoring and statistics.
Well documented README.md with separate sections for API user and service developer audiences. Maybe even include graphviz or mermaidJS UML diagrams.
Clean git history with structured commits and useful messages. No merge master commits.
Passing go fmt, go lint, or better, go-metalinter in the CI.
Show me that, or even half of it, on Github/Gitlab with some open source contributions and I'd be arguing to hire you. Note: I'm not saying this is required to be hired.
Edit: Formatting and links.
https://peter.bourgon.org/go-for-industrial-programming/
https://blog.csnet.me/blog/building-a-go-api-grpc-rest-and-openapi-swagger.1/
EditEdit: You do not need all of this just to get a junior level job. Many developers will not get to all of this at first, especially since when needing to balance time/resources vs shipping. Long term, it should all be there for a production system.
