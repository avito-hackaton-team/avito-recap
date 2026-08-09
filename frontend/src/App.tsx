import { useCallback, useEffect, useState } from 'react';
import { createRecap, getRecap, userMessage } from './api/client';
import type { Profile, Recap } from './api/types';
import { StartScreen } from './screens/StartScreen';
import { StoryScreen } from './screens/StoryScreen';
import { PublicRecapScreen } from './screens/PublicRecapScreen';

const RECAP_YEAR = Number(import.meta.env.VITE_RECAP_YEAR ?? new Date().getFullYear() - 1);

type Screen =
  | { name: 'start' }
  | { name: 'generating'; profile: Profile }
  | { name: 'story'; profile: Profile; recap: Recap }
  | { name: 'failed'; profile: Profile; message: string };

// Моковые данные для PublicRecapScreen
const mockPublicRecap = {
  id: 'mock-recap-123',
  year: 2025,
  displayName: 'Алексей',
  archetype: {
    title: 'Исследователь',
    description:
      'Ты любишь искать новое и находить интересные варианты. Твой год был полон открытий!',
  },
  activeDays: 187,
  views: 3456,
  topCategory: {
    categoryTitle: 'Электроника',
    subcategoryTitle: 'Смартфоны и гаджеты',
  },
  interestSummary:
    'В этом году ты активно интересовался технологиями, искал новый телефон и сравнивал характеристики.',
  badges: [
    {
      id: 'badge-1',
      title: 'Первопроходец',
      description: 'Ты один из первых, кто заглянул в новый раздел "Технологии будущего"',
      level: 'gold',
      iconUrl: '🏆',
    },
  ],
};

export function App() {
  // 👇 ВРЕМЕННО: показываем PublicRecapScreen для отладки
  // После проверки закомментируйте эту строку и раскомментируйте код ниже
  return <PublicRecapScreen recap={mockPublicRecap} isAuthorized={false} />;

  // 👇 ВЕСЬ ОСТАЛЬНОЙ КОД ВРЕМЕННО ЗАКОММЕНТИРОВАН
  /*
  const [screen, setScreen] = useState<Screen>({ name: 'start' });

  const start = useCallback((profile: Profile) => {
    setScreen({ name: 'generating', profile });
  }, []);

  const backToStart = useCallback(() => {
    setScreen({ name: 'start' });
  }, []);

  useEffect(() => {
    if (screen.name !== 'generating') {
      return;
    }

    const { profile } = screen;
    let cancelled = false;

    const generate = async () => {
      try {
        const recapId = await createRecap(profile.id, RECAP_YEAR);
        const recap = await getRecap(recapId);

        if (!cancelled) {
          setScreen({ name: 'story', profile, recap });
        }
      } catch (cause: unknown) {
        console.error(cause);

        if (!cancelled) {
          setScreen({ name: 'failed', profile, message: userMessage(cause) });
        }
      }
    };

    void generate();

    return () => {
      cancelled = true;
    };
  }, [screen]);

  switch (screen.name) {
    case 'start':
      return <StartScreen year={RECAP_YEAR} onStart={start} />;

    case 'generating':
      return (
        <div className="state">
          <div className="state__spinner" aria-hidden="true" />
          <p className="state__title">Собираем ваш {RECAP_YEAR} год…</p>
          <p className="state__note">Считаем просмотры, сделки и любимые категории</p>
        </div>
      );

    case 'failed':
      return (
        <div className="state">
          <p className="state__title">Итоги не собрались</p>
          <p className="state__note">{screen.message}</p>
          <button type="button" className="button button--light" onClick={backToStart}>
            Выбрать другой профиль
          </button>
        </div>
      );

    case 'story':
      return <StoryScreen recap={screen.recap} profile={screen.profile} onExit={backToStart} />;
  }
  */
}
