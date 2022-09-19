from behave import given, when, then
from behave.runner import Context
from hamcrest import *
import time
from features.helpers.player_helper import outfitToInt, intToOutfit, sexToBool


@given('No player with name "{name}" exists')
def step_impl(context: Context, name: str):
    context.query_db(f"DELETE FROM players WHERE name = '{name}'")


@given('A player with name "{name}" exists')
def step_impl(context: Context, name: str):
    context.query_db(
        f"INSERT INTO players (`name`, `account_id`, `conditions`, `sex`, `looktype`, `auction_balance`, `created`, `nick_verify`, `comment`, `signature`, `castDescription`) VALUES ('{name}', {context.account_id}, '', 1, 130, 0, 10000000, '', '', '', '')"
    )


@given('Max number of players created for an account')
def step_impl(context: Context):
    player_names = [
        "Kevin Mammar",
        "Volin Rabei",
        "Oscar Alho",
        "Jacy Borreau",
        "Tommy Lipyca",
        "Oscar Belo do Saco",
        "Carlos Pinto Solto",
        "Paula Traz",
        "Diva Gina Berta",
        "Emma Thomas",
        "Dalva Gina",
        "Aliza Cresceu",
        "Alan Brado",
        "Balan Sarrola",
        "Davi Agra",
        "Kelly Inguissa",
        "ABC",
        "ABCDF",
        "ASDAD"
    ]
    values = []
    for name in player_names:
        values.append(f"('{name}', {context.account_id}, '', 1, 130, 0, 10000000, '', '', '', '')")
    values = ','.join(values)
    context.query_db(
        f"INSERT INTO players (`name`, `account_id`, `conditions`, `sex`, `looktype`, `auction_balance`, `created`, `nick_verify`, `comment`, `signature`, `castDescription`) VALUES {values}"
    )


@when('Client tries to create player with name "{name}", sex "{sex}" and outfit "{outfit}"')
def step_impl(context: Context, name: str, sex: str, outfit: str):
    context.response = context.session.post(
        f"{context.url}/account/player", json={"name": name, "sex": sexToBool(sex), "outfit": outfit}
    )


@then('Player with name "{name}", sex "{sex}" and outfit "{outfit}" is created in logged account')
def step_impl(context: Context, name: str, sex: str, outfit: str):
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
def step_impl(context: Context):
    response = context.response.json()
    print(response)
    assert_that(context.response.status_code, equal_to(400))