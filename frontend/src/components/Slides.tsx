import type {
  Archetype,
  ArchetypeCode,
  Badge,
  Cta,
  ListingRef,
  PeriodInterest,
  Season,
  Slide,
} from '../api/types';

const SEASON_TITLES: Record<Season, string> = {
  winter: 'Зимой',
  spring: 'Весной',
  summer: 'Летом',
  autumn: 'Осенью',
};

/** Все растровые иллюстрации лежат в public/art (источники — в README рядом). */
function artUrl(name: string): string {
  return `/art/${name}.png`;
}

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

function formatPrice(price?: number | null): string {
  if (price == null) {
    return 'Цена не указана';
  }

  return new Intl.NumberFormat('ru-RU', {
    style: 'currency',
    currency: 'RUB',
    maximumFractionDigits: price % 100 === 0 ? 0 : 2,
  }).format(price / 100);
}

export function SlideView({
  slide,
  profileName,
  onShare,
  shareDisabled = false,
  shareLabel,
  archetypeEvidence = false,
  recapYear,
  recapArchetype,
  onExit,
  shareFeedback,
  shareUrl,
}: {
  slide: Slide;
  profileName: string;
  onShare: () => void;
  shareDisabled?: boolean;
  shareLabel?: string;
  archetypeEvidence?: boolean;
  recapYear: number;
  recapArchetype: Archetype;
  onExit: () => void;
  shareFeedback?: { message: string; failed: boolean };
  shareUrl?: string;
}) {
  switch (slide.type) {
    case 'intro':
      return (
        <section className="slide slide--intro" aria-labelledby="intro-title">
          <div className="slide__main intro__copy">
            <p className="intro__act">ВАШ ГОД СОБРАН</p>
            <h1 className="intro__title" id="intro-title">
              {profileName}, ваш {slide.year} год на Авито собран
            </h1>
            <p className="intro__subtitle">
              В конце узнаете свой тип — сначала посмотрим, из чего он сложился.
            </p>
          </div>

          <IntroCover year={slide.year} />
        </section>
      );

    case 'active_days':
      return (
        <section className="slide slide--active-days" aria-labelledby="active-days-title">
          <div className="slide__main active-days__copy">
            <p className="active-days__act">ВАШ РИТМ</p>

            <div className="active-days__metric">
              <strong>{formatNumber(slide.activeDays)}</strong>
              <h1 id="active-days-title">активных дней</h1>
            </div>

            {slide.subtitle ? <p className="active-days__headline">{slide.subtitle}</p> : null}
            <p className="active-days__note">
              Дни, когда вы смотрели, сохраняли, общались или совершали сделки.
            </p>
          </div>

          <div className="active-days__visual" aria-hidden="true">
            <img src={artUrl('active_days')} alt="" />
          </div>
        </section>
      );

    case 'views':
      return (
        <section className="slide slide--views" aria-labelledby="views-title">
          <div className="slide__main views__copy">
            <p className="views__act">МАСШТАБ ПОИСКА</p>

            <div className="views__metric">
              <strong>{formatNumber(slide.views)}</strong>
              <h1 id="views-title">просмотров</h1>
            </div>

            {slide.subtitle ? <p className="views__headline">{slide.subtitle}</p> : null}
            <p className="views__note">
              Столько раз объявления попадали в поле вашего внимания.
            </p>
          </div>

          <div className="views__visual" aria-hidden="true">
            <img className="views__art" src={artUrl('views')} alt="" />
          </div>
        </section>
      );

    case 'messages':
      return (
        <section className="slide slide--messages" aria-labelledby="messages-title">
          <div className="slide__main messages__copy">
            <p className="messages__act">ДИАЛОГИ ГОДА</p>

            <div className="messages__metric">
              <strong>{formatNumber(slide.messages)}</strong>
              <h1 id="messages-title">{slide.subtitle ?? 'сообщений'}</h1>
            </div>

            <p className="messages__note">
              Вопросы, ответы и договорённости, которые помогали двигаться к сделке.
            </p>
          </div>

          <div className="messages__visual" aria-hidden="true">
            <img className="messages__art" src={artUrl('messages')} alt="" />
          </div>
        </section>
      );

    case 'favorites':
      return (
        <section
          className={`slide slide--favorites${slide.oldestFavorite ? ' slide--favorites-featured' : ''}`}
          aria-labelledby="favorites-title"
        >
          <div className="slide__main favorites__copy">
            <p className="favorites__act">ИЗБРАННОЕ</p>

            <div className="favorites__count">
              <strong>{formatNumber(slide.favorites)}</strong>
              {slide.oldestFavorite ? (
                <p>добавлений в избранное</p>
              ) : (
                <h1 id="favorites-title">товаров в избранном</h1>
              )}
            </div>

            {!slide.oldestFavorite && slide.subtitle ? (
              <p className="favorites__headline">{slide.subtitle}</p>
            ) : null}

            {slide.stillAvailable ? (
              <p className="favorites__availability">
                Ещё доступно объявлений: {slide.stillAvailable}
              </p>
            ) : null}
          </div>

          {slide.oldestFavorite ? (
            <div className="favorites__feature">
              <h1 className="favorites__feature-label" id="favorites-title">
                ГЛАВНАЯ НАХОДКА
              </h1>
              <ListingCard
                listing={slide.oldestFavorite}
                caption="Дольше всего в избранном"
                variant="feature"
              />
            </div>
          ) : (
            <div className="favorites__visual" aria-hidden="true">
              <img className="favorites__art" src={artUrl('favorites')} alt="" />
            </div>
          )}
        </section>
      );

    case 'purchases':
      return (
        <section className="slide slide--purchases" aria-labelledby="purchases-title">
          <div className="slide__main purchases__copy">
            <p className="purchases__act">НАХОДКИ СТАЛИ ВАШИМИ</p>

            <div className="purchases__metric">
              <strong>{formatNumber(slide.purchases)}</strong>
              <h1 id="purchases-title">{slide.subtitle ?? 'покупок'}</h1>
            </div>

            <p className="purchases__note">
              Столько находок перешли из поиска в вашу историю.
            </p>
          </div>

          <div className="purchases__visual">
            <AchievementCard
              badge={slide.badge}
              art="badge"
              caption={slide.badge ? 'Ваше достижение' : 'Итог покупок'}
              title={slide.badge ? undefined : `${formatNumber(slide.purchases)} покупок`}
              featured
            />
          </div>
        </section>
      );

    case 'sales':
      return (
        <section className="slide slide--sales" aria-labelledby="sales-title">
          <div className="slide__main sales__copy">
            <p className="sales__act">НОВЫЕ ХОЗЯЕВА</p>

            <div className="sales__metric">
              <strong>{formatNumber(slide.sales)}</strong>
              <h1 id="sales-title">{slide.subtitle ?? 'продаж'}</h1>
            </div>

            <p className="sales__note">Ваши вещи нашли новых владельцев.</p>
          </div>

          <div className="sales__visual">
            <AchievementCard
              badge={slide.badge}
              art="badge-sales"
              caption={slide.badge ? 'Ваше достижение' : 'Итог продаж'}
              title={slide.badge ? undefined : `${formatNumber(slide.sales)} продаж`}
              amount={slide.amountRange?.label}
              featured
            />
          </div>
        </section>
      );

    case 'favorite_category': {
      const recommendations = slide.recommendations?.slice(0, 3) ?? [];
      const hasRecommendations = recommendations.length > 0;

      return (
        <section
          className={`slide slide--category ${hasRecommendations ? 'slide--category-with-recommendations' : 'slide--category-solo'}`}
          aria-labelledby="category-title"
        >
          <div className="category__answer">
            <p className="category__act">ГЛАВНЫЙ ИНТЕРЕС</p>
            <p className="category__lead">Чаще всего вас приводило сюда</p>
            <h1 className="category__title" id="category-title">
              {slide.category.title}
            </h1>

            {slide.subcategory ? (
              <p className="category__subcategory">
                и подкатегория <strong>{slide.subcategory.title}</strong>
              </p>
            ) : null}

            {typeof slide.share === 'number' ? (
              <div className="category__share">
                <p>
                  <strong>{slide.share}%</strong>
                  <span>активности</span>
                </p>
                <div
                  className="category__meter"
                  role="meter"
                  aria-label={`Доля активности ${slide.share}%`}
                  aria-valuenow={slide.share}
                  aria-valuemin={0}
                  aria-valuemax={100}
                >
                  <span style={{ width: `${slide.share}%` }} />
                </div>
              </div>
            ) : null}
          </div>

          {hasRecommendations ? (
            <>
              <div className="category-recommendation-stamp" aria-hidden="true">
                <span>ПОДБОРКА ГОТОВА</span>
                <strong>{recommendations.length}</strong>
                <p>{formatRecommendationCount(recommendations.length)} по вашему интересу</p>
              </div>
              <div className="category__recommendations">
                <div className="category__recommendation-list">
                  {recommendations.map((listing) => (
                    <ListingCard key={listing.id} listing={listing} variant="compact" />
                  ))}
                </div>
                {(slide.recommendations?.length ?? 0) > 1 ? (
                  <p className="category__more">
                    Ещё {Math.min((slide.recommendations?.length ?? 1) - 1, 2)}
                  </p>
                ) : null}
              </div>
            </>
          ) : (
            <div className="category__art" aria-hidden="true">
              <img src={artUrl('favorites')} alt="" />
            </div>
          )}
        </section>
      );
    }

    case 'interests':
      return (
        <section className="slide slide--interests" aria-labelledby="interests-title">
          <header className="interests__header">
            <p className="interests__act">КАК МЕНЯЛИСЬ ИНТЕРЕСЫ</p>
            <h1 id="interests-title">{slide.subtitle ?? 'Год был разным'}</h1>
            <p>Вот что вам было особенно интересно в разное время года.</p>
          </header>

          <ul className="interests__periods">
            {slide.periods.map((period) => (
              <SeasonCard key={period.period} period={period} />
            ))}
          </ul>
        </section>
      );

    case 'archetype':
      return (
        <section
          className={`slide slide--archetype${archetypeEvidence ? ' slide--archetype-evidence' : ''}`}
          aria-labelledby="archetype-title"
        >
          <div className="slide__main archetype__copy">
            <p className="archetype__act">ВАШ АРХЕТИП</p>
            <h1 id="archetype-title">
              {profileName}, вы — <span>{slide.archetype.title}</span>
            </h1>

            {archetypeEvidence ? (
              <div className="archetype__evidence">
                <p className="archetype__description">{slide.archetype.description}</p>
                <ul className="archetype__reasons">
                  {slide.archetype.reasons.map((reason) => (
                    <li key={reason.metric}>
                      <strong>{reason.value}</strong>
                      <span>{reason.explanation}</span>
                    </li>
                  ))}
                </ul>
                <p className="archetype__disclaimer">
                  Это не психологический тест. Тип определён по вашим действиям за год и может
                  меняться.
                </p>
              </div>
            ) : null}
          </div>

          <div className="archetype__visual" aria-hidden="true">
            <div className="archetype__sheet">
              <img src={artUrl(`archetype-${slide.archetype.code}`)} alt="" />
            </div>
          </div>
        </section>
      );

    case 'final': {
      const primaryAction = selectPrimaryAction(slide.actions, recapArchetype.code);
      const shareAction = slide.actions?.find((action) => action.action === 'share_recap');

      return (
        <section className="slide slide--final" aria-labelledby="final-title">
          <div className="final-summary">
            <header className="final-summary__header">
              <p className="final-summary__act">ВАШИ ИТОГИ {recapYear}</p>
              <h1 id="final-title">
                {recapArchetype.title}: вот каким был ваш {recapYear} год
              </h1>
              <p className="final-summary__subtitle">
                {slide.subtitle ?? recapArchetype.description}
              </p>
            </header>

            <div className="final-summary__actions">
              <div className="final-summary__artwork" aria-hidden="true">
                <img
                  className="final-summary__art"
                  src={artUrl(`archetype-${recapArchetype.code}`)}
                  alt=""
                />
              </div>
              <p className="final-summary__context">
                Итоги готовы — можно продолжить с того, что вам ближе.
              </p>

              {primaryAction ? <NextStep cta={primaryAction} /> : null}

              {shareAction ? (
                <ActionButton
                  cta={shareAction}
                  onShare={onShare}
                  disabled={shareDisabled}
                  label={shareLabel ?? 'Поделиться итогами'}
                />
              ) : null}

              <p
                className={`final-summary__feedback${shareFeedback?.failed ? ' final-summary__feedback--error' : ''}`}
                role="status"
                aria-live="polite"
                aria-atomic="true"
              >
                {shareFeedback?.message ?? '\u00a0'}
              </p>

              {shareUrl ? (
                <a className="final-summary__public-link" href={shareUrl}>
                  Открыть публичную карточку
                </a>
              ) : null}

              <button type="button" className="final-summary__exit" onClick={onExit}>
                Выбрать другой профиль
              </button>
            </div>

            {slide.stats?.length ? (
              <ul
                className={`final-summary__stats final-summary__stats--${Math.min(slide.stats.length, 4)}`}
              >
                {slide.stats.slice(0, 4).map((stat) => (
                  <StatTile key={stat.code} stat={stat} />
                ))}
              </ul>
            ) : null}
          </div>
        </section>
      );
    }

    default:
      // Контракт ушёл вперёд: незнакомый слайд лучше пропустить, чем ронять историю.
      return null;
  }
}

