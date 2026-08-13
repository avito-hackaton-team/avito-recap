import type { Archetype, ArchetypeCode, Cta, Slide } from '../api/types';
import styles from './FinalRecapSlide.module.css';

type FinalSlide = Extract<Slide, { type: 'final' }>;
type FavoriteCategorySlide = Extract<Slide, { type: 'favorite_category' }>;

const HEADLINES: Record<ArchetypeCode, string> = {
  collector: 'Год, в котором важные находки оставались рядом',
  dealmaker: 'Год, в котором вещи находили новых владельцев',
  explorer: 'Год, в котором вы не переставали искать новое',
  negotiator: 'Год, в котором вы всегда находили общий язык',
};

const STAT_ART: Record<string, string> = {
  active_days: 'active_days',
  views: 'views',
  favorites: 'favorites',
  messages: 'messages',
  seasons: 'season-spring',
};

function artUrl(name: string): string {
  return `/art/${name}.png`;
}

function formatNumber(value: number): string {
  return new Intl.NumberFormat('ru-RU').format(value);
}

function selectPrimaryAction(
  actions: Cta[] | undefined,
  archetype: ArchetypeCode,
): Cta | undefined {
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

interface FinalRecapSlideProps {
  slide: FinalSlide;
  year: number;
  archetype: Archetype;
  favoriteCategory?: FavoriteCategorySlide;
  interestSummary?: string;
  onShare: () => void;
  shareDisabled: boolean;
  shareLabel?: string;
  shareFeedback?: { message: string; failed: boolean };
  shareUrl?: string;
  onExit: () => void;
}

export function FinalRecapSlide({
  slide,
  year,
  archetype,
  favoriteCategory,
  interestSummary,
  onShare,
  shareDisabled,
  shareLabel,
  shareFeedback,
  shareUrl,
  onExit,
}: FinalRecapSlideProps) {
  const primaryAction = selectPrimaryAction(slide.actions, archetype.code);
  const shareAction = slide.actions?.find((action) => action.action === 'share_recap');

  return (
    <section className={styles.slide} aria-labelledby="final-title">
      <div className={styles.cover}>
        <div className={styles.content}>
          <header className={styles.header}>
            <p className={styles.label}>ВАШИ ИТОГИ {year}</p>
            <h1 id="final-title">{HEADLINES[archetype.code]}</h1>
            <p className={styles.subtitle}>{archetype.description}</p>
          </header>

          {slide.stats?.length ? (
            <ul className={styles.stats} data-count={Math.min(slide.stats.length, 4)}>
              {slide.stats.slice(0, 4).map((stat) => (
                <li key={stat.code}>
                  <div className={styles.statArt} aria-hidden="true">
                    <img src={artUrl(STAT_ART[stat.code] ?? 'final')} alt="" />
                  </div>
                  <strong>{formatNumber(stat.value)}</strong>
                  <span>{stat.label}</span>
                </li>
              ))}
            </ul>
          ) : null}

          <div className={styles.actions}>
            {shareAction ? (
              <button type="button" onClick={onShare} disabled={shareDisabled}>
                {shareLabel ?? 'Поделиться итогами ↗'}
              </button>
            ) : null}

            {primaryAction ? (
              primaryAction.url ? (
                <a href={primaryAction.url}>{primaryAction.title}</a>
              ) : (
                <span role="note">{primaryAction.title}</span>
              )
            ) : null}

            <button type="button" className={styles.exit} onClick={onExit}>
              Другой профиль
            </button>

            <p
              className={shareFeedback?.failed ? styles.feedbackError : styles.feedback}
              role="status"
              aria-live="polite"
              aria-atomic="true"
            >
              {shareFeedback?.message ?? '\u00a0'}
            </p>

            {shareUrl ? (
              <a className={styles.publicLink} href={shareUrl}>
                Открыть публичную карточку
              </a>
            ) : null}
          </div>
        </div>

        <aside className={styles.visualPanel} aria-label="Архетип и главный интерес года">
          <div className={styles.visualHeading}>
            <span>АРХЕТИП ГОДА</span>
            <strong aria-hidden="true">{archetype.title.charAt(0).toLocaleUpperCase('ru-RU')}</strong>
          </div>

          <figure className={styles.archetypeVisual}>
            <img
              src={artUrl(`archetype-${archetype.code}`)}
              alt=""
              aria-hidden="true"
            />
            <figcaption>
              <span>Ваш архетип</span>
              <strong>{archetype.title}</strong>
            </figcaption>
          </figure>

          {favoriteCategory ? (
            <div className={styles.interestCard}>
              <strong>Главный интерес — {favoriteCategory.category.title}</strong>
              <span>{interestSummary ?? favoriteCategory.subcategory?.title}</span>
            </div>
          ) : null}

          <div className={styles.ring} aria-hidden="true" />
        </aside>
      </div>
    </section>
  );
}
