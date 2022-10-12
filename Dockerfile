FROM golang:1.19-alpine as build

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build ./MergeTwoChannelsWithNil/*

FROM scratch as image

COPY --from=build /app/merge_two_channels_with_nil .

CMD ["/merge_two_channels_with_nil"]