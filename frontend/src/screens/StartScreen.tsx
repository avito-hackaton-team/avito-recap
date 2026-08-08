import { useEffect, useState } from 'react';

import { listProfiles, userMessage } from '../api/client';
import type { Profile } from '../api/types';

interface StartScreenProps {
  year: number;
  onStart: (profile: Profile) => void;
}

export function StartScreen({ year, onStart }: StartScreenProps) {
  const [profiles, setProfiles] = useState<Profile[]>([]);
  const [selected, setSelected] = useState<Profile | null>(null);
  const [pickerOpen, setPickerOpen] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let cancelled = false;

    listProfiles()
      .then((items) => {
        if (cancelled) {
          return;
        }

        setProfiles(items);
        setSelected(items[0] ?? null);
      })
      .catch((cause: unknown) => {
        if (!cancelled) {
          console.error(cause);
          setError(userMessage(cause));
        }
      })
      .finally(() => {
        if (!cancelled) {
          setLoading(false);
        }
      });

    return () => {
      cancelled = true;
    };
  }, []);

  return (
    <div className="start">
      <header className="start__top">
        <span className="start__privacy">🔒 Это ваши личные итоги. Никто их не увидит</span>

        <div className="picker">
          <button
            type="button"
            className="picker__toggle"
            onClick={() => setPickerOpen((open) => !open)}
            disabled={loading || profiles.length === 0}
          >
            <span className="picker__avatar" aria-hidden="true">
              {selected?.name.slice(0, 1) ?? '—'}
            </span>
            {selected ? selected.name : 'Профиль'} ⌄
          </button>

          {pickerOpen ? (
            <div className="picker__menu">
              <p className="picker__caption">Тестовые профили</p>

              <ul className="picker__list">
                {profiles.map((profile) => (
                  <li key={profile.id}>
                    <button
                      type="button"
                      className={`picker__item${selected?.id === profile.id ? ' picker__item--active' : ''}`}
                      onClick={() => {
                        setSelected(profile);
                        setPickerOpen(false);
                      }}
                    >
                      <span className="picker__avatar" aria-hidden="true">
                        {profile.name.slice(0, 1)}
                      </span>
                      <span className="picker__text">
                        <span className="picker__name">
                          {profile.name} {profile.surname}
                        </span>
                        {profile.hint ? <span className="picker__hint">{profile.hint}</span> : null}
                      </span>
                      {selected?.id === profile.id ? <span aria-hidden="true">✓</span> : null}
                    </button>
                  </li>
                ))}
              </ul>
            </div>
          ) : null}
        </div>
      </header>

      <div className="start__hero">
        <div className="start__content">
          <h1 className="hero hero--start">
            Вспомни, каким был <span className="hero__year">{year}</span> год вместе с Авито
          </h1>

          <p className="note note--wide">
            Мы собрали для вас самые важные моменты вашей активности на площадке за этот год.
          </p>

          {error ? <p className="start__error">{error}</p> : null}
          {loading ? <p className="note">Загружаем профили…</p> : null}
          {!loading && !error && profiles.length === 0 ? (
            <p className="note">Профилей нет — наполните базу тестовыми данными.</p>
          ) : null}

          <button
            type="button"
            className="button button--light button--large"
            disabled={!selected || loading}
            onClick={() => selected && onStart(selected)}
          >
            Посмотреть итоги года →
          </button>

          <p className="note">🕐 Это займёт около 1–2 минут</p>
        </div>

        <div className="start__art" aria-hidden="true">
          <span className="start__year">{String(year).slice(0, 2)}</span>
          <span className="start__year start__year--second">{String(year).slice(2)}</span>
        </div>
      </div>
    </div>
  );
}
