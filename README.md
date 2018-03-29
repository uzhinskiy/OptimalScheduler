# OptimalScheduler
Задачка из учебника Скиены

Генерация исходных данных:
`go run datagen.go`

На выходе получаем файл OptimalSchedule вида:
    ID      Start   Stop    Name
    1       12      17      Event0
    2       12      19      Event1
    3       9       15      Event2


Обработка данных:
`cat OptimalSchedule | go run optsched.go`

Результат исполнения:
 * Список событий
 * Графическое решение задачи (для этого необходим gnuplot)



