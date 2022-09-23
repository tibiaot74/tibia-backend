from behave import when, then
from behave.runner import Context
from hamcrest import *


@when('Player named "{name}" is deleted')
def step_impl(context: Context, name: str):
    result = context.query_db(
        f"SELECT `id` FROM players WHERE `name` = '{name}'"
    )
    player_id = result[0]["id"] if result else '0'
    context.player_id = player_id
    context.response = context.session.delete(
        f"{context.url}/account/player/{player_id}"
    )


@then('Player "{name}" does not exist anymore')
def step_impl(context: Context, name: str):
    result = context.query_db(
        f"SELECT `id` FROM players WHERE `name` = '{name}'"
    )
    assert_that(len(result), equal_to(0))
    assert_that(context.response.status_code, equal_to(204))


@then('Player "{name}" can not be deleted')
def step_impl(context: Context, name: str):
    assert_that(context.response.status_code, equal_to(401))