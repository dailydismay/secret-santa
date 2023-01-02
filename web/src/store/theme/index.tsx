import { createTheme, Theme } from "@nextui-org/react";
import { createEffect, createStore } from "effector";

export const darkTheme = createTheme({
  type: "dark",
});

export const lightTheme = createTheme({
  type: "default",
});

export enum ThemeKind {
  Dark,
  Light,
}

export interface ActiveTheme {
  kind: ThemeKind;
  value: Theme;
}

export const $themeSwitcherClicked = createEffect();

export const $theme = createStore({
  kind: ThemeKind.Dark,
  value: darkTheme,
}).on($themeSwitcherClicked, ({ kind }) =>
  kind === ThemeKind.Dark
    ? { kind: ThemeKind.Light, value: lightTheme }
    : { kind: ThemeKind.Dark, value: darkTheme }
);
