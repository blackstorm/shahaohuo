FROM golang:1.14.2-buster as BUILD
# 注意 GO 进行 build 会将此目录记录到二级制文件中
WORKDIR /opt/shahaohuo/
COPY ./ /opt/shahaohuo/
RUN export GOPROXY=https://goproxy.io \
    && GOOS=linux go build -o server cmd/server/main.go

FROM debian:buster-slim
ARG APP_ENV
WORKDIR /opt/shahaohuo/

COPY --from=BUILD /opt/shahaohuo/server .
COPY --from=BUILD /opt/shahaohuo/configs/server/${APP_ENV:-prod}.yaml ./configs/
COPY --from=BUILD /opt/shahaohuo/web/server/* ./resources/

ENTRYPOINT ["/opt/shahaohuo/server"]