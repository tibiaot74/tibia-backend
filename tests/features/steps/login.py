import requests
from behave import given, when, then
from hamcrest import *
import jwt


@given("Client is not logged in")
def step_impl(context):
    context.session = requests.Session()


@given('Client is logged into account of name "{name}" with password "{password}"')
def step_impl(context, name, password):
    context.session = requests.Session()
    response = context.session.post(f"{context.url}/login", json={"name": int(name), "password": password})
    context.session.headers = {"Authorization": response.json()["token"]}


@when('Client tries to login with name "{name}" and password "{password}"')
def step_impl(context, name, password):
    context.logged_name = name
    context.response = requests.post(f"{context.url}/login", json={"name": int(name), "password": password})


@when("Client tries to access a secured functionality")
def step_impl(context):
    context.response = context.session.get(f"{context.url}/ping")


@then("Login is successfull")
def step_impl(context):
    response = context.response.json()
    assert_that(context.response.status_code, equal_to(200))
    assert_that(response["token"], not_none())
    decoded_jwt = jwt.decode(response["token"], "JWT_TOKEN", algorithms=["HS256"])

    account = context.query_db(f"SELECT * FROM accounts WHERE name = '{context.logged_name}'")[0]

    assert_that(decoded_jwt["id"], equal_to(account["id"]))
    assert_that(decoded_jwt["name"], equal_to(account["name"]))
    assert_that(decoded_jwt["email"], equal_to(account["email"]))


@then("Login fails for inexistent account")
def step_impl(context):
    assert_that(context.response.status_code, equal_to(404))


@then("Login fails for wrong credentials")
def step_impl(context):
    assert_that(context.response.status_code, equal_to(401))


@then("Login authorization is refreshed")
def step_impl(context):
    response = context.response.json()
    assert_that(response["token"], is_not(equal_to(context.jwt)))


@then("Secured functionality is accessible")
def step_impl(context):
    response = context.response.json()
    assert_that(context.response.status_code, equal_to(200))
    assert_that(response["message"], equal_to("pong"))


@then("Secured functionality is not accessible")
def step_impl(context):
    assert_that(context.response.status_code, equal_to(401))