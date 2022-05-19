FROM golang as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -buildvcs=false -o k8s-app

FROM scratch
COPY --from=build /app/cmd/k8s-app /
ENV GIN_MODE=release
USER 1001
EXPOSE 8080
ENTRYPOINT [ "/k8s-app" ]