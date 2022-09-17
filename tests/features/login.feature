Feature: Login

    Scenario: Successfully login into account
        Given An account with name "744774", email "jaspion@bol.com.br" and password "Senha123" exists
        When Client tries to login with name "744774" and password "Senha123"
        Then Login is successfull

    Scenario: Refresh login
        Given An account with name "744774", email "jaspion@bol.com.br" and password "Senha123" exists
        And Client is logged into account of name "744774" with password "Senha123"
        When Client waits "1" seconds
        And Client tries to login with name "744774" and password "Senha123"
        Then Login authorization is refreshed

    Scenario: Login with inexistent account
        Given No account with name "744774" and email "jaspion@bol.com.br" exists
        When Client tries to login with name "744774" and password "Senha123"
        Then Login fails for inexistent account

    Scenario: Login with wrong credentials
        Given An account with name "744774", email "jaspion@bol.com.br" and password "Senha123" exists
        When Client tries to login with name "744774" and password "Senha"
        Then Login fails for wrong credentials