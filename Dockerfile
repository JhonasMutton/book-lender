##Run Image
FROM alpine
COPY bin/application application
COPY .env .env
ENTRYPOINT ["./application"]
