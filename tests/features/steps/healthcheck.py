from behave import given, when, then
from hamcrest import *
import requests


@when("Client tries to use application")
def step_impl(context):
    context.response = requests.get(f"{context.url}/health")


@then("Application responds Successfully")
def step_impl(context):
    assert_that(context.response.status_code, equal_to(200))