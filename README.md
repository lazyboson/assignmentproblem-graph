### Introduction - Solving classical Assignment Problem with Hopcroft-Karp Algorithm 
    This Codebase generate random tasks, which are submitted to rabbitmq. Further it also generate set of agents.
    Later it runs consumer on rabbitmq who will be picking up these tasks from queue and making a list.
    Later we built a graph with expression matching of agents and tasks
    Finally, it runs Hopcroft-Karp algorithm on the graph and returns which agent is assigned to which task

### Usages - 
    Run RabbitMQ container using docker-compose script with command  -

    docker-compose up --detach 
    go run main.go

