FROM golang:alpine as builder
LABEL maintainer="bennie.anware@gmail.com"

#ARG BUILD_TOKEN
#ARG BUILD_USER
#ARG BUILD_MACHINE
#ARG CI_PROJECT_NAME


ENV GIT_TERMINAL_PROMPT=1
ENV GOBIN /go/bin
ENV GOPATH /app
ENV PATH=$GOPATH/bin:$PATH

#ENV build_token=$BUILD_TOKEN
#ENV build_user=$BUILD_USER
#ENV build_machine=$BUILD_MACHINE
#ENV ci_project_name=$CI_PROJECT_NAME

RUN apk update && \
apk upgrade &&\
apk add --no-cache git ca-certificates tzdata && \
cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
echo "Asia/Jakarta" > /etc/timezone

RUN mkdir -p /app/src/user-svc
ADD . /app/user-svc
WORKDIR /app/user-svc


#RUN echo "machine $build_machine login $build_user password  $build_token" > ~/.netrc && \
#go get . && \

RUN go get . && \
go mod tidy && \
go mod download && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o user-svc

RUN ls -al /app/user-svc/resources

FROM scratch

COPY --from=builder /app/user-svc/resources    /app/user-svc/resources
COPY --from=builder /app/user-svc/user-svc /app/user-svc/user-svc
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

WORKDIR /app/user-svc

EXPOSE 3000
EXPOSE 3030

CMD ["/app/user-svc/user-svc", "rest"]
CMD ["/app/user-svc/user-svc", "grpc"]

