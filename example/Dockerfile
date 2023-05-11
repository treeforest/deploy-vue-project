FROM busybox:latest

COPY ["./server", "/"]

VOLUME ["/dist"]
EXPOSE 80

CMD ["./server"]