import {
  CalendarDays,
  Check,
  Copy,
  Eye,
  Heart,
  Search,
  Sparkles,
  Tag,
  ArrowRight,
} from 'lucide-react';
import { useState, useEffect } from 'react';
import type { SharedRecap } from '../api/types';
import { BadgeCard } from '../components/BadgeCard';
import { MetricCard } from '../components/MetricCard';

export function PublicRecapScreen({
  recap,
  isAuthorized = false,
}: {
  recap: SharedRecap;
  isAuthorized?: boolean;
}) {
  const [copied, setCopied] = useState(false);
  const [url, setUrl] = useState('');

  useEffect(() => {
    if (typeof window !== 'undefined') {
      setUrl(window.location.href);
    }
  }, []);

  const share = async () => {
    if (!url) {
      console.warn('URL ещё не доступен');
      return;
    }

    try {
      if (navigator.share) {
        await navigator.share({
          title: `Итоги ${recap.year} — Авито`,
          text: `${recap.displayName} — ${recap.archetype.title}`,
          url,
        });
        return;
      }

      await navigator.clipboard.writeText(url);
      setCopied(true);
      setTimeout(() => setCopied(false), 2200);
    } catch (err) {
      console.warn('Не удалось поделиться:', err);

      try {
        await navigator.clipboard.writeText(url);
        setCopied(true);
        setTimeout(() => setCopied(false), 2200);
      } catch {
        const userInput = prompt('Скопируйте ссылку на итоги:', url);
        if (userInput !== null) {
          setCopied(true);
          setTimeout(() => setCopied(false), 2200);
        }
      }
    }
  };

  const b = recap.badges?.[0];
  const cat = recap.topCategory;

  return (
    <main className="page">
      <header className="topbar">
        <a className="brand" href="/">
          ●●● <b>Авито</b>
        </a>
        <span>{recap.year}</span>
      </header>

      <section className="hero">
        <div>
          <p className="eyebrow">
            <Sparkles size={16} /> ИТОГИ ГОДА
          </p>
          <h1>
            Привет, {recap.displayName || 'друг'}!<br />
            <span>Вот каким был твой {recap.year}</span>
          </h1>
          <p className="hero__desc">
            Ты искал, выбирал, сравнивал и открывал новое. Собрали несколько моментов, которые
            рассказывают о твоём годе на Авито.
          </p>
          <div className="archetype">
            <div className="archetype__icon">
              <Search />
            </div>
            <div>
              <small>ТВОЙ ТИП</small>
              <strong>{recap.archetype?.title || 'Исследователь'}</strong>
              <p>
                {recap.archetype?.description ||
                  'Ты любишь искать новое и находить интересные варианты.'}
              </p>
            </div>
          </div>
        </div>
        <div className="hero__visual">
          <div className="orb" />
          <div className="hero__card">
            <Heart fill="currentColor" />
            <span>🛍️</span>
            <b>Твой год</b>
          </div>
          <i>✦</i>
        </div>
      </section>

      <section className="grid">
        <div className="main">
          <section className="card">
            <div className="heading">
              <div>
                <small>В ЦИФРАХ</small>
                <h2>Ты был в движении</h2>
              </div>
              <span>📈</span>
            </div>
            <div className="metrics">
              <MetricCard
                icon={<CalendarDays />}
                value={Math.max(0, recap.activeDays ?? 0)}
                label="активных дней"
                accent="orange"
              />
              <MetricCard
                icon={<Eye />}
                value={Math.max(0, recap.views ?? 0)}
                label="просмотров"
                accent="blue"
              />
            </div>
          </section>

          <section className="card">
            <div className="heading">
              <div>
                <small>ГЛАВНЫЙ ИНТЕРЕС</small>
                <h2>Чаще всего ты заглядывал сюда</h2>
              </div>
              <Tag />
            </div>
            {cat?.categoryTitle ? (
              <div className="category">
                <span>📱</span>
                <div>
                  <b>{cat.categoryTitle}</b>
                  <small>{cat.subcategoryTitle || 'Категория года'}</small>
                </div>
              </div>
            ) : (
              <div className="empty">🧭 Категория пока не определена</div>
            )}
          </section>

          <section className="card interest">
            <strong>“</strong>
            <div>
              <small>ТВОИ ИНТЕРЕСЫ</small>
              <h2>Год был разным</h2>
              <p>{recap.interestSummary?.trim() || 'Интересы этого года пока не определены.'}</p>
            </div>
          </section>
        </div>

        <aside>
          <section className="card achievement">
            <div className="heading">
              <div>
                <small>ДОСТИЖЕНИЕ</small>
                <h2>Есть чем поделиться</h2>
              </div>
              <span>🏆</span>
            </div>
            {b ? (
              <BadgeCard
                title={b.title}
                description={b.description}
                level={b.level}
                iconUrl={b.iconUrl}
              />
            ) : (
              <div className="empty">✨ Достижение скоро появится</div>
            )}
            <p className="muted">
              Показываем одно главное достижение. Остальные варианты уже предусмотрены контрактом.
            </p>
          </section>

          <section className="share">
            <small>ПОНРАВИЛОСЬ?</small>
            <h2>Поделись своими итогами</h2>
            <p>Ссылка откроет эту публичную карточку без авторизации.</p>
            <button onClick={share} disabled={!url}>
              {copied ? (
                <>
                  <Check /> Ссылка скопирована
                </>
              ) : (
                <>
                  <Copy /> Поделиться
                </>
              )}
            </button>
            {isAuthorized ? (
              <button className="outline">
                Открыть полную версию <ArrowRight />
              </button>
            ) : null}
          </section>
        </aside>
      </section>

      <footer>
        Итоги {recap.year}
        <span>Публичная карточка</span>
      </footer>
    </main>
  );
}
