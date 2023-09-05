# Build stage
FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# Final minimal stage 
FROM scratch
WORKDIR /app
COPY --from=build /app/main /main
COPY ./public /app/public
EXPOSE 3000
ENTRYPOINT ["/main"]