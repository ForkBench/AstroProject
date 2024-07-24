import { writable } from 'svelte/store';
import type { Stage } from '../bindings/astroproject/astro/structs/models';
import { type Competition } from '../bindings/astroproject/astro/structs/models';

export const SelectedCompetition = writable<Competition>();
export const CurrentStage = writable<Stage>();
export const Competitions = writable<(Competition)[]>([]);