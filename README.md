# CSL2090 Course Project:

## Team Members
- Jyotin Goel (B22AI063)

## Installation instructions for CLI tool

```bash
tar zxvf CSL2090_PCS2_CourseProject_ollama_Linux_x86_64.tar.gz
sudo sh install.sh
```

## Instructions to use CLI tool

```bash
jpt setserver <host>:<port> <model>
```

# Load Balancer Setup Instructions


## Step 1: Start the ollama instances


Here we take a sample of 3
```bash
docker run -d -v ollama:/root/.ollama -p 3030:11434 --name ollama1 ollama/ollama
docker run -d -v ollama:/root/.ollama -p 11432:11434 --name ollama2 ollama/ollama
docker run -d -v ollama:/root/.ollama -p 11433:11434 --name ollama3 ollama/ollama
```

## (Optional)
```bash
docker exec -it <container_id> bash
#/ ollama pull gemma:2b
```

## Step 2: Build the nginx image

```bash
cd nginx_app/
docker build -t loadbalancer .
```

## Step 3: Run the nginx image

```bash
docker run -p 8080:8080 --network="host" loadbalancer
```


## Concepts used:

- Docker Container

- Least number of active connections load balancing

- Round Robin Load balancing


## Step 4: Run the jpt client:

```bash
jpt setserver 127.0.0.1:8080 gemma:2b
```

### In case you want to access another person's hosted llms

```bash
jpt setserver <hosted_ip>:8080 gemma:2b
```

