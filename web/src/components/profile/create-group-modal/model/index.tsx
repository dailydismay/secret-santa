import { createGroup } from "api/groups";
import { createEffect, createEvent, createStore, forward } from "effector";
import { CreateGroupClickedPayload } from "./types";

export const $createGroupClicked = createEvent<CreateGroupClickedPayload>();
export const $groupCreated = createEvent<string>();
export const $groupModalToggled = createEvent<boolean>();

export const $createGroupFx = createEffect(
  async (payload: CreateGroupClickedPayload) => {
    const { id } = await createGroup(payload.title);
    return id;
  }
);
export const $isCreateGroupModalOpen = createStore(false);

$isCreateGroupModalOpen
  .on($groupModalToggled, (_, p) => p)
  .on($groupCreated, () => false);

forward({
  from: $createGroupClicked,
  to: $createGroupFx,
});

forward({
  from: $createGroupFx.doneData,
  to: $groupCreated,
});
