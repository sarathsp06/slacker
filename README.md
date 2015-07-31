###Slacker

[Slacker](https://github.com/sarathsp06/slacker) is a slack message backup service that pushes the slack messages from once channel to elastic search cluster and make it available for search

###What slacker is all about ?

 - Slack message backup
 - Full text search on messages
 - easy access to  message and scale


 ###To Contribute
  - Install elastic search which is used for backing up db
    - install docker
    - docker pull elasticsearch
  - while running for testing specify the ip and port as environment variables
       - ELASTIC_IP `Eg:[:9200]`
       - ELASTIC_PORT `Eg:http:127.0.0.1 `
