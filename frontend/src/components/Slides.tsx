import type { ReactNode } from 'react';

import type { Cta, Season, Slide, SlideType } from '../api/types';
import { Art } from './Art';

const SEASON_TITLES: Record<Season, string> = {
  winter: 'Зимой',
  spring: 'Весной',
  summer: 'Летом',
  autumn: 'Осенью',
};

/** Иконки лежат в public/art: Fluent Emoji 3D, MIT (см. public/art/README.md). */
function artUrl(name: string): string {
  return `/art/${name}.png`;
}

const SLIDE_ICONS: Record<SlideType, string> = {
  intro: '🎊',
  active_days: '📅',
  views: '👁',
  favorites: '♡',
  favorite_category: '🛍',
  purchases: '🛍',
  sales: '₽',
  messages: '💬',
  interests: '📊',
  archetype: '👤',
  final: '✨',
};

/** Плитки финального экрана переиспользуют иконки слайдов. */
const STAT_ART: Record<string, string> = {
  active_days: 'active_days',
  views: 'views',
  favorites: 'favorites',
  messages: 'messages',
  seasons: 'season-spring',
};

function formatNumber(value: number): string {
  return new Intl.NumberFormat('ru-RU').format(value);
}

/** Общая рамка слайда: подпись метрики слева сверху, контент слева, иллюстрация справа. */
function SlideLayout({
  slide,
  children,
  aside,
  centered = false,
}: {
  slide: Slide;
  children: ReactNode;
  aside?: ReactNode;
  centered?: boolean;
}) {
  return (
    <section className={`slide${centered ? ' slide--centered' : ''}`}>
      <div className="slide__main">
        <div className="slide__label">
          <span className="slide__label-icon" aria-hidden="true">
            {SLIDE_ICONS[slide.type]}
          </span>
          <span>{slide.title}</span>
        </div>

        {children}
      </div>

      <div className="slide__aside">{aside ?? <Art type={slide.type} />}</div>
    </section>
  );
}

/** Крупное число, под ним жирная строка и пояснение — как на макетах. */
function Counter({
  value,
  headline,
  note,
  children,
}: {
  value: number;
  headline?: string;
  note?: string;
  children?: ReactNode;
}) {
  return (
    <>
      <p className="counter">{formatNumber(value)}</p>
      {headline ? <p className="headline">{headline}</p> : null}
      {note ? <p className="note">{note}</p> : null}
      {children}
    </>
  );
}

