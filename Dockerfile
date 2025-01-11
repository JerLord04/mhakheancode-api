FROM golang:1.23.4
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd /app/app && go build -o main .
ARG DB_USER
ARG DB_PASS
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG SSL_MODE
ARG TIME_ZONE
ENV DB_USER=$DB_USER
ENV DB_PASS=$DB_PASS
ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV DB_NAME=$DB_NAME
ENV SSL_MODE=$SSL_MODE
ENV TIME_ZONE=$TIME_ZONE
RUN echo "DB_USER=$DB_USER" >> /app/.env && \
    echo "DB_PASS=$DB_PASS" >> /app/.env && \
    echo "DB_HOST=$DB_HOST" >> /app/.env && \
    echo "DB_PORT=$DB_PORT" >> /app/.env && \
    echo "DB_NAME=$DB_NAME" >> /app/.env && \
    echo "SSL_MODE=$SSL_MODE" >> /app/.env && \
    echo "TIME_ZONE=$TIME_ZONE" >> /app/.env
EXPOSE 3000
CMD ["./app/main"]