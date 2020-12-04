FROM scratch

COPY cmd/cmd .

EXPOSE 8098

ENTRYPOINT [ "./cmd" ]