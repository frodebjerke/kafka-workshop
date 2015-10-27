# Solutions

## Challenge 1

Write a message 'hello' to topic ex1
````
echo hello | kafkacat -P -b 192.168.99.100 -t ex1
````

Read messages from topic ex1
````
kafkacat -b 192.168.99.100 -C -t ex1  
````

## Challenge 2

See topics
````
kafkacat -b 192.168.99.100 -L
````

Get last two from starwars topic
````
kafkacat -b 192.168.99.100 -C -t starwars -o -2
````

Get first from starwars topic
````
kafkacat -b 192.168.99.100 -C -t starwars -c 1 -o 0
````

## Challenge 3

Write 'hello' to partition 0 of topic ex3
````
echo hello | kafkacat -b 192.168.99.100 -P -t ex3 -p 0
````

Write 'world' to partition 1 of topic ex3
````
echo world | kafkacat -b 192.168.99.100 -P -t ex3 -p 1
````

Read all messages on the topic ex3 without specifying partition
````
kafkacat -b 192.168.99.100 -C -t ex3
````

Read all messages in partition 1 on the topic ex3
````
kafkacat -b 192.168.99.100 -C -t ex3 -p 1  
````
