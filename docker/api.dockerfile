FROM golang:1.21.5

WORKDIR /app

COPY . .

RUN go mod download &&\
    make api

EXPOSE 8030

ENTRYPOINT [ "/app/tmp/api-server" ]
CMD [ "--host", "0.0.0.0", "--port", "8030" ]
