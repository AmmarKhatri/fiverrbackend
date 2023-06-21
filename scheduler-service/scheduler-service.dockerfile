FROM alpine:latest

RUN mkdir /app

COPY schedulerApp /app

CMD [ "/app/schedulerApp"]