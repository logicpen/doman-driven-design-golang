## Domain Driven Design With Golang
* DDD is a development approach where domain is in the core of the application. Domain the actual or real world use case that needs to be modeled. The system follows ddd considered to represent the real world system.

* It's always recommended engineers discuss about the core domain with the business teams to understand domain as well as subdomains which would help engineering teams to model the domain.

* Let's take an example of authentication/authorization service, in this case our core domain would be `User` related data such as `username`, `password`, `emailId` etc that we want to store and use in the application and ultimately in the business.

* For the implementation in this repository, following pattern has been followed,
    * _Domain :-_ It's in the core and independent of all the layers such as transport layer or `Delivery` layer(i.e REST API, grpc etc), persistent layer or `Repository` layer(database), business implementation layer or `UseCase` layer etc.
    * _Repository :-_ It is persistent layer mainly responsible for interacting with database and caching.
    It is dependent on `Domain` to create models from persistent(db, caching etc) system. It doesn't implement business logic itself but provide an interface that is used to implement business logic. It also abstract away the persistent layer implementation from business logic so business logic doesn't have to know which db is used and how it's been interacted. 
    * _UseCase :-_ This layer implements core business logic with the help of `Repository` layer. `UserCase` layer is independent and oblivious to the inner implementation of `Repository` layer.
    * _Delivery :-_ This layer provides an interface for the outside world to interact with the application via exposing REST endpoints, GRPC endpoint or other any other interface that is suitable.