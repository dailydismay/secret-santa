import { TranslationItem } from "store/i18n";

type NavItem = {
  path: string;
  translationItem: TranslationItem;
};

export const loggedInNavItems: NavItem[] = [
  {
    path: "/groups",
    translationItem: TranslationItem.NavGroups,
  },
];

export const loggedOutNavItems: NavItem[] = [
  {
    path: "/login",
    translationItem: TranslationItem.NavLogin,
  },
];
