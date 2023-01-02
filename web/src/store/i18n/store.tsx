import { Locale } from "date-fns";
import { enUS, ru } from "date-fns/locale";

import { combine, createStore } from "effector";
import { persist } from "effector-storage/local";
import { $setLanguage } from "./events";
import { Language, TranslationItem, TranslationsDict } from "./model";

export const $language = createStore<Language>("en");

persist({
  store: $language,
  key: "language",
});

export const $translations = createStore<TranslationsDict>({
  en: {
    [TranslationItem.NavBrand]: "Secret Santa",
    [TranslationItem.NavCreateGroup]: "Create",
    [TranslationItem.NavGroups]: "Groups",
    [TranslationItem.NavProfile]: "Profile",
    [TranslationItem.NavLogout]: "Logout",
    [TranslationItem.LoginToContinue]: "Login to Continue",
    [TranslationItem.LoginButton]: "Use VK Account",
    [TranslationItem.AccountCreatedAt]: "You joined us",
    [TranslationItem.VisitCreatorGithub]: "Visit Creator GitHub Profile"
  },
  ru: {
    [TranslationItem.NavBrand]: "Тайный Санта",
    [TranslationItem.NavCreateGroup]: "Создать лобби",
    [TranslationItem.NavGroups]: "Ваши лобби",
    [TranslationItem.NavProfile]: "Профиль",
    [TranslationItem.NavLogout]: "Выйти",
    [TranslationItem.LoginToContinue]: "Войдите, чтобы продолжить",
    [TranslationItem.LoginButton]: "Продолжить с  ВКонтакте",
    [TranslationItem.AccountCreatedAt]: "Ты присоединился к нам",
    [TranslationItem.VisitCreatorGithub]: "Посетить GitHub Профиль автора"
  },
});

export const $translation = combine($language, $translations, (a, t) => t[a]);

$translation.on($setLanguage, (_, l) => $translations.getState()[l]);
$language.on($setLanguage, (_, p) => p);

export const $dateFnsLocale = $language.map((x) => {
  const mapping: { [key in Language]: Locale } = {
    en: enUS,
    ru,
  };

  return mapping[x] ?? enUS;
});
