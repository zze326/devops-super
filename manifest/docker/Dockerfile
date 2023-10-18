FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
COPY resource                $WORKDIR/
COPY manifest/config/config-tpl.yaml        $WORKDIR/config/config-tpl.yaml
COPY manifest/docker/entrypoint.sh        $WORKDIR/entrypoint.sh
COPY ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main
RUN chmod +x $WORKDIR/entrypoint.sh


###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
ENTRYPOINT ["./entrypoint.sh"]
