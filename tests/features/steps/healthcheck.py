from behave import when, then
from behave.runner import Context
from hamcrest import *
import requests


@when("Client tries to use application")
def step_impl(context: Context):
    context.response = requests.get(f"{context.url}/health")


@then("Application responds successfully")
def step_impl(context: Context):
    assert_that(context.response.status_code, equal_to(200))