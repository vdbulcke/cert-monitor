FROM alpine:latest  
RUN apk --no-cache add ca-certificates
## non privileged user
USER 1111 
# EXPOSE 9000
WORKDIR /app/
COPY cert-monitor /app/cert-monitor

ENTRYPOINT ["/app/cert-monitor", "server" ,"--config",  "/app/config.yaml"]