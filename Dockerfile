FROM debian

COPY ./api-kubeum /usr/bin/api-kubeum

ENTRYPOINT /usr/bin/api-kubeum