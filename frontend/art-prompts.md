# Avito Soft 3D — manifest иллюстраций

Единый набор 3D-ассетов для desktop-истории. Текущие `views.webp`,
`active_days.webp` и `favorites.webp` — референсы серии, а не объекты для прямого
копирования в следующих генерациях.

## Зафиксированный стиль

- мягкий объёмный 3D, округлые формы без мелкой детализации;
- soft-touch пластик с умеренным сатиновым бликом, без стекла и хрома;
- фронтальный ракурс 3/4, камера чуть сверху, почти без перспективных искажений;
- один крупный объект или компактная группа, читаемая в размере 280–440 px;
- мягкий верхний свет слева, аккуратное ambient occlusion между деталями;
- прозрачный фон без комнаты, пола, рамки и цветной подложки;
- без текста, цифр, логотипов, водяных знаков и фотореализма.

### Палитра

| Роль | Цвет |
|---|---|
| Основной синий | `#00AAFF` |
| Глубокий синий | `#006EDB` |
| Коралловый акцент | `#FF6163` |
| Зелёный акцент | `#97CF26` |
| Фиолетовый акцент | `#A169F7` |
| Светлая поверхность | `#F5F5F7` |
| Тёмные детали | `#15151A` |

На одном ассете использовать один основной цвет и не больше двух акцентов.
Белые поверхности должны оставаться нейтральными, без серой фоновой плашки.

## Референсы

| Файл | Что брать за образец |
|---|---|
| `views.webp` | основной эталон камеры, масштаба, синего пластика и контраста |
| `active_days.webp` | эталон округлой конструкции, светлых поверхностей и зелёного акцента |
| `favorites.webp` | эталон кораллового акцента и стопки карточек; не переносить его тёплый серо-красный ореол в остальные ассеты |

При генерации с референсами прикладывать все три изображения и добавлять:

```text
Use the attached images only as a visual system reference. Match the camera,
rounded geometry, soft-touch plastic material, lighting softness, color depth,
object scale and transparent cutout quality. Create the new subject described
below; do not copy the reference objects or their exact composition.
```

## Master prompt

Подставить описание объекта вместо `[SUBJECT]`:

```text
A premium stylized 3D illustration of [SUBJECT]. Avito Soft 3D visual language:
rounded chunky geometry, simplified but clearly recognizable construction,
soft-touch satin plastic, subtle bevels, controlled highlights and soft ambient
occlusion. Front three-quarter view from slightly above, near-orthographic
camera, consistent proportions and lighting with the attached reference set.
Avito palette: cyan blue #00AAFF, deep blue #006EDB, coral #FF6163, lime green
#97CF26, violet #A169F7, off-white #F5F5F7 and dark details #15151A. Use one
dominant color and no more than two accents. One centered hero object or one
compact object group, strong readable silhouette, fills about 84% of a square
canvas with even safe margins. Transparent background, clean RGBA cutout, no
scene, no floor, no background plate, no text, no letters, no numbers, no logo,
no watermark, no photorealism. 1024x1024 PNG.
```

Если генератор поддерживает negative prompt:

```text
busy composition, tiny details, thin fragile parts, hard black shadow, gray or
colored background, pedestal, room, horizon, excessive glow, glass, chrome,
metallic reflections, photorealism, text, letters, numbers, logo, watermark,
cropped object, multiple camera angles
```

## Manifest

Основные ассеты со статусом `integrated` уже используются в desktop-истории.
Для остальных к Master prompt добавлять описание из таблицы.

| Статус | Файл | `[SUBJECT]` и композиция |
|---|---|---|
| reference | `views.webp` | a blue smartphone with three simple marketplace listing cards on screen and one large dark magnifying glass leaning across its lower-right edge |
| reference | `active_days.webp` | a chunky blue and off-white desk calendar with three rounded green marker pins, one compact object with a strong silhouette |
| reference | `favorites.webp` | a small offset stack of off-white marketplace listing cards with one oversized coral heart-shaped bookmark in front |
| integrated | `messages.webp` | two overlapping rounded chat bubbles, cyan blue and coral, with a check mark and a small marketplace tag |
| integrated | `badge.webp` | one premium rounded hexagonal achievement medal in cyan and deep blue with a simple raised star shape and two short coral ribbons, no lettering |
| integrated | `badge-sales.webp` | one premium rounded sales medal in coral and deep blue with a simple raised price-tag shape and two short cyan ribbons, no lettering |
| integrated | `season-winter.webp` | one chunky rounded snowflake in cyan and off-white, simplified to six broad arms |
| integrated | `season-spring.webp` | one young green sprout with two large leaves growing from a rounded off-white pot with a cyan rim |
| integrated | `season-summer.webp` | one rounded cyan suitcase with a coral handle and a small raised sun shape, no stickers or text |
| integrated | `season-autumn.webp` | one warm coral lounge chair with two broad floating autumn leaves in violet and muted orange |
| integrated | `archetype-collector.webp` | a compact display case holding three carefully arranged rounded finds: a small lamp, camera and vase, with cyan structure and coral accents |
| integrated | `archetype-explorer.webp` | a chunky cyan compass combined with a small magnifying glass, pointing diagonally upward, with a lime-green accent |
| integrated | `archetype-negotiator.webp` | two rounded speech bubbles meeting around a simple handshake shape, cyan and violet with one coral accent |
| integrated | `archetype-dealmaker.webp` | a rounded briefcase with a raised checkmark-shaped clasp and one small price tag, deep blue with lime-green and coral accents |
| integrated | `final.webp` | an open off-white personal album containing four simplified marketplace cards, with colored section tabs and small sparkles |
| integrated | `interests.webp` | three rounded ascending columns with one broad curved upward arrow, cyan with lime-green highlight, no axes or numbers |

## Экспорт и проверка

- исходник генерации: `PNG`, `1024×1024`, цветовое пространство `sRGB`, настоящий alpha;
- delivery-файл: `WebP` с альфа-каналом; исходник хранить отдельно, а файл после оптимизации желательно
  удерживать в пределах 500 KB без заметной деградации краёв и градиентов;
- объект занимает примерно 80–88% кадра и не касается краёв;
- проверить на белом и тёмном фоне: не должно быть прямоугольной подложки,
  грязного серого ореола или жёсткой обрезки тени;
- визуально сравнить рядом с тремя референсами при одинаковом размере;
- имя файла должно точно совпадать с manifest и быть в `public/art/`;
- сохранять исходный prompt и seed рядом с задачей или в описании генерации,
  чтобы можно было воспроизвести серию.

Следующая итерация ассетов: две медали → четыре сезона → четыре архетипа.
