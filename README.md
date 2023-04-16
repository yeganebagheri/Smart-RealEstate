# Smart-RealEstate

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
