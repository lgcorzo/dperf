FROM scratch
LABEL maintainer="Sovereign Ecosystem Maintenance <dev@lgcorzo.dev>"

COPY ./dperf /dperf

ENTRYPOINT ["/dperf"]
