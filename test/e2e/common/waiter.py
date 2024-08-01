"""Utilities for working with API Gateway resources"""

import datetime
import time
import typing

import boto3
import pytest

DEFAULT_WAIT_UNTIL_TIMEOUT_SECONDS = 60*30
DEFAULT_WAIT_UNTIL_INTERVAL_SECONDS = 15
DEFAULT_WAIT_UNTIL_DELETED_TIMEOUT_SECONDS = 60*10
DEFAULT_WAIT_UNTIL_DELETED_INTERVAL_SECONDS = 15


ResourceMatchFunc = typing.NewType(
    'ResourceMatchFunc',
    typing.Callable[[dict], bool],
)

GetResourceFunc = typing.NewType(
    'GetResourceFunc',
    typing.Callable[[], dict],
)


def wait_until_deleted(
        get_resource: GetResourceFunc,
        timeout_seconds: int = DEFAULT_WAIT_UNTIL_TIMEOUT_SECONDS,
        interval_seconds: int = DEFAULT_WAIT_UNTIL_INTERVAL_SECONDS,
) -> None:
    # TODO
    """Waits until a resource is deleted from the API Gateway API

    Usage:
        from e2e.waiter import wait_until, status_matches

        wait_until(
            certificate_arn,
            status_matches("ISSUED"),
        )

    Raises:
        pytest.fail upon timeout
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while safe_get(get_resource) is not None:
        if datetime.datetime.now() >= timeout:
            pytest.fail('Timed out waiting for resource to be deleted in API Gateway')
        time.sleep(interval_seconds)


def safe_get(get_resource: GetResourceFunc):
    client = boto3.client('apigateway')
    try:
        return get_resource()
    except client.exceptions.NotFoundException:
        return None
