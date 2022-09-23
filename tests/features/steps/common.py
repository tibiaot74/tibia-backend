from behave import given, when, then
from behave.runner import Context
import time


@when('Client waits "{seconds}" seconds')
def step_impl(context: Context, seconds: str):
    time.sleep(int(seconds))
