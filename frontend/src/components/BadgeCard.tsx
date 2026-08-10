import { ShieldCheck, Star } from 'lucide-react';
import type { LucideIcon } from 'lucide-react';

import type { BadgeLevel } from '../api/types';

const META: Record<BadgeLevel, { label: string; icon: LucideIcon }> = {
  bronze: { label: 'Бронза', icon: ShieldCheck },
  silver: { label: 'Серебро', icon: ShieldCheck },
  gold: { label: 'Золото', icon: Star },
};

interface BadgeCardProps {
  title: string;
  description?: string | null;
  level?: BadgeLevel | null;
  iconUrl?: string | null;
}

export function BadgeCard({ title, description, level, iconUrl }: BadgeCardProps) {
  const badgeLevel = level ?? 'bronze';
  const meta = META[badgeLevel];
  const Icon = meta.icon;

  return (
    <article className={`badge badge--${badgeLevel}`}>
      <div className="badge__icon">
        {iconUrl ? <img src={iconUrl} alt="" /> : <Icon aria-hidden="true" />}
      </div>
      <div>
        <span className="badge__level">{meta.label}</span>
        <h3>{title || 'Достижение'}</h3>
        {description ? <p>{description}</p> : null}
      </div>
    </article>
  );
}
