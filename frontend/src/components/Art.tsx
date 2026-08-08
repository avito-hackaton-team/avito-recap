import { useState } from 'react';

import type { SlideType } from '../api/types';

/**
 * Заглушки под 3D-иллюстрации из макета. Если в public/art/<type>.png лежит
 * настоящий ассет — показывается он; пока файла нет, рисуется крупный глиф.
 * Так вёрстка не ждёт дизайнера и не ломается, когда ассеты появятся.
 */
const FALLBACK_GLYPH: Record<SlideType, string> = {
  intro: '🎉',
  active_days: '📅',
  views: '🔍',
  favorites: '❤️',
  favorite_category: '📱',
  purchases: '🛍️',
  sales: '🏷️',
  messages: '💬',
  interests: '📊',
  archetype: '🧑‍💻',
  final: '✨',
};

const CONFETTI = ['star', 'cube', 'ball', 'star', 'cube'] as const;

/**
 * Форматы пробуются по очереди: сначала растровый ассет (Fluent Emoji 3D),
 * потом свой SVG, и только если нет ни одного — глиф.
 */
const FORMATS = ['png', 'svg'] as const;

export function Art({ type }: { type: SlideType }) {
  const [attempt, setAttempt] = useState(0);

  const format = FORMATS[attempt];

  return (
    <div className="art" aria-hidden="true">
      <span className="art__glow" />

      {CONFETTI.map((shape, index) => (
        <span key={`${shape}-${index}`} className={`art__confetti art__confetti--${shape}-${index}`} />
      ))}

      {format ? (
        <img
          className="art__image"
          src={`/art/${type}.${format}`}
          alt=""
          onError={() => setAttempt((current) => current + 1)}
        />
      ) : (
        <span className="art__glyph">{FALLBACK_GLYPH[type]}</span>
      )}
    </div>
  );
}
