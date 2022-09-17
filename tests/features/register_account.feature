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
        | name   | email              |
        | 744774 | other@bol.com.br   |
        | 101010 | jaspion@bol.com.br |

Scenario: Register with invalid email
    Given No account with name "744774" and email "jaspion@bol.com.br" exists
    When Client creates account with name "744774", email "jaspion" and password "Senha123"
    Then Account creation fails