function IntroCover({ year }: { year: number }) {
  return (
    <div className="intro__visual" aria-hidden="true">
      <div className="intro__cover">
        <span className="intro__cover-label">Личный альбом</span>
        <strong>{year}</strong>
        <span className="intro__cover-wordmark">Авито</span>
      </div>
    </div>
  );
}

const BADGE_LEVEL_LABELS: Record<NonNullable<Badge['level']>, string> = {
  bronze: 'Бронза',
  silver: 'Серебро',
  gold: 'Золото',
};

function AchievementCard({
  badge,
  art,
  caption = 'Ваше достижение',
  title,
  amount,
  featured = false,
}: {
  badge?: Badge | null;
  art?: string;
  caption?: string;
  title?: string;
  amount?: string;
  featured?: boolean;
}) {
  const cardTitle = badge?.title ?? title;

  return (
    <article
      className={`achievement-card${featured ? ' achievement-card--featured' : ''}${art ? ' achievement-card--with-art' : ''}`}
    >
      <p className="achievement-card__caption">{caption}</p>
      {art ? <img className="achievement-card__art" src={artUrl(art)} alt="" /> : null}
      <div className="achievement-card__copy">
        {badge?.level ? (
          <span className={`achievement-card__level achievement-card__level--${badge.level}`}>
            {BADGE_LEVEL_LABELS[badge.level]}
          </span>
        ) : null}
        {cardTitle ? <h2>{cardTitle}</h2> : null}
        {badge?.description ? <p>{badge.description}</p> : null}
      </div>

      {amount ? <AmountRange value={amount} /> : null}
    </article>
  );
}

