## 1.7.3. Игра без права на ошибку

В совсем маленьких программах обертывать ошибки обычно не требуется. Поэтому сразу скажу, что задачка эта непривычно большая. Заодно попрактикуетесь в работе с унаследованным кодом на Go. Поскольку это финальная задача модуля — можно считать ее толстым боссом ツ

Мы будем писать игру, в которой любая ошибка приводит к трагическому финалу. Сначала пройдем по коду, а затем я расскажу, что требуется сделать.

### Команда, объект, шаг

В нашей игре участник выполняет некоторые команды над объектами игрового мира. Поэтому есть тип command и и список доступных команд:

```Go
// label - уникальное наименование
type label string

// command - команда, которую можно выполнять в игре
type command label

// список доступных команд
var (
    eat  = command("eat")
    take = command("take")
    talk = command("talk to")
)
```

Есть объекты игрового мира — тип `thing`. Каждый объект знает, какие команды для него доступны:

```Go
// thing - объект, который существует в игре
type thing struct {
    name    label
    actions map[command]string
}

// supports() возвращает true, если объект
// поддерживает команду action
func (t thing) supports(action command) bool {
    _, ok := t.actions[action]
    return ok
}
```

Список объектов:

```Go
// полный список объектов в игре
var (
    apple = thing{"apple", map[command]string{
        eat:  "mmm, delicious!",
        take: "you have an apple now",
    }}
    bob = thing{"bob", map[command]string{
        talk: "Bob says hello",
    }}
    coin = thing{"coin", map[command]string{
        take: "you have a coin now",
    }}
    mirror = thing{"mirror", map[command]string{
        take: "you have a mirror now",
        talk: "mirror does not answer",
    }}
    mushroom = thing{"mushroom", map[command]string{
        eat:  "tastes funny",
        take: "you have a mushroom now",
    }}
)
```

Есть тип `step` — шаг игры. Он объединяет команду и объект, над которым необходимо выполнить действие:

```Go
// step описывает шаг игры: сочетание команды и объекта
type step struct {
    cmd command
    obj thing
}

// isValid() возвращает true, если объект
// совместим с командой
func (s step) isValid() bool {
    return s.obj.supports(s.cmd)
}
```

### Игрок

Есть игрок — тип `player`:

```Go
// player - игрок
type player struct {
    // количество съеденного
    nEaten int
    // количество диалогов
    nDialogs int
    // инвентарь
    inventory []thing
}
```
Игрок умеет проверить, есть ли предмет в инвентаре:

```Go
// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj thing) bool {
    for _, got := range p.inventory {
        if got.name == obj.name {
            return true
        }
    }
    return false
}
```

И выполнить действие над указанным объектом:

```Go
// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj thing) error {
    // действуем в соответствии с командой
    switch cmd {
    case eat:
        if p.nEaten > 1 {
            return errors.New("you don't want to eat anymore")
        }
        p.nEaten++
    case take:
        if p.has(obj) {
            return fmt.Errorf("you already have a %s", obj)
        }
        p.inventory = append(p.inventory, obj)
    case talk:
        if p.nDialogs > 0 {
            return errors.New("you don't want to talk anymore")
        }
        p.nDialogs++
    }
    return nil
}
```

Как видите, тут кое-что может пойти не так:

* игрок уже ел или разговаривал, и больше не хочет,
* предмет уже есть в инвентаре.

На каждую такую ситуацию метод `do()` возвращает ошибку.

### Игра

Наконец, есть тип `game` — сама игра:

```Go
// game описывает игру
type game struct {
    // игрок
    player *player
    // объекты игрового мира
    things map[label]int
    // количество успешно выполненных шагов
    nSteps int
}
```

Исходно в игре есть некоторое количество предметов каждого типа (поле `things`). Игра умеет проверить, остались ли они или уже закончились:

```Go
// has() проверяет, остались ли в игровом мире указанные предметы
func (g *game) has(obj thing) bool {
    count := g.things[obj.name]
    return count > 0
}
```

Игра может выполнить шаг:

```Go
// execute() выполняет шаг step
func (g *game) execute(st step) error {
    // проверяем совместимость команды и объекта
    if !st.isValid() {
        return fmt.Errorf("cannot %s", st)
    }

    // когда игрок берет или съедает предмет,
    // тот пропадает из игрового мира
    if st.cmd == take || st.cmd == eat {
        if !g.has(st.obj) {
            return fmt.Errorf("there are no %ss left", st.obj)
        }
        g.things[st.obj.name]--
    }

    // выполняем команду от имени игрока
    if err := g.player.do(st.cmd, st.obj); err != nil {
        return err
    }

    g.nSteps++
    return nil
}
```

В методе `execute()` последовательность действий такая:

* Игра проверяет, что шаг корректный.
* Игра проверяет, остались ли предметы, с которыми требуется выполнить действие.
* Игра делегирует выполнение шага игроку.

Если что-то пошло не так — метод возвращает ошибку.

Новая игра выглядит так:

```Go
// newGame() создает новую игру
func newGame() *game {
    p := newPlayer()
    things := map[label]int{
        apple.name:    2,
        coin.name:     3,
        mirror.name:   1,
        mushroom.name: 1,
    }
    return &game{p, things, 0}
}
```

Пример игры с конкретной последовательностью шагов:

```Go
func main() {
    gm := newGame()
    steps := []step{
        {eat, apple},
        {talk, bob},
        {take, coin},
        {eat, mushroom},
    }

    for _, st := range steps {
        if err := tryStep(gm, st); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
    fmt.Println("You win!")
}

// tryStep() выполняет шаг игры и печатает результат
func tryStep(gm *game, st step) error {
    // ...
}
```

