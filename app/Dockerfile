#

FROM alpine:latest
# FROM scratch
WORKDIR /usr/src/app
COPY bin .

ENV TZ=Asia/Seoul
RUN apk --no-cache add tzdata && \
        cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
        echo $TZ > /etc/timezone
EXPOSE 8080
CMD ["./mango-app"]