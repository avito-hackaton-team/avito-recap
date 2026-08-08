import { useState } from 'react';

import { SlideView } from '../components/Slides';
import type { Profile, Recap } from '../api/types';

interface StoryScreenProps {
  recap: Recap;
  profile: Profile;
  onExit: () => void;
}

export function StoryScreen({ recap, profile, onExit }: StoryScreenProps) {
  const [index, setIndex] = useState(0);
  const [shared, setShared] = useState(false);

  const slides = recap.slides;
  const slide = slides[index];
  const isLast = index === slides.length - 1;

  const share = () => {
    const link = `${window.location.origin}${window.location.pathname}?recap=${recap.id}`;

    navigator.clipboard
      .writeText(link)
      .then(() => setShared(true))
      .catch((cause: unknown) => {
        // Буфер обмена недоступен без https — показываем ссылку, копирует сам пользователь.
        console.error(cause);
        window.prompt('Скопируйте ссылку на итоги', link);
      });
  };

  if (!slide) {
    return null;
  }

  return (
    <div className="story">
      <header className="story__top">
        <button type="button" className="story__back" onClick={onExit}>
          ‹ Назад
        </button>

        <div
          className="story__progress"
          role="progressbar"
          aria-valuenow={index + 1}
          aria-valuemin={1}
          aria-valuemax={slides.length}
        >
          {slides.map((item, position) => (
            <span
              key={`${item.type}-${position}`}
              className={`story__segment${position <= index ? ' story__segment--done' : ''}`}
            />
          ))}
        </div>

        <span className="story__counter">
          {index + 1} / {slides.length}
        </span>
      </header>

      <SlideView slide={slide} onShare={share} />

      <footer className="story__bottom">
        <button
          type="button"
          className="story__circle"
          disabled={index === 0}
          aria-label="Предыдущий слайд"
          onClick={() => setIndex((current) => current - 1)}
        >
          ←
        </button>

        <div className="story__dots" aria-hidden="true">
          {slides.map((item, position) => (
            <span
              key={`dot-${item.type}-${position}`}
              className={`story__dot${position === index ? ' story__dot--active' : ''}`}
            />
          ))}
        </div>

        {shared ? <span className="story__toast">Ссылка скопирована</span> : null}

        <button
          type="button"
          className="button button--light"
          onClick={() => (isLast ? onExit() : setIndex((current) => current + 1))}
        >
          {isLast ? `К профилям, ${profile.name}` : 'Дальше →'}
        </button>
      </footer>
    </div>
  );
}
