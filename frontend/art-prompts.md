# Промпты для 3D-иллюстраций

Как пользоваться: сгенерировать картинку (ChatGPT, Midjourney, Recraft — любой),
сохранить как PNG **с прозрачным фоном**, положить в `public/art/` под именем из
колонки «Файл». Компонент `Art` подхватит её автоматически, код менять не нужно.

Размер: квадрат, от 512×512. Иконки на слайде рисуются в 380px, на карточках
сезонов — 92px, в плитках финала — 40px.

## Общая часть стиля

Добавляй этот хвост к каждому промпту — он держит набор в одной манере:

```
3D rendered icon, isometric view, soft studio lighting, glossy plastic material,
rounded chunky shapes, purple #7C3AED and violet #A78BFA palette with white
accents, subtle drop shadow, centered composition, transparent background,
square 1:1, no text, no watermark
```

## Чего не хватает сейчас

Эти два места закрыты слабее всего — их я бы генерил в первую очередь.

| Файл | Промпт |
|---|---|
| `mascot.png` | `Friendly cartoon young man wearing a bright purple hoodie, winking, making a heart shape with both hands, upper body, floating confetti stars and cubes around him` |
| `badge.png` | `Hexagonal award medal with laurel wreath, purple gradient body, small white bar chart with rising arrow in the center, sparkles around` |

Маскот сейчас не выводится вообще (на финальном экране его нет), медаль —
обычная «1 место» из Fluent Emoji.

## Если захочется свой набор целиком

Сейчас эти места закрыты иконками Fluent Emoji 3D (MIT). Заменять не обязательно,
но если нужен полностью авторский набор — вот объекты:

| Файл | Объект |
|---|---|
| `intro.png` | party popper with confetti |
| `active_days.png` | wall calendar with one day highlighted in green |
| `views.png` | smartphone showing a list of ads, with a magnifying glass |
| `favorites.png` | glossy heart |
| `favorite_category.png` | smartphone with headphones next to it |
| `purchases.png` | two shopping bags |
| `sales.png` | price tag with a ruble sign |
| `messages.png` | two overlapping speech bubbles |
| `interests.png` | bar chart with an upward arrow |
| `archetype.png` | magnifying glass over a laptop |
| `final.png` | sparkles and stars |
| `season-winter.png` | snowflake |
| `season-spring.png` | potted seedling |
| `season-summer.png` | yellow suitcase with a sun |
| `season-autumn.png` | orange armchair with falling leaves |

## Что важно проверить после генерации

- фон именно прозрачный, а не белый: на тёмной сцене белый квадрат видно сразу;
- объект занимает почти весь кадр, без больших полей;
- одинаковый угол обзора у всех иконок, иначе набор рассыпается.