function AmountRange({ value }: { value: string }) {
  return (
    <div className="achievement-card__amount">
      <span>Примерная сумма продаж</span>
      <strong>{value}</strong>
    </div>
  );
}

function SeasonCard({ period }: { period: PeriodInterest }) {
  return (
    <li className={`season-card season-card--${period.period}`}>
      <div className="season-card__media" aria-hidden="true">
        <img src={artUrl(`season-${period.period}`)} alt="" />
      </div>
      <div className="season-card__copy">
        <p className="season-card__season">{SEASON_TITLES[period.period]}</p>
        <h2>{period.category.title}</h2>
        {period.subcategory ? <p className="season-card__subcategory">{period.subcategory.title}</p> : null}
      </div>
      {typeof period.weight === 'number' ? (
        <p className="season-card__weight">&gt; {period.weight}% активности сезона</p>
      ) : null}
    </li>
  );
}

function StatTile({ stat }: { stat: NonNullable<Extract<Slide, { type: 'final' }>['stats']>[number] }) {
  return (
    <li className="final-stat">
      <img src={artUrl(STAT_ART[stat.code] ?? 'final')} alt="" />
      <strong>{formatNumber(stat.value)}</strong>
      <span>{stat.label}</span>
    </li>
  );
}

function selectPrimaryAction(actions: Cta[] | undefined, archetype: ArchetypeCode): Cta | undefined {
  if (!actions?.length) {
    return undefined;
  }

  const productActions = actions.filter((action) => action.action !== 'share_recap');
  const preferredAction: Partial<Record<ArchetypeCode, Cta['action']>> = {
    collector: 'open_favorites',
    dealmaker: 'create_listing',
    explorer: 'open_category',
    negotiator: 'open_category',
  };

  return (
    productActions.find((action) => action.action === preferredAction[archetype]) ??
    productActions[0]
  );
}

