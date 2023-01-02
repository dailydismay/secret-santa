import { createEffect, createStore, guard } from "effector";
import { createGate } from "effector-react";
import { persist } from "effector-storage/local";
import { fetchProfile, loginWithCode } from "api/auth";
import { TokensPair, Visitor } from "./types";
import { setToken } from "api/client";
import { $logoutClicked } from "./events";

export const $fetchProfileFx = createEffect(async () => {
  const profile = await fetchProfile();

  return profile;
});

export const $loginWithCodeGate = createGate<string>();

export const $loginWithCodeFx = createEffect(async (code: string) => {
  const data = await loginWithCode(code);

  return data;
});

guard({
  source: $loginWithCodeGate.open,
  filter: $loginWithCodeFx.pending.map((x) => !x),
  target: $loginWithCodeFx,
});

export const $token = createStore<TokensPair>({
  accessToken: null,
  refreshToken: null,
})
  .on($loginWithCodeFx.doneData, (s, p) => ({
    ...s,
    accessToken: p.accessToken,
  }))
  .reset($loginWithCodeFx.fail)
  .reset($logoutClicked);

$token.watch((x) => {
  setToken(x.accessToken);
});

persist({
  store: $token,
  key: "token",
});

export const $visitor = createStore<Visitor | null>(null)
  .on($loginWithCodeFx.doneData, (_, p) => p.user)
  .on($fetchProfileFx.doneData, (_, p) => p)
  .reset($logoutClicked);

export const $isLoggedIn = $token.map((x) => !!x.accessToken);

export const $fetchVisitor = createGate();

guard({
  clock: $fetchVisitor.open,
  source: $isLoggedIn,
  filter: Boolean,
  target: $fetchProfileFx,
});
