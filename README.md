# Smart-RealEstate

# Hexagonal architecture

Hexagonal architecture (also known as ports and adapters architecture) is a software architecture based on the idea of separating core business from external concerns.

![image](https://user-images.githubusercontent.com/83599883/232283650-97874249-2ab8-4b0d-b1b5-95ec559f88b4.png)

Core:
In this architecture, everything surrounds the core of the application. The presence of other things outside should be completely ignored. In other words, the main core should not be aware of how the application is presented or where the data is stored.
The kernel can be seen as a "box" capable of solving all business logic independently of the application infrastructure. This approach allows us to test the core in isolation and gives us the ability to change infrastructure components easily.

![image](https://user-images.githubusercontent.com/83599883/232283838-24d15299-1e93-4b53-9841-0ac24c5afbef.png)

Actors:

Actors are real-world things that want to interact with the core. These things can be people, databases, or even other applications.
Actors can be categorized into two groups depending on who is causing the interaction:

![image](https://user-images.githubusercontent.com/83599883/232284085-1320c811-7d6c-4378-9441-a6b1cc2947a4.png)


Drivers are those who create communication with the core. They do this to call a specific service in the kernel.
A human or a CLI are perfect examples of driver actors.

![image](https://user-images.githubusercontent.com/83599883/232284122-2161c6d1-9da1-49d1-a7c3-4ee0c722451b.png)


Driven (or secondary), are those who expect the core to be the one who initiates the relationship. 
In this case, it is the kernel that needs something that the Driver provides, so it sends a request to the Driver and invokes a specific action on it. 
For example, if the kernel needs to store data in a postgresql database, the kernel initiates a connection to execute an INSERT query on the postgresql client.

![image](https://user-images.githubusercontent.com/83599883/232284103-5d4c6c87-cff2-4c20-a770-89fa49de1bad.png)

Ports: 

Note that Actors and Core speak different languages. An external application sends a request over http to make a call to the main service (which doesn't understand what http means).
Another example is when the core (which is technology agnostic) wants to store data in a postgresql database.
Then, there must be "something" that can help us make such translations. This is where ports and adapters come into play.

Ports for driver actors:

![image](https://user-images.githubusercontent.com/83599883/232285280-cacb64e7-8fbf-449c-86ea-027cc90b9130.png)

Ports for driven actors:

![image](https://user-images.githubusercontent.com/83599883/232285387-084d451b-1255-4eea-95b7-69888adc5538.png)

Adapters
On the other hand, we have adapters that are responsible for converting a request from actor to core and vice versa. This is necessary, because as we said before, actors and core "speak" different languages.
An adapter for a port driver converts a specific technology request into a call to a core service.

![image](https://user-images.githubusercontent.com/83599883/232285796-264a5539-d328-44a4-83d5-140273bc8685.png)


Dependency Injection
After completing the implementation, it is necessary to somehow connect the adapters to the corresponding ports.
This can be done when the application starts and allows us to decide which adapter should be connected on each port, this is what we call "dependency injection".
For example, if we want to store data in a postgresql database, we only need to connect an adapter for the postgresql database to the corresponding port.

![image](https://user-images.githubusercontent.com/83599883/232286062-eb2ee42d-d529-436e-b005-bbf10dde9db7.png)


