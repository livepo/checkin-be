# 根据项目名称调整APP名称，建议和repo名称保持一致
ARG APP=checkin-be

# 运行阶段
FROM alpine:latest AS final

ARG APP
ENV APP=${APP}
ENV WORKDIR=/data
WORKDIR ${WORKDIR}


CMD ["sh", "-c", "./${APP}"]
