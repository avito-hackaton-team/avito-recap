import { BrandLoader } from '../components/BrandLoader';

export function PublicLoadingScreen() {
  return (
    <main className="public-loading" aria-labelledby="public-loading-title">
      <div className="public-loading__content">
        <BrandLoader />
        <h1 id="public-loading-title">Открываем публичные итоги</h1>
        <p>Проверяем безопасную публичную карточку</p>
        <span className="visually-hidden" role="status">
          Публичные итоги загружаются.
        </span>
      </div>
    </main>
  );
}
