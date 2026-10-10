import { api } from '../../store/api';

export interface MeditationPhase {
  label: string;
  seconds: number;
  speak: string;
}
export interface MeditationScript {
  repeat: number;
  phases: MeditationPhase[];
}
export interface MeditationType {
  id: number;
  slug: string;
  name: string;
  description: string;
  icon: string;
  script: MeditationScript;
}
export interface Session {
  id: number;
  meditationTypeId: number;
  typeName: string;
  durationSeconds: number;
  completedAt: string;
  notes: string;
}
export interface HabitStats {
  totalSessions: number;
  totalMinutes: number;
  currentStreakDays: number;
  longestStreakDays: number;
  byType: Record<string, number>;
  last30Days: { date: string; count: number }[];
}

export const meditationApi = api.injectEndpoints({
  endpoints: (build) => ({
    listTypes: build.query<MeditationType[], void>({
      query: () => '/meditation/types',
      providesTags: ['Types'],
    }),
    listSessions: build.query<Session[], void>({
      query: () => '/meditation/sessions?limit=50',
      providesTags: ['Sessions'],
    }),
    createSession: build.mutation<Session, { meditationTypeId: number; durationSeconds: number; notes?: string }>({
      query: (body) => ({ url: '/meditation/sessions', method: 'POST', body }),
      invalidatesTags: ['Sessions', 'Stats'],
    }),
    getStats: build.query<HabitStats, void>({
      query: () => '/meditation/stats',
      providesTags: ['Stats'],
    }),
  }),
});

export const {
  useListTypesQuery,
  useListSessionsQuery,
  useCreateSessionMutation,
  useGetStatsQuery,
} = meditationApi;