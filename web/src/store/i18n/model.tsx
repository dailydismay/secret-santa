export type Language = "en" | "ru";

export enum TranslationItem {
  NavBrand,
  NavCreateGroup,
  NavGroups,
  NavProfile,
  NavLogout,
  LoginToContinue,
  LoginButton,
  AccountCreatedAt,
  VisitCreatorGithub,
}

export type TranslationsDict = Record<
  Language,
  Record<TranslationItem, string>
>;
