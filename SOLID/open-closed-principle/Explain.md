# Princípio Aberto-Fechado (Open-Closed Principle)

O Princípio Aberto-Fechado é um dos cinco princípios do SOLID, um conjunto de diretrizes para a programação orientada a objetos que visa criar sistemas mais compreensíveis, flexíveis e de fácil manutenção. Este princípio foi formulado por Bertrand Meyer em 1988.

## Definição

O Princípio Aberto-Fechado afirma que:

> "Entidades de software (classes, módulos, funções, etc.) devem estar abertas para implementação, mas fechadas para modificação."

## Explicação

### Aberto para Implementação

Uma entidade de software deve permitir que seu comportamento seja estendido. Isso significa que você deve ser capaz de adicionar novas funcionalidades ou comportamentos ao sistema sem alterar o código existente. Isso é geralmente alcançado através da herança, composição ou implementação de interfaces.

### Fechado para Modificação

Uma entidade de software deve ser fechada para modificação, o que significa que, uma vez que a entidade está implementada e testada, seu código fonte não deve ser alterado. Isso ajuda a evitar a introdução de novos bugs em um sistema que já está funcionando corretamente.