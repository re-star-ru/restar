Сущности для работы

```mermaid
flowchart LR
    Склад
    Ячейка --> Склад
    Товар
    Характеристика --> Товар
    Остаток --> Характеристика
    Остаток --> Ячейка
    
    Работник
    Вид_работы
    Акт_дефектовки --> Работник
    Акт_дефектовки --> Вид_работы
    
    Талон_ремонта --> Акт_дефектовки
    Талон_ремонта --> Работник
    Талон_ремонта --> Вид_работы
    
    Чек --> Талон_ремонта
    Чек --> Продажа
```