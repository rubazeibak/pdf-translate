FROM scratch

COPY ./bin/fb-events-locator /

ENTRYPOINT ["/fb-events-locator"]
