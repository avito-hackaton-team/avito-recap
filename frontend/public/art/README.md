# Иллюстрации слайдов

Файлы `*.png` — это [Fluent Emoji](https://github.com/microsoft/fluentui-emoji)
от Microsoft, 3D-варианты. Лицензия **MIT**, использование в проекте свободное
при сохранении копии лицензии.

Имена файлов совпадают с типами слайдов из контракта, поэтому компонент `Art`
собирает путь как `/art/<type>.png` без отдельной таблицы соответствий:

| Файл | Что это | Оригинал в Fluent Emoji |
|---|---|---|
| `intro.png` | хлопушка | Party popper |
| `active_days.png` | календарь | Calendar |
| `views.png` | лупа | Magnifying glass tilted left |
| `favorites.png` | сердце | Red heart |
| `favorite_category.png` | телефон | Mobile phone |
| `purchases.png` | пакеты | Shopping bags |
| `sales.png` | мешок денег | Money bag |
| `messages.png` | облачко речи | Speech balloon |
| `interests.png` | диаграмма | Bar chart |
| `archetype.png` | лицо с моноклем | Face with monocle |
| `final.png` | искры | Sparkles |
| `season-winter.png` | снежинка | Snowflake |
| `season-spring.png` | росток | Seedling |
| `season-summer.png` | чемодан | Luggage |
| `season-autumn.png` | лист | Fallen leaf |

Два файла сгенерированы отдельно (ChatGPT, промпты в `../art-prompts.md`),
уменьшены до 512×512 и сжаты:

| Файл | Что это |
|---|---|
| `mascot.png` | маскот финального экрана: парень в фиолетовом худи |
| `badge.png` | шестиугольная медаль с лавровыми ветвями |

`active_days.svg` — нарисованный вручную запасной вариант. Компонент `Art`
пробует форматы по очереди `png` → `svg` → эмодзи-глиф, так что если положить
сюда ассеты из дизайн-макета под теми же именами, они заменят текущие без правок
кода.
