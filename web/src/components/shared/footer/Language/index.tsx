import { CSS, Text } from "@nextui-org/react";
import { useEvent, useStore } from "effector-react";
import { $language, $setLanguage, Language } from "store/i18n";

const langCss: CSS = {
  mr: "$3",
  cursor: "pointer",
  fontWeight: "bold",
};

const activeLangCss: CSS = {
  ...langCss,
  textGradient: "45deg, $purple600 -20%, $pink600 100%",
};

export const LanguageButton: React.FC<{ language: Language }> = ({
  language,
}) => {
  const $languageButtonStyle = $language.map((state) =>
    state === language ? activeLangCss : langCss
  );

  const css = useStore($languageButtonStyle);
  const setLanguage = useEvent($setLanguage);

  return (
    <Text onClick={() => setLanguage(language)} h5 css={css}>
      {language.toUpperCase()}
    </Text>
  );
};
