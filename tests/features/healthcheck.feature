Feature: Healthcheck

    Scenario: Check application is online
        When Client tries to use application
        Then Application responds Successfully