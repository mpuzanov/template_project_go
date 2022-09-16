# stage 1: build
FROM golang:1.19 as build
LABEL stage=intermediate
WORKDIR /app
COPY . .
RUN make build

# stage 2: scratch
FROM scratch as scratch
COPY --from=build /app/bin/project /bin/project
CMD ["project"]