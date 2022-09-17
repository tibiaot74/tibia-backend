from behave import given, when, then
from hamcrest import *

from features.helpers.player_helper import outfitToInt, intToOutfit, sexToBool


@given('No player with name "{name}" exists')
def step_impl(context, name):
    context.query_db(f"DELETE FROM players WHERE name = '{name}'")


@given('A player with name "{name}" exists')
def step_impl(context, name):
    context.query_db(
        f"INSERT INTO players (`name`, `account_id`, `conditions`, `sex`, `looktype`, `auction_balance`, `created`, `nick_verify`, `comment`, `signature`, `castDescription`) VALUES ('{name}', 12, '', 1, 130, 0, 10000000, '', '', '', '')"
    )


@when('Client tries to create player with name "{name}", sex "{sex}" and outfit "{outfit}"')
def step_impl(context, name, sex, outfit):
    context.response = context.session.post(
        f"{context.url}/account/player", json={"name": name, "sex": sexToBool(sex), "outfit": outfit}
    )


@then('Player with name "{name}", sex "{sex}" and outfit "{outfit}" is created in logged account')
def step_impl(context, name, sex, outfit):
    player = context.query_db(f"SELECT * FROM players WHERE name = '{name}'")[0]

    assert_that(player, not_none())
    assert_that(player["name"], equal_to(name))
    assert_that(player["sex"], equal_to(sexToBool(sex)))
    assert_that(player["looktype"], equal_to(outfitToInt(outfit, sexToBool(sex))))
    assert_that(player["account_id"], equal_to(context.jwt["id"]))

    response = context.response.json()
    assert_that(context.response.status_code, equal_to(201))
    assert_that(response["id"], equal_to(player["id"]))
    assert_that(response["name"], equal_to(player["name"]))
    assert_that(response["sex"], equal_to(player["sex"]))
    assert_that(response["outfit"], equal_to(intToOutfit(player["looktype"], sexToBool(sex))))


@then("Player creation fails")
def step_impl(context):
    response = context.response.json()
    print(response)
    assert_that(context.response.status_code, equal_to(400))