FROM golang:1.14.2-alpine as BUILD
# 注意 GO 进行 build 会将此目录记录到二级制文件中
WORKDIR /opt/shahaohuo/
COPY ./ /opt/shahaohuo/
RUN GOOS=linux go build -o shahaohuo cmd/shahaohuo/shahaohuo.go

FROM alpine:3.11
ARG APP_ENV
WORKDIR /opt/shahaohuo/

RUN apk add --no-cache ca-certificates tzdata

COPY --from=BUILD /opt/shahaohuo/shahaohuo .
COPY --from=BUILD /opt/shahaohuo/configs/server/${APP_ENV:-prod}.yaml ./configs/
COPY --from=BUILD /opt/shahaohuo/web/server/ ./resources/

ENTRYPOINT ["/opt/shahaohuo/shahaohuo"]