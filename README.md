## Getting Started

### Domain-Driven Design (DDD) Folder Structure & Wiki

This project follows the principles of Domain-Driven Design (DDD) and uses the following folder structure to organize the codebase:
````
- domain
    - entities
            - entity1.go
            - entity2.go
            - ...
    - value_objects
        - value_object1.go
        - value_object2.go
        - ...
- application
    - usecase
        - usecase1.go
        - usecase2.go
        - ...
    - services
        - service1.go
        - service2.go
        - ...
- infrastructure
    - persistence
        - repository
            - repository1.go
            - repository2.go
            - ...
        - db
            - db.go
            - migration
                - migration1.go
                - migration2.go
                - ...
    - messaging
        - kafka
            - consumer.go
            - producer.go
    - external
        - http
            - http_client.go
            - http_server.go
- cmd 
    - main.go
````

- The domain folder contains all the business logic of the application. It is divided into two subfolders: model and services.

- The model folder contains all the entities and value objects that make up the application's domain model.
- Entities are objects that have an identity and encapsulate the business rules of the application.
    - > In Domain-Driven Design (DDD), entities can have business logic, as they represent business concepts or objects and it makes sense that they have their own behavior and logic. Entities are responsible for maintaining the state of the domain and encapsulating the data and methods that are specific to that business concept.
    - > Entities can have methods that implement business rules, validation, or other business-specific logic. For example, an Order entity may have methods for calculating the total cost of the order, checking for discounts, or applying taxes.
    - > However, it's important to note that entities should only contain logic that is directly related to the state of the entity. Business logic that is not directly related to the state of the entity should be handled by services. This helps to separate the concerns of the entities and services, making the system more maintainable and flexible.
    - > So, in short, entities can have business logic, but it should be logic that is directly related to the state of the entity and should not contain any logic that is not directly related to the state of the entity, that should be handled by services.
- Value objects are objects that don't have an identity and are used to represent simple values like money or date.
- The services folder contains domain services that are used to perform complex business logic that doesn't belong to any specific entity or value object.
  - >**Services** are used to encapsulate business logic that does not fit into an entity or value object. 
    > 
    > Services are typically stateless and used for tasks that don't have a lifecycle, like calculating a value or providing information. 
    > 
    > Services are usually stateless, and their methods are usually more procedural than object-oriented. They can also be used to encapsulate cross-cutting concerns, like logging or security. Services can be used by multiple usecases.
- The infrastructure folder contains all the technical logic of the application. It is divided into three subfolders: persistence, messaging, and external.
    - > The **persistence** folder contains the code responsible for storing and retrieving data from a database.
    - > The **messaging** folder contains the code responsible for sending and receiving messages through message queues.
    - > The **external** folder contains the code responsible for communicating with external systems, such as APIs.
- The **application** folder contains the code responsible for orchestrating the usecases of the application.

- The **usecase** folder contains the code for the usecases of the application, which are use to orchestrate the services and entities to accomplish specific tasks.
  - > **Use cases** represent the business logic that handles a specific business goal or objective.
    >
    > These are the interactions between the application and the user. They are used to handle the input from the external services, translate it to the domain model, and handle the output that is sent to external services.
    >
    > They are also responsible for enforcing the business rules and orchestrating the interaction between entities and value objects.
  - > In Domain-Driven Design (DDD), it is generally recommended that use cases (also known as application services) do not access entities directly. Instead, use cases should interact with entities through services.
    >
    > The idea behind this is to separate the concerns of the `use cases` and `entities`, so that changes to the use cases do not affect the entities, and vice versa. By having services as a layer between the use cases and entities, it becomes easier to make changes to the system without affecting the other parts of the system.
    >
    > Services provide a way to encapsulate complex logic and behavior that is not specific to any particular entity. They are responsible for implementing the use cases or business processes of the system and can use one or more entities to perform their operations. Services can also interact with other services to coordinate the overall behavior of the system.
    > 
    > Additionally, use case should not have knowledge about how the entities are stored or retrieved, so it's better to have services handle this type of operation, so the use case can focus on the business process and the service take care of the technical details.
    > 
    > It's worth noting that there are some exceptions, for example, if the use case is a simple CRUD operation, it can be fine to access the entities directly, but in general, it's a good practice to have services as a layer between the use cases and entities.  
     
    
- The cmd/root.go file is the entry point of the application.

This folder structure allows for a clear separation of concerns and makes it easy to understand the different responsibilities of each part of the codebase.

### Dependency Injection
- This project uses dependency injection to manage the dependencies between different parts of the codebase. 
This makes it easy to swap out the implementation of different interfaces with mocks or different implementations during testing and makes it easier to change the implementation of a specific feature.

### Database Setup

#### Docker

- create migration file
```bash
docker compose --profile tools run create-migration
```

<p>
Because we've mounted the <span style="color: darkseagreen; "> <b>/tmp/migrations</b> </span> folder in the container to the local project folder 
<span style="color: darkseagreen;"><b>/migrations</b></span>, we can see two new files appear in migrations folder. 
One for our up migration, and one for down
</p>

- run migration
```bash
docker compose --profile tools run migrate
```

#### Manual Installation

Install migrate CLI tools to manage database versioning

Linux:

```bash
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate
```

Mac:

```bash
brew install golang-migrate
```

Run migration script


```bash
migrate -database "postgresql://username:password@localhost:5432/test?sslmode=disable" -path migrations/ up
