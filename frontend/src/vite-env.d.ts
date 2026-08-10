/// <reference types="vite/client" />

/**
 * Vite объявляет env с индексной сигнатурой `any`, из-за чего любое чтение
 * переменной теряет тип. Здесь перечислено то, что реально используется.
 */
interface ImportMetaEnv {
  readonly VITE_API_BASE_URL?: string;
  readonly VITE_RECAP_YEAR?: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
