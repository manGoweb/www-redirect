FROM alpine:3.5

WORKDIR /app

COPY main main
ENV FORMAT "https://www.%s%s"
ENV PORT 80
EXPOSE 80

CMD ./main
