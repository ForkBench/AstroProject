import { writable } from 'svelte/store';
import type { Stage } from '../bindings/changeme/astro/services/models';
import { type Competition } from '../bindings/changeme/astro/services/models';

export const SelectedCompetition = writable<Competition | undefined>(undefined);
export const CurrentStage = writable<Stage | undefined>(undefined);
export const Competitions = writable<(Competition | undefined)[]>([]);