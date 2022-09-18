Feature: Register Player

    Background: Login into account
        Given An account with name "744774", email "jaspion@bol.com.br" and password "Senha123" exists
        And Client is logged into account of name "744774" with password "Senha123"

    Scenario Outline: Scenario Outline name: Successfully register new player
        Given No player with name <name> exists
        When Client tries to create player with name <name>, sex <sex> and outfit <outfit>
        Then Player with name <name>, sex <sex> and outfit <outfit> is created in logged account

        Examples: Players
            | name           | sex      | outfit   |
            | "Talin Rabano" | "Male"   | "mage"   |
            | "Paula Tejano" | "Female" | "hunter" |
            | "Levi Adão "   | "Male"   | "noble"  |
            | "Cuca Beludo " | "Male"   | "knight" |

    Scenario: Register player that already exists
        Given A player with name "Jacinto Pinto" exists
        When Client tries to create player with name "Jacinto Pinto", sex "Male" and outfit "mage"
        Then Player creation fails

    Scenario Outline: Register player with invalid fields
        Given No player with name <name> exists
        When Client tries to create player with name <name>, sex <sex> and outfit <outfit>
        Then Player creation fails

        Examples: Players
            | name                              | sex    | outfit        |
            | "Jacinto Pinto Leite Aquino Rego" | "Male" | "mage"        |
            | "Melbi Lau"                       | "Male" | "mage_female" |

    Scenario: Create more players than max number of player per account
        Given Max number of players created for an account
        When Client tries to create player with name "Jacinto Pinto", sex "Male" and outfit "mage"
        Then Player creation fails