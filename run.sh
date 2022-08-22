docker stop meetin-auth
docker rm meetin-auth
docker build -t meetin-auth:latest .
docker run -d -p 0.0.0.0:8080:8080/tcp --name meetin-auth meetin-auth:latest