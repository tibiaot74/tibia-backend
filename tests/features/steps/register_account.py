import bcrypt
import requests
from behave import given, when, then
from behave.runner import Context
from hamcrest import *


@given('No account with name "{name}" and email "{email}" exists')
def step_impl(context: Context, name: str, email: str):
    context.query_db(f"DELETE FROM accounts WHERE name = '{name}' OR email = '{email}'")


@given('An account with name "{name}" and email "{email}" exists')
@given('An account with name "{name}", email "{email}" and password "{password}" exists')
def step_impl(context: Context, name: str, email: str, password: str = "Senha123"):
    pass_hash = bcrypt.hashpw(password.encode("utf8"), bcrypt.gensalt()).decode("utf8")
    context.query_db(
        f"INSERT INTO accounts (`name`, `email`, `password`, `premdays`, `lastday`, `key`, `warnings`, `premium_points`, `backup_points`, `guild_points`, `guild_points_stats`, `blocked`, `group_id`, `vip_time`, `email_new`, `email_new_time`, `email_code`, `next_email`, `created`, `page_lastday`, `page_access`, `rlname`, `location`, `flag`, `last_post`, `create_date`, `create_ip`, `vote`) VALUES ('{name}', '{email}', '{pass_hash}', 0, 0, '', 0, 0, 0, 0, 0, 0, 0, 0, '', 0, '', 0, 0, 0, 0, '', '', '', 0, 0, 0, 0)"
    )
    result = context.query_db(f"SELECT `id` FROM accounts WHERE `name` = {name}")[0]
    context.account_id = int(result["id"])


@when('Client creates account with name "{name}", email "{email}" and password "{password}"')
def step_impl(context: Context, name: str, email: str, password: str):
    context.response = requests.post(
        f"{context.url}/account", json={"name": int(name), "email": email, "password": password}
    )


@then('Account with name "{name}", email "{email}" and password "{password}" (encrypted with bcrypt) exists')
def step_impl(context: Context, name: str, email: str, password: str):
    result = context.query_db(f"SELECT * FROM accounts WHERE name = '{name}' AND email = '{email}'")
    account = result[0]

    assert_that(account, not_none())
    assert_that(account["name"], equal_to(name))
    assert_that(account["email"], equal_to(email))
    assert_that(bcrypt.checkpw(password.encode("utf8"), account["password"].encode("utf8")))

    response = context.response.json()
    assert_that(context.response.status_code, equal_to(201))
    assert_that(response["id"], equal_to(account["id"]))
    assert_that(response["name"], equal_to(account["name"]))
    assert_that(response["email"], equal_to(account["email"]))


@then("Account creation fails")
def step_impl(context: Context):
    response = context.response.json()
    print(response)
    assert_that(context.response.status_code, equal_to(400))