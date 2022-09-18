Feature: Delete Players

    Background: Login into account
        Given An account with name "744774", email "jaspion@bol.com.br" and password "Senha123" exists
        And Client is logged into account of name "744774" with password "Senha123"

    Scenario: Delete nonexistent player in account
        Given No player with name "Raley Abunda" exists
        When Player named "Raley Abunda" is deleted
        Then Player "Raley Abunda" does not exist anymore

    Scenario: Delete player in account
        Given A player with name "Yasmin Asbolla" exists
        When Player named "Yasmin Asbolla" is deleted
        Then Player "Yasmin Asbolla" does not exist anymore