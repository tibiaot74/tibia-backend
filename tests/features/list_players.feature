Feature: List Players

    Background: Login into account
        Given An account with name "744774", email "jaspion@bol.com.br" and password "Senha123" exists
        And Client is logged into account of name "744774" with password "Senha123"

    Scenario: Account with no players created
        When Client tries to list all players in their account
        Then No players are retrieved

    Scenario:
        Given A player with name "Jacinto Pinto" exists
        When Client tries to list all players in their account
        Then Player named "Jacinto Pinto" is listed