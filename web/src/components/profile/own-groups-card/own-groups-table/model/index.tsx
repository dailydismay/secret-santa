import { Group, listOwn } from "api/groups";
import { createEffect, createStore, forward } from "effector";
import { createGate } from "effector-react";

export const $fetchOwnGroupsFx = createEffect(async () => {
  const data = await listOwn();

  return data;
});

export const $ownGroups = createStore<Group[] | null>(null).on(
  $fetchOwnGroupsFx.doneData,
  (_, p) => p
);

export const $isLoading = $fetchOwnGroupsFx.pending;

export const $listOwnGroupsGate = createGate();

forward({
  from: $listOwnGroupsGate.open,
  to: $fetchOwnGroupsFx,
});
