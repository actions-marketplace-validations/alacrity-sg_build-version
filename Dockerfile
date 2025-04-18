FROM scratch
COPY build-version .
ENV RUN_ARGS ""
CMD [ "./build-version", "${RUN_ARGS}"]
