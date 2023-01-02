import { useStore } from "effector-react";
import { TranslationItem } from "./model";
import { $translation } from "./store";

export const useTranslation = () => {
  const translation = useStore($translation);

  return translation;
};
