# Very simple REST API server using Docker and python

The server uses python and the FastAPI library to deliver json data. 

## 1. Create the folder structure

We want to use a folder structure where the local directory has a subfolder called "src" where we will place the python file into. 
The needed files will be the Dockerfile, a requirements.txt file and the main.py file in the subfolder.

```
.
+- src/ 
|  |
|  +-- main.py
|
+-- Dockerfile
|
+-- requirements.txt
```

## Building the container

We want the container be build from the commandline using

```
docker build . -t server1
```

naming it server1 with the tag flag '-t'.

## Starting the container 

We want to start the container with

```
docker run -p 8080:8080 server1
```

with exposing the port 8080 inside of the container to the port 8080 on our local machine.
Now, we can test it by opening a web broweser and input the URL http://localhost:8080


       
