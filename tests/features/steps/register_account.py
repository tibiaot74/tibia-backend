import bcrypt
import requests
from behave import given, when, then


@given('No account with name "{name}" and email "{email}" exists')
def step_impl(context, name, email):
    context.query_db(f"DELETE FROM accounts WHERE name = '{name}' OR email = '{email}'")


@when('Client creates account with name "{name}", email "{email}" and password "{password}"')
def step_impl(context, name, email, password):
    requests.post(f"{context.url}/account", json={"name": int(name), "email": email, "password": password})


@then('Account with name "{name}", email "{email}" and password "{password}" (encrypted with bcrypt) exists')
def step_impl(context, name, email, password):
    result = context.query_db(f"SELECT * FROM accounts WHERE name = '{name}' AND email = '{email}'")
    account = result[0]

    assert account
    assert account["name"] == name
    assert account["email"] == email
    assert bcrypt.checkpw(password.encode("utf8"), account["password"].encode("utf8"))
