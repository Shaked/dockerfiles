ARG APP_USER=python
ARG APP_HOME=/home/python
ARG APP_USER_UID=999
FROM python:3.12-slim as default-image
    ARG APP_USER
    ARG APP_HOME
    ARG APP_USER_UID
    ENV PATH "${APP_HOME}/.local/bin:${APP_HOME}/.venv/bin:$PATH"
    WORKDIR $APP_HOME
    RUN groupadd -r ${APP_USER} && \
        useradd -r -g ${APP_USER} -u ${APP_USER_UID} -d $APP_HOME $APP_USER && \
        chown -R python:python $APP_HOME

FROM default-image as builder
    USER ${APP_USER}
    ENV PIP_CACHE_DIR "${APP_HOME}/.cache/pip"
    RUN --mount=type=cache,target=${APP_HOME}/.cache/pip,id=python-pip \
        python -m pip install --upgrade pip && \
        python -m pip install poetry

    COPY --chown=python:python pyproject.toml poetry.lock ./

    ENV POETRY_NO_INTERACTION 1
    ENV POETRY_VIRTUALENVS_IN_PROJECT 1
    ENV POETRY_VIRTUALENVS_CREATE 1
    ENV POETRY_CACHE_DIR /tmp/poetry_cache
    RUN --mount=type=secret,id=POETRY_AUTH_USER,uid=${APP_USER_UID} \
        --mount=type=secret,id=POETRY_AUTH_PASS,uid=${APP_USER_UID} \
        --mount=type=cache,mode=0777,target=${POETRY_CACHE_DIR} \
        poetry config http-basic.private_repo $(cat /run/secrets/POETRY_AUTH_USER) $(cat /run/secrets/POETRY_AUTH_PASS) && \
        poetry install --no-root --no-interaction --without dev

    COPY src/   src

FROM builder as unit-test
    COPY tests/ tests
    RUN --mount=type=secret,id=POETRY_AUTH_USER,uid=${APP_USER_UID} \
        --mount=type=secret,id=POETRY_AUTH_PASS,uid=${APP_USER_UID} \
        --mount=type=cache,target=${APP_HOME}/.cache/pypoetry \
        poetry config http-basic.private_repo $(cat /run/secrets/POETRY_AUTH_USER) $(cat /run/secrets/POETRY_AUTH_PASS) && \
        poetry install --no-root --no-interaction --only dev

    ENV PYTHONPATH "${PYTHONPATH}:${APP_HOME}/src"
    ENTRYPOINT [ "pytest" ]
    CMD [ "tests" ]

FROM default-image as dist-service
    RUN apt-get update && apt-get install -y tini
    USER ${APP_USER}
    COPY --from=builder ${APP_HOME} ${APP_HOME}
    ENTRYPOINT ["tini", "--"]
    CMD ["python", "src/run.py"]

FROM default-image as dist-api
    USER ${APP_USER}
    COPY --from=builder ${APP_HOME} ${APP_HOME}
    ENTRYPOINT ["uvicorn"]
    CMD ["src.main:app"]
