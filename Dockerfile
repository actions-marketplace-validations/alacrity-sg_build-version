FROM scratch
COPY build-version .
ENV RUN_ARGS ""
RUN chmod +x ,/build-version
CMD [ "./build-version", "${RUN_ARGS}"]