export function SlideView({ slide, onShare }: { slide: Slide; onShare: () => void }) {
  switch (slide.type) {
    case 'intro':
      return (
        <SlideLayout slide={slide} centered>
          <h1 className="hero">
            Вспомни, каким был <span className="hero__year">{slide.year}</span> год вместе с Авито
          </h1>
          {slide.subtitle ? <p className="note note--wide">{slide.subtitle}</p> : null}
        </SlideLayout>
      );

    case 'active_days':
      return (
        <SlideLayout slide={slide}>
          <Counter
            value={slide.activeDays}
            headline={slide.subtitle}
            note="Искали, находили, общались и помогали вещам находить новых хозяев."
          />
        </SlideLayout>
      );

    case 'views':
      return (
        <SlideLayout slide={slide}>
          <Counter
            value={slide.views}
            headline={slide.subtitle}
            note="Вы настоящий исследователь: столько интересов помогли находить лучшее."
          />
        </SlideLayout>
      );

    case 'messages':
      return (
        <SlideLayout slide={slide}>
          <Counter
            value={slide.messages}
            headline={slide.subtitle}
            note="Благодаря вам сделки становились быстрее и проще."
          />
        </SlideLayout>
      );

    case 'favorites':
      return (
        <SlideLayout slide={slide}>
          <Counter
            value={slide.favorites}
            headline={slide.subtitle}
            note="Вы умеете находить интересное — часть отложенного всё ещё ждёт вас."
          >
            {slide.stillAvailable ? (
              <p className="pill">Ещё доступны: {slide.stillAvailable}</p>
            ) : null}
          </Counter>
        </SlideLayout>
      );

    case 'purchases':
      return (
        <SlideLayout
          slide={slide}
          aside={
            slide.badge ? (
              <BadgePanel caption="Ваш бейдж" title={slide.badge.title} note={slide.badge.description} />
            ) : undefined
          }
        >
          <Counter
            value={slide.purchases}
            headline={slide.subtitle}
            note="Выбираете с умом! Пусть каждая следующая покупка приносит больше пользы."
          />
        </SlideLayout>
      );

    case 'sales':
      return (
        <SlideLayout
          slide={slide}
          aside={
            slide.badge || slide.amountRange ? (
              <BadgePanel
                caption={slide.badge ? 'Ваш бейдж' : 'Итог года'}
                title={slide.badge?.title ?? 'Продажи года'}
                note={slide.badge?.description}
                amount={slide.amountRange?.label}
              />
            ) : undefined
          }
        >
          <Counter
            value={slide.sales}
            headline={slide.subtitle}
            note="Вы помогли другим найти нужные вещи, а вашим — обрести новых владельцев."
          />
        </SlideLayout>
      );

    case 'favorite_category':
      return (
        <SlideLayout slide={slide}>
          <p className="hero hero--accent">{slide.category.title}</p>

          {slide.subcategory ? (
            <p className="note">
              и подкатегория <span className="accent">{slide.subcategory.title}</span>
            </p>
          ) : null}

          {slide.subtitle ? <p className="headline">{slide.subtitle}</p> : null}

          {typeof slide.share === 'number' ? (
            <div className="meter" aria-label={`Доля активности ${slide.share}%`}>
              <span className="meter__fill" style={{ width: `${slide.share}%` }} />
            </div>
          ) : null}
        </SlideLayout>
      );

    case 'interests':
      return (
        <section className="slide slide--wide">
          <div className="slide__label">
            <span className="slide__label-icon" aria-hidden="true">
              {SLIDE_ICONS.interests}
            </span>
            <span>{slide.title}</span>
          </div>

          {slide.subtitle ? <p className="headline headline--lower">{slide.subtitle}</p> : null}
          <p className="note">Вот что было вам особенно интересно в разное время года</p>

          <ul className="seasons">
            {slide.periods.map((period) => (
              <li key={period.period} className="seasons__card">
                <p className={`seasons__title seasons__title--${period.period}`}>
                  {SEASON_TITLES[period.period]}
                </p>
                <img className="seasons__art" src={artUrl(`season-${period.period}`)} alt="" />
                <p className="seasons__category">{period.category.title}</p>
                {period.subcategory ? (
                  <p className="seasons__note">{period.subcategory.title}</p>
                ) : null}
                {typeof period.weight === 'number' ? (
                  <p className="seasons__weight">↗ {period.weight}% активности сезона</p>
                ) : null}
              </li>
            ))}
          </ul>
        </section>
      );

    case 'archetype':
      return (
        <SlideLayout slide={slide}>
          <p className="hero">
            вы — <span className="accent">{slide.archetype.title.toLowerCase()}</span>
          </p>
          <p className="note note--wide">{slide.archetype.description}</p>

          <p className="section-title">Что это значит</p>
          <ul className="reasons">
            {slide.archetype.reasons.map((reason) => (
              <li key={reason.metric} className="reasons__item">
                <span className="reasons__value">{reason.value}</span>
                <span className="reasons__text">{reason.explanation}</span>
              </li>
            ))}
          </ul>

          <p className="hint">
            <span className="hint__title">Это не тест</span>
            Мы определили ваш тип по вашим действиям за год. И он может меняться — это нормально!
          </p>
        </SlideLayout>
      );

    case 'final':
      return (
        <section className="slide slide--wide">
          <div className="final__top">
            <h2 className="hero">
              Вот это год!
              <br />
              <span className="accent">{slide.subtitle ?? 'Вы всё делали по-своему'} ♡</span>
            </h2>

            <img className="final__mascot" src={artUrl('mascot')} alt="" />
          </div>

          <div className="final">
            {slide.stats?.length ? (
              <ul className="stats">
                {slide.stats.map((stat) => (
                  <li key={stat.code} className="stats__item">
                    <img className="stats__art" src={artUrl(STAT_ART[stat.code] ?? 'final')} alt="" />
                    <span className="stats__value">{formatNumber(stat.value)}</span>
                    <span className="stats__label">{stat.label}</span>
                  </li>
                ))}
              </ul>
            ) : null}

            <div className="final__share">
              <p className="final__share-title">Поделитесь своими итогами года!</p>

              {slide.actions?.map((action) => (
                <ActionButton key={action.action} cta={action} onShare={onShare} />
              ))}

              <p className="note">Похвастайтесь своими результатами 🤩</p>
            </div>
          </div>
        </section>
      );

    default:
      // Контракт ушёл вперёд: незнакомый слайд лучше пропустить, чем ронять историю.
      return null;
  }
}

function ActionButton({ cta, onShare }: { cta: Cta; onShare: () => void }) {
  if (cta.action === 'share_recap') {
    return (
      <button type="button" className="button button--accent" onClick={onShare}>
        ⤴ {cta.title}
      </button>
    );
  }

  return (
    <span className="button button--outline" role="note">
      {cta.title}
    </span>
  );
}

function BadgePanel({
  caption,
  title,
  note,
  amount,
}: {
  caption: string;
  title: string;
  note?: string;
  amount?: string;
}) {
  return (
    <div className="panel">
      <p className="panel__caption">{caption}</p>
      <img className="panel__medal" src={artUrl('badge')} alt="" />
      <p className="panel__title">{title}</p>
      {note ? <p className="panel__note">{note}</p> : null}

      {amount ? (
        <div className="panel__amount">
          <span className="panel__amount-caption">Примерная сумма продаж</span>
          <span className="panel__amount-value">{amount}</span>
        </div>
      ) : null}
    </div>
  );
}
