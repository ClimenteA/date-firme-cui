FROM ubuntu:latest

WORKDIR /home

COPY ./dist/datecuiapi ./datecuiapi

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates

RUN chmod +x /home/datecuiapi

CMD ["/home/datecuiapi"]
