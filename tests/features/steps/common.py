from behave import given, when, then
import time


@when('Client waits "{seconds}" seconds')
def step_impl(context, seconds):
    time.sleep(int(seconds))
