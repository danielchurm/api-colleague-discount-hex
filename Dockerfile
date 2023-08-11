# syntax=docker/dockerfile:1.0.0-experimental

# checkov:skip=CKV_DOCKER_9:apt setup by infra
# checkov:skip=CKV_DOCKER_2:most containers are not servers so dont have a healthcheck

# BUILD STAGE - Stage for building the app

FROM public.ecr.aws/docker/library/golang:1.20-bullseye as build_app

RUN mkdir -p -m 0600 /root/.ssh && ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /smartshop/
COPY . /smartshop

RUN --mount=type=ssh make tools
RUN --mount=type=ssh make deps

RUN make build

# TEST STAGE - Stage for testing, includes dev dependencies
#
FROM build_app as test


# BUILD BASE IMAGE STAGE - Stage for building base image runtime
#
FROM public.ecr.aws/ubuntu/ubuntu:20.04_stable as build_base_image
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN apt update -yq && \
  DEBIAN_FRONTEND=noninteractive apt install --no-install-recommends -yq \
    make \
#TODO Uncomment if using a DB
#    postgresql-client \
    ca-certificates && \
  apt remove -y --purge python* && \
  apt autoremove -y --purge  && \
  apt clean -y  && \
  apt autoclean -y && \
  rm -rf /var/lib/apt/lists/* && \
  groupadd -g 999 smartshop && \
  useradd -r -u 999 -g smartshop smartshop


# RUNTIME STAGE - Stage for running the service
#
FROM build_base_image as runtime

WORKDIR /smartshop/

COPY --from=build_app --chown=smartshop:smartshop /smartshop/Makefile /smartshop/
COPY --from=build_app --chown=smartshop:smartshop /smartshop/migrate /smartshop/
COPY --from=build_app --chown=smartshop:smartshop /smartshop/migrations /smartshop/migrations
COPY --from=build_app --chown=smartshop:smartshop /smartshop/smartshop-service /smartshop/
COPY --from=build_app --chown=smartshop:smartshop /smartshop/docker/docker-entrypoint.sh /smartshop/
#TODO Uncomment if using a DB
# COPY --from=build_app --chown=smartshop:smartshop /smartshop/docker/docker-entrypoint-migrations.sh /smartshop/
# COPY --from=build_app --chown=smartshop:smartshop /smartshop/smartshop-services-tools/bin/wait-for-postgres.sh /smartshop/

USER smartshop

EXPOSE 8080

ENTRYPOINT ["./docker-entrypoint.sh"]
