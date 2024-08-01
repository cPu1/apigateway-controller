# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Utilities for working with Resource resources"""

import datetime
import time

import boto3
import pytest
from e2e import SERVICE_NAME


DEFAULT_WAIT_UNTIL_DELETED_TIMEOUT_SECONDS = 60 * 10
DEFAULT_WAIT_UNTIL_DELETED_INTERVAL_SECONDS = 15


def wait_until_deleted(
        resource_id: str,
        rest_api_id: str,
        timeout_seconds: int = DEFAULT_WAIT_UNTIL_DELETED_TIMEOUT_SECONDS,
        interval_seconds: int = DEFAULT_WAIT_UNTIL_DELETED_INTERVAL_SECONDS,
) -> None:
    """Waits until a Resource with a supplied ID is no longer returned from
    the API Gateway API.

    Usage:
        from e2e.resource import wait_until_deleted
        wait_until_deleted(resource_id)

    Raises:
        pytest.fail if the Resource is not deleted within timeout_timeout_seconds.
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while True:
        if datetime.datetime.now() >= timeout:
            pytest.fail(
                "Timed out waiting for Resource to be "
                "deleted in API Gateway API"
            )
        time.sleep(interval_seconds)

        latest = get(resource_id, rest_api_id)
        if latest is None:
            break


def get(resource_id, rest_api_id):
    """Returns a dict containing the Resource record from the API Gateway API.

    If no such Resource exists, returns None.
    """
    c = boto3.client(SERVICE_NAME)
    try:
        return c.get_resource(resourceId=resource_id, restApiId=rest_api_id)
    except c.exceptions.NotFoundException:
        return None

