FROM scratch
COPY build-version .
ENV RUN_ARGS="-repo-path=."
CMD [ "./build-version", "${RUN_ARGS}"]
