import type { DayActivity, Slide } from '../api/types';

export const WEEKS_IN_MONTH = 5;

export interface YearColumn {
  key: string;
  actions: number;
  level: number;
  peak: boolean;
}

export interface YearMonthRow {
  month: number;
  columns: YearColumn[];
  actions: number;
  activeDays: number;
}

const SUPPORTED_SLIDE_TYPES = new Set<Slide['type']>([
  'intro',
  'active_days',
  'views',
  'favorites',
  'favorite_category',
  'purchases',
  'sales',
  'messages',
  'interests',
  'archetype',
  'final',
]);

export function isSupportedSlide(slide: Slide): boolean {
  return SUPPORTED_SLIDE_TYPES.has(slide.type);
}

export function formatRecommendationFindingCount(value: number): string {
  return ['находок', 'одна находка', 'две находки', 'три находки'][value] ?? `${value} находок`;
}

export function buildYearRows(days: DayActivity[], peak: DayActivity | null): YearMonthRow[] {
  if (days.length === 0) {
    return [];
  }

  const byMonth = new Map<number, Map<number, number>>();

  for (const day of days) {
    const parsed = parseDate(day.date);
    if (!parsed) {
      continue;
    }

    const month = byMonth.get(parsed.month) ?? new Map<number, number>();
    month.set(parsed.period, (month.get(parsed.period) ?? 0) + day.actions);
    byMonth.set(parsed.month, month);
  }

  if (byMonth.size === 0) {
    return [];
  }

  const peakDate = peak ? parseDate(peak.date) : null;
  const max = Math.max(...[...byMonth.values()].flatMap((month) => [...month.values()]));
  const rows: YearMonthRow[] = [];

  for (let month = 0; month < 12; month += 1) {
    const slots = byMonth.get(month);
    const columns: YearColumn[] = [];
    let actions = 0;

    for (let period = 0; period < WEEKS_IN_MONTH; period += 1) {
      const value = slots?.get(period) ?? 0;
      actions += value;
      columns.push({
        key: `${month}-${period}`,
        actions: value,
        level: intensity(value, max),
        peak: peakDate !== null && peakDate.month === month && peakDate.period === period,
      });
    }

    rows.push({
      month,
      columns,
      actions,
      activeDays: countActiveDays(days, month),
    });
  }

  return rows;
}

function countActiveDays(days: DayActivity[], month: number): number {
  return days.reduce((count, day) => {
    const parsed = parseDate(day.date);
    return parsed && parsed.month === month && day.actions > 0 ? count + 1 : count;
  }, 0);
}

function intensity(value: number, max: number): number {
  if (value <= 0 || max <= 0) {
    return 0;
  }

  return Math.min(4, Math.max(1, Math.ceil((value / max) * 4)));
}

function parseDate(date: string): { month: number; period: number } | null {
  const match = /^(\d{4})-(\d{2})-(\d{2})$/.exec(date);
  if (!match?.[2] || !match[3]) {
    return null;
  }

  const year = Number(match[1]);
  const month = Number(match[2]) - 1;
  const day = Number(match[3]);
  const parsed = new Date(Date.UTC(year, month, day));

  if (
    month < 0 ||
    month > 11 ||
    day < 1 ||
    parsed.getUTCFullYear() !== year ||
    parsed.getUTCMonth() !== month ||
    parsed.getUTCDate() !== day
  ) {
    return null;
  }

  return { month, period: Math.min(WEEKS_IN_MONTH - 1, Math.floor((day - 1) / 7)) };
}
