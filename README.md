# Problem Statement:
## You will receive a string as input, potentially a mixture of upper and lower case, numbers, special characters etc. The task is to determine if the string contains at least one of each letter of the alphabet. Return true if all are found and false if not. Write it as a RESTful web service (no authentication necessary) in any language/framework you choose and document the service. Please describe how you would deploy this application into AWS, including which AWS services you would use, and what deployment method or tools.

I. a) What were alternatives for building an api in Golang, and b) why did i choose Mux over them?:
 a) Other popular alternatives to Mux for building REST APIs in Golang are Gin, Echo, Gorilla/Mux, and net/http.
 b) 
  - Mux is a better choice because it's lightweight and extremely fast.
  - It offers many features such as URL routing, URL parameters, middleware support, and a clean syntax.
  - Additionally, it provides great flexibility when it comes to customizing the API endpoints.
  - The code is also well-documented, making it easier to maintain.

II. Why Go over Python for this specific task?:
 a) I chose Go for building the REST API instead of Python because Go provides excellent performance and scalability long-term. Go is also a great language for developing microservices and distributed systems, which makes it a great choice for writing REST APIs. Additionally, Go is known for its concurrency support, which makes it ideal for building applications that require high throughput. Finally, Go is easy to apply this concept in a scalable approach and has a concise syntax, which makes it easier to develop, read and maintain as a micro-service.

III. Go (mux) vs Python (Flask & Django):
  Golang is a great language to use when building a web service because it is fast, lightweight, and has a simple syntax. It is also well suited to concurrent programming tasks, making it ideal for high-performance web applications. Additionally, Golang has excellent support for the RESTful web service architecture, allowing for easy development of services that adhere to the RESTful principles.

  Flask is better suited for applications that need to be built quickly, with minimal setup. It also allows for rapid prototyping, as the development cycle is much faster than Golang. Flask is also well suited for applications that are not as performance-critical and do not require the same type of concurrency as Golang.

  Django, on the other hand, is well-suited for applications that need to be built quickly and easily, while still offering a robust well-rounded feature set. It is also well-suited for applications that need to scale easily, as it has built-in support for scalability and high-availability architectures. Additionally, Django offers a wide range of additional features, an authentication system, and an administrative interface. The celery library supports asynchronous tasks and combined with the threading library concurrency is achievable in python.

IV. How to tackle the challenge:
  - The algorithim must traverse the string at least once to check if all the letters of the alphabet are present.
  - As a result the algorithim will take an amount of time proportional to the length of the string.
  - The most straightforward algorithm for this challenge is a linear search algorithm. A linear search can be used to traverse the string and check if all letters of the alphabet are present. I am choosing the linear search algorithm due to it being relatively simple and efficient in terms of time complexity (O(n)).
  - An alternative algorithm that could be used to solve this problem is a hash table. This data structure stores the letters of the alphabet as keys and checks if each key is present in the string. The time complexity of this algorithm would be O(1).
  - Another alternative algorithm is a trie. This data structure stores the letters of the alphabet as nodes and checks if the string contains all of them. The time complexity of a trie algorithm would also be O(1).

Deployment into AWS:

 - To deploy this application into AWS, the following services could be used:
  - Amazon Elastic Compute Cloud (EC2) for hosting the application.
  - Amazon Elastic Container Service (ECS) for containerizing and deploying the application (local containerization testing with Docker and Kubernetes)
  - Amazon Relational Database Service (RDS) or Storage Buckets (S3) for a more advanced configuration.
  - Amazon API Gateway, To create a REST API for the application at a highly scalable system to connect to.
  - AWS Cloud Formation 
 - Ideal Tools via AWS would be:
  - CloudFormation to create a template that defines the environment you need and then deploy a stack of resources with a few clicks.
  - CodeDeploy to deploy an automation process by deploying code to Amazon EC2 instances, on-premises instances, or serverless Lambda functions.
  - CodePipeline to automate the continuous delivery process to define a CI/CD (continuous integration and continuous delivery) pipeline that automates the build, test, and deployment of the microservice.

Differnent time complexities:
  - O(n): is a linear time complexity, meaning that its performance is directly proportional to the size of the input. It means that the time taken to complete a task increases linearly with the input size.
  - O(n2): is a quadratic time complexity, meaning that its performance is directly proportional to the square of the size of the input. It means that the time taken to complete a task increases exponentially with the input size.
  - O(log n): Logarithmic time complexity, meaning that its performance increases at a logarithmic rate with the increase in input size.
  - O(1): Constant time complexity, meaning that its performance does not depend on the size of the input and is always the same no matter the size.
  - O(n log n): Linearithmic time complexity, meaning that its performance is directly proportional to the product of the input size and the logarithm of the input size.
  - O(2n): Exponential time complexity, meaning that its performance increases exponentially with the increase in input size.

aws team

TESTTT
aws
send finished