function NextStep({ cta }: { cta: Cta }) {
  return (
    <div className="final-summary__next-step">
      <span>Следующий шаг</span>
      {cta.url ? (
        <a className="button final-summary__primary" href={cta.url}>
          {cta.title}
        </a>
      ) : (
        <span className="button final-summary__primary" role="note">
          {cta.title}
        </span>
      )}
    </div>
  );
}

function ActionButton({
  cta,
  onShare,
  disabled,
  label,
}: {
  cta: Cta;
  onShare: () => void;
  disabled: boolean;
  label?: string;
}) {
  if (cta.action === 'share_recap') {
    return (
      <button type="button" className="button button--accent" onClick={onShare} disabled={disabled}>
        ⤴ {label ?? cta.title}
      </button>
    );
  }

  return (
    <span className="button button--outline" role="note">
      {cta.title}
    </span>
  );
}

function ListingCard({
  listing,
  caption,
  variant,
}: {
  listing: ListingRef;
  caption?: string;
  variant: 'feature' | 'compact';
}) {
  const addedAt = formatListingDate(listing.addedAt);

  return (
    <article className={`listing-card listing-card--${variant}`}>
      {caption ? <p className="listing-card__caption">{caption}</p> : null}
      <div className="listing-card__image">
        {listing.imageUrl ? (
          <img src={listing.imageUrl} alt={listing.title} />
        ) : (
          <ListingImageFallback title={listing.title} />
        )}
      </div>
      <p className="listing-card__title">{listing.title}</p>
      <p className="listing-card__price">{formatPrice(listing.price)}</p>
      {addedAt ? <p className="listing-card__note">В избранном с {addedAt}</p> : null}
    </article>
  );
}

function formatListingDate(value?: string): string | null {
  if (!value) {
    return null;
  }

  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return null;
  }

  return new Intl.DateTimeFormat('ru-RU').format(date);
}

function formatRecommendationCount(value: number): string {
  const form = new Intl.PluralRules('ru-RU').select(value);

  if (form === 'one') {
    return 'свежее объявление';
  }
  if (form === 'few') {
    return 'свежих объявления';
  }

  return 'свежих объявлений';
}

function ListingImageFallback({ title }: { title: string }) {
  const monogram = title
    .split(/\s+/)
    .filter(Boolean)
    .slice(0, 2)
    .map((part) => part[0]?.toLocaleUpperCase('ru-RU') ?? '')
    .join('');

  return (
    <span className="listing-card__image-fallback">
      <span className="listing-card__image-monogram" aria-hidden="true">
        {monogram || 'А'}
      </span>
      <small>объявление</small>
    </span>
  );
}
