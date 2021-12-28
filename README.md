# Docker memory issue reproduction

Running application inside 4 docker containers optaineed from same image creates same values instead of unique values when generating randdom.

The same application with different instances run correct when trying in bare metal, without containerization.

Tested this application developed with PHP, Node.js, Golang. 

To reproduce this issue, just run:

```
git clone https://github.com/duktig-dev/docker-memory-issue-reproduction.git
docker-compose up -d
```

Containers will stop running automatically after 2500 values generated each inside `./log` directlry.

Each application appends logs data into own file. 
 
You can analyze created values duplication with running:

```
php checker.php
```


