FROM scratch

COPY cmd/cmd .

EXPOSE 8102

ENTRYPOINT [ "./cmd" ]