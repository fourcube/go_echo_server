FROM centurylink/ca-certs
WORKDIR /app
COPY go_echo_server /app/
CMD ["/app/go_echo_server"]