Здесь шаги приводят к успеху:

```
trying to eat apple... OK
trying to talk to bob... OK
trying to take coin... OK
trying to eat mushroom... OK
You win!
```

А здесь нет:

```Go
func main() {
    gm := newGame()
    steps := []step{
        {talk, bob},
        {talk, bob},
    }
    // ...
}
```

```
trying to talk to bob... OK
trying to talk to bob... FAIL
you don't want to talk anymore
exit status 1
```

### Задание

Как видно по коду, ошибки создаются через `errors.New()` и `fmt.Errorf()`. Хочется больше структуры, поэтому добавьте отдельный тип для каждого вида ошибок:

```Go
// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError
```

Метод `player.do()` должен возвращать либо одну из этих ошибок, либо `nil` (если ошибок не было).

Кроме того, создайте ошибку верхнего уровня:

```Go
// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
    // количество шагов, успешно выполненных
    // до того, как произошла ошибка
    nSteps int
    // ...
}
```

Метод `game.execute()` должен возвращать либо ошибку типа `gameOverError`, либо `nil` (если ошибок не было).

* Если метод получил ошибку от `player.do()` — пусть обернет ее в `gameOverError`.
* Если ошибка произошла в самом `game.execute()` — пусть создаст ошибку подходящего типа, а затем обернет ее в `gameOverError`.

И последнее. Создайте функцию `giveAdvice()`, которая дает игроку совет, как избежать случившейся ошибки в будущем:

```Go
func giveAdvice(err error) string {
    // ...
}
```

Правила работы `giveAdvice()`:

* Если команда не совместима с объектом, возвращает `things like 'COMMAND OBJECT' are impossible`, где `COMMAND` — название команды, а `OBJECT` — название объекта.
* Если в игре закончились объекты определенного типа, возвращает `be careful with scarce OBJECTs`, где `OBJECT` — название объекта.
* Если игрок слишком много ел, возвращает `eat less`. Если игрок слишком много говорил, возвращает `talk to less`.
* Если игрок превысил лимит на количество объектов определенного типа в инвентаре, возвращает `don't be greedy, LIMIT OBJECT is enough`, где `LIMIT` — значение лимита, а `OBJECT` — название объекта.

Например:

```
things like 'eat bob' are impossible
be careful with scarce mirrors
don't be greedy, 1 apple is enough
```

### Итого

1. Создайте отдельные типы ошибок и возвращайте ошибки этих типов в подходящих случаях.
2. Создайте тип `gameOverError` и возвращайте ошибку этого типа из `game.execute()`.
3. Создайте функцию `giveAdvice()`, которая возвращает совет на основе ошибки.

Полный код программы — в файле [fullcode.go](https://github.com/BalamutAndrey/Stepik-Thank-Go-Golang-in-practice/blob/master/1.7.3.%20A%20game%20without%20the%20right%20to%20make%20mistake/fullcode.go).

**Важно**: в качестве решения отправляйте не весь код, а только фрагмент, отмеченный комментариями «начало решения» и «конец решения».

___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**

**Sample Output:**
> **PASS**
___

```Go
// начало решения

// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError any

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError any

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError any

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError any

// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
    // количество шагов, успешно выполненных
    // до того, как произошла ошибка
    nSteps int
    // ...
}

// player - игрок
type player struct {
    // количество съеденного
    nEaten int
    // количество диалогов
    nDialogs int
    // инвентарь
    inventory []thing
}

// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj thing) bool {
    for _, got := range p.inventory {
        if got.name == obj.name {
            return true
        }
    }
    return false
}

// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj thing) error {
    // действуем в соответствии с командой
    switch cmd {
    case eat:
        if p.nEaten > 1 {
            return errors.New("you don't want to eat anymore")
        }
        p.nEaten++
    case take:
        if p.has(obj) {
            return fmt.Errorf("you already have a %s", obj)
        }
        p.inventory = append(p.inventory, obj)
    case talk:
        if p.nDialogs > 0 {
            return errors.New("you don't want to talk anymore")
        }
        p.nDialogs++
    }
    return nil
}

// newPlayer создает нового игрока
func newPlayer() *player {
    return &player{inventory: []thing{}}
}

// game описывает игру
type game struct {
    // игрок
    player *player
    // объекты игрового мира
    things map[label]int
    // количество успешно выполненных шагов
    nSteps int
}

// has() проверяет, остались ли в игровом мире указанные предметы
func (g *game) has(obj thing) bool {
    count := g.things[obj.name]
    return count > 0
}

// execute() выполняет шаг step
func (g *game) execute(st step) error {
    // проверяем совместимость команды и объекта
    if !st.isValid() {
        return fmt.Errorf("cannot %s", st)
    }

    // когда игрок берет или съедает предмет,
    // тот пропадает из игрового мира
    if st.cmd == take || st.cmd == eat {
        if !g.has(st.obj) {
            return fmt.Errorf("there are no %ss left", st.obj)
        }
        g.things[st.obj.name]--
    }

    // выполняем команду от имени игрока
    if err := g.player.do(st.cmd, st.obj); err != nil {
        return err
    }

    g.nSteps++
    return nil
}

// newGame() создает новую игру
func newGame() *game {
    p := newPlayer()
    things := map[label]int{
        apple.name:    2,
        coin.name:     3,
        mirror.name:   1,
        mushroom.name: 1,
    }
    return &game{p, things, 0}
}

// giveAdvice() возвращает совет, который
// поможет игроку избежать ошибки err в будущем
func giveAdvice(err error) string {
    // ...
    return ""
}

// конец решения
```