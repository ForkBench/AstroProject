import { writable } from 'svelte/store';
import type { Competition } from '../bindings/changeme/astro/services/models';

export const SelectedCompetition = writable<Competition | undefined>(undefined);