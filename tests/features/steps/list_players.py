from behave import when, then
from behave.runner import Context
from hamcrest import *


@when('Client tries to list all players in their account')
def step_impl(context: Context):
    context.response = context.session.get(
        f"{context.url}/account/player"
    )


@then('No players are retrieved')
def step_impl(context: Context):
    response = context.response.json()
    print(response)
    assert_that(context.response.status_code, equal_to(200))
    assert_that(response["players"], equal_to([]))


@then('Player named "{name}" is listed')
def step_impl(context: Context, name: str):
    response = context.response.json()
    print(response)
    assert_that(context.response.status_code, equal_to(200))
    assert_that(len(response), equal_to(1))
    assert_that(response["players"][0]["name"], equal_to(name))
