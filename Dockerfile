FROM public.ecr.aws/bitnami/golang:latest

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd
RUN go build -buildvcs=false -o /ecs-sample-app

EXPOSE 8080

CMD [ "/ecs-sample-app" ]