Feature: Delete Players

    Background: Login into account
        Given An account with name "744774", email "jaspion@bol.com.br" and password "Senha123" exists
        And Client is logged into account of name "744774" with password "Senha123"

    Scenario: Delete nonexistent player in account
        Given No player with name "Raley Abunda" exists
        When Player named "Raley Abunda" is deleted
        Then Player "Raley Abunda" can not be deleted

    Scenario: Delete player in account
        Given A player with name "Yasmin Asbolla" exists
        When Player named "Yasmin Asbolla" is deleted
        Then Player "Yasmin Asbolla" does not exist anymore

    Scenario: Delete player from another account
        Given A player with name "Torei Ukano" exists
        And An account with name "102030", email "qualquercoisa@uol.com.br" and password "321123" exists
        And Client is logged into account of name "102030" with password "321123"
        When Player named "Torei Ukano" is deleted
        Then Player "Torei Ukano" can not be deleted