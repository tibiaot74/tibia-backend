Feature: Register Account


Scenario: Successfully register account
    Given No account with name "744774" and email "jaspion@bol.com.br" exists
    When Client creates account with name "744774", email "jaspion@bol.com.br" and password "Senha123"
    Then Account with name "744774", email "jaspion@bol.com.br" and password "Senha123" (encrypted with bcrypt) exists