FROM golang:1.22

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# TODO 
RUN go install github.com/spf13/cobra-cli@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0


RUN apt-get update && apt-get install sqlite3 -y

RUN apt-get update && apt-get install -y gcc libc6-dev

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]