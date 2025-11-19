# lensz-server-web
Lens'z backend server using Gin (Go framework) utilizing HTTP and websockets for protocols of communication.

docker run -p 8080:8080 -e DB_HOST=host.docker.internal -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=1111 -e DB_NAME=lenszdb -e SERVER_PORT=8080 -e JWT_SECRET=veryveryconfidentialcode -e ADMIN_EMAIL=deswandy88@gmail.com -e ADMIN_PASSWORD=010325 -e ADMIN_NAME=SuperAdmin lensz-backend