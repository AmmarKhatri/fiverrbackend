FROM alpine:latest

RUN mkdir /app

COPY catalogApp /app

CMD [ "/app/catalogApp"]