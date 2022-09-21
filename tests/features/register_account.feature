Feature: Register Account

    Scenario: Successfully register account
        Given No account with name "744774" and email "jaspion@bol.com.br" exists
        When Client creates account with name "744774", email "jaspion@bol.com.br" and password "Senha123"
        Then Account with name "744774", email "jaspion@bol.com.br" and password "Senha123" (encrypted with bcrypt) exists

    Scenario Outline: Register account that already exists
        Given An account with name "744774" and email "jaspion@bol.com.br" exists
        When Client creates account with name <name>, email <email> and password "Senha123"
        Then Account creation fails

        Examples: Accounts
            | name     | email                |
            | "744774" | "other@bol.com.br"   |
            | "101010" | "jaspion@bol.com.br" |

    Scenario Outline: Register account with invalid fields
        Given No account with name <name> and email <email> exists
        When Client creates account with name <name>, email <email> and password <password>
        Then Account creation fails

        Examples: Accounts
            | name     | email                | password   |
            | "74"     | "jaspion@bol.com.br" | "Senha123" |
            | "744774" | "jaspion"            | "Senha123" |
            | "744774" | "jaspion@bol.com.br" | "a"        |

    Scenario Outline: Register account with non compliance password
        Given No account with name <name> and email <email> exists
        When Client creates account with name <name>, email <email> and password <password>
        Then Account creation fails

    Examples: Accounts
        | name     | email                | password                                            |
        | "74"     | "jaspion@bol.com.br" | "curta"                                             |
        | "744774" | "jaspion"            | "senhaextremamentegrandequenaofazsentidomastudobem" |
        | "744774" | "jaspion@bol.com.br" | "Asenha!"                                           |